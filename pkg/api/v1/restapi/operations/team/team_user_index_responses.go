// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// TeamUserIndexOKCode is the HTTP code returned for type TeamUserIndexOK
const TeamUserIndexOKCode int = 200

/*TeamUserIndexOK A collection of team users

swagger:response teamUserIndexOK
*/
type TeamUserIndexOK struct {

	/*
	  In: Body
	*/
	Payload []*models.TeamUser `json:"body,omitempty"`
}

// NewTeamUserIndexOK creates TeamUserIndexOK with default headers values
func NewTeamUserIndexOK() *TeamUserIndexOK {

	return &TeamUserIndexOK{}
}

// WithPayload adds the payload to the team user index o k response
func (o *TeamUserIndexOK) WithPayload(payload []*models.TeamUser) *TeamUserIndexOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user index o k response
func (o *TeamUserIndexOK) SetPayload(payload []*models.TeamUser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserIndexOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// TeamUserIndexForbiddenCode is the HTTP code returned for type TeamUserIndexForbidden
const TeamUserIndexForbiddenCode int = 403

/*TeamUserIndexForbidden User is not authorized

swagger:response teamUserIndexForbidden
*/
type TeamUserIndexForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserIndexForbidden creates TeamUserIndexForbidden with default headers values
func NewTeamUserIndexForbidden() *TeamUserIndexForbidden {

	return &TeamUserIndexForbidden{}
}

// WithPayload adds the payload to the team user index forbidden response
func (o *TeamUserIndexForbidden) WithPayload(payload *models.GeneralError) *TeamUserIndexForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user index forbidden response
func (o *TeamUserIndexForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserIndexForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TeamUserIndexDefault Some error unrelated to the handler

swagger:response teamUserIndexDefault
*/
type TeamUserIndexDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTeamUserIndexDefault creates TeamUserIndexDefault with default headers values
func NewTeamUserIndexDefault(code int) *TeamUserIndexDefault {
	if code <= 0 {
		code = 500
	}

	return &TeamUserIndexDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the team user index default response
func (o *TeamUserIndexDefault) WithStatusCode(code int) *TeamUserIndexDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the team user index default response
func (o *TeamUserIndexDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the team user index default response
func (o *TeamUserIndexDefault) WithPayload(payload *models.GeneralError) *TeamUserIndexDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the team user index default response
func (o *TeamUserIndexDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TeamUserIndexDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
