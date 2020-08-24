// Code generated by go-swagger; DO NOT EDIT.

package trading

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

// NewGetPrivateGetMarginsParams creates a new GetPrivateGetMarginsParams object
// with the default values initialized.
func NewGetPrivateGetMarginsParams() *GetPrivateGetMarginsParams {
	var ()
	return &GetPrivateGetMarginsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateGetMarginsParamsWithTimeout creates a new GetPrivateGetMarginsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateGetMarginsParamsWithTimeout(timeout time.Duration) *GetPrivateGetMarginsParams {
	var ()
	return &GetPrivateGetMarginsParams{

		timeout: timeout,
	}
}

// NewGetPrivateGetMarginsParamsWithContext creates a new GetPrivateGetMarginsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateGetMarginsParamsWithContext(ctx context.Context) *GetPrivateGetMarginsParams {
	var ()
	return &GetPrivateGetMarginsParams{

		Context: ctx,
	}
}

// NewGetPrivateGetMarginsParamsWithHTTPClient creates a new GetPrivateGetMarginsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateGetMarginsParamsWithHTTPClient(client *http.Client) *GetPrivateGetMarginsParams {
	var ()
	return &GetPrivateGetMarginsParams{
		HTTPClient: client,
	}
}

/*GetPrivateGetMarginsParams contains all the parameters to send to the API endpoint
for the get private get margins operation typically these are written to a http.Request
*/
type GetPrivateGetMarginsParams struct {

	/*Amount
	  Amount, integer for future, float for option. For perpetual and futures the amount is in USD units, for options it is amount of corresponding cryptocurrency contracts, e.g., BTC or ETH.

	*/
	Amount float64
	/*InstrumentName
	  Instrument name

	*/
	InstrumentName string
	/*Price
	  Price

	*/
	Price float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private get margins params
func (o *GetPrivateGetMarginsParams) WithTimeout(timeout time.Duration) *GetPrivateGetMarginsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private get margins params
func (o *GetPrivateGetMarginsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private get margins params
func (o *GetPrivateGetMarginsParams) WithContext(ctx context.Context) *GetPrivateGetMarginsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private get margins params
func (o *GetPrivateGetMarginsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private get margins params
func (o *GetPrivateGetMarginsParams) WithHTTPClient(client *http.Client) *GetPrivateGetMarginsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private get margins params
func (o *GetPrivateGetMarginsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmount adds the amount to the get private get margins params
func (o *GetPrivateGetMarginsParams) WithAmount(amount float64) *GetPrivateGetMarginsParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the get private get margins params
func (o *GetPrivateGetMarginsParams) SetAmount(amount float64) {
	o.Amount = amount
}

// WithInstrumentName adds the instrumentName to the get private get margins params
func (o *GetPrivateGetMarginsParams) WithInstrumentName(instrumentName string) *GetPrivateGetMarginsParams {
	o.SetInstrumentName(instrumentName)
	return o
}

// SetInstrumentName adds the instrumentName to the get private get margins params
func (o *GetPrivateGetMarginsParams) SetInstrumentName(instrumentName string) {
	o.InstrumentName = instrumentName
}

// WithPrice adds the price to the get private get margins params
func (o *GetPrivateGetMarginsParams) WithPrice(price float64) *GetPrivateGetMarginsParams {
	o.SetPrice(price)
	return o
}

// SetPrice adds the price to the get private get margins params
func (o *GetPrivateGetMarginsParams) SetPrice(price float64) {
	o.Price = price
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateGetMarginsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param amount
	qrAmount := o.Amount
	qAmount := swag.FormatFloat64(qrAmount)
	if qAmount != "" {
		if err := r.SetQueryParam("amount", qAmount); err != nil {
			return err
		}
	}

	// query param instrument_name
	qrInstrumentName := o.InstrumentName
	qInstrumentName := qrInstrumentName
	if qInstrumentName != "" {
		if err := r.SetQueryParam("instrument_name", qInstrumentName); err != nil {
			return err
		}
	}

	// query param price
	qrPrice := o.Price
	qPrice := swag.FormatFloat64(qrPrice)
	if qPrice != "" {
		if err := r.SetQueryParam("price", qPrice); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
