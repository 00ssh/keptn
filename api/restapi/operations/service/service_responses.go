// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/keptn/keptn/api/models"
)

// ServiceCreatedCode is the HTTP code returned for type ServiceCreated
const ServiceCreatedCode int = 201

/*ServiceCreated service onboarded

swagger:response serviceCreated
*/
type ServiceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ChannelInfo `json:"body,omitempty"`
}

// NewServiceCreated creates ServiceCreated with default headers values
func NewServiceCreated() *ServiceCreated {

	return &ServiceCreated{}
}

// WithPayload adds the payload to the service created response
func (o *ServiceCreated) WithPayload(payload *models.ChannelInfo) *ServiceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service created response
func (o *ServiceCreated) SetPayload(payload *models.ChannelInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ServiceDefault error

swagger:response serviceDefault
*/
type ServiceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceDefault creates ServiceDefault with default headers values
func NewServiceDefault(code int) *ServiceDefault {
	if code <= 0 {
		code = 500
	}

	return &ServiceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the service default response
func (o *ServiceDefault) WithStatusCode(code int) *ServiceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the service default response
func (o *ServiceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the service default response
func (o *ServiceDefault) WithPayload(payload *models.Error) *ServiceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service default response
func (o *ServiceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
