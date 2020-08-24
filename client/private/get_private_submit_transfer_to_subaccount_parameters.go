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

// NewGetPrivateSubmitTransferToSubaccountParams creates a new GetPrivateSubmitTransferToSubaccountParams object
// with the default values initialized.
func NewGetPrivateSubmitTransferToSubaccountParams() *GetPrivateSubmitTransferToSubaccountParams {
	var ()
	return &GetPrivateSubmitTransferToSubaccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateSubmitTransferToSubaccountParamsWithTimeout creates a new GetPrivateSubmitTransferToSubaccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateSubmitTransferToSubaccountParamsWithTimeout(timeout time.Duration) *GetPrivateSubmitTransferToSubaccountParams {
	var ()
	return &GetPrivateSubmitTransferToSubaccountParams{

		timeout: timeout,
	}
}

// NewGetPrivateSubmitTransferToSubaccountParamsWithContext creates a new GetPrivateSubmitTransferToSubaccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateSubmitTransferToSubaccountParamsWithContext(ctx context.Context) *GetPrivateSubmitTransferToSubaccountParams {
	var ()
	return &GetPrivateSubmitTransferToSubaccountParams{

		Context: ctx,
	}
}

// NewGetPrivateSubmitTransferToSubaccountParamsWithHTTPClient creates a new GetPrivateSubmitTransferToSubaccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateSubmitTransferToSubaccountParamsWithHTTPClient(client *http.Client) *GetPrivateSubmitTransferToSubaccountParams {
	var ()
	return &GetPrivateSubmitTransferToSubaccountParams{
		HTTPClient: client,
	}
}

/*GetPrivateSubmitTransferToSubaccountParams contains all the parameters to send to the API endpoint
for the get private submit transfer to subaccount operation typically these are written to a http.Request
*/
type GetPrivateSubmitTransferToSubaccountParams struct {

	/*Amount
	  Amount of funds to be transferred

	*/
	Amount float64
	/*Currency
	  The currency symbol

	*/
	Currency string
	/*Destination
	  Id of destination subaccount

	*/
	Destination int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) WithTimeout(timeout time.Duration) *GetPrivateSubmitTransferToSubaccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) WithContext(ctx context.Context) *GetPrivateSubmitTransferToSubaccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) WithHTTPClient(client *http.Client) *GetPrivateSubmitTransferToSubaccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmount adds the amount to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) WithAmount(amount float64) *GetPrivateSubmitTransferToSubaccountParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) SetAmount(amount float64) {
	o.Amount = amount
}

// WithCurrency adds the currency to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) WithCurrency(currency string) *GetPrivateSubmitTransferToSubaccountParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithDestination adds the destination to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) WithDestination(destination int64) *GetPrivateSubmitTransferToSubaccountParams {
	o.SetDestination(destination)
	return o
}

// SetDestination adds the destination to the get private submit transfer to subaccount params
func (o *GetPrivateSubmitTransferToSubaccountParams) SetDestination(destination int64) {
	o.Destination = destination
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateSubmitTransferToSubaccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param currency
	qrCurrency := o.Currency
	qCurrency := qrCurrency
	if qCurrency != "" {
		if err := r.SetQueryParam("currency", qCurrency); err != nil {
			return err
		}
	}

	// query param destination
	qrDestination := o.Destination
	qDestination := swag.FormatInt64(qrDestination)
	if qDestination != "" {
		if err := r.SetQueryParam("destination", qDestination); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
