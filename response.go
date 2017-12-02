package infoblox

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	STATUS_OK           = 200
	STATUS_CREATED      = 201
	STATUS_INVALID      = 400
	STATUS_UNAUTHORIZED = 401
	STATUS_FORBIDDEN    = 403
	STATUS_NOTFOUND     = 404
	STATUS_LIMIT        = 429
	STATUS_GATEWAY      = 502
)

// APIResponse is used to parse the response from Infoblox.
// GET requests tend to respond with Objects or lists of Objects
// while POST,PUT,DELETE returns the Object References as a string
// https://192.168.2.200/wapidoc/#get
type APIResponse http.Response

func (r APIResponse) readBody() ([]byte, error) {
	defer r.Body.Close()

	// The default HTTP client now handles decompression for us, but
	// this is left for backwards compatability
	header := strings.ToLower(r.Header.Get("Content-Encoding"))
	if header == "" || strings.Index(header, "gzip") == -1 {
		return ioutil.ReadAll(r.Body)
	}

	reader, err := gzip.NewReader(r.Body)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
}

func (r APIResponse) ReadBody() string {
	if b, err := r.readBody(); err == nil {
		return string(b)
	}
	return ""
}

// Parses a JSON encoded HTTP response into the supplied interface.
func (r APIResponse) Parse(out interface{}) error {
	switch r.StatusCode {
	case STATUS_UNAUTHORIZED:
		fallthrough
	case STATUS_NOTFOUND:
		fallthrough
	case STATUS_GATEWAY:
		fallthrough
	case STATUS_FORBIDDEN:
		fallthrough
	case STATUS_INVALID:
		b, err := r.readBody()
		if err != nil {
			return err
		}
		e := Error{}
		if err := json.Unmarshal(b, &e); err != nil {
			return fmt.Errorf("Error parsing error response: %v", string(b))
		}
		return e
	//case STATUS_LIMIT:
	//  err = RateLimitError{
	//    Limit:     r.RateLimit(),
	//    Remaining: r.RateLimitRemaining(),
	//    Reset:     r.RateLimitReset(),
	//  }
	//  return
	case STATUS_CREATED:
		fallthrough
	case STATUS_OK:
		b, err := r.readBody()
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, out); err != nil && err != io.EOF {
			return err
		}
	default:
		b, err := r.readBody()
		if err != nil {
			return err
		}
		return fmt.Errorf("%v", string(b))
	}
	return nil
}
