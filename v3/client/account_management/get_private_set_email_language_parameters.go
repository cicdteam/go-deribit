// Code generated by go-swagger; DO NOT EDIT.

package account_management

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

// NewGetPrivateSetEmailLanguageParams creates a new GetPrivateSetEmailLanguageParams object
// with the default values initialized.
func NewGetPrivateSetEmailLanguageParams() *GetPrivateSetEmailLanguageParams {
	var ()
	return &GetPrivateSetEmailLanguageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateSetEmailLanguageParamsWithTimeout creates a new GetPrivateSetEmailLanguageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateSetEmailLanguageParamsWithTimeout(timeout time.Duration) *GetPrivateSetEmailLanguageParams {
	var ()
	return &GetPrivateSetEmailLanguageParams{

		timeout: timeout,
	}
}

// NewGetPrivateSetEmailLanguageParamsWithContext creates a new GetPrivateSetEmailLanguageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateSetEmailLanguageParamsWithContext(ctx context.Context) *GetPrivateSetEmailLanguageParams {
	var ()
	return &GetPrivateSetEmailLanguageParams{

		Context: ctx,
	}
}

// NewGetPrivateSetEmailLanguageParamsWithHTTPClient creates a new GetPrivateSetEmailLanguageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateSetEmailLanguageParamsWithHTTPClient(client *http.Client) *GetPrivateSetEmailLanguageParams {
	var ()
	return &GetPrivateSetEmailLanguageParams{
		HTTPClient: client,
	}
}

/*GetPrivateSetEmailLanguageParams contains all the parameters to send to the API endpoint
for the get private set email language operation typically these are written to a http.Request
*/
type GetPrivateSetEmailLanguageParams struct {

	/*Language
	  The abbreviated language name. Valid values include `"en"`, `"ko"`, `"zh"`

	*/
	Language string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private set email language params
func (o *GetPrivateSetEmailLanguageParams) WithTimeout(timeout time.Duration) *GetPrivateSetEmailLanguageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private set email language params
func (o *GetPrivateSetEmailLanguageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private set email language params
func (o *GetPrivateSetEmailLanguageParams) WithContext(ctx context.Context) *GetPrivateSetEmailLanguageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private set email language params
func (o *GetPrivateSetEmailLanguageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private set email language params
func (o *GetPrivateSetEmailLanguageParams) WithHTTPClient(client *http.Client) *GetPrivateSetEmailLanguageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private set email language params
func (o *GetPrivateSetEmailLanguageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguage adds the language to the get private set email language params
func (o *GetPrivateSetEmailLanguageParams) WithLanguage(language string) *GetPrivateSetEmailLanguageParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get private set email language params
func (o *GetPrivateSetEmailLanguageParams) SetLanguage(language string) {
	o.Language = language
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateSetEmailLanguageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param language
	qrLanguage := o.Language
	qLanguage := qrLanguage
	if qLanguage != "" {
		if err := r.SetQueryParam("language", qLanguage); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
