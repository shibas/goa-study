package showtaplist

import (
	"errors"
	"goa-study/app"
	"goa-study/env"
	"goa-study/models/queries/taps"
	rt "goa-study/models/receivers/taps"
)

type Result struct {
	Envelope []*app.GoaStudyTap
}

func Serve(i interface{}) (err error) {
	res, ok := i.(*Result)
	if !ok {
		return errors.New("invalid interface")
	}

	taps, err := taps.FindTapInfoAll(env.ReadOnlyConnection)
	if err != nil {
		return
	}

	tpack := rt.FormatTapPackAll(taps)
	res.Envelope = tpack
	return
}
