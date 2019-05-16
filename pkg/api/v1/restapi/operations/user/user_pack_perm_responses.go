// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// UserPackPermOKCode is the HTTP code returned for type UserPackPermOK
const UserPackPermOKCode int = 200

/*UserPackPermOK Plain success message

swagger:response userPackPermOK
*/
type UserPackPermOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserPackPermOK creates UserPackPermOK with default headers values
func NewUserPackPermOK() *UserPackPermOK {

	return &UserPackPermOK{}
}

// WithPayload adds the payload to the user pack perm o k response
func (o *UserPackPermOK) WithPayload(payload *models.GeneralError) *UserPackPermOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user pack perm o k response
func (o *UserPackPermOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserPackPermOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserPackPermForbiddenCode is the HTTP code returned for type UserPackPermForbidden
const UserPackPermForbiddenCode int = 403

/*UserPackPermForbidden User is not authorized

swagger:response userPackPermForbidden
*/
type UserPackPermForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserPackPermForbidden creates UserPackPermForbidden with default headers values
func NewUserPackPermForbidden() *UserPackPermForbidden {

	return &UserPackPermForbidden{}
}

// WithPayload adds the payload to the user pack perm forbidden response
func (o *UserPackPermForbidden) WithPayload(payload *models.GeneralError) *UserPackPermForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user pack perm forbidden response
func (o *UserPackPermForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserPackPermForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserPackPermPreconditionFailedCode is the HTTP code returned for type UserPackPermPreconditionFailed
const UserPackPermPreconditionFailedCode int = 412

/*UserPackPermPreconditionFailed Failed to parse request body

swagger:response userPackPermPreconditionFailed
*/
type UserPackPermPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserPackPermPreconditionFailed creates UserPackPermPreconditionFailed with default headers values
func NewUserPackPermPreconditionFailed() *UserPackPermPreconditionFailed {

	return &UserPackPermPreconditionFailed{}
}

// WithPayload adds the payload to the user pack perm precondition failed response
func (o *UserPackPermPreconditionFailed) WithPayload(payload *models.GeneralError) *UserPackPermPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user pack perm precondition failed response
func (o *UserPackPermPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserPackPermPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserPackPermUnprocessableEntityCode is the HTTP code returned for type UserPackPermUnprocessableEntity
const UserPackPermUnprocessableEntityCode int = 422

/*UserPackPermUnprocessableEntity Pack is not assigned

swagger:response userPackPermUnprocessableEntity
*/
type UserPackPermUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserPackPermUnprocessableEntity creates UserPackPermUnprocessableEntity with default headers values
func NewUserPackPermUnprocessableEntity() *UserPackPermUnprocessableEntity {

	return &UserPackPermUnprocessableEntity{}
}

// WithPayload adds the payload to the user pack perm unprocessable entity response
func (o *UserPackPermUnprocessableEntity) WithPayload(payload *models.GeneralError) *UserPackPermUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user pack perm unprocessable entity response
func (o *UserPackPermUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserPackPermUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UserPackPermDefault Some error unrelated to the handler

swagger:response userPackPermDefault
*/
type UserPackPermDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserPackPermDefault creates UserPackPermDefault with default headers values
func NewUserPackPermDefault(code int) *UserPackPermDefault {
	if code <= 0 {
		code = 500
	}

	return &UserPackPermDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user pack perm default response
func (o *UserPackPermDefault) WithStatusCode(code int) *UserPackPermDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user pack perm default response
func (o *UserPackPermDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user pack perm default response
func (o *UserPackPermDefault) WithPayload(payload *models.GeneralError) *UserPackPermDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user pack perm default response
func (o *UserPackPermDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserPackPermDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
