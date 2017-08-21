package showbeerlist

import (
	"errors"
	"goa-study/app"
	"goa-study/env"
	"goa-study/models/queries/beers"
	rb "goa-study/models/receivers/beers"
)

type Result struct {
	Envelope []*app.GoaStudyBeer
}

func Serve(i interface{}) (err error) {
	res, ok := i.(*Result)
	if !ok {
		return errors.New("invalid interface")
	}

	beers, err := beers.FindBeerInfoAll(env.ReadOnlyConnection)
	if err != nil {
		return
	}

	bpack := rb.FormatBeerPackAll(beers)
	res.Envelope = bpack
	return
}
