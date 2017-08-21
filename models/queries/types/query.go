package types

import (
	"time"

	db "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

// Types DB objects
type Types struct {
	ID        *int      `db:"id,omitempty"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}

const tableName string = "types"

func FindTypeIDByTitle(s sqlbuilder.Database, title string) (Types, error) {
	var t Types
	coll := s.Collection(tableName)
	err := coll.Find(db.Cond{"title": title}).One(&t)
	return t, err
}

func FindOrCreateTypeIDByTitle(tx sqlbuilder.Tx, title string) (Types, error) {
	var t Types
	coll := tx.Collection(tableName)
	err := coll.Find(db.Cond{"title": title}).One(&t)
	if err == nil && t.ID != nil {
		return t, err
	}
	if err != nil && err != db.ErrNoMoreRows {
		return t, err
	}
	data := Types{
		Title: title,
	}
	lastID, err := coll.Insert(&data)
	err = coll.Find(lastID).One(&t)
	return t, err
}
