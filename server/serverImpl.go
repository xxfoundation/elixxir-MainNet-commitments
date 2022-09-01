////////////////////////////////////////////////////////////////////////////////
// Copyright © 2022 xx foundation                                             //
//                                                                            //
// Use of this source code is governed by a license that can be found in the  //
// LICENSE file.                                                              //
////////////////////////////////////////////////////////////////////////////////

package server

import (
	"context"
	gorsa "crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"git.xx.network/elixxir/mainnet-commitments/storage"
	"git.xx.network/elixxir/mainnet-commitments/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/xx-labs/sleeve/wallet"
	"gitlab.com/xx_network/crypto/signature/rsa"
	"gitlab.com/xx_network/primitives/id"
	"gitlab.com/xx_network/primitives/id/idf"
	"net/http"
	"testing"
	"time"
)

// Params struct holds data needed to create a server Impl
type Params struct {
	KeyPath      string
	CertPath     string
	ContractHash string
	Port         string
}

// StartServer creates a server object from params
func StartServer(params Params, s *storage.Storage) error {
	impl := &Impl{
		s:            s,
		contractHash: params.ContractHash,
	}

	// Build gin server, link to verify code
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "access-control-allow-origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.POST("/commitment", func(c *gin.Context) {
		jww.DEBUG.Printf("Received commitment request %+v...", c.Request)
		var newCommitment messages.Commitment
		if err := c.BindJSON(&newCommitment); err != nil {
			jww.INFO.Printf("Failed to bind JSON: %+v", err)
			wrappedErr := c.Error(err)
			c.JSON(http.StatusBadRequest, wrappedErr.JSON())
			return
		}
		jww.INFO.Printf("Received commitment request %+v", newCommitment)
		err := impl.Verify(c, newCommitment)
		if err != nil {
			jww.INFO.Printf("Failed to verify commitment: %+v", err)
			wrappedErr := c.Error(err)
			c.JSON(http.StatusForbidden, wrappedErr.JSON())
			return
		}
		c.JSON(http.StatusAccepted, newCommitment)
	})
	r.GET("/info", func(c *gin.Context) {
		jww.DEBUG.Printf("Received info request %+v...", c.Request)
		fmt.Println(c.Request.URL.Query())
		nid := c.Request.URL.Query().Get("id")
		if nid == "" {
			jww.ERROR.Printf("No ID in received request")
			wrappedErr := c.Error(errors.New("No ID in received request"))
			c.JSON(http.StatusBadRequest, wrappedErr.JSON())
			return
		}
		convertedID := "\\" + nid[1:]
		commitment, err := impl.s.GetCommitment(convertedID)
		if err != nil {
			jww.ERROR.Printf("Failed to get commitment for nid %s: %+v", nid, err)
			wrappedErr := c.Error(err)
			c.JSON(http.StatusBadRequest, wrappedErr.JSON())
			return
		}
		c.JSON(http.StatusOK, messages.CommitmentInfo{
			ValidatorWallet: commitment.Wallet,
			NominatorWallet: commitment.NominatorWallet,
			SelectedStake:   commitment.SelectedStake,
			MaxStake:        commitment.MaxStake,
			Email:           commitment.Email,
		})
	})
	impl.comms = r
	// Run with TLS
	if params.KeyPath == "" && params.CertPath == "" {
		jww.WARN.Println("NO TLS CONFIGURED")
		return r.Run(fmt.Sprintf("0.0.0.0:%s", params.Port))
	} else {
		return r.RunTLS(fmt.Sprintf("0.0.0.0:%s", params.Port), params.CertPath, params.KeyPath)
	}
}

// Impl struct stores protocomms & storage for server implementation
type Impl struct {
	comms        *gin.Engine
	s            *storage.Storage
	contractHash string
}

