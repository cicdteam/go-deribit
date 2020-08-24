// Code generated by go-swagger; DO NOT EDIT.

package private

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cicdteam/go-deribit/v3/models"
)

// GetPrivateGetPositionsReader is a Reader for the GetPrivateGetPositions structure.
type GetPrivateGetPositionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetPositionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateGetPositionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPrivateGetPositionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateGetPositionsOK creates a GetPrivateGetPositionsOK with default headers values
func NewGetPrivateGetPositionsOK() *GetPrivateGetPositionsOK {
	return &GetPrivateGetPositionsOK{}
}

/*GetPrivateGetPositionsOK handles this case with default header values.

When successful returns array of positions
*/
type GetPrivateGetPositionsOK struct {
	Payload *models.PrivateGetPositionsResponse
}

func (o *GetPrivateGetPositionsOK) Error() string {
	return fmt.Sprintf("[GET /private/get_positions][%d] getPrivateGetPositionsOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetPositionsOK) GetPayload() *models.PrivateGetPositionsResponse {
	return o.Payload
}

func (o *GetPrivateGetPositionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetPositionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateGetPositionsBadRequest creates a GetPrivateGetPositionsBadRequest with default headers values
func NewGetPrivateGetPositionsBadRequest() *GetPrivateGetPositionsBadRequest {
	return &GetPrivateGetPositionsBadRequest{}
}

/*GetPrivateGetPositionsBadRequest handles this case with default header values.

When some error occurs
*/
type GetPrivateGetPositionsBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetPrivateGetPositionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /private/get_positions][%d] getPrivateGetPositionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetPrivateGetPositionsBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetPrivateGetPositionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
