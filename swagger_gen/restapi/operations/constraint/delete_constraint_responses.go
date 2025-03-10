// Code generated by go-swagger; DO NOT EDIT.

package constraint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Allen-Career-Institute/flagr/swagger_gen/models"
)

// DeleteConstraintOKCode is the HTTP code returned for type DeleteConstraintOK
const DeleteConstraintOKCode int = 200

/*
DeleteConstraintOK deleted

swagger:response deleteConstraintOK
*/
type DeleteConstraintOK struct {
}

// NewDeleteConstraintOK creates DeleteConstraintOK with default headers values
func NewDeleteConstraintOK() *DeleteConstraintOK {

	return &DeleteConstraintOK{}
}

// WriteResponse to the client
func (o *DeleteConstraintOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*
DeleteConstraintDefault generic error response

swagger:response deleteConstraintDefault
*/
type DeleteConstraintDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConstraintDefault creates DeleteConstraintDefault with default headers values
func NewDeleteConstraintDefault(code int) *DeleteConstraintDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteConstraintDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete constraint default response
func (o *DeleteConstraintDefault) WithStatusCode(code int) *DeleteConstraintDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete constraint default response
func (o *DeleteConstraintDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete constraint default response
func (o *DeleteConstraintDefault) WithPayload(payload *models.Error) *DeleteConstraintDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete constraint default response
func (o *DeleteConstraintDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConstraintDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
