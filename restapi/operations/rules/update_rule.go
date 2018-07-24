// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/projectodd/kwsk/models"
)

// UpdateRuleHandlerFunc turns a function with the right signature into a update rule handler
type UpdateRuleHandlerFunc func(UpdateRuleParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateRuleHandlerFunc) Handle(params UpdateRuleParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateRuleHandler interface for that can handle valid update rule params
type UpdateRuleHandler interface {
	Handle(UpdateRuleParams, *models.Principal) middleware.Responder
}

// NewUpdateRule creates a new http.Handler for the update rule operation
func NewUpdateRule(ctx *middleware.Context, handler UpdateRuleHandler) *UpdateRule {
	return &UpdateRule{Context: ctx, Handler: handler}
}

/*UpdateRule swagger:route PUT /namespaces/{namespace}/rules/{ruleName} Rules updateRule

Update a rule

Update a rule

*/
type UpdateRule struct {
	Context *middleware.Context
	Handler UpdateRuleHandler
}

func (o *UpdateRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateRuleParams()

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
