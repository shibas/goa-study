package showtap

import (
	"errors"
	"goa-study/app"
	"goa-study/env"
	"goa-study/models/queries/taps"

	rt "goa-study/models/receivers/taps"
)

type Server struct {
	TapID int
}

type Result struct {
	Envelope *app.GoaStudyTap
}

func (s *Server) Serve(i interface{}) (err error) {
	res, ok := i.(*Result)
	if !ok {
		return errors.New("invalid interface")
	}
	if s.TapID == 0 {
		return errors.New("invalid parameters")
	}

	tap, err := taps.FindTapInfoByID(env.ReadOnlyConnection, s.TapID)
	if err != nil {
		return
	}

	tpack := rt.FormatTapPack(tap)
	res.Envelope = &tpack
	return
}
