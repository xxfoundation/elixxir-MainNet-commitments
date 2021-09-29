package client

import (
	"crypto"
	"crypto/sha256"
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"github.com/pkg/errors"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/crypto/csprng"
	"gitlab.com/xx_network/crypto/signature/rsa"
)

// SignAndTransmit creates a Client object & transmits commitment info to the server
func SignAndTransmit(pk, idfBytes []byte, wallet string, h *connect.Host, s Sender) error {
	key, err := rsa.LoadPrivateKeyFromPem(pk)
	if err != nil {
		return errors.WithMessage(err, "Failed to load private key")
	}
	hasher := sha256.New()
	_, err = hasher.Write(idfBytes)
	if err != nil {
		return errors.WithMessage(err, "Failed to write IDF to hash")
	}
	_, err = hasher.Write([]byte(wallet))
	if err != nil {
		return errors.WithMessage(err, "Failed to write wallet to hash")
	}
	sig, err := rsa.Sign(csprng.NewSystemRNG(), key, crypto.SHA256, hasher.Sum(nil), nil)
	if err != nil {
		return errors.WithMessage(err, "Failed to sign node info")
	}

	err = s.SignAndTransmit(h, &messages.Commitment{
		PrivateKey: pk,
		IDF:        idfBytes,
		Wallet:     wallet,
		Signature:  sig,
	})

	if err != nil {
		return errors.WithMessage(err, "Failed to register commitment")
	}
	return nil
}
