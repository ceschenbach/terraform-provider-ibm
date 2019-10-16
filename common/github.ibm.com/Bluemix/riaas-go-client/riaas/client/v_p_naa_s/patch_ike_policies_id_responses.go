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

// PatchIkePoliciesIDReader is a Reader for the PatchIkePoliciesID structure.
type PatchIkePoliciesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchIkePoliciesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchIkePoliciesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchIkePoliciesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchIkePoliciesIDOK creates a PatchIkePoliciesIDOK with default headers values
func NewPatchIkePoliciesIDOK() *PatchIkePoliciesIDOK {
	return &PatchIkePoliciesIDOK{}
}

/*PatchIkePoliciesIDOK handles this case with default header values.

The IKE policy was updated successfully.
*/
type PatchIkePoliciesIDOK struct {
	Payload *models.IKEPolicy
}

func (o *PatchIkePoliciesIDOK) Error() string {
	return fmt.Sprintf("[PATCH /ike_policies/{id}][%d] patchIkePoliciesIdOK  %+v", 200, o.Payload)
}

func (o *PatchIkePoliciesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IKEPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIkePoliciesIDBadRequest creates a PatchIkePoliciesIDBadRequest with default headers values
func NewPatchIkePoliciesIDBadRequest() *PatchIkePoliciesIDBadRequest {
	return &PatchIkePoliciesIDBadRequest{}
}

/*PatchIkePoliciesIDBadRequest handles this case with default header values.

An invalid IKE policy patch was provided.
*/
type PatchIkePoliciesIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchIkePoliciesIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /ike_policies/{id}][%d] patchIkePoliciesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchIkePoliciesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
