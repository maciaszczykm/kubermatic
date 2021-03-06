// Code generated by go-swagger; DO NOT EDIT.

package whitelistedregistry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListWhitelistedRegistriesParams creates a new ListWhitelistedRegistriesParams object
// with the default values initialized.
func NewListWhitelistedRegistriesParams() *ListWhitelistedRegistriesParams {

	return &ListWhitelistedRegistriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListWhitelistedRegistriesParamsWithTimeout creates a new ListWhitelistedRegistriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListWhitelistedRegistriesParamsWithTimeout(timeout time.Duration) *ListWhitelistedRegistriesParams {

	return &ListWhitelistedRegistriesParams{

		timeout: timeout,
	}
}

// NewListWhitelistedRegistriesParamsWithContext creates a new ListWhitelistedRegistriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListWhitelistedRegistriesParamsWithContext(ctx context.Context) *ListWhitelistedRegistriesParams {

	return &ListWhitelistedRegistriesParams{

		Context: ctx,
	}
}

// NewListWhitelistedRegistriesParamsWithHTTPClient creates a new ListWhitelistedRegistriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListWhitelistedRegistriesParamsWithHTTPClient(client *http.Client) *ListWhitelistedRegistriesParams {

	return &ListWhitelistedRegistriesParams{
		HTTPClient: client,
	}
}

/*ListWhitelistedRegistriesParams contains all the parameters to send to the API endpoint
for the list whitelisted registries operation typically these are written to a http.Request
*/
type ListWhitelistedRegistriesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list whitelisted registries params
func (o *ListWhitelistedRegistriesParams) WithTimeout(timeout time.Duration) *ListWhitelistedRegistriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list whitelisted registries params
func (o *ListWhitelistedRegistriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list whitelisted registries params
func (o *ListWhitelistedRegistriesParams) WithContext(ctx context.Context) *ListWhitelistedRegistriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list whitelisted registries params
func (o *ListWhitelistedRegistriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list whitelisted registries params
func (o *ListWhitelistedRegistriesParams) WithHTTPClient(client *http.Client) *ListWhitelistedRegistriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list whitelisted registries params
func (o *ListWhitelistedRegistriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListWhitelistedRegistriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
