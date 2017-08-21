package controllers

import (
	"goa-study/app"
	"goa-study/models/servers/changebeer"
	"goa-study/models/servers/createbeer"
	"goa-study/models/servers/removebeer"
	"goa-study/models/servers/showbeer"
	"goa-study/models/servers/showbeerlist"

	db "upper.io/db.v3"

	"github.com/goadesign/goa"
)

// BeerController implements the beer resource.
type BeerController struct {
	*goa.Controller
}

// NewBeerController creates a beer controller.
func NewBeerController(service *goa.Service) *BeerController {
	return &BeerController{Controller: service.NewController("BeerController")}
}

// Change runs the change action.
func (c *BeerController) Change(ctx *app.ChangeBeerContext) error {
	// BeerController_Change: start_implement

	// Put your logic here
	res := &changebeer.Result{}
	cbs := changebeer.Server{
		BeerID: ctx.BeerID,
		Name:   ctx.Payload.Title,
		Type:   ctx.Payload.Type,
	}
	err := cbs.Serve(res)
	if err != nil && err == db.ErrNoMoreRows {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "change beer api error", err)
		return ctx.InternalServerError()
	}
	// BeerController_Change: end_implement
	return ctx.OK(res.Envelope)
}

// Create runs the create action.
func (c *BeerController) Create(ctx *app.CreateBeerContext) error {
	// BeerController_Create: start_implement

	// Put your logic here
	res := &createbeer.Result{}
	cbs := createbeer.Server{
		Name: ctx.Payload.Title,
		Type: ctx.Payload.Type,
	}
	err := cbs.Serve(res)
	if err != nil {
		goa.LogError(ctx, "create beer api error", err)
		return ctx.InternalServerError()
	}
	// BeerController_Create: end_implement
	return ctx.Created(res.Envelope)
}

// List runs the list action.
func (c *BeerController) List(ctx *app.ListBeerContext) error {
	// BeerController_List: start_implement

	// Put your logic here
	res := &showbeerlist.Result{}
	err := showbeerlist.Serve(res)
	if err != nil && err == db.ErrNoMoreRows {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "get beer list api error", err)
		return ctx.InternalServerError()
	}

	// BeerController_List: end_implement
	return ctx.OK(res.Envelope)
}

// Remove runs the remove action.
func (c *BeerController) Remove(ctx *app.RemoveBeerContext) error {
	// BeerController_Remove: start_implement

	// Put your logic here
	sbs := removebeer.Server{
		BeerID: ctx.BeerID,
	}
	err := sbs.Serve()
	if err != nil && err == db.ErrNoMoreRows {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "remove beer api error", err)
		return ctx.InternalServerError()
	}

	// BeerController_Remove: end_implement
	return ctx.NoContent()
}

// Show runs the show action.
func (c *BeerController) Show(ctx *app.ShowBeerContext) error {
	// BeerController_Show: start_implement

	// Put your logic here
	res := &showbeer.Result{}
	sbs := showbeer.Server{
		BeerID: ctx.BeerID,
	}
	err := sbs.Serve(res)
	if err != nil && err == db.ErrNoMoreRows {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "get beer api error", err)
		return ctx.InternalServerError()
	}

	// BeerController_Show: end_implement
	return ctx.OK(res.Envelope)
}
