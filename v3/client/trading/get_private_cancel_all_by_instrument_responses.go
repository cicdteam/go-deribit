// Code generated by go-swagger; DO NOT EDIT.

package trading

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cicdteam/go-deribit/v3/models"
)

// GetPrivateCancelAllByInstrumentReader is a Reader for the GetPrivateCancelAllByInstrument structure.
type GetPrivateCancelAllByInstrumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateCancelAllByInstrumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateCancelAllByInstrumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateCancelAllByInstrumentOK creates a GetPrivateCancelAllByInstrumentOK with default headers values
func NewGetPrivateCancelAllByInstrumentOK() *GetPrivateCancelAllByInstrumentOK {
	return &GetPrivateCancelAllByInstrumentOK{}
}

/*GetPrivateCancelAllByInstrumentOK handles this case with default header values.

GetPrivateCancelAllByInstrumentOK get private cancel all by instrument o k
*/
type GetPrivateCancelAllByInstrumentOK struct {
	Payload *models.PrivateCancelAllResponse
}

func (o *GetPrivateCancelAllByInstrumentOK) Error() string {
	return fmt.Sprintf("[GET /private/cancel_all_by_instrument][%d] getPrivateCancelAllByInstrumentOK  %+v", 200, o.Payload)
}

func (o *GetPrivateCancelAllByInstrumentOK) GetPayload() *models.PrivateCancelAllResponse {
	return o.Payload
}

func (o *GetPrivateCancelAllByInstrumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateCancelAllResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
