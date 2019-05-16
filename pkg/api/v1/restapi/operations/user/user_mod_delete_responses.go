// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// UserModDeleteOKCode is the HTTP code returned for type UserModDeleteOK
const UserModDeleteOKCode int = 200

/*UserModDeleteOK Plain success message

swagger:response userModDeleteOK
*/
type UserModDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserModDeleteOK creates UserModDeleteOK with default headers values
func NewUserModDeleteOK() *UserModDeleteOK {

	return &UserModDeleteOK{}
}

// WithPayload adds the payload to the user mod delete o k response
func (o *UserModDeleteOK) WithPayload(payload *models.GeneralError) *UserModDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user mod delete o k response
func (o *UserModDeleteOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserModDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserModDeleteForbiddenCode is the HTTP code returned for type UserModDeleteForbidden
const UserModDeleteForbiddenCode int = 403

/*UserModDeleteForbidden User is not authorized

swagger:response userModDeleteForbidden
*/
type UserModDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserModDeleteForbidden creates UserModDeleteForbidden with default headers values
func NewUserModDeleteForbidden() *UserModDeleteForbidden {

	return &UserModDeleteForbidden{}
}

// WithPayload adds the payload to the user mod delete forbidden response
func (o *UserModDeleteForbidden) WithPayload(payload *models.GeneralError) *UserModDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user mod delete forbidden response
func (o *UserModDeleteForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserModDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserModDeletePreconditionFailedCode is the HTTP code returned for type UserModDeletePreconditionFailed
const UserModDeletePreconditionFailedCode int = 412

/*UserModDeletePreconditionFailed Failed to parse request body

swagger:response userModDeletePreconditionFailed
*/
type UserModDeletePreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserModDeletePreconditionFailed creates UserModDeletePreconditionFailed with default headers values
func NewUserModDeletePreconditionFailed() *UserModDeletePreconditionFailed {

	return &UserModDeletePreconditionFailed{}
}

// WithPayload adds the payload to the user mod delete precondition failed response
func (o *UserModDeletePreconditionFailed) WithPayload(payload *models.GeneralError) *UserModDeletePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user mod delete precondition failed response
func (o *UserModDeletePreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserModDeletePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserModDeleteUnprocessableEntityCode is the HTTP code returned for type UserModDeleteUnprocessableEntity
const UserModDeleteUnprocessableEntityCode int = 422

/*UserModDeleteUnprocessableEntity Mod is not assigned

swagger:response userModDeleteUnprocessableEntity
*/
type UserModDeleteUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserModDeleteUnprocessableEntity creates UserModDeleteUnprocessableEntity with default headers values
func NewUserModDeleteUnprocessableEntity() *UserModDeleteUnprocessableEntity {

	return &UserModDeleteUnprocessableEntity{}
}

// WithPayload adds the payload to the user mod delete unprocessable entity response
func (o *UserModDeleteUnprocessableEntity) WithPayload(payload *models.GeneralError) *UserModDeleteUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user mod delete unprocessable entity response
func (o *UserModDeleteUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserModDeleteUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UserModDeleteDefault Some error unrelated to the handler

swagger:response userModDeleteDefault
*/
type UserModDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserModDeleteDefault creates UserModDeleteDefault with default headers values
func NewUserModDeleteDefault(code int) *UserModDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &UserModDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user mod delete default response
func (o *UserModDeleteDefault) WithStatusCode(code int) *UserModDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user mod delete default response
func (o *UserModDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user mod delete default response
func (o *UserModDeleteDefault) WithPayload(payload *models.GeneralError) *UserModDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user mod delete default response
func (o *UserModDeleteDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserModDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
