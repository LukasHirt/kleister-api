// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewModUserIndexParams creates a new ModUserIndexParams object
// no default values defined in spec.
func NewModUserIndexParams() ModUserIndexParams {

	return ModUserIndexParams{}
}

// ModUserIndexParams contains all the bound params for the mod user index operation
// typically these are obtained from a http.Request
//
// swagger:parameters ModUserIndex
type ModUserIndexParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A mod UUID or slug
	  Required: true
	  In: path
	*/
	ModID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewModUserIndexParams() beforehand.
func (o *ModUserIndexParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rModID, rhkModID, _ := route.Params.GetOK("modID")
	if err := o.bindModID(rModID, rhkModID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindModID binds and validates parameter ModID from path.
func (o *ModUserIndexParams) bindModID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ModID = raw

	return nil
}
