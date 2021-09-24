package client

import (
	"crypto"
	"crypto/sha256"
	"git.xx.network/elixxir/mainnet-commitments/messages"
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/crypto/csprng"
	"gitlab.com/xx_network/crypto/signature/rsa"
)

// SignAndTransmit creates a Client object & transmits commitment info to the server
func SignAndTransmit(pk, idfBytes []byte, wallet string, h *connect.Host, s Sender) error {
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

	err = s.SignAndTransmit(h, &messages.Commitment{
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
