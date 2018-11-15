package requests

import (
	"net/url"
	"io"
	"bytes"
)

type RequestOptionFunc func(*Request) error

// WithHeader applies the header to the request.
func WithHeader(key, value string) RequestOptionFunc {
	return func(r *Request) error {
		r.Headers = append(r.Headers, Header{
			Key:    key,
			Values: []string{value},
		})
		return nil
	}
}

// WithQueryParam adds a key, value pair to the `QueryParam` struct.
func WithQueryParam(key string, values ...string) RequestOptionFunc {
	return func(r *Request) error {
		r.QueryParams[key] = append(r.QueryParams[key], values...)
		return nil
	}
}

// WithQueryParams sets the `QueryParams` field on the `Request` struct.
func WithQueryParams(params url.Values) RequestOptionFunc {
	return func(r *Request) error {
		r.QueryParams = params
		return nil
	}
}

// WithBasicAuth attaches an Authorization header to the `Request` struct that
// uses basic auth.
func WithBasicAuth(token string) RequestOptionFunc {
	return func(request *Request) error {
		request.Headers = append(request.Headers, Header{Key:"Authorization",  Values: []string{"Basic " + token}})
		return nil
	}
}

// WithBearerToken attaches an Authorization header to the `Request` struct that
// uses a bearer token.
func WithBearerToken(token string) RequestOptionFunc {
	return func(request *Request) error {
		request.Headers = append(request.Headers, Header{Key:"Authorization",  Values: []string{"Bearer " + token}})
		return nil
	}
}

// Request is a HTTP request.
type Request struct {
	Method  string
	URL     string
	Headers []Header
	Body    io.Reader
	// QueryParams is a map of query parameters. Uses url.Values
	QueryParams url.Values
	// Fragment is the bit after a #
	Fragment string
	// User info
	User string
}

// GetUrl constructs the final URL from the components on the `Request` struct.
func (r Request) GetUrl() string {
	var buf bytes.Buffer
	buf.WriteString(r.URL)
	if r.QueryParams != nil {
		buf.WriteByte('?')
		buf.WriteString(r.QueryParams.Encode())
	}
	if r.Fragment != "" {
		buf.WriteByte('#')
		buf.WriteString(r.Fragment)
	}
	return buf.String()
}
