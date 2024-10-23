// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Allen-Career-Institute/flagr/swagger_gen/models"
)

// DeleteSegmentOKCode is the HTTP code returned for type DeleteSegmentOK
const DeleteSegmentOKCode int = 200

/*
DeleteSegmentOK deleted

swagger:response deleteSegmentOK
*/
type DeleteSegmentOK struct {
}

// NewDeleteSegmentOK creates DeleteSegmentOK with default headers values
func NewDeleteSegmentOK() *DeleteSegmentOK {

	return &DeleteSegmentOK{}
}

// WriteResponse to the client
func (o *DeleteSegmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*
DeleteSegmentDefault generic error response

swagger:response deleteSegmentDefault
*/
type DeleteSegmentDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSegmentDefault creates DeleteSegmentDefault with default headers values
func NewDeleteSegmentDefault(code int) *DeleteSegmentDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSegmentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete segment default response
func (o *DeleteSegmentDefault) WithStatusCode(code int) *DeleteSegmentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete segment default response
func (o *DeleteSegmentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete segment default response
func (o *DeleteSegmentDefault) WithPayload(payload *models.Error) *DeleteSegmentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete segment default response
func (o *DeleteSegmentDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSegmentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
