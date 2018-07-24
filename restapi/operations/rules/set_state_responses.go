// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/projectodd/kwsk/models"
)

// SetStateOKCode is the HTTP code returned for type SetStateOK
const SetStateOKCode int = 200

/*SetStateOK Rule has been enabled or disabled

swagger:response setStateOK
*/
type SetStateOK struct {
}

// NewSetStateOK creates SetStateOK with default headers values
func NewSetStateOK() *SetStateOK {

	return &SetStateOK{}
}

// WriteResponse to the client
func (o *SetStateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SetStateAcceptedCode is the HTTP code returned for type SetStateAccepted
const SetStateAcceptedCode int = 202

/*SetStateAccepted Rule has been enabled or disabled

swagger:response setStateAccepted
*/
type SetStateAccepted struct {
}

// NewSetStateAccepted creates SetStateAccepted with default headers values
func NewSetStateAccepted() *SetStateAccepted {

	return &SetStateAccepted{}
}

// WriteResponse to the client
func (o *SetStateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// SetStateBadRequestCode is the HTTP code returned for type SetStateBadRequest
const SetStateBadRequestCode int = 400

/*SetStateBadRequest Bad request

swagger:response setStateBadRequest
*/
type SetStateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewSetStateBadRequest creates SetStateBadRequest with default headers values
func NewSetStateBadRequest() *SetStateBadRequest {

	return &SetStateBadRequest{}
}

// WithPayload adds the payload to the set state bad request response
func (o *SetStateBadRequest) WithPayload(payload *models.ErrorMessage) *SetStateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set state bad request response
func (o *SetStateBadRequest) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetStateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetStateUnauthorizedCode is the HTTP code returned for type SetStateUnauthorized
const SetStateUnauthorizedCode int = 401

/*SetStateUnauthorized Unauthorized request

swagger:response setStateUnauthorized
*/
type SetStateUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewSetStateUnauthorized creates SetStateUnauthorized with default headers values
func NewSetStateUnauthorized() *SetStateUnauthorized {

	return &SetStateUnauthorized{}
}

// WithPayload adds the payload to the set state unauthorized response
func (o *SetStateUnauthorized) WithPayload(payload *models.ErrorMessage) *SetStateUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set state unauthorized response
func (o *SetStateUnauthorized) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetStateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetStateNotFoundCode is the HTTP code returned for type SetStateNotFound
const SetStateNotFoundCode int = 404

/*SetStateNotFound Item not found

swagger:response setStateNotFound
*/
type SetStateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewSetStateNotFound creates SetStateNotFound with default headers values
func NewSetStateNotFound() *SetStateNotFound {

	return &SetStateNotFound{}
}

// WithPayload adds the payload to the set state not found response
func (o *SetStateNotFound) WithPayload(payload *models.ErrorMessage) *SetStateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set state not found response
func (o *SetStateNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetStateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SetStateInternalServerErrorCode is the HTTP code returned for type SetStateInternalServerError
const SetStateInternalServerErrorCode int = 500

/*SetStateInternalServerError Server error

swagger:response setStateInternalServerError
*/
type SetStateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewSetStateInternalServerError creates SetStateInternalServerError with default headers values
func NewSetStateInternalServerError() *SetStateInternalServerError {

	return &SetStateInternalServerError{}
}

// WithPayload adds the payload to the set state internal server error response
func (o *SetStateInternalServerError) WithPayload(payload *models.ErrorMessage) *SetStateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set state internal server error response
func (o *SetStateInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetStateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
