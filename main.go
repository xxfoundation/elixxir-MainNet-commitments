///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package main

import (
	"git.xx.network/elixxir/mainnet-commitments/client"
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
// Accepts args nodeKeyPath, idfPath, contractPath, wallet, commitmentsAddress, commitmentsCertPath
func SignAndTransmit(this js.Value, inputs []js.Value) interface{} {
	keyPath := inputs[0].String()
	idfPath := inputs[1].String()
	contractPath := inputs[2].String()
	wallet := inputs[3].String()
	commitmentServerAddress := inputs[4].String()
	commitmentsCertPath := inputs[5].String()

	err := client.SignAndTransmit(keyPath, idfPath, contractPath, wallet, commitmentServerAddress, commitmentsCertPath)
	if err != nil {
		return map[string]interface{}{"Error": err.Error()}
	}
	return map[string]interface{}{"OK": true}
}
