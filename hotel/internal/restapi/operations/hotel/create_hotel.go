// Code generated by go-swagger; DO NOT EDIT.

package hotel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
)

// CreateHotelHandlerFunc turns a function with the right signature into a create hotel handler
type CreateHotelHandlerFunc func(CreateHotelParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateHotelHandlerFunc) Handle(params CreateHotelParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// CreateHotelHandler interface for that can handle valid create hotel params
type CreateHotelHandler interface {
	Handle(CreateHotelParams, *models.User) middleware.Responder
}

// NewCreateHotel creates a new http.Handler for the create hotel operation
func NewCreateHotel(ctx *middleware.Context, handler CreateHotelHandler) *CreateHotel {
	return &CreateHotel{Context: ctx, Handler: handler}
}

/*
	CreateHotel swagger:route POST /hotel hotel createHotel

Create hotel
*/
type CreateHotel struct {
	Context *middleware.Context
	Handler CreateHotelHandler
}

func (o *CreateHotel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateHotelParams()
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

// CreateHotelOKBody create hotel o k body
//
// swagger:model CreateHotelOKBody
type CreateHotelOKBody struct {

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this create hotel o k body
func (o *CreateHotelOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create hotel o k body based on context it is used
func (o *CreateHotelOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateHotelOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateHotelOKBody) UnmarshalBinary(b []byte) error {
	var res CreateHotelOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
