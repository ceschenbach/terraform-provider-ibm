// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetInstancesIDTagsTagNameReader is a Reader for the GetInstancesIDTagsTagName structure.
type GetInstancesIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewGetInstancesIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetInstancesIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetInstancesIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetInstancesIDTagsTagNameNoContent creates a GetInstancesIDTagsTagNameNoContent with default headers values
func NewGetInstancesIDTagsTagNameNoContent() *GetInstancesIDTagsTagNameNoContent {
	return &GetInstancesIDTagsTagNameNoContent{}
}

/*GetInstancesIDTagsTagNameNoContent handles this case with default header values.

error
*/
type GetInstancesIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *GetInstancesIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/tags/{tag_name}][%d] getInstancesIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *GetInstancesIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesIDTagsTagNameNotFound creates a GetInstancesIDTagsTagNameNotFound with default headers values
func NewGetInstancesIDTagsTagNameNotFound() *GetInstancesIDTagsTagNameNotFound {
	return &GetInstancesIDTagsTagNameNotFound{}
}

/*GetInstancesIDTagsTagNameNotFound handles this case with default header values.

error
*/
type GetInstancesIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetInstancesIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/tags/{tag_name}][%d] getInstancesIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *GetInstancesIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesIDTagsTagNameDefault creates a GetInstancesIDTagsTagNameDefault with default headers values
func NewGetInstancesIDTagsTagNameDefault(code int) *GetInstancesIDTagsTagNameDefault {
	return &GetInstancesIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*GetInstancesIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type GetInstancesIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get instances ID tags tag name default response
func (o *GetInstancesIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *GetInstancesIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[GET /instances/{id}/tags/{tag_name}][%d] GetInstancesIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstancesIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}