// Code generated by go-swagger; DO NOT EDIT.

package market_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cicdteam/go-deribit/v3/models"
)

// GetPublicGetIndexReader is a Reader for the GetPublicGetIndex structure.
type GetPublicGetIndexReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicGetIndexReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicGetIndexOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPublicGetIndexOK creates a GetPublicGetIndexOK with default headers values
func NewGetPublicGetIndexOK() *GetPublicGetIndexOK {
	return &GetPublicGetIndexOK{}
}

/*GetPublicGetIndexOK handles this case with default header values.

GetPublicGetIndexOK get public get index o k
*/
type GetPublicGetIndexOK struct {
	Payload *models.PublicIndexResponse
}

func (o *GetPublicGetIndexOK) Error() string {
	return fmt.Sprintf("[GET /public/get_index][%d] getPublicGetIndexOK  %+v", 200, o.Payload)
}

func (o *GetPublicGetIndexOK) GetPayload() *models.PublicIndexResponse {
	return o.Payload
}

func (o *GetPublicGetIndexOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicIndexResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
