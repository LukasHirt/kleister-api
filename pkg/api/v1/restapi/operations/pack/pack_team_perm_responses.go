// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// PackTeamPermOKCode is the HTTP code returned for type PackTeamPermOK
const PackTeamPermOKCode int = 200

/*PackTeamPermOK Plain success message

swagger:response packTeamPermOK
*/
type PackTeamPermOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPackTeamPermOK creates PackTeamPermOK with default headers values
func NewPackTeamPermOK() *PackTeamPermOK {

	return &PackTeamPermOK{}
}

// WithPayload adds the payload to the pack team perm o k response
func (o *PackTeamPermOK) WithPayload(payload *models.GeneralError) *PackTeamPermOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pack team perm o k response
func (o *PackTeamPermOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PackTeamPermOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PackTeamPermForbiddenCode is the HTTP code returned for type PackTeamPermForbidden
const PackTeamPermForbiddenCode int = 403

/*PackTeamPermForbidden User is not authorized

swagger:response packTeamPermForbidden
*/
type PackTeamPermForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPackTeamPermForbidden creates PackTeamPermForbidden with default headers values
func NewPackTeamPermForbidden() *PackTeamPermForbidden {

	return &PackTeamPermForbidden{}
}

// WithPayload adds the payload to the pack team perm forbidden response
func (o *PackTeamPermForbidden) WithPayload(payload *models.GeneralError) *PackTeamPermForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pack team perm forbidden response
func (o *PackTeamPermForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PackTeamPermForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PackTeamPermPreconditionFailedCode is the HTTP code returned for type PackTeamPermPreconditionFailed
const PackTeamPermPreconditionFailedCode int = 412

/*PackTeamPermPreconditionFailed Failed to parse request body

swagger:response packTeamPermPreconditionFailed
*/
type PackTeamPermPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPackTeamPermPreconditionFailed creates PackTeamPermPreconditionFailed with default headers values
func NewPackTeamPermPreconditionFailed() *PackTeamPermPreconditionFailed {

	return &PackTeamPermPreconditionFailed{}
}

// WithPayload adds the payload to the pack team perm precondition failed response
func (o *PackTeamPermPreconditionFailed) WithPayload(payload *models.GeneralError) *PackTeamPermPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pack team perm precondition failed response
func (o *PackTeamPermPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PackTeamPermPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PackTeamPermUnprocessableEntityCode is the HTTP code returned for type PackTeamPermUnprocessableEntity
const PackTeamPermUnprocessableEntityCode int = 422

/*PackTeamPermUnprocessableEntity Team is not assigned

swagger:response packTeamPermUnprocessableEntity
*/
type PackTeamPermUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPackTeamPermUnprocessableEntity creates PackTeamPermUnprocessableEntity with default headers values
func NewPackTeamPermUnprocessableEntity() *PackTeamPermUnprocessableEntity {

	return &PackTeamPermUnprocessableEntity{}
}

// WithPayload adds the payload to the pack team perm unprocessable entity response
func (o *PackTeamPermUnprocessableEntity) WithPayload(payload *models.GeneralError) *PackTeamPermUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pack team perm unprocessable entity response
func (o *PackTeamPermUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PackTeamPermUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PackTeamPermDefault Some error unrelated to the handler

swagger:response packTeamPermDefault
*/
type PackTeamPermDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewPackTeamPermDefault creates PackTeamPermDefault with default headers values
func NewPackTeamPermDefault(code int) *PackTeamPermDefault {
	if code <= 0 {
		code = 500
	}

	return &PackTeamPermDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the pack team perm default response
func (o *PackTeamPermDefault) WithStatusCode(code int) *PackTeamPermDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the pack team perm default response
func (o *PackTeamPermDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the pack team perm default response
func (o *PackTeamPermDefault) WithPayload(payload *models.GeneralError) *PackTeamPermDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pack team perm default response
func (o *PackTeamPermDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PackTeamPermDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
