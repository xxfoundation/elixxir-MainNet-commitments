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
	"fmt"
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"git.xx.network/elixxir/mainnet-commitments/storage"
	"git.xx.network/elixxir/mainnet-commitments/utils"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/xx-labs/sleeve/wallet"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/crypto/signature/rsa"
	"gitlab.com/xx_network/primitives/id/idf"
	"google.golang.org/grpc/reflection"
	"testing"
)

// Params struct holds data needed to create a server Impl
type Params struct {
	Key           []byte
	Cert          []byte
	Port          string
	StorageParams storage.Params
}

// StartServer creates a server object from params
func StartServer(params Params) (*Impl, error) {
	// Create grpc server
	addr := fmt.Sprintf("0.0.0.0:%s", params.Port)
	pc, lis, err := connect.StartCommServer(&utils.ServerID, addr,
		params.Cert, params.Key, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to start comms server")
	}

	// initialize storage
	s, err := storage.NewStorage(params.StorageParams)
	impl := &Impl{
		pc: pc,
		s:  s,
	}

	go func() {
		messages.RegisterCommitmentsServer(pc.LocalServer, impl)

		// Register reflection service on gRPC server.
		reflection.Register(pc.LocalServer)
		if err := pc.LocalServer.Serve(lis); err != nil {
			err = errors.New(err.Error())
			jww.FATAL.Panicf("Failed to serve: %+v", err)
		}
		jww.INFO.Printf("Shutting down registration server listener:"+
			" %s", lis)
	}()

	return impl, nil
}

// Impl struct stores protocomms & storage for server implementation
type Impl struct {
	pc *connect.ProtoComms
	s  *storage.Storage
}

// Verify func is the main endpoint for the mainnet-commitments server
func (i *Impl) Verify(_ context.Context, msg *messages.Commitment) (*messages.CommitmentResponse, error) {
	// Load IDF from JSON bytes
	idfStruct := &idf.IdFile{}
	err := json.Unmarshal(msg.IDF, idfStruct)
	if err != nil {
		err = errors.WithMessage(err, "Failed to unmarshal IDF json")
		jww.ERROR.Printf(err.Error())
		return nil, err
	}

	jww.INFO.Printf("Received verification request from %+v", idfStruct.ID)

	ok, err := wallet.ValidateXXNetworkAddress(msg.Wallet)
	if err != nil {
		err = errors.WithMessage(err, "Failed to validate wallet address")
		jww.ERROR.Printf(err.Error())
		return nil, err
	}
	if !ok {
		err = errors.New("Wallet validation returned false")
		jww.ERROR.Printf(err.Error())
		return nil, err
	}

	// Hash node info from message
	hashed, hash, err := utils.HashNodeInfo(msg.Wallet, msg.IDF, msg.Contract)
	if err != nil {
		err = errors.WithMessage(err, "Failed to hash node info")
		jww.ERROR.Printf(err.Error())
		return nil, err
	}

	// Get member info from database
	hexId := "\\" + idfStruct.HexNodeID[1:]
	m, err := i.s.GetMember(hexId)
	if err != nil {
		err = errors.WithMessagef(err, "Member %s [%+v] not found", idfStruct.ID, idfStruct.IdBytes)
		jww.ERROR.Printf(err.Error())
		return nil, err
	}

	block, rest := pem.Decode(m.Cert)
	jww.INFO.Printf("Decoded cert into block: %+v, rest: %+v", block, rest)
	var cert *x509.Certificate
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		err = errors.WithMessage(err, "Failed to load certificate")
		jww.ERROR.Printf(err.Error())
		return nil, err
	}
	rsaPublicKey := cert.PublicKey.(*gorsa.PublicKey)

	// Attempt to verify signature
	err = rsa.Verify(&rsa.PublicKey{PublicKey: *rsaPublicKey}, hash, hashed, msg.Signature, nil)
	if err != nil {
		err = errors.WithMessage(err, "Could not verify node commitment info signature")
		jww.ERROR.Printf(err.Error())
		return nil, err
	}

	// Insert commitment info to the database once verified
	err = i.s.InsertCommitment(storage.Commitment{
		Id:        m.Id,
		Contract:  msg.Contract,
		Wallet:    msg.Wallet,
		Signature: msg.Signature,
	})
	jww.INFO.Printf("Registered commitment from %+v [%+v]", idfStruct.ID, msg.Wallet)
	return &messages.CommitmentResponse{}, nil
}

func (i *Impl) Stop() {
	i.pc.Shutdown()
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
