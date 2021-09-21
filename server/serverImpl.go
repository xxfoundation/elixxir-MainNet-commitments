package server

import (
	"context"
	"crypto"
	"crypto/sha256"
	"encoding/json"
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"git.xx.network/elixxir/mainnet-commitments/storage"
	"github.com/pkg/errors"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/crypto/signature/rsa"
	"gitlab.com/xx_network/primitives/id"
	"gitlab.com/xx_network/primitives/id/idf"
)

// Params struct holds data needed to create a server Impl
type Params struct {
	Key           []byte
	Cert          []byte
	Address       string
	StorageParams storage.Params
}

// StartServer creates a server object from params
func StartServer(params Params) (*Impl, error) {
	pc, _, err := connect.StartCommServer(&id.Permissioning, params.Address, params.Cert, params.Key, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to start comms server")
	}
	s, err := storage.NewStorage(params.StorageParams)
	impl := &Impl{
		pc: pc,
		s:  s,
	}
	messages.RegisterCommitmentsServer(pc.LocalServer, impl)
	return impl, nil
}

// Impl struct stores protocomms & storage for server implementation
type Impl struct {
	pc *connect.ProtoComms
	s  *storage.Storage
}

// Verify func is the main endpoint for the mainnet-commitments server
func (i *Impl) Verify(_ context.Context, msg *messages.Commitment) (*messages.Ack, error) {
	// Load IDF from JSON bytes
	idfStruct := &idf.IdFile{}
	err := json.Unmarshal(msg.IDF, idfStruct)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to unmarshal IDF json")
	}

	// Get member info from database
	m, err := i.s.GetMember(idfStruct.IdBytes[:])
	if err != nil {
		return nil, errors.WithMessagef(err, "Member %s [%+v] not found", idfStruct.ID, idfStruct.IdBytes)
	}

	// Hash IDF & Wallet
	hasher := sha256.New()
	_, err = hasher.Write(msg.IDF)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to write IDF to hasher")
	}
	_, err = hasher.Write([]byte(msg.Wallet))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to write wallet to hasher")
	}

	// Load certificate from database
	cert, err := rsa.LoadPublicKeyFromPem(m.Cert)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to load certificate")
	}

	// Attempt to verify signature
	err = rsa.Verify(cert, crypto.SHA256, hasher.Sum(nil), msg.Signature, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Could not verify node commitment info signature")
	}

	// Insert commitment info to the database once verified
	err = i.s.InsertCommitment(storage.Commitment{
		Id:        m.Id,
		Wallet:    msg.Wallet,
		Signature: msg.Signature,
	})
	return &messages.Ack{}, nil
}
