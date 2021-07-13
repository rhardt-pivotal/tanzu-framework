// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GenerateTKGConfigForAWSHandlerFunc turns a function with the right signature into a generate t k g config for a w s handler
type GenerateTKGConfigForAWSHandlerFunc func(GenerateTKGConfigForAWSParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GenerateTKGConfigForAWSHandlerFunc) Handle(params GenerateTKGConfigForAWSParams) middleware.Responder {
	return fn(params)
}

// GenerateTKGConfigForAWSHandler interface for that can handle valid generate t k g config for a w s params
type GenerateTKGConfigForAWSHandler interface {
	Handle(GenerateTKGConfigForAWSParams) middleware.Responder
}

// NewGenerateTKGConfigForAWS creates a new http.Handler for the generate t k g config for a w s operation
func NewGenerateTKGConfigForAWS(ctx *middleware.Context, handler GenerateTKGConfigForAWSHandler) *GenerateTKGConfigForAWS {
	return &GenerateTKGConfigForAWS{Context: ctx, Handler: handler}
}

/*GenerateTKGConfigForAWS swagger:route POST /api/providers/aws/generate aws generateTKGConfigForAWS

Generate TKG configuration file for AWS"

*/
type GenerateTKGConfigForAWS struct {
	Context *middleware.Context
	Handler GenerateTKGConfigForAWSHandler
}

func (o *GenerateTKGConfigForAWS) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGenerateTKGConfigForAWSParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}