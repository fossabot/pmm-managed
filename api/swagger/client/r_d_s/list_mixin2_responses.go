// Code generated by go-swagger; DO NOT EDIT.

package r_d_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/percona/pmm-managed/api/swagger/models"
)

// ListMixin2Reader is a Reader for the ListMixin2 structure.
type ListMixin2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMixin2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListMixin2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListMixin2OK creates a ListMixin2OK with default headers values
func NewListMixin2OK() *ListMixin2OK {
	return &ListMixin2OK{}
}

/*ListMixin2OK handles this case with default header values.

(empty)
*/
type ListMixin2OK struct {
	Payload *models.APIRDSListResponse
}

func (o *ListMixin2OK) Error() string {
	return fmt.Sprintf("[GET /v0/rds][%d] listMixin2OK  %+v", 200, o.Payload)
}

func (o *ListMixin2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIRDSListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
