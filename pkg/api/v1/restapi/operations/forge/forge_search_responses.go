// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kleister/kleister-api/pkg/api/v1/models"
)

// ForgeSearchOKCode is the HTTP code returned for type ForgeSearchOK
const ForgeSearchOKCode int = 200

/*ForgeSearchOK A collection of Forge versions

swagger:response forgeSearchOK
*/
type ForgeSearchOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Forge `json:"body,omitempty"`
}

// NewForgeSearchOK creates ForgeSearchOK with default headers values
func NewForgeSearchOK() *ForgeSearchOK {

	return &ForgeSearchOK{}
}

// WithPayload adds the payload to the forge search o k response
func (o *ForgeSearchOK) WithPayload(payload []*models.Forge) *ForgeSearchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the forge search o k response
func (o *ForgeSearchOK) SetPayload(payload []*models.Forge) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ForgeSearchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Forge, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ForgeSearchForbiddenCode is the HTTP code returned for type ForgeSearchForbidden
const ForgeSearchForbiddenCode int = 403

/*ForgeSearchForbidden User is not authorized

swagger:response forgeSearchForbidden
*/
type ForgeSearchForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewForgeSearchForbidden creates ForgeSearchForbidden with default headers values
func NewForgeSearchForbidden() *ForgeSearchForbidden {

	return &ForgeSearchForbidden{}
}

// WithPayload adds the payload to the forge search forbidden response
func (o *ForgeSearchForbidden) WithPayload(payload *models.GeneralError) *ForgeSearchForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the forge search forbidden response
func (o *ForgeSearchForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ForgeSearchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ForgeSearchDefault Some error unrelated to the handler

swagger:response forgeSearchDefault
*/
type ForgeSearchDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewForgeSearchDefault creates ForgeSearchDefault with default headers values
func NewForgeSearchDefault(code int) *ForgeSearchDefault {
	if code <= 0 {
		code = 500
	}

	return &ForgeSearchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the forge search default response
func (o *ForgeSearchDefault) WithStatusCode(code int) *ForgeSearchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the forge search default response
func (o *ForgeSearchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the forge search default response
func (o *ForgeSearchDefault) WithPayload(payload *models.GeneralError) *ForgeSearchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the forge search default response
func (o *ForgeSearchDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ForgeSearchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
