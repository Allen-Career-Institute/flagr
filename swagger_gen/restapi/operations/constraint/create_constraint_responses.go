// Code generated by go-swagger; DO NOT EDIT.

package constraint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Allen-Career-Institute/flagr/swagger_gen/models"
)

// CreateConstraintOKCode is the HTTP code returned for type CreateConstraintOK
const CreateConstraintOKCode int = 200

/*
CreateConstraintOK the constraint created

swagger:response createConstraintOK
*/
type CreateConstraintOK struct {

	/*
	  In: Body
	*/
	Payload *models.Constraint `json:"body,omitempty"`
}

// NewCreateConstraintOK creates CreateConstraintOK with default headers values
func NewCreateConstraintOK() *CreateConstraintOK {

	return &CreateConstraintOK{}
}

// WithPayload adds the payload to the create constraint o k response
func (o *CreateConstraintOK) WithPayload(payload *models.Constraint) *CreateConstraintOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create constraint o k response
func (o *CreateConstraintOK) SetPayload(payload *models.Constraint) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConstraintOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
CreateConstraintDefault generic error response

swagger:response createConstraintDefault
*/
type CreateConstraintDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateConstraintDefault creates CreateConstraintDefault with default headers values
func NewCreateConstraintDefault(code int) *CreateConstraintDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateConstraintDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create constraint default response
func (o *CreateConstraintDefault) WithStatusCode(code int) *CreateConstraintDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create constraint default response
func (o *CreateConstraintDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create constraint default response
func (o *CreateConstraintDefault) WithPayload(payload *models.Error) *CreateConstraintDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create constraint default response
func (o *CreateConstraintDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateConstraintDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
