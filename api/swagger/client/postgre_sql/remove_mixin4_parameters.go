// Code generated by go-swagger; DO NOT EDIT.

package postgre_sql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRemoveMixin4Params creates a new RemoveMixin4Params object
// with the default values initialized.
func NewRemoveMixin4Params() *RemoveMixin4Params {
	var ()
	return &RemoveMixin4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveMixin4ParamsWithTimeout creates a new RemoveMixin4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveMixin4ParamsWithTimeout(timeout time.Duration) *RemoveMixin4Params {
	var ()
	return &RemoveMixin4Params{

		timeout: timeout,
	}
}

// NewRemoveMixin4ParamsWithContext creates a new RemoveMixin4Params object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveMixin4ParamsWithContext(ctx context.Context) *RemoveMixin4Params {
	var ()
	return &RemoveMixin4Params{

		Context: ctx,
	}
}

// NewRemoveMixin4ParamsWithHTTPClient creates a new RemoveMixin4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveMixin4ParamsWithHTTPClient(client *http.Client) *RemoveMixin4Params {
	var ()
	return &RemoveMixin4Params{
		HTTPClient: client,
	}
}

/*RemoveMixin4Params contains all the parameters to send to the API endpoint
for the remove mixin4 operation typically these are written to a http.Request
*/
type RemoveMixin4Params struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove mixin4 params
func (o *RemoveMixin4Params) WithTimeout(timeout time.Duration) *RemoveMixin4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove mixin4 params
func (o *RemoveMixin4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove mixin4 params
func (o *RemoveMixin4Params) WithContext(ctx context.Context) *RemoveMixin4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove mixin4 params
func (o *RemoveMixin4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove mixin4 params
func (o *RemoveMixin4Params) WithHTTPClient(client *http.Client) *RemoveMixin4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove mixin4 params
func (o *RemoveMixin4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove mixin4 params
func (o *RemoveMixin4Params) WithID(id int32) *RemoveMixin4Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove mixin4 params
func (o *RemoveMixin4Params) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveMixin4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
