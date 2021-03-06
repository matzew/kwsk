// Code generated by go-swagger; DO NOT EDIT.

package activations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNamespacesNamespaceActivationsActivationidResultParams creates a new GetNamespacesNamespaceActivationsActivationidResultParams object
// no default values defined in spec.
func NewGetNamespacesNamespaceActivationsActivationidResultParams() GetNamespacesNamespaceActivationsActivationidResultParams {

	return GetNamespacesNamespaceActivationsActivationidResultParams{}
}

// GetNamespacesNamespaceActivationsActivationidResultParams contains all the bound params for the get namespaces namespace activations activationid result operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetNamespacesNamespaceActivationsActivationidResult
type GetNamespacesNamespaceActivationsActivationidResultParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Name of activation to fetch
	  Required: true
	  In: path
	*/
	Activationid string
	/*The entity namespace
	  Required: true
	  In: path
	*/
	Namespace string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetNamespacesNamespaceActivationsActivationidResultParams() beforehand.
func (o *GetNamespacesNamespaceActivationsActivationidResultParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rActivationid, rhkActivationid, _ := route.Params.GetOK("activationid")
	if err := o.bindActivationid(rActivationid, rhkActivationid, route.Formats); err != nil {
		res = append(res, err)
	}

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNamespacesNamespaceActivationsActivationidResultParams) bindActivationid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Activationid = raw

	return nil
}

func (o *GetNamespacesNamespaceActivationsActivationidResultParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	return nil
}
