///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package client

import (
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"git.xx.network/elixxir/mainnet-commitments/server"
	"git.xx.network/elixxir/mainnet-commitments/storage"
	//"github.com/xx-labs/sleeve/wallet"
	"gitlab.com/xx_network/comms/connect"
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

//func TestSignAndTransmit(t *testing.T) {
//	pk, err := rsa.GenerateKey(csprng.NewSystemRNG(), 2048)
//	if err != nil {
//		t.Errorf("Failed to gen key: %+v", err)
//	}
//	nid := id.NewIdFromString("zezima", id.Node, t)
//	idb := [33]byte{}
//	copy(idb[:], nid.Marshal())
//	idFile := idf.IdFile{
//		ID:        nid.String(),
//		Type:      nid.GetType().String(),
//		Salt:      [32]byte{},
//		IdBytes:   idb,
//		HexNodeID: nid.HexEncode(),
//	}
//	idfBytes, err := json.Marshal(idFile)
//	if err != nil {
//		t.Errorf("Failed to marshal IDF: %+v", err)
//	}
//
//	s, err := wallet.NewSleeve(rand.Reader, "password")
//	if err != nil {
//		t.Errorf("Failed to create sleeve: %+v", err)
//	}
//	waddr := wallet.XXNetworkAddressFromMnemonic(s.GetOutputMnemonic())
//
//	contractBytes := []byte("I solemnly swear that I am up to no good")
//	err = SignAndTransmit(rsa.CreatePrivateKeyPem(pk), idfBytes, contractBytes, waddr, nil, &MockSender{t, nid.Bytes(), rsa.CreatePublicKeyPem(pk.GetPublic())})
//	if err != nil {
//		t.Errorf("Failed to sign & transmit: %+v", err)
//	}
//}

func TestSignAndTransmit_big(t *testing.T) {
	keyPath := "/Users/jonahhusson/gitlab.com/elixxir/wrapper/creds/1/cmix-key.key"
	idfPath := "/Users/jonahhusson/gitlab.com/elixxir/wrapper/creds/1/cmix-IDF.json"
	contractPath := "/tmp/contract.txt"
	wallet := "wallethere"
	addr := "35.81.172.71:11420"
	certPath := "/Users/jonahhusson/gitlab.com/elixxir/deployment/scripts/commitments/commitments.release.elixxir.io.crt"
	err := SignAndTransmit(keyPath, idfPath, contractPath, wallet, addr, certPath)
	if err != nil {
		t.Errorf("Failed to register commitment: %+v", err)
	}
}
