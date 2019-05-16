// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ModUpdateHandlerFunc turns a function with the right signature into a mod update handler
type ModUpdateHandlerFunc func(ModUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ModUpdateHandlerFunc) Handle(params ModUpdateParams) middleware.Responder {
	return fn(params)
}

// ModUpdateHandler interface for that can handle valid mod update params
type ModUpdateHandler interface {
	Handle(ModUpdateParams) middleware.Responder
}

// NewModUpdate creates a new http.Handler for the mod update operation
func NewModUpdate(ctx *middleware.Context, handler ModUpdateHandler) *ModUpdate {
	return &ModUpdate{Context: ctx, Handler: handler}
}

/*ModUpdate swagger:route PUT /mods/{modID} mod modUpdate

Update a specific mod

*/
type ModUpdate struct {
	Context *middleware.Context
	Handler ModUpdateHandler
}

func (o *ModUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewModUpdateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
