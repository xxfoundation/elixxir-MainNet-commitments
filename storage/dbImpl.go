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
	return db.db.Create(&members).Error
}

func (db *DatabaseImpl) InsertCommitment(commitment Commitment) error {
	return db.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, DoUpdates: clause.AssignmentColumns([]string{"wallet", "nominator_wallet", "email", "selected_multiplier"})}).Create(&commitment).Error
}

func (db *DatabaseImpl) GetMember(id string) (*Member, error) {
	jww.INFO.Printf("Getting member with id %+v", id)
	m := Member{}
	return &m, db.db.First(&m, "id = ?", id).Error
}

func (db *DatabaseImpl) GetCommitment(id string) (*Commitment, error) {
	jww.INFO.Printf("Getting member with id %+v", id)
	c := Commitment{}
	return &c, db.db.First(&c, "id = ?", id).Error
}
