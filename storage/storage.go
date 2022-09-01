////////////////////////////////////////////////////////////////////////////////
// Copyright © 2022 xx foundation                                             //
//                                                                            //
// Use of this source code is governed by a license that can be found in the  //
// LICENSE file.                                                              //
////////////////////////////////////////////////////////////////////////////////

// Handles the high level storage API.

package storage

import "testing"

type Params struct {
	Username string
	Password string
	DBName   string
	Address  string
	Port     string
}

// Storage struct is the API for the storage package
type Storage struct {
	// Stored Database interface
	database
}

// NewStorage creates a new Storage object wrapping a database interface
// Returns a Storage object, and error
func NewStorage(params Params) (*Storage, error) {
	db, err := newDatabase(params.Username, params.Password, params.DBName, params.Address, params.Port)
	storage := &Storage{db}
	return storage, err
}

func (s *Storage) GetMapImpl(t *testing.T) *MapImpl {
	if t == nil {
		return nil
	}
	return s.database.(*MapImpl)
}
