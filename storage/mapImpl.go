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
	for _, m := range members {
		db.members[base64.StdEncoding.EncodeToString(m.Id)] = m
	}
	return nil
}

func (db *MapImpl) InsertCommitment(commitment Commitment) error {
	db.commitments[base64.StdEncoding.EncodeToString(commitment.Id)] = commitment
	return nil
}

func (db *MapImpl) GetMember(id []byte) (*Member, error) {
	encoded := base64.StdEncoding.EncodeToString(id)
	m, ok := db.members[encoded]
	if !ok {
		return nil, errors.Errorf("No member in MapImpl with id %+v", id)
	}
	return &m, nil
}
