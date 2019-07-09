package infoblox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"strings"
)

//Resource represents a WAPI object type
type Resource struct {
	conn       *Client
	wapiObject string
}

// NewResource creates new resource object from the passed parameters.
// As Infoblox contains hundreds of different objects, this function
// will allow the users to adopt new ones vary easy.
func NewResource(c *Client, wapiObject string) *Resource {
	return &Resource{
		conn:       c,
		wapiObject: wapiObject,
	}
}

// Options represents the Options to be passed to the Infoblox WAPI
type Options struct {
	//The maximum number of objects to be returned.  If set to a negative
	//number the appliance will return an error when the number of returned
	//objects would exceed the setting. The default is -1000. If this is
	//set to a positive number, the results will be truncated when necessary.
	MaxResults *int

	ReturnFields      []string //A list of returned fields
	ReturnBasicFields bool     // Return basic fields in addition to ReturnFields
}

// A Condition is used for searching
type Condition struct {
	Field     *string // EITHER A documented field of the object (only set one)
	Attribute *string // OR the name of an extensible attribute (only set one)
	Modifiers string  // Optional search modifiers "!:~<>" (otherwise exact match)
	Value     string  // Value or regular expression to search for
}

// All returns an array of all records for this resource
func (r Resource) All(opts *Options) ([]map[string]interface{}, error) {
	return r.Find([]Condition{}, opts)
}

func (r Resource) Delete(ref string) (string, error) {
	uri := r.resourceBase() + ref
	resp, err := r.conn.SendRequest("DELETE", uri, "", nil)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bodyOut := buf.String()
	return bodyOut, nil
}

// Find resources with query parameters. Conditions are combined with AND
// logic.  When a field is a list of extensible attribute that can have multiple
// values, the condition is true if any value in the list matches.
func (r Resource) Find(query []Condition, opts *Options) ([]map[string]interface{}, error) {
	resp, err := r.find(query, opts)
	if err != nil {
		return nil, err
	}

	var out []map[string]interface{}
	err = resp.Parse(&out)
	if err != nil {
		return nil, fmt.Errorf("%+v", err)
	}
	return out, nil
}

// Query retrieves objects from an Infoblox instance that meet specific
// conditions and options. The return object is defined by the user and
// passed to the function by the "out" parameter.
func (r Resource) Query(query []Condition, opts *Options, out interface{}) error {
	resp, err := r.find(query, opts)
	if err != nil {
		return err
	}

	err = resp.Parse(&out)
	if err != nil {
		return fmt.Errorf("%+v", err)
	}
	return nil
}

func (r Resource) find(query []Condition, opts *Options) (*APIResponse, error) {
	q := r.getQuery(opts, query, url.Values{})

	resp, err := r.conn.SendRequest("GET", r.resourceURI()+"?"+q.Encode(), "", nil)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %v", err)
	}

	return resp, nil
}

type APIErrorResponse struct {
	Text string `json:"text"`
}

func (r Resource) JsonAction(url string, actionType string, data string) (*APIResponse, error) {
	var err error
	head := make(map[string]string)

	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v\n", err)
	}

	head["Content-Type"] = "application/json"

	resp, err := r.conn.SendRequest(actionType, url, data, head)
	return resp, err
}

func (r Resource) CreateJson(url string, opts *Options, data []byte) (string, error) {
	var urlStr, bodyJSON string

	urlStr = r.resourceURI()
	bodyJSON = string(data[:])

	resp, _ := r.JsonAction(urlStr, "POST", bodyJSON)
	if resp.StatusCode == 400 {
		var t APIErrorResponse
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &t)
		return t.Text, errors.New(t.Text)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body[:]), nil
	}
}

func (r Resource) UpdateJson(url string, opts *Options, data []byte) (string, error) {
	var urlStr, bodyJSON string

	urlStr = r.resourceBase() + url
	bodyJSON = string(data[:])

	resp, _ := r.JsonAction(urlStr, "PUT", bodyJSON)
	if resp.StatusCode == 400 {
		var t APIErrorResponse
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &t)
		return t.Text, errors.New(t.Text)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body[:]), nil
	}
}

// Create creates a resource. Returns the ref of the created resource.
func (r Resource) Create(data url.Values, opts *Options, body interface{}) (string, error) {
	q := r.getQuery(opts, []Condition{}, data)
	q.Set("_return_fields", "") //Force object response

	var err error
	head := make(map[string]string)
	var bodyStr, urlStr string
	if body == nil {
		// Send URL-encoded data in the request body
		urlStr = r.resourceURI()
		bodyStr = q.Encode()
		head["Content-Type"] = "application/x-www-form-urlencoded"
	} else {
		// Put url-encoded data in the URL and send the body parameter as a JSON body.
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return "", fmt.Errorf("Error creating request: %v", err)
		}
		log.Printf("POST body: %s\n", bodyJSON)
		urlStr = r.resourceURI() + "?" + q.Encode()
		bodyStr = string(bodyJSON)
		head["Content-Type"] = "application/json"
	}

	resp, err := r.conn.SendRequest("POST", urlStr, bodyStr, head)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %v", err)
	}

	//fmt.Printf("%v", resp.ReadBody())

	// If you POST to record:host with a scheduled creation time, it sends back a string regardless of the presence of _return_fields
	var responseData interface{}
	var ret string
	if err := resp.Parse(&responseData); err != nil {
		return "", fmt.Errorf("%+v", err)
	}
	switch s := responseData.(type) {
	case string:
		ret = s
	case map[string]interface{}:
		ret = s["_ref"].(string)
	default:
		return "", fmt.Errorf("Invalid return type %T", s)
	}

	return ret, nil
}

func (r Resource) getQuery(opts *Options, query []Condition, extra url.Values) url.Values {
	v := extra

	returnFieldOption := "_return_fields"
	if opts != nil && opts.ReturnBasicFields {
		returnFieldOption = "_return_fields+"
	}

	if opts != nil && opts.ReturnFields != nil {
		v.Set(returnFieldOption, strings.Join(opts.ReturnFields, ","))
	}

	if opts != nil && opts.MaxResults != nil {
		v.Set("_max_results", strconv.Itoa(*opts.MaxResults))
	}

	for _, cond := range query {
		search := ""
		if cond.Field != nil {
			search += *cond.Field
		} else if cond.Attribute != nil {
			search += "*" + *cond.Attribute
		}
		search += cond.Modifiers
		v.Set(search, cond.Value)
	}

	return v
}

func (r Resource) resourceBase() string {
	return BasePath
}
func (r Resource) resourceURI() string {
	return BasePath + r.wapiObject
}
