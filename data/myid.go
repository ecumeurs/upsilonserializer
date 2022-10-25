package data

import (
	"github.com/google/uuid"
)

//go:generate msgp

type MyID string

// NewId creates a new MyID
func NewId() MyID {
	return FromUUID(uuid.New())
}

// FromUUID creates a MyID from a UUID
func FromUUID(id uuid.UUID) MyID {
	return MyID(id.String())
}

// ToUUID converts a MyID to a UUID
func (m MyID) ToUUID() uuid.UUID {
	return uuid.MustParse(string(m))
}
