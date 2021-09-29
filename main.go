// +build js
// +build wasm

package main

import (
	"encoding/json"
	"git.xx.network/elixxir/mainnet-commitments/client"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/primitives/id"
	"gitlab.com/xx_network/primitives/id/idf"
	"gitlab.com/xx_network/primitives/utils"
	"syscall/js"
)

var done chan bool

// wasm compilations assigns SignAndTransmit func to signAndTransmit js global
func main() {
	f := js.FuncOf(SignAndTransmit)
	js.Global().Set("signAndTransmit", f)
	done = make(chan bool)
	<-done
}

// SignAndTransmit signs & transmits info to the commitments server
// Accepts args nodeCertPath, nodeKeyPath, idfPath, wallet, commitmentsAddress, commitmentsCertPath
func SignAndTransmit(this js.Value, inputs []js.Value) interface{} {
	certPath := inputs[0].String()
	keyPath := inputs[1].String()
	idfPath := inputs[2].String()
	wallet := inputs[3].String()
	address := inputs[4].String()
	commitmentsCertPath := inputs[5].String()

	var cert, key, idfBytes, commitmentCert []byte
	var err error
	var ep string
	// Read certificate file
	if ep, err = utils.ExpandPath(certPath); err == nil {
		cert, err = utils.ReadFile(ep)
		if err != nil {
			return map[string]interface{}{"Error": err.Error()}
		}
	} else {
		return map[string]interface{}{"Error": err.Error()}
	}

	// Read key file
	if ep, err = utils.ExpandPath(keyPath); err == nil {
		key, err = utils.ReadFile(ep)
		if err != nil {
			return map[string]interface{}{"Error": err.Error()}
		}
	} else {
		return map[string]interface{}{"Error": err.Error()}
	}

	// Read id file
	if ep, err = utils.ExpandPath(idfPath); err == nil {
		idfBytes, err = utils.ReadFile(ep)
		if err != nil {
			return map[string]interface{}{"Error": err.Error()}
		}
	} else {
		return map[string]interface{}{"Error": err.Error()}
	}

	if ep, err = utils.ExpandPath(commitmentsCertPath); err == nil {
		commitmentCert, err = utils.ReadFile(ep)
		if err != nil {
			return map[string]interface{}{"Error": err.Error()}
		}
	} else {
		return map[string]interface{}{"Error": err.Error()}
	}

	idfStruct := &idf.IdFile{}
	err = json.Unmarshal(idfBytes, idfStruct)
	if err != nil {
		return map[string]interface{}{"Error": err.Error()}
	}

	nodeID, err := id.Unmarshal(idfStruct.IdBytes[:])
	if err != nil {
		return map[string]interface{}{"Error": err.Error()}
	}

	cl, err := client.StartClient(key, cert, idfStruct.Salt[:], nodeID)
	if err != nil {
		return map[string]interface{}{"Error": err.Error()}
	}

	h, err := connect.NewHost(&id.Permissioning, address, commitmentCert, connect.GetDefaultHostParams())
	if err != nil {
		return map[string]interface{}{"Error": err.Error()}
	}

	// Sign & transmit information
	err = client.SignAndTransmit(key, idfBytes, wallet, h, cl)
	if err != nil {
		return map[string]interface{}{"Error": err.Error()}
	}
	return map[string]interface{}{}
}
