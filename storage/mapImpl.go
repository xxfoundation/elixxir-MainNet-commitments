///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package storage

import (
	"encoding/base64"
	"github.com/pkg/errors"
	"sync"
)

// MapImpl struct implements the database interface with an underlying Map
type MapImpl struct {
	members     map[string]Member
	commitments map[string]Commitment
	sync.RWMutex
}

func (db *MapImpl) InsertMembers(members []Member) error {
	db.Lock()
	defer db.Unlock()
	for _, m := range members {
		db.members[base64.StdEncoding.EncodeToString(m.Id)] = m
	}
	return nil
}

func (db *MapImpl) InsertCommitment(commitment Commitment) error {
	db.Lock()
	defer db.Unlock()
	db.commitments[string(commitment.Id)] = commitment
	return nil
}

func (db *MapImpl) GetMember(id string) (*Member, error) {
	db.RLock()
	defer db.RUnlock()
	m, ok := db.members[id]
	if !ok {
		return nil, errors.Errorf("No member in MapImpl with id %+v", id)
	}
	return &m, nil
}
