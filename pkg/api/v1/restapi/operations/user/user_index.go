// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UserIndexHandlerFunc turns a function with the right signature into a user index handler
type UserIndexHandlerFunc func(UserIndexParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserIndexHandlerFunc) Handle(params UserIndexParams) middleware.Responder {
	return fn(params)
}

// UserIndexHandler interface for that can handle valid user index params
type UserIndexHandler interface {
	Handle(UserIndexParams) middleware.Responder
}

// NewUserIndex creates a new http.Handler for the user index operation
func NewUserIndex(ctx *middleware.Context, handler UserIndexHandler) *UserIndex {
	return &UserIndex{Context: ctx, Handler: handler}
}

/*UserIndex swagger:route GET /users user userIndex

Fetch all available users

*/
type UserIndex struct {
	Context *middleware.Context
	Handler UserIndexHandler
}

func (o *UserIndex) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserIndexParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
