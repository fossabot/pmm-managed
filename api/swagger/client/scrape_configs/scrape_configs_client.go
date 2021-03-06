// Code generated by go-swagger; DO NOT EDIT.

package scrape_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new scrape configs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scrape configs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateMixin7 creates creates a new scrape config errors invalid argument 3 if some argument is not valid already exists 6 if scrape config with that job name is already present failed precondition 9 if reachability check was requested and some scrape target can t be reached
*/
func (a *Client) CreateMixin7(params *CreateMixin7Params) (*CreateMixin7OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMixin7Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateMixin7",
		Method:             "POST",
		PathPattern:        "/v0/scrape-configs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateMixin7Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateMixin7OK), nil

}

/*
Delete deletes removes existing scrape config by job name errors not found 5 if no such scrape config is present
*/
func (a *Client) Delete(params *DeleteParams) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Delete",
		Method:             "DELETE",
		PathPattern:        "/v0/scrape-configs/{job_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOK), nil

}

/*
Get gets returns a scrape config by job name errors not found 5 if no such scrape config is present
*/
func (a *Client) Get(params *GetParams) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get",
		Method:             "GET",
		PathPattern:        "/v0/scrape-configs/{job_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOK), nil

}

/*
ListMixin7 lists returns all scrape configs
*/
func (a *Client) ListMixin7(params *ListMixin7Params) (*ListMixin7OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMixin7Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListMixin7",
		Method:             "GET",
		PathPattern:        "/v0/scrape-configs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListMixin7Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListMixin7OK), nil

}

/*
Update updates updates existing scrape config by job name errors invalid argument 3 if some argument is not valid not found 5 if no such scrape config is present failed precondition 9 if reachability check was requested and some scrape target can t be reached
*/
func (a *Client) Update(params *UpdateParams) (*UpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Update",
		Method:             "PUT",
		PathPattern:        "/v0/scrape-configs/{scrape_config.job_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