// Verify func is the main endpoint for the mainnet-commitments server
func (i *Impl) Verify(_ context.Context, msg messages.Commitment) error {
	// Load IDF from JSON
	idfStruct := &idf.IdFile{}
	idfBytes, err := base64.URLEncoding.DecodeString(msg.IDF)
	if err != nil {
		err = errors.WithMessage(err, "Failed to decode IDF string")
		jww.ERROR.Println(err)
		return err
	}
	err = json.Unmarshal(idfBytes, idfStruct)
	if err != nil {
		err = errors.WithMessage(err, "Failed to unmarshal IDF json")
		jww.ERROR.Println(err)
		return err
	}
	jww.INFO.Printf("Received verification request from %+v", idfStruct.ID)

	if msg.Contract != i.contractHash {
		err = errors.Errorf("Contract hash received [%+v] does not match valid hash on server [%+v]", msg.Contract, i.contractHash)
		jww.ERROR.Println(err)
		return err
	}

	// Validate wallets
	if msg.NominatorWallet == msg.ValidatorWallet {
		return errors.New("Nominator wallet and validator wallet cannot be the same")
	}

	if msg.NominatorWallet != "" {
		ok, err := wallet.ValidateXXNetworkAddress(msg.NominatorWallet)
		if err != nil {
			err = errors.WithMessage(err, "Failed to validate nominator wallet address")
			jww.ERROR.Println(err)
			return err
		}
		if !ok {
			err = errors.New("Nominator wallet validation returned false")
			jww.ERROR.Println(err)
			return err
		}

	}

	ok, err := wallet.ValidateXXNetworkAddress(msg.ValidatorWallet)
	if err != nil {
		err = errors.WithMessage(err, "Failed to validate validator wallet address")
		jww.ERROR.Println(err)
		return err
	}
	if !ok {
		err = errors.New("Validator wallet validation returned false")
		jww.ERROR.Println(err)
		return err
	}

	// Check hex node ID (betanet nodes don't have this)
	nid, err := id.Unmarshal(idfStruct.IdBytes[:])
	if err != nil {
		err = errors.WithMessage(err, "Failed to unmarshal ID")
		jww.ERROR.Println(err)
		return err
	}

	hexNodeID := nid.HexEncode()

	// Get member info from database
	hexId := "\\" + hexNodeID[1:]
	m, err := i.s.GetMember(hexId)
	if err != nil {
		err = errors.WithMessagef(err, "Member %s [%+v] not found", idfStruct.ID, idfStruct.IdBytes)
		jww.ERROR.Println(err)
		return err
	}

	contractBytes, err := base64.URLEncoding.DecodeString(msg.Contract)
	if err != nil {
		err = errors.WithMessage(err, "Failed to decode contract hash from base64")
		jww.ERROR.Println(err)
		return err
	}

	// Hash node info from message
	hashed, hash, err := utils.HashNodeInfo(msg.NominatorWallet, msg.ValidatorWallet, idfBytes, contractBytes)
	if err != nil {
		err = errors.WithMessage(err, "Failed to hash node info")
		jww.ERROR.Println(err)
		return err
	}

	// Decode certificate & extract public component
	block, _ := pem.Decode(m.Cert)
	var cert *x509.Certificate
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		err = errors.WithMessage(err, "Failed to load certificate")
		jww.ERROR.Println(err)
		return err
	}
	rsaPublicKey := cert.PublicKey.(*gorsa.PublicKey)

	// Decode signature
	sigBytes, err := base64.URLEncoding.DecodeString(msg.Signature)
	if err != nil {
		err = errors.WithMessage(err, "Failed to decode signature from base64")
		jww.ERROR.Println(err)
		return err
	}

	// Attempt to verify signature
	err = rsa.Verify(&rsa.PublicKey{PublicKey: *rsaPublicKey}, hash, hashed, sigBytes, nil)
	if err != nil {
		err = errors.WithMessage(err, "Could not verify node commitment info signature")
		jww.ERROR.Println(err)
		return err
	}

	// Insert commitment info to the database once verified
	c := storage.Commitment{
		Id:        m.Id,
		Contract:  contractBytes,
		Wallet:    msg.ValidatorWallet,
		Signature: sigBytes,
	}
	if msg.SelectedStake != 0 {
		c.SelectedStake = msg.SelectedStake
	}
	if msg.Email != "" {
		c.Email = msg.Email
	}
	if msg.NominatorWallet != "" {
		c.NominatorWallet = msg.NominatorWallet
	}
	err = i.s.InsertCommitment(c)
	if err != nil {
		err = errors.WithMessage(err, "Failed to insert commitment")
		jww.ERROR.Println(err)
		return err
	}

	jww.INFO.Printf("Registered commitment from %+v [Nominator: %+s, Validator: %s]", idfStruct.ID, msg.NominatorWallet, msg.ValidatorWallet)
	return nil
}

func (i *Impl) GetStorage() *storage.Storage {
	return i.s
}

func (i *Impl) SetStorage(t *testing.T, s *storage.Storage) {
	if t == nil {
		panic("Cannot set storage on impl outside of testing")
	}
	i.s = s
}
