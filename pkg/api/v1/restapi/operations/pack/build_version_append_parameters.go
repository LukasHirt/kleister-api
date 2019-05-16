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

// NewBuildVersionAppendParams creates a new BuildVersionAppendParams object
// no default values defined in spec.
func NewBuildVersionAppendParams() BuildVersionAppendParams {

	return BuildVersionAppendParams{}
}

// BuildVersionAppendParams contains all the bound params for the build version append operation
// typically these are obtained from a http.Request
//
// swagger:parameters BuildVersionAppend
type BuildVersionAppendParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A build UUID or slug
	  Required: true
	  In: path
	*/
	BuildID string
	/*A pack UUID or slug
	  Required: true
	  In: path
	*/
	PackID string
	/*The version data to append to build
	  Required: true
	  In: body
	*/
	Params *models.BuildVersionParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewBuildVersionAppendParams() beforehand.
func (o *BuildVersionAppendParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBuildID, rhkBuildID, _ := route.Params.GetOK("buildID")
	if err := o.bindBuildID(rBuildID, rhkBuildID, route.Formats); err != nil {
		res = append(res, err)
	}

	rPackID, rhkPackID, _ := route.Params.GetOK("packID")
	if err := o.bindPackID(rPackID, rhkPackID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.BuildVersionParams
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

// bindBuildID binds and validates parameter BuildID from path.
func (o *BuildVersionAppendParams) bindBuildID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.BuildID = raw

	return nil
}

// bindPackID binds and validates parameter PackID from path.
func (o *BuildVersionAppendParams) bindPackID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PackID = raw

	return nil
}
