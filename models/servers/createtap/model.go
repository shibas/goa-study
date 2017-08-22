package createtap

import (
	"context"
	"errors"
	"goa-study/app"
	"goa-study/env"
	"goa-study/models/queries/beers"
	"goa-study/models/queries/taps"
	rt "goa-study/models/receivers/taps"
)

type Server struct {
	BeerID int
}

type Result struct {
	Envelope *app.GoaStudyTap
}

func (s *Server) Serve(i interface{}) (err error) {
	res, ok := i.(*Result)
	if !ok {
		return errors.New("invalid interface")
	}
	if s.BeerID == 0 {
		return errors.New("invalid parameters")
	}
	// beerIDがなければエラー
	br, err := beers.FindBeerInfoByID(env.ReadOnlyConnection, s.BeerID)
	if err != nil {
		return
	}

	tx, err := env.Connection.NewTx(context.Background())
	tx.SetLogging(env.OnDevelopment)
	if err != nil {
		return
	}

	defer func() {
		tx.Rollback()
		tx.Close()
	}()

	tap, err := taps.CreateNewTap(tx, s.BeerID)
	if err != nil {
		return
	}
	tx.Commit()

	tpack := rt.CreateTapPack(br, tap)
	res.Envelope = &tpack
	return
}
