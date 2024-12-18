// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
)

// UpdateBookingHandlerFunc turns a function with the right signature into a update booking handler
type UpdateBookingHandlerFunc func(UpdateBookingParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateBookingHandlerFunc) Handle(params UpdateBookingParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// UpdateBookingHandler interface for that can handle valid update booking params
type UpdateBookingHandler interface {
	Handle(UpdateBookingParams, *models.User) middleware.Responder
}

// NewUpdateBooking creates a new http.Handler for the update booking operation
func NewUpdateBooking(ctx *middleware.Context, handler UpdateBookingHandler) *UpdateBooking {
	return &UpdateBooking{Context: ctx, Handler: handler}
}

/*
	UpdateBooking swagger:route PUT /booking/{booking_id} customer hotelier updateBooking

Update booking
*/
type UpdateBooking struct {
	Context *middleware.Context
	Handler UpdateBookingHandler
}

func (o *UpdateBooking) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateBookingParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
