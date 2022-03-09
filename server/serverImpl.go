///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package server

import (
	"bytes"
	"context"
	gorsa "crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/csv"
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
	pathutils "gitlab.com/xx_network/primitives/utils"
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
	IDListPath   string
}

// StartServer creates a server object from params
func StartServer(params Params, s *storage.Storage) error {
	impl := &Impl{
		s:            s,
		contractHash: params.ContractHash,
		idList:       map[string]interface{}{},
	}

	// Attempt to load in list of node IDs exempt from duplicate wallet checking
	if p, err := pathutils.ExpandPath(params.IDListPath); err == nil {
		idList, err := pathutils.ReadFile(p)
		if err != nil {
			return errors.WithMessage(err, "Failed to read ID list path")
		}
		r := csv.NewReader(bytes.NewReader(idList))
		records, err := r.ReadAll()
		for _, r := range records {
			nid := r[0]
			impl.idList[nid] = true
		}
	} else {
		return errors.WithMessage(err, "Failed to expand ID list path")
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
	idList       map[string]interface{}
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
	ok, err := wallet.ValidateXXNetworkAddress(msg.PaymentWallet)
	if err != nil {
		err = errors.WithMessage(err, "Failed to validate payment wallet address")
		jww.ERROR.Println(err)
		return err
	}
	if !ok {
		err = errors.New("Payment wallet validation returned false")
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
	hashed, hash, err := utils.HashNodeInfo(msg.PaymentWallet, idfBytes, contractBytes)
	if err != nil {
		err = errors.WithMessage(err, "Failed to hash node info")
		jww.ERROR.Println(err)
		return err
	}

	// Decode certificate & extract public component
	block, rest := pem.Decode(m.Cert)
	jww.INFO.Printf("Decoded cert into block: %+v, rest: %+v", block, rest)
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

	// Check if wallet is in old commitments
	// Only check this if node ID is not in exempt list
	if _, ok := i.idList[nid.String()]; !ok {
		ok, err = i.s.CheckWallet(msg.PaymentWallet)
		if err != nil {
			err = errors.WithMessage(err, "Failed to check wallet status in old commitments")
			jww.ERROR.Println(err)
			return err
		}
		if !ok {
			err = errors.Errorf("Cannot add wallet %s: wallet already used in the old commitments system", msg.PaymentWallet)
			jww.ERROR.Println(err)
			return err
		}
	}

	// Insert commitment info to the database once verified
	err = i.s.InsertCommitment(storage.Commitment{
		Id:        m.Id,
		Contract:  contractBytes,
		Wallet:    msg.PaymentWallet,
		Signature: sigBytes,
	})
	if err != nil {
		err = errors.WithMessage(err, "Failed to insert commitment")
		jww.ERROR.Println(err)
		return err
	}

	jww.INFO.Printf("Registered commitment from %+v [Wallet: %s]", idfStruct.ID, msg.PaymentWallet)
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
