package taps

import (
	"time"

	"upper.io/db.v3/lib/sqlbuilder"
)

// Taps DB objects
type Taps struct {
	ID        *int      `db:"id,omitempty"`
	BeerID    int       `db:"beer_id"`
	CreatedAt time.Time `db:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}

type TapFull struct {
	ID        int       `db:"tap_id"`
	BeerID    int       `db:"beer_id"`
	BeerTitle string    `db:"beer_title"`
	CreatedAt time.Time `db:"tap_created_at,omitempty"`
	UpdatedAt time.Time `db:"tap_updated_at,omitempty"`
}

const tableName string = "taps"

func CreateNewTap(tx sqlbuilder.Tx, beerID int) (Taps, error) {
	var t Taps
	coll := tx.Collection(tableName)
	data := Taps{
		BeerID: beerID,
	}
	lastID, err := coll.Insert(&data)
	err = coll.Find(lastID).One(&t)
	return t, err
}

func ChangeTapInfo(tx sqlbuilder.Tx, tapID int, beerID int) (Taps, error) {
	var t Taps
	coll := tx.Collection(tableName)
	data := Taps{
		BeerID: beerID,
	}
	err := coll.Find(tapID).Update(data)
	if err != nil {
		return t, err
	}
	err = coll.Find(tapID).One(&t)
	return t, err
}

func RemoveTapByID(tx sqlbuilder.Tx, tapID int) error {
	coll := tx.Collection(tableName)
	err := coll.Find(tapID).Delete()
	return err
}

func FindTapInfoByID(s sqlbuilder.Database, id int) (TapFull, error) {
	var tf TapFull
	q := s.Select(selectQueriesTapPack...).
		From("taps as t").
		Join("beers as b").On("b.id = t.beer_id").
		Where("t.id", id)
	err := q.One(&tf)
	return tf, err
}

func FindTapInfoAll(s sqlbuilder.Database) ([]*TapFull, error) {
	var tfl []*TapFull
	q := s.Select(selectQueriesTapPack...).
		From("taps as t").
		Join("beers as b").On("b.id = t.beer_id")
	err := q.All(&tfl)
	return tfl, err
}

var (
	selectQueriesTapPack = []interface{}{
		"t.id as tap_id",
		"b.id as beer_id",
		"b.title as beer_title",
		"t.created_at as tap_created_at",
		"t.updated_at as tap_updated_at",
	}
)
