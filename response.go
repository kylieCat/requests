package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

// Response is a HTTP response.
type Response struct {
	*Request
	Status
	Headers []Header
	Body
}

// Header returns the canonicalised version of a response header as a string
// If there is no key present in the response the empty string is returned.
// If multiple headers are present, they are canonicalised into as single string
// by joining them with a comma. See RFC 2616 § 4.2.
func (r *Response) Header(key string) string {
	var vals []string
	for _, h := range r.Headers {

		// TODO(dfc) § 4.2 states that not all header values can be combined, but equally those
		// that cannot be combined with a comma may not be present more than once in a
		// header block.
		if h.Key == key {
			vals = append(vals, h.Values...)
		}
	}
	return strings.Join(vals, ",")
}

type Body struct {
	io.ReadCloser
	json    *json.Decoder
	content string
}

// JSON decodes the next JSON encoded object in the body to v.
func (b *Body) JSON(v interface{}) error {
	if b.json == nil {
		b.json = json.NewDecoder(b)
	}
	return b.json.Decode(v)
}

// Content returns the bosy of the response in it's raw sting format
func (b *Body) Content() string {
	if b.content == "" {
		content, err := ioutil.ReadAll(b)
		if err != nil {
			return ""
		}
		b.content = string(content)
	}
	return b.content
}

// return the body as a string, or bytes, or something
func headers(h map[string][]string) []Header {
	headers := make([]Header, 0, len(h))
	for k, v := range h {
		headers = append(headers, Header{
			Key:    k,
			Values: v,
		})
	}
	return headers
}
