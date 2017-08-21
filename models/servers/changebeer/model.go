package changebeer

import (
	"context"
	"errors"
	"goa-study/app"
	"goa-study/env"
	"goa-study/models/queries/beers"
	"goa-study/models/queries/types"
	rb "goa-study/models/receivers/beers"
)

type Server struct {
	BeerID int
	Name   string
	Type   string
}

type Result struct {
	Envelope *app.GoaStudyBeer
}

func (s *Server) Serve(i interface{}) (err error) {
	res, ok := i.(*Result)
	if !ok {
		return errors.New("invalid interface")
	}
	if s.BeerID == 0 || s.Name == "" || s.Type == "" {
		return errors.New("invalid parameters")
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

	// IDがなければエラー
	_, err = beers.FindBeerInfoByID(env.ReadOnlyConnection, s.BeerID)
	if err != nil {
		return
	}

	// type id を取得なければ作成
	tp, err := types.FindOrCreateTypeIDByTitle(tx, s.Type)
	if err != nil {
		return
	}

	cbr, err := beers.ChangeBeerInfo(tx, s.Name, *tp.ID, s.BeerID)
	if err != nil {
		return
	}
	tx.Commit()

	bpack := rb.CreateBeerPack(cbr, tp)
	res.Envelope = &bpack
	return
}
