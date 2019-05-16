// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// BuildVersionAppendOKCode is the HTTP code returned for type BuildVersionAppendOK
const BuildVersionAppendOKCode int = 200

/*BuildVersionAppendOK Plain success message

swagger:response buildVersionAppendOK
*/
type BuildVersionAppendOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewBuildVersionAppendOK creates BuildVersionAppendOK with default headers values
func NewBuildVersionAppendOK() *BuildVersionAppendOK {

	return &BuildVersionAppendOK{}
}

// WithPayload adds the payload to the build version append o k response
func (o *BuildVersionAppendOK) WithPayload(payload *models.GeneralError) *BuildVersionAppendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the build version append o k response
func (o *BuildVersionAppendOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BuildVersionAppendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BuildVersionAppendForbiddenCode is the HTTP code returned for type BuildVersionAppendForbidden
const BuildVersionAppendForbiddenCode int = 403

/*BuildVersionAppendForbidden User is not authorized

swagger:response buildVersionAppendForbidden
*/
type BuildVersionAppendForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewBuildVersionAppendForbidden creates BuildVersionAppendForbidden with default headers values
func NewBuildVersionAppendForbidden() *BuildVersionAppendForbidden {

	return &BuildVersionAppendForbidden{}
}

// WithPayload adds the payload to the build version append forbidden response
func (o *BuildVersionAppendForbidden) WithPayload(payload *models.GeneralError) *BuildVersionAppendForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the build version append forbidden response
func (o *BuildVersionAppendForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BuildVersionAppendForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BuildVersionAppendPreconditionFailedCode is the HTTP code returned for type BuildVersionAppendPreconditionFailed
const BuildVersionAppendPreconditionFailedCode int = 412

/*BuildVersionAppendPreconditionFailed Failed to parse request body

swagger:response buildVersionAppendPreconditionFailed
*/
type BuildVersionAppendPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewBuildVersionAppendPreconditionFailed creates BuildVersionAppendPreconditionFailed with default headers values
func NewBuildVersionAppendPreconditionFailed() *BuildVersionAppendPreconditionFailed {

	return &BuildVersionAppendPreconditionFailed{}
}

// WithPayload adds the payload to the build version append precondition failed response
func (o *BuildVersionAppendPreconditionFailed) WithPayload(payload *models.GeneralError) *BuildVersionAppendPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the build version append precondition failed response
func (o *BuildVersionAppendPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BuildVersionAppendPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BuildVersionAppendUnprocessableEntityCode is the HTTP code returned for type BuildVersionAppendUnprocessableEntity
const BuildVersionAppendUnprocessableEntityCode int = 422

/*BuildVersionAppendUnprocessableEntity Version is already appended

swagger:response buildVersionAppendUnprocessableEntity
*/
type BuildVersionAppendUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewBuildVersionAppendUnprocessableEntity creates BuildVersionAppendUnprocessableEntity with default headers values
func NewBuildVersionAppendUnprocessableEntity() *BuildVersionAppendUnprocessableEntity {

	return &BuildVersionAppendUnprocessableEntity{}
}

// WithPayload adds the payload to the build version append unprocessable entity response
func (o *BuildVersionAppendUnprocessableEntity) WithPayload(payload *models.ValidationError) *BuildVersionAppendUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the build version append unprocessable entity response
func (o *BuildVersionAppendUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BuildVersionAppendUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*BuildVersionAppendDefault Some error unrelated to the handler

swagger:response buildVersionAppendDefault
*/
type BuildVersionAppendDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewBuildVersionAppendDefault creates BuildVersionAppendDefault with default headers values
func NewBuildVersionAppendDefault(code int) *BuildVersionAppendDefault {
	if code <= 0 {
		code = 500
	}

	return &BuildVersionAppendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the build version append default response
func (o *BuildVersionAppendDefault) WithStatusCode(code int) *BuildVersionAppendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the build version append default response
func (o *BuildVersionAppendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the build version append default response
func (o *BuildVersionAppendDefault) WithPayload(payload *models.GeneralError) *BuildVersionAppendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the build version append default response
func (o *BuildVersionAppendDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BuildVersionAppendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
