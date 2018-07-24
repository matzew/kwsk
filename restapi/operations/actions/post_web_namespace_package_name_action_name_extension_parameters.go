// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostWebNamespacePackageNameActionNameExtensionParams creates a new PostWebNamespacePackageNameActionNameExtensionParams object
// no default values defined in spec.
func NewPostWebNamespacePackageNameActionNameExtensionParams() PostWebNamespacePackageNameActionNameExtensionParams {

	return PostWebNamespacePackageNameActionNameExtensionParams{}
}

// PostWebNamespacePackageNameActionNameExtensionParams contains all the bound params for the post web namespace package name action name extension operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostWebNamespacePackageNameActionNameExtension
type PostWebNamespacePackageNameActionNameExtensionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ActionName string
	/*
	  Required: true
	  In: path
	*/
	Extension string
	/*
	  Required: true
	  In: path
	*/
	Namespace string
	/*
	  Required: true
	  In: path
	*/
	PackageName string
	/*The parameters for the action being invoked
	  In: body
	*/
	Payload interface{}
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostWebNamespacePackageNameActionNameExtensionParams() beforehand.
func (o *PostWebNamespacePackageNameActionNameExtensionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rActionName, rhkActionName, _ := route.Params.GetOK("actionName")
	if err := o.bindActionName(rActionName, rhkActionName, route.Formats); err != nil {
		res = append(res, err)
	}

	rExtension, rhkExtension, _ := route.Params.GetOK("extension")
	if err := o.bindExtension(rExtension, rhkExtension, route.Formats); err != nil {
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

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body interface{}
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("payload", "body", "", err))
		} else {

			// no validation on generic interface
			o.Payload = body
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWebNamespacePackageNameActionNameExtensionParams) bindActionName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ActionName = raw

	return nil
}

func (o *PostWebNamespacePackageNameActionNameExtensionParams) bindExtension(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Extension = raw

	return nil
}

func (o *PostWebNamespacePackageNameActionNameExtensionParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	return nil
}

func (o *PostWebNamespacePackageNameActionNameExtensionParams) bindPackageName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PackageName = raw

	return nil
}
