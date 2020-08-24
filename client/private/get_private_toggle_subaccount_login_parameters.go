// Code generated by go-swagger; DO NOT EDIT.

package private

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
	"github.com/go-openapi/swag"
)

// NewGetPrivateToggleSubaccountLoginParams creates a new GetPrivateToggleSubaccountLoginParams object
// with the default values initialized.
func NewGetPrivateToggleSubaccountLoginParams() *GetPrivateToggleSubaccountLoginParams {
	var ()
	return &GetPrivateToggleSubaccountLoginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateToggleSubaccountLoginParamsWithTimeout creates a new GetPrivateToggleSubaccountLoginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateToggleSubaccountLoginParamsWithTimeout(timeout time.Duration) *GetPrivateToggleSubaccountLoginParams {
	var ()
	return &GetPrivateToggleSubaccountLoginParams{

		timeout: timeout,
	}
}

// NewGetPrivateToggleSubaccountLoginParamsWithContext creates a new GetPrivateToggleSubaccountLoginParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateToggleSubaccountLoginParamsWithContext(ctx context.Context) *GetPrivateToggleSubaccountLoginParams {
	var ()
	return &GetPrivateToggleSubaccountLoginParams{

		Context: ctx,
	}
}

// NewGetPrivateToggleSubaccountLoginParamsWithHTTPClient creates a new GetPrivateToggleSubaccountLoginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateToggleSubaccountLoginParamsWithHTTPClient(client *http.Client) *GetPrivateToggleSubaccountLoginParams {
	var ()
	return &GetPrivateToggleSubaccountLoginParams{
		HTTPClient: client,
	}
}

/*GetPrivateToggleSubaccountLoginParams contains all the parameters to send to the API endpoint
for the get private toggle subaccount login operation typically these are written to a http.Request
*/
type GetPrivateToggleSubaccountLoginParams struct {

	/*Sid
	  The user id for the subaccount

	*/
	Sid int64
	/*State
	  enable or disable login.

	*/
	State string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private toggle subaccount login params
func (o *GetPrivateToggleSubaccountLoginParams) WithTimeout(timeout time.Duration) *GetPrivateToggleSubaccountLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private toggle subaccount login params
func (o *GetPrivateToggleSubaccountLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private toggle subaccount login params
func (o *GetPrivateToggleSubaccountLoginParams) WithContext(ctx context.Context) *GetPrivateToggleSubaccountLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private toggle subaccount login params
func (o *GetPrivateToggleSubaccountLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private toggle subaccount login params
func (o *GetPrivateToggleSubaccountLoginParams) WithHTTPClient(client *http.Client) *GetPrivateToggleSubaccountLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private toggle subaccount login params
func (o *GetPrivateToggleSubaccountLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSid adds the sid to the get private toggle subaccount login params
func (o *GetPrivateToggleSubaccountLoginParams) WithSid(sid int64) *GetPrivateToggleSubaccountLoginParams {
	o.SetSid(sid)
	return o
}

// SetSid adds the sid to the get private toggle subaccount login params
func (o *GetPrivateToggleSubaccountLoginParams) SetSid(sid int64) {
	o.Sid = sid
}

// WithState adds the state to the get private toggle subaccount login params
func (o *GetPrivateToggleSubaccountLoginParams) WithState(state string) *GetPrivateToggleSubaccountLoginParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get private toggle subaccount login params
func (o *GetPrivateToggleSubaccountLoginParams) SetState(state string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateToggleSubaccountLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param sid
	qrSid := o.Sid
	qSid := swag.FormatInt64(qrSid)
	if qSid != "" {
		if err := r.SetQueryParam("sid", qSid); err != nil {
			return err
		}
	}

	// query param state
	qrState := o.State
	qState := qrState
	if qState != "" {
		if err := r.SetQueryParam("state", qState); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}