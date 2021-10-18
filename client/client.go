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
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"gitlab.com/xx_network/crypto/csprng"
	"gitlab.com/xx_network/crypto/signature/rsa"
)

// SignAndTransmit creates a Client object & transmits commitment info to the server
func SignAndTransmit(pk, idfBytes, contractBytes []byte, wallet, serverCertPath, serverAddress string) error {
	cl := resty.New()
	cl.SetRootCertificate(serverCertPath)
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
	body := messages.Commitment{
		IDF:       idfBytes,
		Contract:  contractBytes,
		Wallet:    wallet,
		Signature: sig,
	}
	resp, err := cl.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(messages.Commitment{}).
		SetError(nil).
		Post(serverAddress)

	if err != nil {
		return errors.WithMessagef(err, "Failed to register commitment, received response: %+v", resp)
	}
	return nil
}
