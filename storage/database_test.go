package storage

import (
	"bytes"
	"gitlab.com/xx_network/primitives/id"
	"testing"
)

func setup(t *testing.T) (*Storage, error) {
	if t == nil {
		panic("Cannot run setup outside of testing")
	}
	p := Params{
		Username: "jonahhusson",
		Password: "",
		DBName:   "commitments",
		Address:  "0.0.0.0",
		Port:     "5432",
	}

	return NewStorage(p)
}

func TestDatabase(t *testing.T) {
	s, err := setup(t)
	if err != nil {
		t.Fatalf("Failed to setup storage: %+v", err)
	}
	id1 := id.NewIdFromString("zezimaone", id.Node, t)
	m1 := Member{
		Id:   id1.Bytes(),
		Cert: []byte("cert1"),
	}
	m2 := Member{
		Id:   id.NewIdFromString("zezimatwo", id.Node, t).Bytes(),
		Cert: []byte("cert2"),
	}
	err = s.InsertMembers([]Member{m1, m2})
	if err != nil {
		t.Errorf("Failed to insert members: %+v", err)
	}
	rm, err := s.GetMember(id1.Bytes())
	if err != nil {
		t.Errorf("Failed to get member: %+v", err)
	}
	if bytes.Compare(rm.Cert, []byte("cert1")) != 0 {
		t.Errorf("Members didn't match")
	}
	err = s.InsertCommitment(Commitment{
		Id:        id1.Bytes(),
		Wallet:    "wallet1",
		Signature: []byte("sig1"),
	})
	if err != nil {
		t.Errorf("Failed to insert commitment for member 1: %+v", err)
	}
}
