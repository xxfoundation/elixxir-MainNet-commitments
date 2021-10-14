///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package client

import (
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"git.xx.network/elixxir/mainnet-commitments/utils"
	"github.com/pkg/errors"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/crypto/csprng"
	"gitlab.com/xx_network/crypto/signature/rsa"
)

// SignAndTransmit creates a Client object & transmits commitment info to the server
func SignAndTransmit(pk, idfBytes, contractBytes []byte, wallet string, h *connect.Host, s Sender) error {
	key, err := rsa.LoadPrivateKeyFromPem(pk)
	if err != nil {
		return errors.WithMessage(err, "Failed to load private key")
	}

	hashed, hash, err := utils.HashNodeInfo(wallet, idfBytes, contractBytes)
	if err != nil {
		return errors.WithMessage(err, "Failed to hash node info")
	}

	sig, err := rsa.Sign(csprng.NewSystemRNG(), key, hash, hashed, nil)
	if err != nil {
		return errors.WithMessage(err, "Failed to sign node info")
	}

	err = s.TransmitSignature(h, &messages.Commitment{
		IDF:       idfBytes,
		Contract:  contractBytes,
		Wallet:    wallet,
		Signature: sig,
	})

	if err != nil {
		return errors.WithMessage(err, "Failed to register commitment")
	}
	return nil
}
