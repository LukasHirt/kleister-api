// Code generated by go-swagger; DO NOT EDIT.

package pack

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

// NewPackUserDeleteParams creates a new PackUserDeleteParams object
// no default values defined in spec.
func NewPackUserDeleteParams() PackUserDeleteParams {

	return PackUserDeleteParams{}
}

// PackUserDeleteParams contains all the bound params for the pack user delete operation
// typically these are obtained from a http.Request
//
// swagger:parameters PackUserDelete
type PackUserDeleteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A pack UUID or slug
	  Required: true
	  In: path
	*/
	PackID string
	/*The pack user data to delete
	  Required: true
	  In: body
	*/
	Params *models.PackUserParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPackUserDeleteParams() beforehand.
func (o *PackUserDeleteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPackID, rhkPackID, _ := route.Params.GetOK("packID")
	if err := o.bindPackID(rPackID, rhkPackID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PackUserParams
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

// bindPackID binds and validates parameter PackID from path.
func (o *PackUserDeleteParams) bindPackID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PackID = raw

	return nil
}
