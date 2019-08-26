// Code generated by go-swagger; DO NOT EDIT.

package hetzner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListHetznerCredentialsReader is a Reader for the ListHetznerCredentials structure.
type ListHetznerCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListHetznerCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListHetznerCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListHetznerCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListHetznerCredentialsOK creates a ListHetznerCredentialsOK with default headers values
func NewListHetznerCredentialsOK() *ListHetznerCredentialsOK {
	return &ListHetznerCredentialsOK{}
}

/*ListHetznerCredentialsOK handles this case with default header values.

CredentialList
*/
type ListHetznerCredentialsOK struct {
	Payload *models.CredentialList
}

func (o *ListHetznerCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/hetzner/credentials][%d] listHetznerCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListHetznerCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHetznerCredentialsDefault creates a ListHetznerCredentialsDefault with default headers values
func NewListHetznerCredentialsDefault(code int) *ListHetznerCredentialsDefault {
	return &ListHetznerCredentialsDefault{
		_statusCode: code,
	}
}

/*ListHetznerCredentialsDefault handles this case with default header values.

ErrorResponse is the default representation of an error
*/
type ListHetznerCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorDetails
}

// Code gets the status code for the list hetzner credentials default response
func (o *ListHetznerCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListHetznerCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/hetzner/credentials][%d] listHetznerCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListHetznerCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}