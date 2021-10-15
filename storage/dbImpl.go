///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package storage

import (
	jww "github.com/spf13/jwalterweatherman"
	"gorm.io/gorm/clause"
)

func (db *DatabaseImpl) InsertMembers(members []Member) error {
	db.Lock()
	defer db.Unlock()
	return db.db.Create(&members).Error
}

func (db *DatabaseImpl) InsertCommitment(commitment Commitment) error {
	db.Lock()
	defer db.Unlock()
	return db.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&commitment).Error
}

func (db *DatabaseImpl) GetMember(id string) (*Member, error) {
	db.RLock()
	defer db.RUnlock()
	jww.INFO.Printf("Getting member with id %+v", id)
	m := Member{}
	return &m, db.db.First(&m, "id = ?", id).Error
}
