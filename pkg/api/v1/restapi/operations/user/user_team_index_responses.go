// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// UserTeamIndexOKCode is the HTTP code returned for type UserTeamIndexOK
const UserTeamIndexOKCode int = 200

/*UserTeamIndexOK A collection of user teams

swagger:response userTeamIndexOK
*/
type UserTeamIndexOK struct {

	/*
	  In: Body
	*/
	Payload []*models.TeamUser `json:"body,omitempty"`
}

// NewUserTeamIndexOK creates UserTeamIndexOK with default headers values
func NewUserTeamIndexOK() *UserTeamIndexOK {

	return &UserTeamIndexOK{}
}

// WithPayload adds the payload to the user team index o k response
func (o *UserTeamIndexOK) WithPayload(payload []*models.TeamUser) *UserTeamIndexOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user team index o k response
func (o *UserTeamIndexOK) SetPayload(payload []*models.TeamUser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserTeamIndexOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.TeamUser, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UserTeamIndexForbiddenCode is the HTTP code returned for type UserTeamIndexForbidden
const UserTeamIndexForbiddenCode int = 403

/*UserTeamIndexForbidden User is not authorized

swagger:response userTeamIndexForbidden
*/
type UserTeamIndexForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserTeamIndexForbidden creates UserTeamIndexForbidden with default headers values
func NewUserTeamIndexForbidden() *UserTeamIndexForbidden {

	return &UserTeamIndexForbidden{}
}

// WithPayload adds the payload to the user team index forbidden response
func (o *UserTeamIndexForbidden) WithPayload(payload *models.GeneralError) *UserTeamIndexForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user team index forbidden response
func (o *UserTeamIndexForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserTeamIndexForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UserTeamIndexDefault Some error unrelated to the handler

swagger:response userTeamIndexDefault
*/
type UserTeamIndexDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewUserTeamIndexDefault creates UserTeamIndexDefault with default headers values
func NewUserTeamIndexDefault(code int) *UserTeamIndexDefault {
	if code <= 0 {
		code = 500
	}

	return &UserTeamIndexDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user team index default response
func (o *UserTeamIndexDefault) WithStatusCode(code int) *UserTeamIndexDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user team index default response
func (o *UserTeamIndexDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user team index default response
func (o *UserTeamIndexDefault) WithPayload(payload *models.GeneralError) *UserTeamIndexDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user team index default response
func (o *UserTeamIndexDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserTeamIndexDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
