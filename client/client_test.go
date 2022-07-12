///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package client

import (
	"crypto/rand"
	gorsa "crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"git.xx.network/elixxir/mainnet-commitments/server"
	"git.xx.network/elixxir/mainnet-commitments/storage"
	"github.com/xx-labs/sleeve/wallet"
	"gitlab.com/xx_network/crypto/csprng"
	"gitlab.com/xx_network/crypto/signature/rsa"
	"gitlab.com/xx_network/primitives/id"
	"gitlab.com/xx_network/primitives/id/idf"
	"math/big"
	"os"
	"testing"
	"time"
)

func TestSignAndTransmit(t *testing.T) {
	rng := csprng.NewSystemRNG()
	rng.SetSeed([]byte("start"))
	pk, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		t.Errorf("Failed to gen key: %+v", err)
	}
	nid := id.NewIdFromString("jonah", id.Node, t)
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

	s, err := wallet.NewSleeve(rand.Reader, "password")
	if err != nil {
		t.Errorf("Failed to create sleeve: %+v", err)
	}
	waddr := wallet.XXNetworkAddressFromMnemonic(s.GetOutputMnemonic())
	waddr2 := wallet.XXNetworkAddressFromMnemonic(s.GetMnemonic())

	testKeyPath := "/tmp/commitmenttestkey.key"
	testIDFPath := "/tmp/testidf.json"
	err = os.WriteFile(testKeyPath, rsa.CreatePrivateKeyPem(pk), os.ModePerm)
	if err != nil {
		t.Errorf("Failed to write test key: %+v", err)
	}
	err = os.WriteFile(testIDFPath, idfBytes, os.ModePerm)
	if err != nil {
		t.Errorf("Failed to write test idf: %+v", err)
	}

	certBytes, err := makeCert(&pk.PrivateKey)
	if err != nil {
		t.Errorf("Failed to create test cert: %+v", err)
	}

	mapImpl, err := storage.NewStorage(storage.Params{})
	if err != nil {
		t.Error("Failed to init storage for mock server")
	}
	err = mapImpl.InsertMembers([]storage.Member{{
		Id:   nid.Bytes(),
		Cert: certBytes,
	},
	})
	if err != nil {
		t.Errorf("Failed to insert members: %+v", err)
	}

	var errChan = make(chan error)
	var doneChan = make(chan bool)

	go func() {
		err := server.StartServer(server.Params{
			KeyPath:      "",
			CertPath:     "",
			ContractHash: "eGoC90IBWQPGxv2FJVLScpEvR0DhWEdhiobiF_cfVBnSXhAxr-5YUxOJZESTTrBLkDpoWxRIt1XVb3Aa_pvizg==",
			Port:         "11420",
		}, mapImpl)
		if err != nil {
			t.Errorf("Failed to start dummy server")
			errChan <- err
		} else {
			doneChan <- true
		}
	}()
	time.Sleep(time.Millisecond * 100)

	err = SignAndTransmit(testKeyPath, testIDFPath, waddr, waddr2, "http://localhost:11420", "", "", "", 0.0)
	if err != nil {
		t.Errorf("Failed to sign & transmit: %+v", err)
	}
}

func makeCert(pk *gorsa.PrivateKey) ([]byte, error) {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{"Company, INC."},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{"Golden Gate Bridge"},
			PostalCode:    []string{"94016"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &pk.PublicKey, pk)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:    "",
		Headers: nil,
		Bytes:   caBytes,
	}
	return pem.EncodeToMemory(block), nil
}

func TestGetInfo(t *testing.T) {
	ret, err := GetInfo("\\x616263313233", "", "http://0.0.0.0:11420")
	if err != nil {
		t.Errorf("Failed to get info: %+v", err)
	}
	t.Log(string(ret))
}
