// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeletePackHandlerFunc turns a function with the right signature into a delete pack handler
type DeletePackHandlerFunc func(DeletePackParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePackHandlerFunc) Handle(params DeletePackParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeletePackHandler interface for that can handle valid delete pack params
type DeletePackHandler interface {
	Handle(DeletePackParams, *models.User) middleware.Responder
}

// NewDeletePack creates a new http.Handler for the delete pack operation
func NewDeletePack(ctx *middleware.Context, handler DeletePackHandler) *DeletePack {
	return &DeletePack{Context: ctx, Handler: handler}
}

/* DeletePack swagger:route DELETE /packs/{pack_id} pack deletePack

Delete a specific pack

*/
type DeletePack struct {
	Context *middleware.Context
	Handler DeletePackHandler
}

func (o *DeletePack) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeletePackParams()
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
