package controllers

import (
	"goa-study/app"
	"goa-study/models/servers/changetap"
	"goa-study/models/servers/createtap"
	"goa-study/models/servers/removetap"
	"goa-study/models/servers/showtap"
	"goa-study/models/servers/showtaplist"

	db "upper.io/db.v3"

	"github.com/goadesign/goa"
)

// TapController implements the tap resource.
type TapController struct {
	*goa.Controller
}

// NewTapController creates a tap controller.
func NewTapController(service *goa.Service) *TapController {
	return &TapController{Controller: service.NewController("TapController")}
}

// Change runs the change action.
func (c *TapController) Change(ctx *app.ChangeTapContext) error {
	// TapController_Change: start_implement

	// Put your logic here
	res := &changetap.Result{}
	cts := changetap.Server{
		TapID:  ctx.TapID,
		BeerID: ctx.Payload.BeerID,
	}
	err := cts.Serve(res)
	if err != nil && err == db.ErrNoMoreRows {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "change tap api error", err)
		return ctx.InternalServerError()
	}

	// TapController_Change: end_implement
	return ctx.OK(res.Envelope)
}

// Create runs the create action.
func (c *TapController) Create(ctx *app.CreateTapContext) error {
	// TapController_Create: start_implement

	// Put your logic here
	res := &createtap.Result{}
	cts := createtap.Server{
		BeerID: ctx.Payload.BeerID,
	}
	err := cts.Serve(res)
	if err != nil && err == db.ErrNoMoreRows {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "create tap api error", err)
		return ctx.InternalServerError()
	}

	// TapController_Create: end_implement
	return ctx.OK(res.Envelope)
}

// List runs the list action.
func (c *TapController) List(ctx *app.ListTapContext) error {
	// TapController_List: start_implement

	// Put your logic here
	res := &showtaplist.Result{}
	err := showtaplist.Serve(res)
	if err != nil && err == db.ErrNoMoreRows {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "get beer list api error", err)
		return ctx.InternalServerError()
	}

	// TapController_List: end_implement
	return ctx.OK(res.Envelope)
}

// Remove runs the remove action.
func (c *TapController) Remove(ctx *app.RemoveTapContext) error {
	// TapController_Remove: start_implement

	// Put your logic here
	rts := removetap.Server{
		TapID: ctx.TapID,
	}
	err := rts.Serve()
	if err != nil && err == db.ErrNoMoreRows {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "remove beer api error", err)
		return ctx.InternalServerError()
	}

	// TapController_Remove: end_implement
	return ctx.NoContent()
}

// Show runs the show action.
func (c *TapController) Show(ctx *app.ShowTapContext) error {
	// TapController_Show: start_implement

	// Put your logic here
	res := &showtap.Result{}
	sts := showtap.Server{
		TapID: ctx.TapID,
	}
	err := sts.Serve(res)
	if err != nil && err == db.ErrNoMoreRows {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "get beer api error", err)
		return ctx.InternalServerError()
	}

	// TapController_Show: end_implement
	return ctx.OK(res.Envelope)
}
