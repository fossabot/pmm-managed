// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	models "github.com/Percona-Lab/pmm-api/http/models"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// AddBareMetalReader is a Reader for the AddBareMetal structure.
type AddBareMetalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddBareMetalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddBareMetalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddBareMetalOK creates a AddBareMetalOK with default headers values
func NewAddBareMetalOK() *AddBareMetalOK {
	return &AddBareMetalOK{}
}

/*AddBareMetalOK handles this case with default header values.

(empty)
*/
type AddBareMetalOK struct {
	Payload *models.InventoryAddBareMetalNodeResponse
}

func (o *AddBareMetalOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Nodes/AddBareMetal][%d] addBareMetalOK  %+v", 200, o.Payload)
}

func (o *AddBareMetalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryAddBareMetalNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
