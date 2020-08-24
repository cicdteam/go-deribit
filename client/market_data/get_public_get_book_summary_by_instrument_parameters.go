// Code generated by go-swagger; DO NOT EDIT.

package market_data

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

// NewGetPublicGetBookSummaryByInstrumentParams creates a new GetPublicGetBookSummaryByInstrumentParams object
// with the default values initialized.
func NewGetPublicGetBookSummaryByInstrumentParams() *GetPublicGetBookSummaryByInstrumentParams {
	var ()
	return &GetPublicGetBookSummaryByInstrumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicGetBookSummaryByInstrumentParamsWithTimeout creates a new GetPublicGetBookSummaryByInstrumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicGetBookSummaryByInstrumentParamsWithTimeout(timeout time.Duration) *GetPublicGetBookSummaryByInstrumentParams {
	var ()
	return &GetPublicGetBookSummaryByInstrumentParams{

		timeout: timeout,
	}
}

// NewGetPublicGetBookSummaryByInstrumentParamsWithContext creates a new GetPublicGetBookSummaryByInstrumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicGetBookSummaryByInstrumentParamsWithContext(ctx context.Context) *GetPublicGetBookSummaryByInstrumentParams {
	var ()
	return &GetPublicGetBookSummaryByInstrumentParams{

		Context: ctx,
	}
}

// NewGetPublicGetBookSummaryByInstrumentParamsWithHTTPClient creates a new GetPublicGetBookSummaryByInstrumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicGetBookSummaryByInstrumentParamsWithHTTPClient(client *http.Client) *GetPublicGetBookSummaryByInstrumentParams {
	var ()
	return &GetPublicGetBookSummaryByInstrumentParams{
		HTTPClient: client,
	}
}

/*GetPublicGetBookSummaryByInstrumentParams contains all the parameters to send to the API endpoint
for the get public get book summary by instrument operation typically these are written to a http.Request
*/
type GetPublicGetBookSummaryByInstrumentParams struct {

	/*InstrumentName
	  Instrument name

	*/
	InstrumentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public get book summary by instrument params
func (o *GetPublicGetBookSummaryByInstrumentParams) WithTimeout(timeout time.Duration) *GetPublicGetBookSummaryByInstrumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public get book summary by instrument params
func (o *GetPublicGetBookSummaryByInstrumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public get book summary by instrument params
func (o *GetPublicGetBookSummaryByInstrumentParams) WithContext(ctx context.Context) *GetPublicGetBookSummaryByInstrumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public get book summary by instrument params
func (o *GetPublicGetBookSummaryByInstrumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public get book summary by instrument params
func (o *GetPublicGetBookSummaryByInstrumentParams) WithHTTPClient(client *http.Client) *GetPublicGetBookSummaryByInstrumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public get book summary by instrument params
func (o *GetPublicGetBookSummaryByInstrumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstrumentName adds the instrumentName to the get public get book summary by instrument params
func (o *GetPublicGetBookSummaryByInstrumentParams) WithInstrumentName(instrumentName string) *GetPublicGetBookSummaryByInstrumentParams {
	o.SetInstrumentName(instrumentName)
	return o
}

// SetInstrumentName adds the instrumentName to the get public get book summary by instrument params
func (o *GetPublicGetBookSummaryByInstrumentParams) SetInstrumentName(instrumentName string) {
	o.InstrumentName = instrumentName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicGetBookSummaryByInstrumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param instrument_name
	qrInstrumentName := o.InstrumentName
	qInstrumentName := qrInstrumentName
	if qInstrumentName != "" {
		if err := r.SetQueryParam("instrument_name", qInstrumentName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
