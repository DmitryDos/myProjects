// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
)

// NewUpdateBookingParams creates a new UpdateBookingParams object
//
// There are no default values defined in the spec.
func NewUpdateBookingParams() UpdateBookingParams {

	return UpdateBookingParams{}
}

// UpdateBookingParams contains all the bound params for the update booking operation
// typically these are obtained from a http.Request
//
// swagger:parameters update_booking
type UpdateBookingParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of booking to change
	  Required: true
	  In: path
	*/
	BookingID int64
	/*
	  Required: true
	  In: body
	*/
	Object *models.Booking
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateBookingParams() beforehand.
func (o *UpdateBookingParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBookingID, rhkBookingID, _ := route.Params.GetOK("booking_id")
	if err := o.bindBookingID(rBookingID, rhkBookingID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Booking
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("object", "body", ""))
			} else {
				res = append(res, errors.NewParseError("object", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Object = &body
			}
		}
	} else {
		res = append(res, errors.Required("object", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBookingID binds and validates parameter BookingID from path.
func (o *UpdateBookingParams) bindBookingID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("booking_id", "path", "int64", raw)
	}
	o.BookingID = value

	return nil
}
