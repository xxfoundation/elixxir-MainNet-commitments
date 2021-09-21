package storage

func (db *DatabaseImpl) InsertMembers(members []Member) error {
	return db.db.Create(&members).Error
}

func (db *DatabaseImpl) InsertCommitment(commitment Commitment) error {
	return db.db.Create(&commitment).Error
}

func (db *DatabaseImpl) GetMember(id []byte) (*Member, error) {
	m := Member{}
	return &m, db.db.First(&m, "id = ?", id).Error
}
