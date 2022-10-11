// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteModHandlerFunc turns a function with the right signature into a delete mod handler
type DeleteModHandlerFunc func(DeleteModParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteModHandlerFunc) Handle(params DeleteModParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteModHandler interface for that can handle valid delete mod params
type DeleteModHandler interface {
	Handle(DeleteModParams, *models.User) middleware.Responder
}

// NewDeleteMod creates a new http.Handler for the delete mod operation
func NewDeleteMod(ctx *middleware.Context, handler DeleteModHandler) *DeleteMod {
	return &DeleteMod{Context: ctx, Handler: handler}
}

/* DeleteMod swagger:route DELETE /mods/{mod_id} mod deleteMod

Delete a specific mod

*/
type DeleteMod struct {
	Context *middleware.Context
	Handler DeleteModHandler
}

func (o *DeleteMod) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteModParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
