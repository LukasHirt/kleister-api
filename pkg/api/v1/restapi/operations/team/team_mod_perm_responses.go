// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// TeamModPermOKCode is the HTTP code returned for type TeamModPermOK
const TeamModPermOKCode int = 200

/*TeamModPermOK Plain success message

swagger:response teamModPermOK
*/
type TeamModPermOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamModPermOK creates TeamModPermOK with default headers values
func NewTeamModPermOK() *TeamModPermOK {

	return &TeamModPermOK{}
}

// WithPayload adds the payload to the team mod perm o k response
func (o *TeamModPermOK) WithPayload(payload *models.GeneralError) *TeamModPermOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team mod perm o k response
func (o *TeamModPermOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamModPermOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamModPermForbiddenCode is the HTTP code returned for type TeamModPermForbidden
const TeamModPermForbiddenCode int = 403

/*TeamModPermForbidden User is not authorized

swagger:response teamModPermForbidden
*/
type TeamModPermForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamModPermForbidden creates TeamModPermForbidden with default headers values
func NewTeamModPermForbidden() *TeamModPermForbidden {

	return &TeamModPermForbidden{}
}

// WithPayload adds the payload to the team mod perm forbidden response
func (o *TeamModPermForbidden) WithPayload(payload *models.GeneralError) *TeamModPermForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team mod perm forbidden response
func (o *TeamModPermForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamModPermForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamModPermPreconditionFailedCode is the HTTP code returned for type TeamModPermPreconditionFailed
const TeamModPermPreconditionFailedCode int = 412

/*TeamModPermPreconditionFailed Failed to parse request body

swagger:response teamModPermPreconditionFailed
*/
type TeamModPermPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamModPermPreconditionFailed creates TeamModPermPreconditionFailed with default headers values
func NewTeamModPermPreconditionFailed() *TeamModPermPreconditionFailed {

	return &TeamModPermPreconditionFailed{}
}

// WithPayload adds the payload to the team mod perm precondition failed response
func (o *TeamModPermPreconditionFailed) WithPayload(payload *models.GeneralError) *TeamModPermPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team mod perm precondition failed response
func (o *TeamModPermPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamModPermPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamModPermUnprocessableEntityCode is the HTTP code returned for type TeamModPermUnprocessableEntity
const TeamModPermUnprocessableEntityCode int = 422

/*TeamModPermUnprocessableEntity Mod is not assigned

swagger:response teamModPermUnprocessableEntity
*/
type TeamModPermUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamModPermUnprocessableEntity creates TeamModPermUnprocessableEntity with default headers values
func NewTeamModPermUnprocessableEntity() *TeamModPermUnprocessableEntity {

	return &TeamModPermUnprocessableEntity{}
}

// WithPayload adds the payload to the team mod perm unprocessable entity response
func (o *TeamModPermUnprocessableEntity) WithPayload(payload *models.GeneralError) *TeamModPermUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team mod perm unprocessable entity response
func (o *TeamModPermUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamModPermUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TeamModPermDefault Some error unrelated to the handler

swagger:response teamModPermDefault
*/
type TeamModPermDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamModPermDefault creates TeamModPermDefault with default headers values
func NewTeamModPermDefault(code int) *TeamModPermDefault {
	if code <= 0 {
		code = 500
	}

	return &TeamModPermDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the team mod perm default response
func (o *TeamModPermDefault) WithStatusCode(code int) *TeamModPermDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the team mod perm default response
func (o *TeamModPermDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the team mod perm default response
func (o *TeamModPermDefault) WithPayload(payload *models.GeneralError) *TeamModPermDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team mod perm default response
func (o *TeamModPermDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamModPermDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
