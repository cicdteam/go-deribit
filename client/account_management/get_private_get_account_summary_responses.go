// Code generated by go-swagger; DO NOT EDIT.

package account_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cicdteam/go-deribit/models"
)

// GetPrivateGetAccountSummaryReader is a Reader for the GetPrivateGetAccountSummary structure.
type GetPrivateGetAccountSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetAccountSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateGetAccountSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateGetAccountSummaryOK creates a GetPrivateGetAccountSummaryOK with default headers values
func NewGetPrivateGetAccountSummaryOK() *GetPrivateGetAccountSummaryOK {
	return &GetPrivateGetAccountSummaryOK{}
}

/*GetPrivateGetAccountSummaryOK handles this case with default header values.

GetPrivateGetAccountSummaryOK get private get account summary o k
*/
type GetPrivateGetAccountSummaryOK struct {
	Payload *models.PrivateAccountResponse
}

func (o *GetPrivateGetAccountSummaryOK) Error() string {
	return fmt.Sprintf("[GET /private/get_account_summary][%d] getPrivateGetAccountSummaryOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetAccountSummaryOK) GetPayload() *models.PrivateAccountResponse {
	return o.Payload
}

func (o *GetPrivateGetAccountSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
