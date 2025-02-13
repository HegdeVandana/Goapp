// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/models"
)

// FindonetodosOKCode is the HTTP code returned for type FindonetodosOK
const FindonetodosOKCode int = 200

/*FindonetodosOK OK

swagger:response findonetodosOK
*/
type FindonetodosOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewFindonetodosOK creates FindonetodosOK with default headers values
func NewFindonetodosOK() *FindonetodosOK {

	return &FindonetodosOK{}
}

// WithPayload adds the payload to the findonetodos o k response
func (o *FindonetodosOK) WithPayload(payload *models.Item) *FindonetodosOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the findonetodos o k response
func (o *FindonetodosOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindonetodosOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindonetodosDefault error

swagger:response findonetodosDefault
*/
type FindonetodosDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindonetodosDefault creates FindonetodosDefault with default headers values
func NewFindonetodosDefault(code int) *FindonetodosDefault {
	if code <= 0 {
		code = 500
	}

	return &FindonetodosDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the findonetodos default response
func (o *FindonetodosDefault) WithStatusCode(code int) *FindonetodosDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the findonetodos default response
func (o *FindonetodosDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the findonetodos default response
func (o *FindonetodosDefault) WithPayload(payload *models.Error) *FindonetodosDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the findonetodos default response
func (o *FindonetodosDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindonetodosDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
