// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mituan8/pay/pkg/api-kms/v1/model"
)

// CreateEthereumTransactionReader is a Reader for the CreateEthereumTransaction structure.
type CreateEthereumTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEthereumTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEthereumTransactionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEthereumTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEthereumTransactionCreated creates a CreateEthereumTransactionCreated with default headers values
func NewCreateEthereumTransactionCreated() *CreateEthereumTransactionCreated {
	return &CreateEthereumTransactionCreated{}
}

/* CreateEthereumTransactionCreated describes a response with status code 201, with default header values.

Transaction Created
*/
type CreateEthereumTransactionCreated struct {
	Payload *model.EthereumTransaction
}

func (o *CreateEthereumTransactionCreated) Error() string {
	return fmt.Sprintf("[POST /wallet/{walletId}/transaction/eth][%d] createEthereumTransactionCreated  %+v", 201, o.Payload)
}
func (o *CreateEthereumTransactionCreated) GetPayload() *model.EthereumTransaction {
	return o.Payload
}

func (o *CreateEthereumTransactionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.EthereumTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEthereumTransactionBadRequest creates a CreateEthereumTransactionBadRequest with default headers values
func NewCreateEthereumTransactionBadRequest() *CreateEthereumTransactionBadRequest {
	return &CreateEthereumTransactionBadRequest{}
}

/* CreateEthereumTransactionBadRequest describes a response with status code 400, with default header values.

Validation error / Not found
*/
type CreateEthereumTransactionBadRequest struct {
	Payload *model.ErrorResponse
}

func (o *CreateEthereumTransactionBadRequest) Error() string {
	return fmt.Sprintf("[POST /wallet/{walletId}/transaction/eth][%d] createEthereumTransactionBadRequest  %+v", 400, o.Payload)
}
func (o *CreateEthereumTransactionBadRequest) GetPayload() *model.ErrorResponse {
	return o.Payload
}

func (o *CreateEthereumTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
