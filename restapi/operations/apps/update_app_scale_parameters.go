package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateAppScaleParams creates a new UpdateAppScaleParams object
// with the default values initialized.
func NewUpdateAppScaleParams() UpdateAppScaleParams {
	var ()
	return UpdateAppScaleParams{}
}

// UpdateAppScaleParams contains all the bound params for the update app scale operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateAppScale
type UpdateAppScaleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*App name
	  Required: true
	  Pattern: ^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	  In: path
	*/
	AppName string
	/*
	  In: body
	*/
	Body UpdateAppScaleBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateAppScaleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rAppName, rhkAppName, _ := route.Params.GetOK("app_name")
	if err := o.bindAppName(rAppName, rhkAppName, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body UpdateAppScaleBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {

			if len(res) == 0 {
				o.Body = body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAppScaleParams) bindAppName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.AppName = raw

	if err := o.validateAppName(formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateAppScaleParams) validateAppName(formats strfmt.Registry) error {

	if err := validate.Pattern("app_name", "path", string(o.AppName), `^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`); err != nil {
		return err
	}

	return nil
}
