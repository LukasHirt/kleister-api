// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// TeamUserAppendOKCode is the HTTP code returned for type TeamUserAppendOK
const TeamUserAppendOKCode int = 200

/*TeamUserAppendOK Plain success message

swagger:response teamUserAppendOK
*/
type TeamUserAppendOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserAppendOK creates TeamUserAppendOK with default headers values
func NewTeamUserAppendOK() *TeamUserAppendOK {

	return &TeamUserAppendOK{}
}

// WithPayload adds the payload to the team user append o k response
func (o *TeamUserAppendOK) WithPayload(payload *models.GeneralError) *TeamUserAppendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user append o k response
func (o *TeamUserAppendOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserAppendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamUserAppendForbiddenCode is the HTTP code returned for type TeamUserAppendForbidden
const TeamUserAppendForbiddenCode int = 403

/*TeamUserAppendForbidden User is not authorized

swagger:response teamUserAppendForbidden
*/
type TeamUserAppendForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserAppendForbidden creates TeamUserAppendForbidden with default headers values
func NewTeamUserAppendForbidden() *TeamUserAppendForbidden {

	return &TeamUserAppendForbidden{}
}

// WithPayload adds the payload to the team user append forbidden response
func (o *TeamUserAppendForbidden) WithPayload(payload *models.GeneralError) *TeamUserAppendForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user append forbidden response
func (o *TeamUserAppendForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserAppendForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamUserAppendPreconditionFailedCode is the HTTP code returned for type TeamUserAppendPreconditionFailed
const TeamUserAppendPreconditionFailedCode int = 412

/*TeamUserAppendPreconditionFailed Failed to parse request body

swagger:response teamUserAppendPreconditionFailed
*/
type TeamUserAppendPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserAppendPreconditionFailed creates TeamUserAppendPreconditionFailed with default headers values
func NewTeamUserAppendPreconditionFailed() *TeamUserAppendPreconditionFailed {

	return &TeamUserAppendPreconditionFailed{}
}

// WithPayload adds the payload to the team user append precondition failed response
func (o *TeamUserAppendPreconditionFailed) WithPayload(payload *models.GeneralError) *TeamUserAppendPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user append precondition failed response
func (o *TeamUserAppendPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserAppendPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TeamUserAppendUnprocessableEntityCode is the HTTP code returned for type TeamUserAppendUnprocessableEntity
const TeamUserAppendUnprocessableEntityCode int = 422

/*TeamUserAppendUnprocessableEntity User is already assigned

swagger:response teamUserAppendUnprocessableEntity
*/
type TeamUserAppendUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserAppendUnprocessableEntity creates TeamUserAppendUnprocessableEntity with default headers values
func NewTeamUserAppendUnprocessableEntity() *TeamUserAppendUnprocessableEntity {

	return &TeamUserAppendUnprocessableEntity{}
}

// WithPayload adds the payload to the team user append unprocessable entity response
func (o *TeamUserAppendUnprocessableEntity) WithPayload(payload *models.GeneralError) *TeamUserAppendUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user append unprocessable entity response
func (o *TeamUserAppendUnprocessableEntity) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserAppendUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TeamUserAppendDefault Some error unrelated to the handler

swagger:response teamUserAppendDefault
*/
type TeamUserAppendDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserAppendDefault creates TeamUserAppendDefault with default headers values
func NewTeamUserAppendDefault(code int) *TeamUserAppendDefault {
	if code <= 0 {
		code = 500
	}

	return &TeamUserAppendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the team user append default response
func (o *TeamUserAppendDefault) WithStatusCode(code int) *TeamUserAppendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the team user append default response
func (o *TeamUserAppendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the team user append default response
func (o *TeamUserAppendDefault) WithPayload(payload *models.GeneralError) *TeamUserAppendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user append default response
func (o *TeamUserAppendDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserAppendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
