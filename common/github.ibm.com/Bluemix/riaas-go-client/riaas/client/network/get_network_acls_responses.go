// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetNetworkAclsReader is a Reader for the GetNetworkAcls structure.
type GetNetworkAclsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkAclsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworkAclsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetNetworkAclsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkAclsOK creates a GetNetworkAclsOK with default headers values
func NewGetNetworkAclsOK() *GetNetworkAclsOK {
	return &GetNetworkAclsOK{}
}

/*GetNetworkAclsOK handles this case with default header values.

dummy
*/
type GetNetworkAclsOK struct {
	Payload *models.GetNetworkAclsOKBody
}

func (o *GetNetworkAclsOK) Error() string {
	return fmt.Sprintf("[GET /network_acls][%d] getNetworkAclsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkAclsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetNetworkAclsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkAclsInternalServerError creates a GetNetworkAclsInternalServerError with default headers values
func NewGetNetworkAclsInternalServerError() *GetNetworkAclsInternalServerError {
	return &GetNetworkAclsInternalServerError{}
}

/*GetNetworkAclsInternalServerError handles this case with default header values.

error
*/
type GetNetworkAclsInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetNetworkAclsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /network_acls][%d] getNetworkAclsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNetworkAclsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
