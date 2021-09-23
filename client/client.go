package client

import (
	"crypto"
	"crypto/sha256"
	"encoding/json"
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/crypto/csprng"
	"gitlab.com/xx_network/crypto/signature/rsa"
	"gitlab.com/xx_network/primitives/id"
	"gitlab.com/xx_network/primitives/id/idf"
)

// SignAndTransmit creates a Client object & transmits commitment info to the server
func SignAndTransmit(pk, cert, idfBytes, commitmentCert []byte, wallet, address string) error {
	idfStruct := &idf.IdFile{}
	err := json.Unmarshal(idfBytes, idfStruct)
	if err != nil {
		return err
	}
	nodeID, err := id.Unmarshal(idfStruct.IdBytes[:])
	if err != nil {
		return errors.WithMessage(err, "Failed to unmarshal ID from IDF")
	}

	h, err := connect.NewHost(nil, address, commitmentCert, connect.GetDefaultHostParams())
	if err != nil {
		return err
	}

	cl, err := StartClient(pk, cert, idfStruct.Salt[:], nodeID)
	if err != nil {
		jww.FATAL.Panicf("Failed to start client: %+v", err)
	}

	key, err := rsa.LoadPrivateKeyFromPem(pk)
	if err != nil {
		jww.FATAL.Panicf("Failed to load private key: %+v", err)
	}
	hasher := sha256.New()
	_, err = hasher.Write(idfBytes)
	if err != nil {
		jww.FATAL.Panicf("Failed to write idf to hasher: %+v", err)
	}
	_, err = hasher.Write([]byte(wallet))
	if err != nil {
		jww.FATAL.Panicf("Failed to write wallet to hasher: %+v", err)
	}
	sig, err := rsa.Sign(csprng.NewSystemRNG(), key, crypto.SHA256, hasher.Sum(nil), nil)
	if err != nil {
		jww.FATAL.Panicf("Failed to sign node info: %+v", err)
	}

	err = cl.SignAndTransmit(h, &messages.Commitment{
		PrivateKey: pk,
		IDF:        idfBytes,
		Wallet:     wallet,
		Signature:  sig,
	})
	if err != nil {
		jww.FATAL.Panicf("Error in registering commitment: %+v", err)
	}
	return nil
}
