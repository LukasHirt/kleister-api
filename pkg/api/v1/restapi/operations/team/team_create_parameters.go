// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// NewTeamCreateParams creates a new TeamCreateParams object
// no default values defined in spec.
func NewTeamCreateParams() TeamCreateParams {

	return TeamCreateParams{}
}

// TeamCreateParams contains all the bound params for the team create operation
// typically these are obtained from a http.Request
//
// swagger:parameters TeamCreate
type TeamCreateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The team data to create
	  Required: true
	  In: body
	*/
	Params *models.Team
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewTeamCreateParams() beforehand.
func (o *TeamCreateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Team
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("params", "body"))
			} else {
				res = append(res, errors.NewParseError("params", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Params = &body
			}
		}
	} else {
		res = append(res, errors.Required("params", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
