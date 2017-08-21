package removebeer

import (
	"context"
	"errors"
	"goa-study/env"
	"goa-study/models/queries/beers"
)

type Server struct {
	BeerID int
}

func (s *Server) Serve() (err error) {
	if s.BeerID == 0 {
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

	err = beers.RemoveBeerByID(tx, s.BeerID)
	if err != nil {
		return
	}
	tx.Commit()
	return
}
