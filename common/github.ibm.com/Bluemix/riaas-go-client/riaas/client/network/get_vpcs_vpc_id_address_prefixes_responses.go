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

// GetVpcsVpcIDAddressPrefixesReader is a Reader for the GetVpcsVpcIDAddressPrefixes structure.
type GetVpcsVpcIDAddressPrefixesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpcsVpcIDAddressPrefixesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpcsVpcIDAddressPrefixesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetVpcsVpcIDAddressPrefixesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVpcsVpcIDAddressPrefixesOK creates a GetVpcsVpcIDAddressPrefixesOK with default headers values
func NewGetVpcsVpcIDAddressPrefixesOK() *GetVpcsVpcIDAddressPrefixesOK {
	return &GetVpcsVpcIDAddressPrefixesOK{}
}

/*GetVpcsVpcIDAddressPrefixesOK handles this case with default header values.

dummy
*/
type GetVpcsVpcIDAddressPrefixesOK struct {
	Payload *models.GetVpcsVpcIDAddressPrefixesOKBody
}

func (o *GetVpcsVpcIDAddressPrefixesOK) Error() string {
	return fmt.Sprintf("[GET /vpcs/{vpc_id}/address_prefixes][%d] getVpcsVpcIdAddressPrefixesOK  %+v", 200, o.Payload)
}

func (o *GetVpcsVpcIDAddressPrefixesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVpcsVpcIDAddressPrefixesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpcsVpcIDAddressPrefixesNotFound creates a GetVpcsVpcIDAddressPrefixesNotFound with default headers values
func NewGetVpcsVpcIDAddressPrefixesNotFound() *GetVpcsVpcIDAddressPrefixesNotFound {
	return &GetVpcsVpcIDAddressPrefixesNotFound{}
}

/*GetVpcsVpcIDAddressPrefixesNotFound handles this case with default header values.

error
*/
type GetVpcsVpcIDAddressPrefixesNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetVpcsVpcIDAddressPrefixesNotFound) Error() string {
	return fmt.Sprintf("[GET /vpcs/{vpc_id}/address_prefixes][%d] getVpcsVpcIdAddressPrefixesNotFound  %+v", 404, o.Payload)
}

func (o *GetVpcsVpcIDAddressPrefixesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
