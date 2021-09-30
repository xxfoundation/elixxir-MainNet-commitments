package utils

import (
	"crypto"
	"github.com/pkg/errors"
)

func HashNodeInfo(wallet string, idfBytes []byte) ([]byte, crypto.Hash, error) {
	h := crypto.BLAKE2b_512 // Define & return this here so we aren't defining hash type in 3 places for sign/verify calls
	hasher := h.New()
	_, err := hasher.Write(idfBytes)
	if err != nil {
		return nil, h, errors.WithMessage(err, "Failed to write IDF to hash")
	}
	_, err = hasher.Write([]byte(wallet))
	if err != nil {
		return nil, h, errors.WithMessage(err, "Failed to write wallet to hash")
	}
	return hasher.Sum(nil), h, nil
}
