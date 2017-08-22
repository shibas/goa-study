package removetap

import (
	"context"
	"errors"
	"goa-study/env"
	"goa-study/models/queries/taps"
)

type Server struct {
	TapID int
}

func (s *Server) Serve() (err error) {
	if s.TapID == 0 {
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

	err = taps.RemoveTapByID(tx, s.TapID)
	if err != nil {
		return
	}
	tx.Commit()
	return
}
