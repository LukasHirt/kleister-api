// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ShowProfileHandlerFunc turns a function with the right signature into a show profile handler
type ShowProfileHandlerFunc func(ShowProfileParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn ShowProfileHandlerFunc) Handle(params ShowProfileParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// ShowProfileHandler interface for that can handle valid show profile params
type ShowProfileHandler interface {
	Handle(ShowProfileParams, *models.User) middleware.Responder
}

// NewShowProfile creates a new http.Handler for the show profile operation
func NewShowProfile(ctx *middleware.Context, handler ShowProfileHandler) *ShowProfile {
	return &ShowProfile{Context: ctx, Handler: handler}
}

/* ShowProfile swagger:route GET /profile/self profile showProfile

Fetch profile details of the personal account

*/
type ShowProfile struct {
	Context *middleware.Context
	Handler ShowProfileHandler
}

func (o *ShowProfile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewShowProfileParams()
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
