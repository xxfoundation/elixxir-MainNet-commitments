///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

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
	KeyPath       string
	CertPath      string
	Port          string
	StorageParams storage.Params
}

// StartServer creates a server object from params
func StartServer(params Params) error {
	// initialize storage
	s, err := storage.NewStorage(params.StorageParams)
	if err != nil {
		return err
	}
	impl := &Impl{
		s: s,
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
			_ = c.Error(err)
			c.Status(http.StatusBadRequest)
			return
		}
		jww.INFO.Printf("Received commitment request %+v", newCommitment)
		err := impl.Verify(c, newCommitment)
		if err != nil {
			jww.INFO.Printf("Failed to verify commitment: %+v", err)
			_ = c.Error(err)
			c.Status(http.StatusForbidden)
			return
		}
		c.IndentedJSON(http.StatusAccepted, newCommitment)
	})
	impl.comms = r

	// Run with TLS
	return r.RunTLS(fmt.Sprintf("0.0.0.0:%s", params.Port), params.CertPath, params.KeyPath)
}

// Impl struct stores protocomms & storage for server implementation
type Impl struct {
	comms *gin.Engine
	s     *storage.Storage
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

	ok, err := wallet.ValidateXXNetworkAddress(msg.Wallet)
	if err != nil {
		err = errors.WithMessage(err, "Failed to validate wallet address")
		jww.ERROR.Println(err)
		return err
	}
	if !ok {
		err = errors.New("Wallet validation returned false")
		jww.ERROR.Println(err)
		return err
	}

	// Hash node info from message
	contractBytes, err := base64.URLEncoding.DecodeString(msg.Contract)
	if err != nil {
		err = errors.WithMessage(err, "Failed to decode contract from base64")
	}
	hashed, hash, err := utils.HashNodeInfo(msg.Wallet, idfBytes, contractBytes)
	if err != nil {
		err = errors.WithMessage(err, "Failed to hash node info")
		jww.ERROR.Println(err)
		return err
	}

	if idfStruct.HexNodeID == "" {
		nid, err := id.Unmarshal(idfStruct.IdBytes[:])
		if err != nil {
			err = errors.WithMessage(err, "Failed to unmarshal ID")
			jww.ERROR.Println(err)
			return err
		}

		idfStruct.HexNodeID = nid.HexEncode()
	}

	// Get member info from database
	hexId := "\\" + idfStruct.HexNodeID[1:]
	m, err := i.s.GetMember(hexId)
	if err != nil {
		err = errors.WithMessagef(err, "Member %s [%+v] not found", idfStruct.ID, idfStruct.IdBytes)
		jww.ERROR.Println(err)
		return err
	}

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
	err = i.s.InsertCommitment(storage.Commitment{
		Id:        m.Id,
		Contract:  contractBytes,
		Wallet:    msg.Wallet,
		Signature: sigBytes,
	})
	jww.INFO.Printf("Registered commitment from %+v [%+v]", idfStruct.ID, msg.Wallet)
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
