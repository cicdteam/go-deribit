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

// GetPrivateGetOpenOrdersByInstrumentReader is a Reader for the GetPrivateGetOpenOrdersByInstrument structure.
type GetPrivateGetOpenOrdersByInstrumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetOpenOrdersByInstrumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateGetOpenOrdersByInstrumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateGetOpenOrdersByInstrumentOK creates a GetPrivateGetOpenOrdersByInstrumentOK with default headers values
func NewGetPrivateGetOpenOrdersByInstrumentOK() *GetPrivateGetOpenOrdersByInstrumentOK {
	return &GetPrivateGetOpenOrdersByInstrumentOK{}
}

/*GetPrivateGetOpenOrdersByInstrumentOK handles this case with default header values.

GetPrivateGetOpenOrdersByInstrumentOK get private get open orders by instrument o k
*/
type GetPrivateGetOpenOrdersByInstrumentOK struct {
	Payload *models.PrivateGetOpenOrdersResponse
}

func (o *GetPrivateGetOpenOrdersByInstrumentOK) Error() string {
	return fmt.Sprintf("[GET /private/get_open_orders_by_instrument][%d] getPrivateGetOpenOrdersByInstrumentOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetOpenOrdersByInstrumentOK) GetPayload() *models.PrivateGetOpenOrdersResponse {
	return o.Payload
}

func (o *GetPrivateGetOpenOrdersByInstrumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetOpenOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
