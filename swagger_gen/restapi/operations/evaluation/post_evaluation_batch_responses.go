// Code generated by go-swagger; DO NOT EDIT.

package evaluation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Allen-Career-Institute/flagr/swagger_gen/models"
)

// PostEvaluationBatchOKCode is the HTTP code returned for type PostEvaluationBatchOK
const PostEvaluationBatchOKCode int = 200

/*
PostEvaluationBatchOK evaluation batch result

swagger:response postEvaluationBatchOK
*/
type PostEvaluationBatchOK struct {

	/*
	  In: Body
	*/
	Payload *models.EvaluationBatchResponse `json:"body,omitempty"`
}

// NewPostEvaluationBatchOK creates PostEvaluationBatchOK with default headers values
func NewPostEvaluationBatchOK() *PostEvaluationBatchOK {

	return &PostEvaluationBatchOK{}
}

// WithPayload adds the payload to the post evaluation batch o k response
func (o *PostEvaluationBatchOK) WithPayload(payload *models.EvaluationBatchResponse) *PostEvaluationBatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post evaluation batch o k response
func (o *PostEvaluationBatchOK) SetPayload(payload *models.EvaluationBatchResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostEvaluationBatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostEvaluationBatchDefault generic error response

swagger:response postEvaluationBatchDefault
*/
type PostEvaluationBatchDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostEvaluationBatchDefault creates PostEvaluationBatchDefault with default headers values
func NewPostEvaluationBatchDefault(code int) *PostEvaluationBatchDefault {
	if code <= 0 {
		code = 500
	}

	return &PostEvaluationBatchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post evaluation batch default response
func (o *PostEvaluationBatchDefault) WithStatusCode(code int) *PostEvaluationBatchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post evaluation batch default response
func (o *PostEvaluationBatchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post evaluation batch default response
func (o *PostEvaluationBatchDefault) WithPayload(payload *models.Error) *PostEvaluationBatchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post evaluation batch default response
func (o *PostEvaluationBatchDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostEvaluationBatchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
