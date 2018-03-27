// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSearchableHandlerFunc turns a function with the right signature into a get searchable handler
type GetSearchableHandlerFunc func(GetSearchableParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSearchableHandlerFunc) Handle(params GetSearchableParams) middleware.Responder {
	return fn(params)
}

// GetSearchableHandler interface for that can handle valid get searchable params
type GetSearchableHandler interface {
	Handle(GetSearchableParams) middleware.Responder
}

// NewGetSearchable creates a new http.Handler for the get searchable operation
func NewGetSearchable(ctx *middleware.Context, handler GetSearchableHandler) *GetSearchable {
	return &GetSearchable{Context: ctx, Handler: handler}
}

/*GetSearchable swagger:route GET /searchable getSearchable

gets columns

gets searchable

*/
type GetSearchable struct {
	Context *middleware.Context
	Handler GetSearchableHandler
}

func (o *GetSearchable) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSearchableParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
