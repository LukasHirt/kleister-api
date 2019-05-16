// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ModIndexHandlerFunc turns a function with the right signature into a mod index handler
type ModIndexHandlerFunc func(ModIndexParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ModIndexHandlerFunc) Handle(params ModIndexParams) middleware.Responder {
	return fn(params)
}

// ModIndexHandler interface for that can handle valid mod index params
type ModIndexHandler interface {
	Handle(ModIndexParams) middleware.Responder
}

// NewModIndex creates a new http.Handler for the mod index operation
func NewModIndex(ctx *middleware.Context, handler ModIndexHandler) *ModIndex {
	return &ModIndex{Context: ctx, Handler: handler}
}

/*ModIndex swagger:route GET /mods mod modIndex

Fetch all available mods

*/
type ModIndex struct {
	Context *middleware.Context
	Handler ModIndexHandler
}

func (o *ModIndex) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewModIndexParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
