// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/projectodd/kwsk/models"
)

// DeleteActionInPackageHandlerFunc turns a function with the right signature into a delete action in package handler
type DeleteActionInPackageHandlerFunc func(DeleteActionInPackageParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteActionInPackageHandlerFunc) Handle(params DeleteActionInPackageParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteActionInPackageHandler interface for that can handle valid delete action in package params
type DeleteActionInPackageHandler interface {
	Handle(DeleteActionInPackageParams, *models.Principal) middleware.Responder
}

// NewDeleteActionInPackage creates a new http.Handler for the delete action in package operation
func NewDeleteActionInPackage(ctx *middleware.Context, handler DeleteActionInPackageHandler) *DeleteActionInPackage {
	return &DeleteActionInPackage{Context: ctx, Handler: handler}
}

/*DeleteActionInPackage swagger:route DELETE /namespaces/{namespace}/actions/{packageName}/{actionName} Actions deleteActionInPackage

Delete an action

Delete an action

*/
type DeleteActionInPackage struct {
	Context *middleware.Context
	Handler DeleteActionInPackageHandler
}

func (o *DeleteActionInPackage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteActionInPackageParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
