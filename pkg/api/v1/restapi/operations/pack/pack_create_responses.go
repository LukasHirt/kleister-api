// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// PackCreateOKCode is the HTTP code returned for type PackCreateOK
const PackCreateOKCode int = 200

/*PackCreateOK The created pack data

swagger:response packCreateOK
*/
type PackCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Pack `json:"body,omitempty"`
}

// NewPackCreateOK creates PackCreateOK with default headers values
func NewPackCreateOK() *PackCreateOK {

	return &PackCreateOK{}
}

// WithPayload adds the payload to the pack create o k response
func (o *PackCreateOK) WithPayload(payload *models.Pack) *PackCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pack create o k response
func (o *PackCreateOK) SetPayload(payload *models.Pack) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PackCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PackCreateForbiddenCode is the HTTP code returned for type PackCreateForbidden
const PackCreateForbiddenCode int = 403

/*PackCreateForbidden User is not authorized

swagger:response packCreateForbidden
*/
type PackCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPackCreateForbidden creates PackCreateForbidden with default headers values
func NewPackCreateForbidden() *PackCreateForbidden {

	return &PackCreateForbidden{}
}

// WithPayload adds the payload to the pack create forbidden response
func (o *PackCreateForbidden) WithPayload(payload *models.GeneralError) *PackCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pack create forbidden response
func (o *PackCreateForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PackCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PackCreatePreconditionFailedCode is the HTTP code returned for type PackCreatePreconditionFailed
const PackCreatePreconditionFailedCode int = 412

/*PackCreatePreconditionFailed Failed to parse request body

swagger:response packCreatePreconditionFailed
*/
type PackCreatePreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPackCreatePreconditionFailed creates PackCreatePreconditionFailed with default headers values
func NewPackCreatePreconditionFailed() *PackCreatePreconditionFailed {

	return &PackCreatePreconditionFailed{}
}

// WithPayload adds the payload to the pack create precondition failed response
func (o *PackCreatePreconditionFailed) WithPayload(payload *models.GeneralError) *PackCreatePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pack create precondition failed response
func (o *PackCreatePreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PackCreatePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PackCreateUnprocessableEntityCode is the HTTP code returned for type PackCreateUnprocessableEntity
const PackCreateUnprocessableEntityCode int = 422

/*PackCreateUnprocessableEntity Failed to validate request

swagger:response packCreateUnprocessableEntity
*/
type PackCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewPackCreateUnprocessableEntity creates PackCreateUnprocessableEntity with default headers values
func NewPackCreateUnprocessableEntity() *PackCreateUnprocessableEntity {

	return &PackCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the pack create unprocessable entity response
func (o *PackCreateUnprocessableEntity) WithPayload(payload *models.ValidationError) *PackCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pack create unprocessable entity response
func (o *PackCreateUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PackCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PackCreateDefault Some error unrelated to the handler

swagger:response packCreateDefault
*/
type PackCreateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPackCreateDefault creates PackCreateDefault with default headers values
func NewPackCreateDefault(code int) *PackCreateDefault {
	if code <= 0 {
		code = 500
	}

	return &PackCreateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the pack create default response
func (o *PackCreateDefault) WithStatusCode(code int) *PackCreateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the pack create default response
func (o *PackCreateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the pack create default response
func (o *PackCreateDefault) WithPayload(payload *models.GeneralError) *PackCreateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pack create default response
func (o *PackCreateDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PackCreateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
