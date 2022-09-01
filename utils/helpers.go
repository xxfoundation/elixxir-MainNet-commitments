////////////////////////////////////////////////////////////////////////////////
// Copyright © 2022 xx foundation                                             //
//                                                                            //
// Use of this source code is governed by a license that can be found in the  //
// LICENSE file.                                                              //
////////////////////////////////////////////////////////////////////////////////

package utils

import (
	"crypto"
	"github.com/pkg/errors"
)

func HashNodeInfo(nominatorWallet, validatorWallet string, idfBytes, contractBytes []byte) ([]byte, crypto.Hash, error) {
	h := crypto.BLAKE2b_512 // Define & return this here so we aren't defining hash type in 3 places for sign/verify calls
	hasher := h.New()
	_, err := hasher.Write(idfBytes)
	if err != nil {
		return nil, h, errors.WithMessage(err, "Failed to write IDF to hash")
	}
	_, err = hasher.Write(contractBytes)
	if err != nil {
		return nil, h, errors.WithMessage(err, "Failed to write contract to hash")
	}
	_, err = hasher.Write([]byte(nominatorWallet))
	if err != nil {
		return nil, h, errors.WithMessage(err, "Failed to write nominator wallet to hash")
	}
	_, err = hasher.Write([]byte(validatorWallet))
	if err != nil {
		return nil, h, errors.WithMessage(err, "Failed to write validator wallet to hash")
	}
	return hasher.Sum(nil), h, nil
}
