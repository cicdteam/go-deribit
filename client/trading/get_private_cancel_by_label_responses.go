// Code generated by go-swagger; DO NOT EDIT.

package trading

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cicdteam/go-deribit/models"
)

// GetPrivateCancelByLabelReader is a Reader for the GetPrivateCancelByLabel structure.
type GetPrivateCancelByLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateCancelByLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateCancelByLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateCancelByLabelOK creates a GetPrivateCancelByLabelOK with default headers values
func NewGetPrivateCancelByLabelOK() *GetPrivateCancelByLabelOK {
	return &GetPrivateCancelByLabelOK{}
}

/*GetPrivateCancelByLabelOK handles this case with default header values.

GetPrivateCancelByLabelOK get private cancel by label o k
*/
type GetPrivateCancelByLabelOK struct {
	Payload *models.PrivateCancelAllResponse
}

func (o *GetPrivateCancelByLabelOK) Error() string {
	return fmt.Sprintf("[GET /private/cancel_by_label][%d] getPrivateCancelByLabelOK  %+v", 200, o.Payload)
}

func (o *GetPrivateCancelByLabelOK) GetPayload() *models.PrivateCancelAllResponse {
	return o.Payload
}

func (o *GetPrivateCancelByLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateCancelAllResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
