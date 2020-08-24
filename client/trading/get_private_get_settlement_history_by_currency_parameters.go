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

// NewGetPrivateGetSettlementHistoryByCurrencyParams creates a new GetPrivateGetSettlementHistoryByCurrencyParams object
// with the default values initialized.
func NewGetPrivateGetSettlementHistoryByCurrencyParams() *GetPrivateGetSettlementHistoryByCurrencyParams {
	var ()
	return &GetPrivateGetSettlementHistoryByCurrencyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateGetSettlementHistoryByCurrencyParamsWithTimeout creates a new GetPrivateGetSettlementHistoryByCurrencyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateGetSettlementHistoryByCurrencyParamsWithTimeout(timeout time.Duration) *GetPrivateGetSettlementHistoryByCurrencyParams {
	var ()
	return &GetPrivateGetSettlementHistoryByCurrencyParams{

		timeout: timeout,
	}
}

// NewGetPrivateGetSettlementHistoryByCurrencyParamsWithContext creates a new GetPrivateGetSettlementHistoryByCurrencyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateGetSettlementHistoryByCurrencyParamsWithContext(ctx context.Context) *GetPrivateGetSettlementHistoryByCurrencyParams {
	var ()
	return &GetPrivateGetSettlementHistoryByCurrencyParams{

		Context: ctx,
	}
}

// NewGetPrivateGetSettlementHistoryByCurrencyParamsWithHTTPClient creates a new GetPrivateGetSettlementHistoryByCurrencyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateGetSettlementHistoryByCurrencyParamsWithHTTPClient(client *http.Client) *GetPrivateGetSettlementHistoryByCurrencyParams {
	var ()
	return &GetPrivateGetSettlementHistoryByCurrencyParams{
		HTTPClient: client,
	}
}

/*GetPrivateGetSettlementHistoryByCurrencyParams contains all the parameters to send to the API endpoint
for the get private get settlement history by currency operation typically these are written to a http.Request
*/
type GetPrivateGetSettlementHistoryByCurrencyParams struct {

	/*Count
	  Number of requested items, default - `20`

	*/
	Count *int64
	/*Currency
	  The currency symbol

	*/
	Currency string
	/*Type
	  Settlement type

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) WithTimeout(timeout time.Duration) *GetPrivateGetSettlementHistoryByCurrencyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) WithContext(ctx context.Context) *GetPrivateGetSettlementHistoryByCurrencyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) WithHTTPClient(client *http.Client) *GetPrivateGetSettlementHistoryByCurrencyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) WithCount(count *int64) *GetPrivateGetSettlementHistoryByCurrencyParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) SetCount(count *int64) {
	o.Count = count
}

// WithCurrency adds the currency to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) WithCurrency(currency string) *GetPrivateGetSettlementHistoryByCurrencyParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithType adds the typeVar to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) WithType(typeVar *string) *GetPrivateGetSettlementHistoryByCurrencyParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get private get settlement history by currency params
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateGetSettlementHistoryByCurrencyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	// query param currency
	qrCurrency := o.Currency
	qCurrency := qrCurrency
	if qCurrency != "" {
		if err := r.SetQueryParam("currency", qCurrency); err != nil {
			return err
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
