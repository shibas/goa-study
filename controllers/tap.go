package controllers

import (
	"goa-study/app"

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

	// TapController_Change: end_implement
	res := &app.GoaStudyTap{}
	return ctx.OK(res)
}

// Create runs the create action.
func (c *TapController) Create(ctx *app.CreateTapContext) error {
	// TapController_Create: start_implement

	// Put your logic here

	// TapController_Create: end_implement
	res := &app.GoaStudyTap{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *TapController) List(ctx *app.ListTapContext) error {
	// TapController_List: start_implement

	// Put your logic here

	// TapController_List: end_implement
	res := app.GoaStudyTapCollection{}
	return ctx.OK(res)
}

// Remove runs the remove action.
func (c *TapController) Remove(ctx *app.RemoveTapContext) error {
	// TapController_Remove: start_implement

	// Put your logic here

	// TapController_Remove: end_implement
	return nil
}

// Show runs the show action.
func (c *TapController) Show(ctx *app.ShowTapContext) error {
	// TapController_Show: start_implement

	// Put your logic here

	// TapController_Show: end_implement
	res := &app.GoaStudyTap{}
	return ctx.OK(res)
}
