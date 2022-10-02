// Package articles provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package articles

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetTasks request
	GetTasks(ctx context.Context, params *GetTasksParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateTask request with any body
	CreateTaskWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateTask(ctx context.Context, body CreateTaskJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// MarkTaskAsComplete request with any body
	MarkTaskAsCompleteWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	MarkTaskAsComplete(ctx context.Context, body MarkTaskAsCompleteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ReassignTasks request
	ReassignTasks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetTasks(ctx context.Context, params *GetTasksParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTasksRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateTaskWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateTaskRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateTask(ctx context.Context, body CreateTaskJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateTaskRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) MarkTaskAsCompleteWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewMarkTaskAsCompleteRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) MarkTaskAsComplete(ctx context.Context, body MarkTaskAsCompleteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewMarkTaskAsCompleteRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ReassignTasks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewReassignTasksRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetTasksRequest generates requests for GetTasks
func NewGetTasksRequest(server string, params *GetTasksParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tasks")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "offset", runtime.ParamLocationQuery, params.Offset); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, params.Limit); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateTaskRequest calls the generic CreateTask builder with application/json body
func NewCreateTaskRequest(server string, body CreateTaskJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateTaskRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateTaskRequestWithBody generates requests for CreateTask with any type of body
func NewCreateTaskRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tasks")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewMarkTaskAsCompleteRequest calls the generic MarkTaskAsComplete builder with application/json body
func NewMarkTaskAsCompleteRequest(server string, body MarkTaskAsCompleteJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewMarkTaskAsCompleteRequestWithBody(server, "application/json", bodyReader)
}

// NewMarkTaskAsCompleteRequestWithBody generates requests for MarkTaskAsComplete with any type of body
func NewMarkTaskAsCompleteRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tasks/mark-as-complete")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewReassignTasksRequest generates requests for ReassignTasks
func NewReassignTasksRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tasks/reassign")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetTasks request
	GetTasksWithResponse(ctx context.Context, params *GetTasksParams, reqEditors ...RequestEditorFn) (*GetTasksResponse, error)

	// CreateTask request with any body
	CreateTaskWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateTaskResponse, error)

	CreateTaskWithResponse(ctx context.Context, body CreateTaskJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateTaskResponse, error)

	// MarkTaskAsComplete request with any body
	MarkTaskAsCompleteWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*MarkTaskAsCompleteResponse, error)

	MarkTaskAsCompleteWithResponse(ctx context.Context, body MarkTaskAsCompleteJSONRequestBody, reqEditors ...RequestEditorFn) (*MarkTaskAsCompleteResponse, error)

	// ReassignTasks request
	ReassignTasksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ReassignTasksResponse, error)
}

type GetTasksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Tasks
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GetTasksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTasksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateTaskResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r CreateTaskResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateTaskResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type MarkTaskAsCompleteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON204      *Task
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r MarkTaskAsCompleteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r MarkTaskAsCompleteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ReassignTasksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r ReassignTasksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ReassignTasksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetTasksWithResponse request returning *GetTasksResponse
func (c *ClientWithResponses) GetTasksWithResponse(ctx context.Context, params *GetTasksParams, reqEditors ...RequestEditorFn) (*GetTasksResponse, error) {
	rsp, err := c.GetTasks(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTasksResponse(rsp)
}

// CreateTaskWithBodyWithResponse request with arbitrary body returning *CreateTaskResponse
func (c *ClientWithResponses) CreateTaskWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateTaskResponse, error) {
	rsp, err := c.CreateTaskWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateTaskResponse(rsp)
}

func (c *ClientWithResponses) CreateTaskWithResponse(ctx context.Context, body CreateTaskJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateTaskResponse, error) {
	rsp, err := c.CreateTask(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateTaskResponse(rsp)
}

// MarkTaskAsCompleteWithBodyWithResponse request with arbitrary body returning *MarkTaskAsCompleteResponse
func (c *ClientWithResponses) MarkTaskAsCompleteWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*MarkTaskAsCompleteResponse, error) {
	rsp, err := c.MarkTaskAsCompleteWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseMarkTaskAsCompleteResponse(rsp)
}

func (c *ClientWithResponses) MarkTaskAsCompleteWithResponse(ctx context.Context, body MarkTaskAsCompleteJSONRequestBody, reqEditors ...RequestEditorFn) (*MarkTaskAsCompleteResponse, error) {
	rsp, err := c.MarkTaskAsComplete(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseMarkTaskAsCompleteResponse(rsp)
}

// ReassignTasksWithResponse request returning *ReassignTasksResponse
func (c *ClientWithResponses) ReassignTasksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ReassignTasksResponse, error) {
	rsp, err := c.ReassignTasks(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseReassignTasksResponse(rsp)
}

// ParseGetTasksResponse parses an HTTP response from a GetTasksWithResponse call
func ParseGetTasksResponse(rsp *http.Response) (*GetTasksResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTasksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Tasks
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseCreateTaskResponse parses an HTTP response from a CreateTaskWithResponse call
func ParseCreateTaskResponse(rsp *http.Response) (*CreateTaskResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateTaskResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseMarkTaskAsCompleteResponse parses an HTTP response from a MarkTaskAsCompleteWithResponse call
func ParseMarkTaskAsCompleteResponse(rsp *http.Response) (*MarkTaskAsCompleteResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &MarkTaskAsCompleteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 204:
		var dest Task
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON204 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseReassignTasksResponse parses an HTTP response from a ReassignTasksWithResponse call
func ParseReassignTasksResponse(rsp *http.Response) (*ReassignTasksResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ReassignTasksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
