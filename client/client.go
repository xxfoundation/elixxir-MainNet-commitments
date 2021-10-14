///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package client

import (
	"encoding/json"
	"fmt"

	"gitlab.com/xx_network/comms/connect"
	utils2 "gitlab.com/xx_network/primitives/utils"

	"git.xx.network/elixxir/mainnet-commitments/messages"
	"git.xx.network/elixxir/mainnet-commitments/utils"

	"github.com/pkg/errors"

	"gitlab.com/xx_network/crypto/csprng"
	"gitlab.com/xx_network/crypto/signature/rsa"
	"gitlab.com/xx_network/primitives/id"
	"gitlab.com/xx_network/primitives/id/idf"
)

// SignAndTransmit creates a Client object & transmits commitment info to the server
func SignAndTransmit(keyPath, idfPath, contractPath, wallet, commitmentsAddress, commitmentsCertPath string) error {
	var pk, idfBytes, commitmentCert, contractBytes []byte
	var err error
	var ep string

	// Read key file
	if ep, err = utils2.ExpandPath(keyPath); err == nil {
		pk, err = utils2.ReadFile(ep)
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

	if ep, err = utils2.ExpandPath(commitmentsCertPath); err == nil {
		commitmentCert, err = utils2.ReadFile(ep)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	idfStruct := &idf.IdFile{}
	err = json.Unmarshal(idfBytes, idfStruct)
	if err != nil {
		return err
	}

	nodeID, err := id.Unmarshal(idfStruct.IdBytes[:])
	if err != nil {
		return err
	}

	cl, err := StartClient(pk, idfStruct.Salt[:], nodeID)
	if err != nil {
		return err
	}

	hp := connect.GetDefaultHostParams()
	hp.AuthEnabled = false
	h, err := cl.pc.AddHost(&utils.ServerID, commitmentsAddress, commitmentCert, hp)
	if err != nil {
		return err
	}

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

	fmt.Println("Transmitting signature...")
	err = cl.TransmitSignature(h, &messages.Commitment{
		PrivateKey: pk,
		IDF:        idfBytes,
		Contract:   contractBytes,
		Wallet:     wallet,
		Signature:  sig,
	})

	if err != nil {
		return errors.WithMessage(err, "Failed to register commitment")
	}
	return nil
}
