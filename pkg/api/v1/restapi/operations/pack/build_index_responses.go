// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// BuildIndexOKCode is the HTTP code returned for type BuildIndexOK
const BuildIndexOKCode int = 200

/*BuildIndexOK A collection of builds

swagger:response buildIndexOK
*/
type BuildIndexOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Build `json:"body,omitempty"`
}

// NewBuildIndexOK creates BuildIndexOK with default headers values
func NewBuildIndexOK() *BuildIndexOK {

	return &BuildIndexOK{}
}

// WithPayload adds the payload to the build index o k response
func (o *BuildIndexOK) WithPayload(payload []*models.Build) *BuildIndexOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the build index o k response
func (o *BuildIndexOK) SetPayload(payload []*models.Build) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BuildIndexOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Build, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// BuildIndexForbiddenCode is the HTTP code returned for type BuildIndexForbidden
const BuildIndexForbiddenCode int = 403

/*BuildIndexForbidden User is not authorized

swagger:response buildIndexForbidden
*/
type BuildIndexForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewBuildIndexForbidden creates BuildIndexForbidden with default headers values
func NewBuildIndexForbidden() *BuildIndexForbidden {

	return &BuildIndexForbidden{}
}

// WithPayload adds the payload to the build index forbidden response
func (o *BuildIndexForbidden) WithPayload(payload *models.GeneralError) *BuildIndexForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the build index forbidden response
func (o *BuildIndexForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BuildIndexForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*BuildIndexDefault Some error unrelated to the handler

swagger:response buildIndexDefault
*/
type BuildIndexDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewBuildIndexDefault creates BuildIndexDefault with default headers values
func NewBuildIndexDefault(code int) *BuildIndexDefault {
	if code <= 0 {
		code = 500
	}

	return &BuildIndexDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the build index default response
func (o *BuildIndexDefault) WithStatusCode(code int) *BuildIndexDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the build index default response
func (o *BuildIndexDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the build index default response
func (o *BuildIndexDefault) WithPayload(payload *models.GeneralError) *BuildIndexDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the build index default response
func (o *BuildIndexDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BuildIndexDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
