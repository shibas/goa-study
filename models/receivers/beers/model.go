package beers

import (
	"goa-study/app"
	"goa-study/models/queries/beers"
	"goa-study/models/queries/types"
)

type BeerPack struct {
	ID       int
	Title    string
	TypeID   int
	Type     string
	CreateAt string
	UpdateAt string
}

func CreateBeerPack(b beers.Beers, t types.Types) app.GoaStudyBeer {
	cr := b.CreatedAt.String()
	up := b.UpdatedAt.String()
	return app.GoaStudyBeer{
		ID:        *b.ID,
		Title:     b.Title,
		Type:      t.Title,
		CreatedAt: &cr,
		UpdatedAt: &up,
	}
}

func FormatBeerPack(b beers.BeerFull) app.GoaStudyBeer {
	cr := b.CreatedAt.String()
	up := b.UpdatedAt.String()
	return app.GoaStudyBeer{
		ID:        b.ID,
		Title:     b.Title,
		Type:      b.Type,
		CreatedAt: &cr,
		UpdatedAt: &up,
	}
}

func FormatBeerPackAll(b []*beers.BeerFull) []*app.GoaStudyBeer {
	var res []*app.GoaStudyBeer
	for _, beer := range b {
		cr := beer.CreatedAt.String()
		up := beer.UpdatedAt.String()
		res = append(res, &app.GoaStudyBeer{
			ID:        beer.ID,
			Title:     beer.Title,
			Type:      beer.Type,
			CreatedAt: &cr,
			UpdatedAt: &up,
		})
	}
	return res
}
