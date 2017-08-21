package showbeer

import (
	"errors"
	"goa-study/app"
	"goa-study/env"
	"goa-study/models/queries/beers"
	rb "goa-study/models/receivers/beers"
)

type Server struct {
	BeerID int
}

type Result struct {
	Envelope *app.GoaStudyBeer
}

func (s *Server) Serve(i interface{}) (err error) {
	res, ok := i.(*Result)
	if !ok {
		return errors.New("invalid interface")
	}
	if s.BeerID == 0 {
		return errors.New("invalid parameters")
	}

	beer, err := beers.FindBeerInfoByID(env.ReadOnlyConnection, s.BeerID)
	if err != nil {
		return
	}

	bpack := rb.FormatBeerPack(beer)
	res.Envelope = &bpack
	return
}
