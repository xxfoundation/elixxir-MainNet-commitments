///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package client

import (
	"encoding/json"
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"git.xx.network/elixxir/mainnet-commitments/server"
	"git.xx.network/elixxir/mainnet-commitments/storage"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/crypto/csprng"
	"gitlab.com/xx_network/crypto/signature/rsa"
	"gitlab.com/xx_network/primitives/id"
	"gitlab.com/xx_network/primitives/id/idf"
	"testing"
)

type MockSender struct {
	t        *testing.T
	id, cert []byte
}

func (ms *MockSender) TransmitSignature(host *connect.Host, message *messages.Commitment) error {
	s, err := storage.NewStorage(storage.Params{})
	if err != nil {
		ms.t.Error("Failed to init storage for mock server")
	}
	err = s.InsertMembers([]storage.Member{{
		Id:   ms.id,
		Cert: ms.cert,
	},
	})
	if err != nil {
		ms.t.Errorf("Failed to insert members: %+v", err)
	}
	impl := server.Impl{}
	impl.SetStorage(ms.t, s)
	_, err = impl.Verify(nil, message)
	if err != nil {
		ms.t.Errorf("Failed to verify: %+v", err)
	}
	return nil
}

func TestSignAndTransmit(t *testing.T) {
	pk, err := rsa.GenerateKey(csprng.NewSystemRNG(), 2048)
	if err != nil {
		t.Errorf("Failed to gen key: %+v", err)
	}
	nid := id.NewIdFromString("zezima", id.Node, t)
	idb := [33]byte{}
	copy(idb[:], nid.Marshal())
	idFile := idf.IdFile{
		ID:        nid.String(),
		Type:      nid.GetType().String(),
		Salt:      [32]byte{},
		IdBytes:   idb,
		HexNodeID: nid.HexEncode(),
	}
	idfBytes, err := json.Marshal(idFile)
	if err != nil {
		t.Errorf("Failed to marshal IDF: %+v", err)
	}

	contractBytes := []byte("I solemnly swear that I am up to no good")
	err = SignAndTransmit(rsa.CreatePrivateKeyPem(pk), idfBytes, contractBytes, "wallet", nil, &MockSender{t, nid.Bytes(), rsa.CreatePublicKeyPem(pk.GetPublic())})
	if err != nil {
		t.Errorf("Failed to sign & transmit: %+v", err)
	}
}
