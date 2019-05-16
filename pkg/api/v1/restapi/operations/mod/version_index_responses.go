// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// VersionIndexOKCode is the HTTP code returned for type VersionIndexOK
const VersionIndexOKCode int = 200

/*VersionIndexOK A collection of versions

swagger:response versionIndexOK
*/
type VersionIndexOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Version `json:"body,omitempty"`
}

// NewVersionIndexOK creates VersionIndexOK with default headers values
func NewVersionIndexOK() *VersionIndexOK {

	return &VersionIndexOK{}
}

// WithPayload adds the payload to the version index o k response
func (o *VersionIndexOK) WithPayload(payload []*models.Version) *VersionIndexOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the version index o k response
func (o *VersionIndexOK) SetPayload(payload []*models.Version) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VersionIndexOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// VersionIndexForbiddenCode is the HTTP code returned for type VersionIndexForbidden
const VersionIndexForbiddenCode int = 403

/*VersionIndexForbidden User is not authorized

swagger:response versionIndexForbidden
*/
type VersionIndexForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewVersionIndexForbidden creates VersionIndexForbidden with default headers values
func NewVersionIndexForbidden() *VersionIndexForbidden {

	return &VersionIndexForbidden{}
}

// WithPayload adds the payload to the version index forbidden response
func (o *VersionIndexForbidden) WithPayload(payload *models.GeneralError) *VersionIndexForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the version index forbidden response
func (o *VersionIndexForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VersionIndexForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*VersionIndexDefault Some error unrelated to the handler

swagger:response versionIndexDefault
*/
type VersionIndexDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewVersionIndexDefault creates VersionIndexDefault with default headers values
func NewVersionIndexDefault(code int) *VersionIndexDefault {
	if code <= 0 {
		code = 500
	}

	return &VersionIndexDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the version index default response
func (o *VersionIndexDefault) WithStatusCode(code int) *VersionIndexDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the version index default response
func (o *VersionIndexDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the version index default response
func (o *VersionIndexDefault) WithPayload(payload *models.GeneralError) *VersionIndexDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the version index default response
func (o *VersionIndexDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VersionIndexDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
