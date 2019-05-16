// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ModShowHandlerFunc turns a function with the right signature into a mod show handler
type ModShowHandlerFunc func(ModShowParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ModShowHandlerFunc) Handle(params ModShowParams) middleware.Responder {
	return fn(params)
}

// ModShowHandler interface for that can handle valid mod show params
type ModShowHandler interface {
	Handle(ModShowParams) middleware.Responder
}

// NewModShow creates a new http.Handler for the mod show operation
func NewModShow(ctx *middleware.Context, handler ModShowHandler) *ModShow {
	return &ModShow{Context: ctx, Handler: handler}
}

/*ModShow swagger:route GET /mods/{modID} mod modShow

Fetch a specific mod

*/
type ModShow struct {
	Context *middleware.Context
	Handler ModShowHandler
}

func (o *ModShow) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewModShowParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
