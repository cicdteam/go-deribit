// Code generated by go-swagger; DO NOT EDIT.

package internal_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cicdteam/go-deribit/models"
)

// GetPublicGetFooterReader is a Reader for the GetPublicGetFooter structure.
type GetPublicGetFooterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicGetFooterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicGetFooterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPublicGetFooterOK creates a GetPublicGetFooterOK with default headers values
func NewGetPublicGetFooterOK() *GetPublicGetFooterOK {
	return &GetPublicGetFooterOK{}
}

/*GetPublicGetFooterOK handles this case with default header values.

ok response
*/
type GetPublicGetFooterOK struct {
	Payload *models.PublicGetFooterResponse
}

func (o *GetPublicGetFooterOK) Error() string {
	return fmt.Sprintf("[GET /public/get_footer][%d] getPublicGetFooterOK  %+v", 200, o.Payload)
}

func (o *GetPublicGetFooterOK) GetPayload() *models.PublicGetFooterResponse {
	return o.Payload
}

func (o *GetPublicGetFooterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicGetFooterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
