package requests

import (
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/pkg/errors"
)

type HTTPClient interface {
	Do(r *http.Request) (*http.Response, error)
}

func newDefaultClient() *http.Client {
	c := http.DefaultClient
	jar, _ := cookiejar.New(new(cookiejar.Options))
	c.Jar = jar
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return c
}

// Client is a HTTP Client.
type Client struct {
	client HTTPClient
}

type ClientOptionFunc func(*Client)

func CustomClient(c HTTPClient) ClientOptionFunc {
	return func(client *Client) {
		client.client = c
	}
}

// NewClient returns a `Client` struct. If no arguments are passed a default HTTP client is used.
// A custom HTTP client can also be passed in via a functional parameter.
func NewClient(opts ...ClientOptionFunc) *Client {
	c := &Client{
		client: newDefaultClient(),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Get issues a GET to the specified URL.
func (c *Client) Get(url string, options ...RequestOptionFunc) (*Response, error) {
	req := Request{
		Method: http.MethodGet,
		URL:    url,
	}
	if err := applyOptions(&req, options...); err != nil {
		return nil, err
	}
	return c.do(&req)
}

func (c *Client) Delete(url string, options ...RequestOptionFunc) (*Response, error) {
	req := Request{
		Method: http.MethodDelete,
		URL:    url,
	}
	if err := applyOptions(&req, options...); err != nil {
		return nil, err
	}
	return c.do(&req)
}

func (c *Client) Head(url string, options ...RequestOptionFunc) (*Response, error) {
	req := Request{
		Method: http.MethodHead,
		URL:    url,
	}
	if err := applyOptions(&req, options...); err != nil {
		return nil, err
	}
	return c.do(&req)
}

func (c *Client) Trace(url string, options ...RequestOptionFunc) (*Response, error) {
	req := Request{
		Method: http.MethodTrace,
		URL:    url,
	}
	if err := applyOptions(&req, options...); err != nil {
		return nil, err
	}
	return c.do(&req)
}

// Post issues a POST request to the specified URL.
func (c *Client) Post(url string, body io.Reader, options ...RequestOptionFunc) (*Response, error) {
	req := Request{
		Method: http.MethodPost,
		URL:    url,
		Body:   body,
	}
	if err := applyOptions(&req, options...); err != nil {
		return nil, err
	}
	return c.do(&req)
}

func (c *Client) Put(url string, body io.Reader, options ...RequestOptionFunc) (*Response, error) {
	req := Request{
		Method: http.MethodPut,
		URL:    url,
		Body:   body,
	}
	if err := applyOptions(&req, options...); err != nil {
		return nil, err
	}
	return c.do(&req)
}

func (c *Client) do(request *Request) (*Response, error) {
	req, err := newHttpRequest(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	r := Response{
		Request: request,
		Status: Status{
			Code:   resp.StatusCode,
			Reason: resp.Status[4:],
		},
		Headers: headers(resp.Header),
		Body: Body{
			ReadCloser: resp.Body,
		},
	}
	return &r, nil
}

func applyOptions(req *Request, options ...RequestOptionFunc) error {
	for _, opt := range options {
		if err := opt(req); err != nil {
			return err
		}
	}
	return nil
}

// newHttpRequest converts a *requests.Request into a *http.Request
func newHttpRequest(request *Request) (*http.Request, error) {
	req, err := http.NewRequest(request.Method, request.GetUrl(), request.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header = toHeaders(request.Headers)
	return req, nil
}

// toHeaders converts from Request's Headers slice to http.Request's map[string][]string
func toHeaders(headers []Header) map[string][]string {
	if len(headers) == 0 {
		return nil
	}

	m := make(map[string][]string)
	for _, h := range headers {
		m[h.Key] = h.Values
	}
	return m
}

// Header is a HTTP header.
type Header struct {
	Key    string
	Values []string
}
