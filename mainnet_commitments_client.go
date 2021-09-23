// +build js
// +build wasm

package mainnet_commitments

import (
	"git.xx.network/elixxir/mainnet-commitments/client"
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/xx_network/primitives/utils"
	"syscall/js"
)

var done chan bool

func main() {
	f := js.FuncOf(SignAndTransmit)
	js.Global().Set("signAndTransmit", f)
	done = make(chan bool)
	<-done
}

// SignAndTransmit signs & transmits info to the commitments server
func SignAndTransmit(this js.Value, inputs []js.Value) interface{} {
	commitmentsCertPath := inputs[0].String()
	certPath := inputs[1].String()
	keyPath := inputs[2].String()
	idfPath := inputs[3].String()
	wallet := inputs[4].String()
	address := inputs[5].String()

	var cert, key, idf, commitmentCert []byte
	var err error
	var ep string
	// Read certificate file
	if ep, err = utils.ExpandPath(certPath); err == nil {
		cert, err = utils.ReadFile(ep)
		if err != nil {
			jww.FATAL.Panicf("Failed to read cert file at path %s: %+v", ep, err)
		}
	} else {
		jww.FATAL.Panicf("Failed to expand certificate path: %+v", err)
	}

	// Read key file
	if ep, err = utils.ExpandPath(keyPath); err == nil {
		key, err = utils.ReadFile(ep)
		if err != nil {
			jww.FATAL.Panicf("Failed to read key file at path %s: %+v", ep, err)
		}
	} else {
		jww.FATAL.Panicf("Failed to expand key path: %+v", err)
	}

	// Read id file
	if ep, err = utils.ExpandPath(idfPath); err == nil {
		idf, err = utils.ReadFile(ep)
		if err != nil {
			jww.FATAL.Panicf("Failed to read id file at path %s: %+v", ep, err)
		}
	} else {
		jww.FATAL.Panicf("Failed to expand id path: %+v", err)
	}

	if ep, err = utils.ExpandPath(commitmentsCertPath); err == nil {
		commitmentCert, err = utils.ReadFile(ep)
		if err != nil {
			jww.FATAL.Panicf("Failed to read commitments cert file at path %s: %+v", ep, err)
		}
	} else {
		jww.FATAL.Panicf("Failed to expand commitments certificate path: %+v", err)
	}

	// Sign & transmit information
	err = client.SignAndTransmit(key, cert, idf, commitmentCert, wallet, address)
	if err != nil {
		jww.FATAL.Panicf("Failed to sign & transmit node info: %+v", err)
	}
	return nil
}
