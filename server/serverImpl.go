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
	"encoding/json"
	"encoding/pem"
	"git.xx.network/elixxir/mainnet-commitments/storage"
	"git.xx.network/elixxir/mainnet-commitments/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/xx-labs/sleeve/wallet"
	"gitlab.com/xx_network/crypto/signature/rsa"
	"gitlab.com/xx_network/primitives/id/idf"
	"net/http"
	"testing"
)

// Params struct holds data needed to create a server Impl
type Params struct {
	Key           []byte
	Cert          []byte
	Port          string
	StorageParams storage.Params
}

type Commitment struct {
	IDF       []byte `json:"idf"`
	Contract  []byte `json:"contract"`
	Wallet    string `json:"wallet"`
	Signature []byte `json:"signature"`
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

	r := gin.Default()
	r.POST("/commitment", func(c *gin.Context) {
		var newCommitment Commitment
		if err := c.BindJSON(&newCommitment); err != nil {
			return
		}
		err := impl.Verify(c, newCommitment)
		if err != nil {
			return
		}
		c.IndentedJSON(http.StatusAccepted, newCommitment)
	})
	impl.comms = r
	return r.Run()
}

// Impl struct stores protocomms & storage for server implementation
type Impl struct {
	comms *gin.Engine
	s     *storage.Storage
}

// Verify func is the main endpoint for the mainnet-commitments server
func (i *Impl) Verify(_ context.Context, msg Commitment) error {
	// Load IDF from JSON bytes
	idfStruct := &idf.IdFile{}
	err := json.Unmarshal(msg.IDF, idfStruct)
	if err != nil {
		err = errors.WithMessage(err, "Failed to unmarshal IDF json")
		jww.ERROR.Printf(err.Error())
		return err
	}

	jww.INFO.Printf("Received verification request from %+v", idfStruct.ID)

	ok, err := wallet.ValidateXXNetworkAddress(msg.Wallet)
	if err != nil {
		err = errors.WithMessage(err, "Failed to validate wallet address")
		jww.ERROR.Printf(err.Error())
		return err
	}
	if !ok {
		err = errors.New("Wallet validation returned false")
		jww.ERROR.Printf(err.Error())
		return err
	}

	// Hash node info from message
	hashed, hash, err := utils.HashNodeInfo(msg.Wallet, msg.IDF, msg.Contract)
	if err != nil {
		err = errors.WithMessage(err, "Failed to hash node info")
		jww.ERROR.Printf(err.Error())
		return err
	}

	// Get member info from database
	hexId := "\\" + idfStruct.HexNodeID[1:]
	m, err := i.s.GetMember(hexId)
	if err != nil {
		err = errors.WithMessagef(err, "Member %s [%+v] not found", idfStruct.ID, idfStruct.IdBytes)
		jww.ERROR.Printf(err.Error())
		return err
	}

	block, rest := pem.Decode(m.Cert)
	jww.INFO.Printf("Decoded cert into block: %+v, rest: %+v", block, rest)
	var cert *x509.Certificate
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		err = errors.WithMessage(err, "Failed to load certificate")
		jww.ERROR.Printf(err.Error())
		return err
	}
	rsaPublicKey := cert.PublicKey.(*gorsa.PublicKey)

	// Attempt to verify signature
	err = rsa.Verify(&rsa.PublicKey{PublicKey: *rsaPublicKey}, hash, hashed, msg.Signature, nil)
	if err != nil {
		err = errors.WithMessage(err, "Could not verify node commitment info signature")
		jww.ERROR.Printf(err.Error())
		return err
	}

	// Insert commitment info to the database once verified
	err = i.s.InsertCommitment(storage.Commitment{
		Id:        m.Id,
		Contract:  msg.Contract,
		Wallet:    msg.Wallet,
		Signature: msg.Signature,
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
