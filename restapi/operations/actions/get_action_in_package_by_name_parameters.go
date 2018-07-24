// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetActionInPackageByNameParams creates a new GetActionInPackageByNameParams object
// no default values defined in spec.
func NewGetActionInPackageByNameParams() GetActionInPackageByNameParams {

	return GetActionInPackageByNameParams{}
}

// GetActionInPackageByNameParams contains all the bound params for the get action in package by name operation
// typically these are obtained from a http.Request
//
// swagger:parameters getActionInPackageByName
type GetActionInPackageByNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Name of action to fetch
	  Required: true
	  In: path
	*/
	ActionName string
	/*Include action code in the result
	  In: query
	*/
	Code *bool
	/*The entity namespace
	  Required: true
	  In: path
	*/
	Namespace string
	/*Name of package that contains action
	  Required: true
	  In: path
	*/
	PackageName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetActionInPackageByNameParams() beforehand.
func (o *GetActionInPackageByNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rActionName, rhkActionName, _ := route.Params.GetOK("actionName")
	if err := o.bindActionName(rActionName, rhkActionName, route.Formats); err != nil {
		res = append(res, err)
	}

	qCode, qhkCode, _ := qs.GetOK("code")
	if err := o.bindCode(qCode, qhkCode, route.Formats); err != nil {
		res = append(res, err)
	}

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	rPackageName, rhkPackageName, _ := route.Params.GetOK("packageName")
	if err := o.bindPackageName(rPackageName, rhkPackageName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetActionInPackageByNameParams) bindActionName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ActionName = raw

	return nil
}

func (o *GetActionInPackageByNameParams) bindCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("code", "query", "bool", raw)
	}
	o.Code = &value

	return nil
}

func (o *GetActionInPackageByNameParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	return nil
}

func (o *GetActionInPackageByNameParams) bindPackageName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PackageName = raw

	return nil
}
