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

// GetPrivateGetOrderMarginByIdsReader is a Reader for the GetPrivateGetOrderMarginByIds structure.
type GetPrivateGetOrderMarginByIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetOrderMarginByIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateGetOrderMarginByIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateGetOrderMarginByIdsOK creates a GetPrivateGetOrderMarginByIdsOK with default headers values
func NewGetPrivateGetOrderMarginByIdsOK() *GetPrivateGetOrderMarginByIdsOK {
	return &GetPrivateGetOrderMarginByIdsOK{}
}

/*GetPrivateGetOrderMarginByIdsOK handles this case with default header values.

GetPrivateGetOrderMarginByIdsOK get private get order margin by ids o k
*/
type GetPrivateGetOrderMarginByIdsOK struct {
	Payload *models.PrivateGetOrderMarginByIdsResponse
}

func (o *GetPrivateGetOrderMarginByIdsOK) Error() string {
	return fmt.Sprintf("[GET /private/get_order_margin_by_ids][%d] getPrivateGetOrderMarginByIdsOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetOrderMarginByIdsOK) GetPayload() *models.PrivateGetOrderMarginByIdsResponse {
	return o.Payload
}

func (o *GetPrivateGetOrderMarginByIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetOrderMarginByIdsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
