///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package server

import (
	"context"
	"encoding/json"
	"fmt"
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"git.xx.network/elixxir/mainnet-commitments/storage"
	"git.xx.network/elixxir/mainnet-commitments/utils"
	"github.com/pkg/errors"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/crypto/signature/rsa"
	"gitlab.com/xx_network/primitives/id/idf"
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
	pc, _, err := connect.StartCommServer(&utils.ServerID, fmt.Sprintf("0.0.0.0:%s", params.Port),
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

	// Register verify implementation
	messages.RegisterCommitmentsServer(pc.LocalServer, impl)
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
		return nil, errors.WithMessage(err, "Failed to unmarshal IDF json")
	}

	// Hash node info from message
	hashed, hash, err := utils.HashNodeInfo(msg.Wallet, msg.IDF, msg.Contract)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to hash node info")
	}

	// Get member info from database
	m, err := i.s.GetMember(idfStruct.IdBytes[:])
	if err != nil {
		return nil, errors.WithMessagef(err, "Member %s [%+v] not found", idfStruct.ID, idfStruct.IdBytes)
	}

	// Load member certificate
	cert, err := rsa.LoadPublicKeyFromPem(m.Cert)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to load certificate")
	}

	// Attempt to verify signature
	err = rsa.Verify(cert, hash, hashed, msg.Signature, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Could not verify node commitment info signature")
	}

	// Insert commitment info to the database once verified
	err = i.s.InsertCommitment(storage.Commitment{
		Id:        m.Id,
		Contract:  msg.Contract,
		Wallet:    msg.Wallet,
		Signature: msg.Signature,
	})
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
