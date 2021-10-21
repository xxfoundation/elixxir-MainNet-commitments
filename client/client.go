///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package client

import (
	"crypto/tls"
	"encoding/base64"
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"git.xx.network/elixxir/mainnet-commitments/utils"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"gitlab.com/xx_network/crypto/csprng"
	"gitlab.com/xx_network/crypto/signature/rsa"
	utils2 "gitlab.com/xx_network/primitives/utils"
)

func SignAndTransmit(keyPath, idfPath, contractPath, wallet, serverAddress, serverCertPath string) error {
	var key, idfBytes, contractBytes []byte
	var err error
	var ep string

	// Read key file
	if ep, err = utils2.ExpandPath(keyPath); err == nil {
		key, err = utils2.ReadFile(ep)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	// Read id file
	if ep, err = utils2.ExpandPath(idfPath); err == nil {
		idfBytes, err = utils2.ReadFile(ep)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	if ep, err = utils2.ExpandPath(contractPath); err == nil {
		contractBytes, err = utils2.ReadFile(ep)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	return signAndTransmit(key, idfBytes, contractBytes, wallet, serverAddress, serverCertPath)
}

// SignAndTransmit creates a Client object & transmits commitment info to the server
func signAndTransmit(pk, idfBytes, contractBytes []byte, wallet, serverAddress, serverCertPath string) error {
	// Create new resty client
	cl := resty.New()
	cl.SetRootCertificate(serverCertPath) // Set commitments root certificate
	cl.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	// Hash & sign node info
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

	// Build message body & post to server
	body := messages.Commitment{
		IDF:       base64.URLEncoding.EncodeToString(idfBytes),
		Contract:  base64.URLEncoding.EncodeToString(contractBytes),
		Wallet:    wallet,
		Signature: base64.URLEncoding.EncodeToString(sig),
	}
	resp, err := cl.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(messages.Commitment{}).
		Post(serverAddress + "/commitment")

	if err != nil {
		return errors.WithMessagef(err, "Failed to register commitment, received response: %+v", resp)
	}
	return nil
}
