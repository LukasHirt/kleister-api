// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// AppendTeamToUserHandlerFunc turns a function with the right signature into a append team to user handler
type AppendTeamToUserHandlerFunc func(AppendTeamToUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn AppendTeamToUserHandlerFunc) Handle(params AppendTeamToUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// AppendTeamToUserHandler interface for that can handle valid append team to user params
type AppendTeamToUserHandler interface {
	Handle(AppendTeamToUserParams, *models.User) middleware.Responder
}

// NewAppendTeamToUser creates a new http.Handler for the append team to user operation
func NewAppendTeamToUser(ctx *middleware.Context, handler AppendTeamToUserHandler) *AppendTeamToUser {
	return &AppendTeamToUser{Context: ctx, Handler: handler}
}

/* AppendTeamToUser swagger:route POST /teams/{team_id}/users team appendTeamToUser

Assign a user to team

*/
type AppendTeamToUser struct {
	Context *middleware.Context
	Handler AppendTeamToUserHandler
}

func (o *AppendTeamToUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAppendTeamToUserParams()
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
