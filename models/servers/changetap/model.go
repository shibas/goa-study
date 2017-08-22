package changetap

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
	TapID  int
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
	if s.TapID == 0 || s.BeerID == 0 {
		return errors.New("invalid parameters")
	}
	rodb := env.ReadOnlyConnection
	tx, err := env.Connection.NewTx(context.Background())
	tx.SetLogging(env.OnDevelopment)
	if err != nil {
		return
	}

	defer func() {
		tx.Rollback()
		tx.Close()
	}()
	// tapIDがなければエラー
	_, err = taps.FindTapInfoByID(rodb, s.TapID)
	if err != nil {
		return
	}

	// beerIDがなければエラー
	br, err := beers.FindBeerInfoByID(rodb, s.BeerID)
	if err != nil {
		return
	}

	ctr, err := taps.ChangeTapInfo(tx, s.TapID, s.BeerID)
	if err != nil {
		return
	}
	tx.Commit()

	tpack := rt.CreateTapPack(br, ctr)
	res.Envelope = &tpack
	return
}
