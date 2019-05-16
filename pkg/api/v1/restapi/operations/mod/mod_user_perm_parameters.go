// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// NewModUserPermParams creates a new ModUserPermParams object
// no default values defined in spec.
func NewModUserPermParams() ModUserPermParams {

	return ModUserPermParams{}
}

// ModUserPermParams contains all the bound params for the mod user perm operation
// typically these are obtained from a http.Request
//
// swagger:parameters ModUserPerm
type ModUserPermParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A mod UUID or slug
	  Required: true
	  In: path
	*/
	ModID string
	/*The mod user data to update
	  Required: true
	  In: body
	*/
	Params *models.ModUserParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewModUserPermParams() beforehand.
func (o *ModUserPermParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rModID, rhkModID, _ := route.Params.GetOK("modID")
	if err := o.bindModID(rModID, rhkModID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ModUserParams
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

// bindModID binds and validates parameter ModID from path.
func (o *ModUserPermParams) bindModID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ModID = raw

	return nil
}
