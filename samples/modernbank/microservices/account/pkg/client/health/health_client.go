// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new health API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for health API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
HealthCheck returns 200 to prove the service is alive

returns 200
*/
func (a *Client) HealthCheck(params *HealthCheckParams) (*HealthCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "healthCheck",
		Method:             "POST",
		PathPattern:        "/health",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HealthCheckReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HealthCheckOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}