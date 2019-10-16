// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// DeleteIpsecPoliciesIDReader is a Reader for the DeleteIpsecPoliciesID structure.
type DeleteIpsecPoliciesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIpsecPoliciesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteIpsecPoliciesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteIpsecPoliciesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteIpsecPoliciesIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteIpsecPoliciesIDNoContent creates a DeleteIpsecPoliciesIDNoContent with default headers values
func NewDeleteIpsecPoliciesIDNoContent() *DeleteIpsecPoliciesIDNoContent {
	return &DeleteIpsecPoliciesIDNoContent{}
}

/*DeleteIpsecPoliciesIDNoContent handles this case with default header values.

The IPsec policy was deleted successfully.
*/
type DeleteIpsecPoliciesIDNoContent struct {
}

func (o *DeleteIpsecPoliciesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipsec_policies/{id}][%d] deleteIpsecPoliciesIdNoContent ", 204)
}

func (o *DeleteIpsecPoliciesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIpsecPoliciesIDNotFound creates a DeleteIpsecPoliciesIDNotFound with default headers values
func NewDeleteIpsecPoliciesIDNotFound() *DeleteIpsecPoliciesIDNotFound {
	return &DeleteIpsecPoliciesIDNotFound{}
}

/*DeleteIpsecPoliciesIDNotFound handles this case with default header values.

An IPsec policy with the specified identifier could not be found.
*/
type DeleteIpsecPoliciesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteIpsecPoliciesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ipsec_policies/{id}][%d] deleteIpsecPoliciesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIpsecPoliciesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIpsecPoliciesIDConflict creates a DeleteIpsecPoliciesIDConflict with default headers values
func NewDeleteIpsecPoliciesIDConflict() *DeleteIpsecPoliciesIDConflict {
	return &DeleteIpsecPoliciesIDConflict{}
}

/*DeleteIpsecPoliciesIDConflict handles this case with default header values.

The IPsec policy is in use and cannot be removed.
*/
type DeleteIpsecPoliciesIDConflict struct {
	Payload *models.Riaaserror
}

func (o *DeleteIpsecPoliciesIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /ipsec_policies/{id}][%d] deleteIpsecPoliciesIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteIpsecPoliciesIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
