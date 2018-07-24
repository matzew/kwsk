// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/projectodd/kwsk/models"
)

// DeleteActionInPackageOKCode is the HTTP code returned for type DeleteActionInPackageOK
const DeleteActionInPackageOKCode int = 200

/*DeleteActionInPackageOK Deleted Item

swagger:response deleteActionInPackageOK
*/
type DeleteActionInPackageOK struct {
}

// NewDeleteActionInPackageOK creates DeleteActionInPackageOK with default headers values
func NewDeleteActionInPackageOK() *DeleteActionInPackageOK {

	return &DeleteActionInPackageOK{}
}

// WriteResponse to the client
func (o *DeleteActionInPackageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteActionInPackageBadRequestCode is the HTTP code returned for type DeleteActionInPackageBadRequest
const DeleteActionInPackageBadRequestCode int = 400

/*DeleteActionInPackageBadRequest Bad request

swagger:response deleteActionInPackageBadRequest
*/
type DeleteActionInPackageBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionInPackageBadRequest creates DeleteActionInPackageBadRequest with default headers values
func NewDeleteActionInPackageBadRequest() *DeleteActionInPackageBadRequest {

	return &DeleteActionInPackageBadRequest{}
}

// WithPayload adds the payload to the delete action in package bad request response
func (o *DeleteActionInPackageBadRequest) WithPayload(payload *models.ErrorMessage) *DeleteActionInPackageBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action in package bad request response
func (o *DeleteActionInPackageBadRequest) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionInPackageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteActionInPackageUnauthorizedCode is the HTTP code returned for type DeleteActionInPackageUnauthorized
const DeleteActionInPackageUnauthorizedCode int = 401

/*DeleteActionInPackageUnauthorized Unauthorized request

swagger:response deleteActionInPackageUnauthorized
*/
type DeleteActionInPackageUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionInPackageUnauthorized creates DeleteActionInPackageUnauthorized with default headers values
func NewDeleteActionInPackageUnauthorized() *DeleteActionInPackageUnauthorized {

	return &DeleteActionInPackageUnauthorized{}
}

// WithPayload adds the payload to the delete action in package unauthorized response
func (o *DeleteActionInPackageUnauthorized) WithPayload(payload *models.ErrorMessage) *DeleteActionInPackageUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action in package unauthorized response
func (o *DeleteActionInPackageUnauthorized) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionInPackageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteActionInPackageForbiddenCode is the HTTP code returned for type DeleteActionInPackageForbidden
const DeleteActionInPackageForbiddenCode int = 403

/*DeleteActionInPackageForbidden Unauthorized request

swagger:response deleteActionInPackageForbidden
*/
type DeleteActionInPackageForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionInPackageForbidden creates DeleteActionInPackageForbidden with default headers values
func NewDeleteActionInPackageForbidden() *DeleteActionInPackageForbidden {

	return &DeleteActionInPackageForbidden{}
}

// WithPayload adds the payload to the delete action in package forbidden response
func (o *DeleteActionInPackageForbidden) WithPayload(payload *models.ErrorMessage) *DeleteActionInPackageForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action in package forbidden response
func (o *DeleteActionInPackageForbidden) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionInPackageForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteActionInPackageNotFoundCode is the HTTP code returned for type DeleteActionInPackageNotFound
const DeleteActionInPackageNotFoundCode int = 404

/*DeleteActionInPackageNotFound Item not found

swagger:response deleteActionInPackageNotFound
*/
type DeleteActionInPackageNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionInPackageNotFound creates DeleteActionInPackageNotFound with default headers values
func NewDeleteActionInPackageNotFound() *DeleteActionInPackageNotFound {

	return &DeleteActionInPackageNotFound{}
}

// WithPayload adds the payload to the delete action in package not found response
func (o *DeleteActionInPackageNotFound) WithPayload(payload *models.ErrorMessage) *DeleteActionInPackageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action in package not found response
func (o *DeleteActionInPackageNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionInPackageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteActionInPackageConflictCode is the HTTP code returned for type DeleteActionInPackageConflict
const DeleteActionInPackageConflictCode int = 409

/*DeleteActionInPackageConflict Conflicting item already exists

swagger:response deleteActionInPackageConflict
*/
type DeleteActionInPackageConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionInPackageConflict creates DeleteActionInPackageConflict with default headers values
func NewDeleteActionInPackageConflict() *DeleteActionInPackageConflict {

	return &DeleteActionInPackageConflict{}
}

// WithPayload adds the payload to the delete action in package conflict response
func (o *DeleteActionInPackageConflict) WithPayload(payload *models.ErrorMessage) *DeleteActionInPackageConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action in package conflict response
func (o *DeleteActionInPackageConflict) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionInPackageConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteActionInPackageInternalServerErrorCode is the HTTP code returned for type DeleteActionInPackageInternalServerError
const DeleteActionInPackageInternalServerErrorCode int = 500

/*DeleteActionInPackageInternalServerError Server error

swagger:response deleteActionInPackageInternalServerError
*/
type DeleteActionInPackageInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionInPackageInternalServerError creates DeleteActionInPackageInternalServerError with default headers values
func NewDeleteActionInPackageInternalServerError() *DeleteActionInPackageInternalServerError {

	return &DeleteActionInPackageInternalServerError{}
}

// WithPayload adds the payload to the delete action in package internal server error response
func (o *DeleteActionInPackageInternalServerError) WithPayload(payload *models.ErrorMessage) *DeleteActionInPackageInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action in package internal server error response
func (o *DeleteActionInPackageInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionInPackageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
