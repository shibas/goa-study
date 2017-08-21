package beers

import (
	"time"

	"upper.io/db.v3/lib/sqlbuilder"
)

// Beers DB objects
type Beers struct {
	ID        *int      `db:"id,omitempty"`
	Title     string    `db:"title"`
	Type      int       `db:"type"`
	CreatedAt time.Time `db:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}

type BeerFull struct {
	ID        int       `db:"beer_id"`
	Title     string    `db:"beer_title"`
	TypeID    int       `db:"beer_type_id"`
	Type      string    `db:"beer_type"`
	CreatedAt time.Time `db:"beer_created_at,omitempty"`
	UpdatedAt time.Time `db:"beer_updated_at,omitempty"`
}

const tableName string = "beers"

func CreateNewBeer(tx sqlbuilder.Tx, title string, typeID int) (Beers, error) {
	var b Beers
	coll := tx.Collection(tableName)
	data := Beers{
		Title: title,
		Type:  typeID,
	}
	lastID, err := coll.Insert(&data)
	err = coll.Find(lastID).One(&b)
	return b, err
}

func ChangeBeerInfo(tx sqlbuilder.Tx, title string, typeID int, beerID int) (Beers, error) {
	var b Beers
	coll := tx.Collection(tableName)
	data := Beers{
		Title: title,
		Type:  typeID,
	}
	err := coll.Find(beerID).Update(data)
	if err != nil {
		return b, err
	}
	err = coll.Find(beerID).One(&b)
	return b, err
}

func RemoveBeerByID(tx sqlbuilder.Tx, beerID int) error {
	coll := tx.Collection(tableName)
	err := coll.Find(beerID).Delete()
	return err
}

func FindBeerInfoByID(s sqlbuilder.Database, id int) (BeerFull, error) {
	var bf BeerFull
	q := s.Select(selectQueriesBeerPack...).
		From("beers as b").
		Join("types as t").On("t.id = b.type").
		Where("b.id", id)
	err := q.One(&bf)
	return bf, err
}

func FindBeerInfoAll(s sqlbuilder.Database) ([]*BeerFull, error) {
	var bfl []*BeerFull
	q := s.Select(selectQueriesBeerPack...).
		From("beers as b").
		Join("types as t").On("t.id = b.type")
	err := q.All(&bfl)
	return bfl, err
}

var (
	selectQueriesBeerPack = []interface{}{
		"b.id as beer_id",
		"b.title as beer_title",
		"b.type as beer_type_id",
		"t.title as beer_type",
		"b.created_at as beer_created_at",
		"b.updated_at as beer_updated_at",
	}
)
