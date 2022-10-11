// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ListVersionsOKCode is the HTTP code returned for type ListVersionsOK
const ListVersionsOKCode int = 200

/*ListVersionsOK A collection of versions

swagger:response listVersionsOK
*/
type ListVersionsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Version `json:"body,omitempty"`
}

// NewListVersionsOK creates ListVersionsOK with default headers values
func NewListVersionsOK() *ListVersionsOK {

	return &ListVersionsOK{}
}

// WithPayload adds the payload to the list versions o k response
func (o *ListVersionsOK) WithPayload(payload []*models.Version) *ListVersionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list versions o k response
func (o *ListVersionsOK) SetPayload(payload []*models.Version) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVersionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Version, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListVersionsForbiddenCode is the HTTP code returned for type ListVersionsForbidden
const ListVersionsForbiddenCode int = 403

/*ListVersionsForbidden User is not authorized

swagger:response listVersionsForbidden
*/
type ListVersionsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListVersionsForbidden creates ListVersionsForbidden with default headers values
func NewListVersionsForbidden() *ListVersionsForbidden {

	return &ListVersionsForbidden{}
}

// WithPayload adds the payload to the list versions forbidden response
func (o *ListVersionsForbidden) WithPayload(payload *models.GeneralError) *ListVersionsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list versions forbidden response
func (o *ListVersionsForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVersionsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListVersionsNotFoundCode is the HTTP code returned for type ListVersionsNotFound
const ListVersionsNotFoundCode int = 404

/*ListVersionsNotFound Mod not found

swagger:response listVersionsNotFound
*/
type ListVersionsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListVersionsNotFound creates ListVersionsNotFound with default headers values
func NewListVersionsNotFound() *ListVersionsNotFound {

	return &ListVersionsNotFound{}
}

// WithPayload adds the payload to the list versions not found response
func (o *ListVersionsNotFound) WithPayload(payload *models.GeneralError) *ListVersionsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list versions not found response
func (o *ListVersionsNotFound) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVersionsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListVersionsDefault Some error unrelated to the handler

swagger:response listVersionsDefault
*/
type ListVersionsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewListVersionsDefault creates ListVersionsDefault with default headers values
func NewListVersionsDefault(code int) *ListVersionsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListVersionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list versions default response
func (o *ListVersionsDefault) WithStatusCode(code int) *ListVersionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list versions default response
func (o *ListVersionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list versions default response
func (o *ListVersionsDefault) WithPayload(payload *models.GeneralError) *ListVersionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list versions default response
func (o *ListVersionsDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVersionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
