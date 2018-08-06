package requests

import (
	"io"
)

var c = NewClient()

// Get issues a GET to the specified URL.
func Get(url string, options ...RequestOptionFunc) (*Response, error) {
	return c.Get(url, options...)
}

func Delete(url string, options ...RequestOptionFunc) (*Response, error) {
	return c.Delete(url, options...)
}

func Head(url string, options ...RequestOptionFunc) (*Response, error) {
	return c.Head(url, options...)
}

func Trace(url string, options ...RequestOptionFunc) (*Response, error) {
	return c.Trace(url, options...)
}

// Post issues a POST request to the specified URL.
func Post(url string, body io.Reader, options ...RequestOptionFunc) (*Response, error) {
	return c.Post(url, body, options...)
}

func Put(url string, body io.Reader, options ...RequestOptionFunc) (*Response, error) {
	return c.Put(url, body, options...)
}
