// Code generated by go-swagger; DO NOT EDIT.

package bootstrap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bootstrap API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bootstrap API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BootstrapCreateBootstrapAgent(params *BootstrapCreateBootstrapAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapCreateBootstrapAgentOK, error)

	BootstrapCreateRelayNetwork(params *BootstrapCreateRelayNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapCreateRelayNetworkOK, error)

	BootstrapDeleteBootstrapAgent(params *BootstrapDeleteBootstrapAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapDeleteBootstrapAgentOK, error)

	BootstrapDeleteRelayNetwork(params *BootstrapDeleteRelayNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapDeleteRelayNetworkOK, error)

	BootstrapDownloadRelayNetwork(params *BootstrapDownloadRelayNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapDownloadRelayNetworkOK, error)

	BootstrapDownloadRelayNetworkAgent(params *BootstrapDownloadRelayNetworkAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapDownloadRelayNetworkAgentOK, error)

	BootstrapGetBootstrapAgent(params *BootstrapGetBootstrapAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapAgentOK, error)

	BootstrapGetBootstrapAgentConfig(params *BootstrapGetBootstrapAgentConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapAgentConfigOK, error)

	BootstrapGetBootstrapAgentTemplate(params *BootstrapGetBootstrapAgentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapAgentTemplateOK, error)

	BootstrapGetBootstrapAgentTemplates(params *BootstrapGetBootstrapAgentTemplatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapAgentTemplatesOK, error)

	BootstrapGetBootstrapAgents(params *BootstrapGetBootstrapAgentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapAgentsOK, error)

	BootstrapGetBootstrapInfra(params *BootstrapGetBootstrapInfraParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapInfraOK, error)

	BootstrapGetRelayNetwork(params *BootstrapGetRelayNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetRelayNetworkOK, error)

	BootstrapGetRelayNetworks(params *BootstrapGetRelayNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetRelayNetworksOK, error)

	BootstrapPatchBootstrapAgentTemplate(params *BootstrapPatchBootstrapAgentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapPatchBootstrapAgentTemplateOK, error)

	BootstrapPatchBootstrapInfra(params *BootstrapPatchBootstrapInfraParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapPatchBootstrapInfraOK, error)

	BootstrapPatchRelayNetwork(params *BootstrapPatchRelayNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapPatchRelayNetworkOK, error)

	BootstrapRegisterBootstrapAgent(params *BootstrapRegisterBootstrapAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapRegisterBootstrapAgentOK, error)

	BootstrapUpdateBootstrapAgent(params *BootstrapUpdateBootstrapAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapUpdateBootstrapAgentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BootstrapCreateBootstrapAgent bootstrap create bootstrap agent API
*/
func (a *Client) BootstrapCreateBootstrapAgent(params *BootstrapCreateBootstrapAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapCreateBootstrapAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapCreateBootstrapAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_CreateBootstrapAgent",
		Method:             "POST",
		PathPattern:        "/v2/sentry/bootstrap/{spec.templateRef}/agent",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapCreateBootstrapAgentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapCreateBootstrapAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapCreateBootstrapAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapCreateRelayNetwork bootstrap create relay network API
*/
func (a *Client) BootstrapCreateRelayNetwork(params *BootstrapCreateRelayNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapCreateRelayNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapCreateRelayNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_CreateRelayNetwork",
		Method:             "POST",
		PathPattern:        "/v2/sentry/relaynetwork",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapCreateRelayNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapCreateRelayNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapCreateRelayNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapDeleteBootstrapAgent bootstrap delete bootstrap agent API
*/
func (a *Client) BootstrapDeleteBootstrapAgent(params *BootstrapDeleteBootstrapAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapDeleteBootstrapAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapDeleteBootstrapAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_DeleteBootstrapAgent",
		Method:             "DELETE",
		PathPattern:        "/v2/sentry/bootstrap/{spec.templateRef}/agent/{metadata.name}",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapDeleteBootstrapAgentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapDeleteBootstrapAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapDeleteBootstrapAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapDeleteRelayNetwork bootstrap delete relay network API
*/
func (a *Client) BootstrapDeleteRelayNetwork(params *BootstrapDeleteRelayNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapDeleteRelayNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapDeleteRelayNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_DeleteRelayNetwork",
		Method:             "DELETE",
		PathPattern:        "/v2/sentry/relaynetwork/{metadata.name}",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapDeleteRelayNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapDeleteRelayNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapDeleteRelayNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapDownloadRelayNetwork bootstrap download relay network API
*/
func (a *Client) BootstrapDownloadRelayNetwork(params *BootstrapDownloadRelayNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapDownloadRelayNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapDownloadRelayNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_DownloadRelayNetwork",
		Method:             "GET",
		PathPattern:        "/v2/sentry/relaynetwork/{metadata.name}/download",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapDownloadRelayNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapDownloadRelayNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapDownloadRelayNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapDownloadRelayNetworkAgent bootstrap download relay network agent API
*/
func (a *Client) BootstrapDownloadRelayNetworkAgent(params *BootstrapDownloadRelayNetworkAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapDownloadRelayNetworkAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapDownloadRelayNetworkAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_DownloadRelayNetworkAgent",
		Method:             "GET",
		PathPattern:        "/v2/sentry/relaynetwork/{metadata.name}/{clusterScope}/agentdownload",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapDownloadRelayNetworkAgentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapDownloadRelayNetworkAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapDownloadRelayNetworkAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapGetBootstrapAgent bootstrap get bootstrap agent API
*/
func (a *Client) BootstrapGetBootstrapAgent(params *BootstrapGetBootstrapAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapGetBootstrapAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_GetBootstrapAgent",
		Method:             "GET",
		PathPattern:        "/v2/sentry/bootstrap/{spec.templateRef}/agent/{metadata.name}",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapGetBootstrapAgentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapGetBootstrapAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapGetBootstrapAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapGetBootstrapAgentConfig bootstrap get bootstrap agent config API
*/
func (a *Client) BootstrapGetBootstrapAgentConfig(params *BootstrapGetBootstrapAgentConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapAgentConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapGetBootstrapAgentConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_GetBootstrapAgentConfig",
		Method:             "GET",
		PathPattern:        "/v2/sentry/bootstrap/{spec.templateRef}/agent/{metadata.name}/config",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapGetBootstrapAgentConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapGetBootstrapAgentConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapGetBootstrapAgentConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapGetBootstrapAgentTemplate bootstrap get bootstrap agent template API
*/
func (a *Client) BootstrapGetBootstrapAgentTemplate(params *BootstrapGetBootstrapAgentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapAgentTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapGetBootstrapAgentTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_GetBootstrapAgentTemplate",
		Method:             "GET",
		PathPattern:        "/v2/sentry/bootstrap/template/{metadata.name}",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapGetBootstrapAgentTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapGetBootstrapAgentTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapGetBootstrapAgentTemplateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapGetBootstrapAgentTemplates bootstrap get bootstrap agent templates API
*/
func (a *Client) BootstrapGetBootstrapAgentTemplates(params *BootstrapGetBootstrapAgentTemplatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapAgentTemplatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapGetBootstrapAgentTemplatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_GetBootstrapAgentTemplates",
		Method:             "GET",
		PathPattern:        "/v2/sentry/bootstrap/template",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapGetBootstrapAgentTemplatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapGetBootstrapAgentTemplatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapGetBootstrapAgentTemplatesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapGetBootstrapAgents bootstrap get bootstrap agents API
*/
func (a *Client) BootstrapGetBootstrapAgents(params *BootstrapGetBootstrapAgentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapGetBootstrapAgentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_GetBootstrapAgents",
		Method:             "GET",
		PathPattern:        "/v2/sentry/bootstrap/{templateScope}/agent",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapGetBootstrapAgentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapGetBootstrapAgentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapGetBootstrapAgentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapGetBootstrapInfra bootstrap get bootstrap infra API
*/
func (a *Client) BootstrapGetBootstrapInfra(params *BootstrapGetBootstrapInfraParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetBootstrapInfraOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapGetBootstrapInfraParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_GetBootstrapInfra",
		Method:             "GET",
		PathPattern:        "/v2/sentry/bootstrap/infra/{metadata.name}",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapGetBootstrapInfraReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapGetBootstrapInfraOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapGetBootstrapInfraDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapGetRelayNetwork bootstrap get relay network API
*/
func (a *Client) BootstrapGetRelayNetwork(params *BootstrapGetRelayNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetRelayNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapGetRelayNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_GetRelayNetwork",
		Method:             "GET",
		PathPattern:        "/v2/sentry/relaynetwork/{metadata.name}",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapGetRelayNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapGetRelayNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapGetRelayNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapGetRelayNetworks bootstrap get relay networks API
*/
func (a *Client) BootstrapGetRelayNetworks(params *BootstrapGetRelayNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapGetRelayNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapGetRelayNetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_GetRelayNetworks",
		Method:             "GET",
		PathPattern:        "/v2/sentry/relaynetwork",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapGetRelayNetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapGetRelayNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapGetRelayNetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapPatchBootstrapAgentTemplate bootstrap patch bootstrap agent template API
*/
func (a *Client) BootstrapPatchBootstrapAgentTemplate(params *BootstrapPatchBootstrapAgentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapPatchBootstrapAgentTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapPatchBootstrapAgentTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_PatchBootstrapAgentTemplate",
		Method:             "PUT",
		PathPattern:        "/v2/sentry/bootstrap/template/{metadata.name}",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapPatchBootstrapAgentTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapPatchBootstrapAgentTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapPatchBootstrapAgentTemplateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapPatchBootstrapInfra bootstrap patch bootstrap infra API
*/
func (a *Client) BootstrapPatchBootstrapInfra(params *BootstrapPatchBootstrapInfraParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapPatchBootstrapInfraOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapPatchBootstrapInfraParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_PatchBootstrapInfra",
		Method:             "PUT",
		PathPattern:        "/v2/sentry/bootstrap/infra/{metadata.name}",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapPatchBootstrapInfraReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapPatchBootstrapInfraOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapPatchBootstrapInfraDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapPatchRelayNetwork bootstrap patch relay network API
*/
func (a *Client) BootstrapPatchRelayNetwork(params *BootstrapPatchRelayNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapPatchRelayNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapPatchRelayNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_PatchRelayNetwork",
		Method:             "PUT",
		PathPattern:        "/v2/sentry/relaynetwork/{metadata.name}",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapPatchRelayNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapPatchRelayNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapPatchRelayNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapRegisterBootstrapAgent bootstrap register bootstrap agent API
*/
func (a *Client) BootstrapRegisterBootstrapAgent(params *BootstrapRegisterBootstrapAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapRegisterBootstrapAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapRegisterBootstrapAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_RegisterBootstrapAgent",
		Method:             "POST",
		PathPattern:        "/v2/sentry/bootstrap/{templateToken}/register",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapRegisterBootstrapAgentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapRegisterBootstrapAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapRegisterBootstrapAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BootstrapUpdateBootstrapAgent bootstrap update bootstrap agent API
*/
func (a *Client) BootstrapUpdateBootstrapAgent(params *BootstrapUpdateBootstrapAgentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BootstrapUpdateBootstrapAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBootstrapUpdateBootstrapAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Bootstrap_UpdateBootstrapAgent",
		Method:             "PUT",
		PathPattern:        "/v2/sentry/bootstrap/{spec.templateRef}/agent/{metadata.name}",
		ProducesMediaTypes: []string{"application/json", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BootstrapUpdateBootstrapAgentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BootstrapUpdateBootstrapAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BootstrapUpdateBootstrapAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}