package taps

import (
	"goa-study/app"
	"goa-study/models/queries/beers"
	"goa-study/models/queries/taps"
)

type TapPack struct {
	ID        int
	BeerID    int
	BeerTitle string
	CreateAt  string
	UpdateAt  string
}

func CreateTapPack(b beers.BeerFull, t taps.Taps) app.GoaStudyTap {
	cr := t.CreatedAt.String()
	up := t.UpdatedAt.String()
	return app.GoaStudyTap{
		ID:        *t.ID,
		BeerID:    &b.ID,
		BeerTitle: b.Title,
		CreatedAt: &cr,
		UpdatedAt: &up,
	}
}

func FormatTapPack(t taps.TapFull) app.GoaStudyTap {
	cr := t.CreatedAt.String()
	up := t.UpdatedAt.String()
	return app.GoaStudyTap{
		ID:        t.ID,
		BeerID:    &t.BeerID,
		BeerTitle: t.BeerTitle,
		CreatedAt: &cr,
		UpdatedAt: &up,
	}
}

func FormatTapPackAll(t []*taps.TapFull) []*app.GoaStudyTap {
	var res []*app.GoaStudyTap
	for _, tap := range t {
		cr := tap.CreatedAt.String()
		up := tap.UpdatedAt.String()
		res = append(res, &app.GoaStudyTap{
			ID:        tap.ID,
			BeerID:    &tap.BeerID,
			BeerTitle: tap.BeerTitle,
			CreatedAt: &cr,
			UpdatedAt: &up,
		})
	}
	return res
}
