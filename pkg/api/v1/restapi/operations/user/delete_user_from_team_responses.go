// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// DeleteUserFromTeamOKCode is the HTTP code returned for type DeleteUserFromTeamOK
const DeleteUserFromTeamOKCode int = 200

/*DeleteUserFromTeamOK Plain success message

swagger:response deleteUserFromTeamOK
*/
type DeleteUserFromTeamOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserFromTeamOK creates DeleteUserFromTeamOK with default headers values
func NewDeleteUserFromTeamOK() *DeleteUserFromTeamOK {

	return &DeleteUserFromTeamOK{}
}

// WithPayload adds the payload to the delete user from team o k response
func (o *DeleteUserFromTeamOK) WithPayload(payload *models.GeneralError) *DeleteUserFromTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user from team o k response
func (o *DeleteUserFromTeamOK) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserFromTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserFromTeamForbiddenCode is the HTTP code returned for type DeleteUserFromTeamForbidden
const DeleteUserFromTeamForbiddenCode int = 403

/*DeleteUserFromTeamForbidden User is not authorized

swagger:response deleteUserFromTeamForbidden
*/
type DeleteUserFromTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserFromTeamForbidden creates DeleteUserFromTeamForbidden with default headers values
func NewDeleteUserFromTeamForbidden() *DeleteUserFromTeamForbidden {

	return &DeleteUserFromTeamForbidden{}
}

// WithPayload adds the payload to the delete user from team forbidden response
func (o *DeleteUserFromTeamForbidden) WithPayload(payload *models.GeneralError) *DeleteUserFromTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user from team forbidden response
func (o *DeleteUserFromTeamForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserFromTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserFromTeamNotFoundCode is the HTTP code returned for type DeleteUserFromTeamNotFound
const DeleteUserFromTeamNotFoundCode int = 404

/*DeleteUserFromTeamNotFound User or team not found

swagger:response deleteUserFromTeamNotFound
*/
type DeleteUserFromTeamNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserFromTeamNotFound creates DeleteUserFromTeamNotFound with default headers values
func NewDeleteUserFromTeamNotFound() *DeleteUserFromTeamNotFound {

	return &DeleteUserFromTeamNotFound{}
}

// WithPayload adds the payload to the delete user from team not found response
func (o *DeleteUserFromTeamNotFound) WithPayload(payload *models.GeneralError) *DeleteUserFromTeamNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user from team not found response
func (o *DeleteUserFromTeamNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserFromTeamNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserFromTeamPreconditionFailedCode is the HTTP code returned for type DeleteUserFromTeamPreconditionFailed
const DeleteUserFromTeamPreconditionFailedCode int = 412

/*DeleteUserFromTeamPreconditionFailed Team is not assigned

swagger:response deleteUserFromTeamPreconditionFailed
*/
type DeleteUserFromTeamPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserFromTeamPreconditionFailed creates DeleteUserFromTeamPreconditionFailed with default headers values
func NewDeleteUserFromTeamPreconditionFailed() *DeleteUserFromTeamPreconditionFailed {

	return &DeleteUserFromTeamPreconditionFailed{}
}

// WithPayload adds the payload to the delete user from team precondition failed response
func (o *DeleteUserFromTeamPreconditionFailed) WithPayload(payload *models.GeneralError) *DeleteUserFromTeamPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user from team precondition failed response
func (o *DeleteUserFromTeamPreconditionFailed) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserFromTeamPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteUserFromTeamDefault Some error unrelated to the handler

swagger:response deleteUserFromTeamDefault
*/
type DeleteUserFromTeamDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewDeleteUserFromTeamDefault creates DeleteUserFromTeamDefault with default headers values
func NewDeleteUserFromTeamDefault(code int) *DeleteUserFromTeamDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteUserFromTeamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete user from team default response
func (o *DeleteUserFromTeamDefault) WithStatusCode(code int) *DeleteUserFromTeamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete user from team default response
func (o *DeleteUserFromTeamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete user from team default response
func (o *DeleteUserFromTeamDefault) WithPayload(payload *models.GeneralError) *DeleteUserFromTeamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user from team default response
func (o *DeleteUserFromTeamDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserFromTeamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
