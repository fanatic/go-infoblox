package infoblox

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "strconv"
  "strings"
)

type Resource struct {
  conn  *Client
  child ResourceImpl
  _ref  string
}

type ResourceImpl interface {
  WAPIObject() string
  RemoteAttributeAccessors() []string
  RemotePostAccessors() []string
  RemoteAttributeWriters() []string
}

type Options struct {
  //The maximum number of objects to be returned.  If set to a negative
  //number the appliance will return an error when the number of returned
  //objects would exceed the setting. The default is -1000. If this is
  //set to a positive number, the results will be truncated when necessary.
  MaxResults *int

  ReturnFields      []string //A list of returned fields
  ReturnBasicFields bool     // Return basic fields in addition to ReturnFields
}

// Conditions are used for searching
type Condition struct {
  Field     *string // EITHER A documented field of the object (only set one)
  Attribute *string // OR the name of an extensible attribute (only set one)
  Modifiers string  // Optional search modifiers "!:~<>" (otherwise exact match)
  Value     string  // Value or regular expression to search for
}

// All returns an array of all records for this resource
func (r Resource) All(opts *Options) ([]interface{}, error) {
  return r.Find([]Condition{}, opts)
}

// Find resources with query parameters. Conditions are combined with AND
// logic.  When a field is a list of extensible attribute that can have multiple
// values, the condition is true if any value in the list matches.
func (r Resource) Find(query []Condition, opts *Options) ([]interface{}, error) {
  q := r.getQuery(opts, query)

  log.Printf("GET %s\n", r.resourceURI("")+"?"+q.Encode())
  req, err := http.NewRequest("GET", r.resourceURI("")+"?"+q.Encode(), nil)
  if err != nil {
    return nil, fmt.Errorf("Error creating request: %v\n", err)
  }

  resp, err := r.conn.SendRequest(req)
  if err != nil {
    return nil, fmt.Errorf("Error sending request: %v\n", err)
  }

  //fmt.Printf("%v", resp.ReadBody())

  var out []interface{}
  err = resp.Parse(&out)
  if err != nil {
    return nil, fmt.Errorf("%+v\n", err)
  }
  return out, nil
}

// Get referenced object
func (r Resource) Get(objectRef string, opts *Options) (interface{}, error) {
  q := r.getQuery(opts, []Condition{})

  log.Printf("GET %s\n", r.resourceURI(objectRef)+"?"+q.Encode())
  req, err := http.NewRequest("GET", r.resourceURI(objectRef)+"?"+q.Encode(), nil)
  if err != nil {
    return nil, fmt.Errorf("Error creating request: %v\n", err)
  }

  resp, err := r.conn.SendRequest(req)
  if err != nil {
    return nil, fmt.Errorf("Error sending request: %v\n", err)
  }

  //fmt.Printf("%v", resp.ReadBody())

  var out interface{}
  err = resp.Parse(&out)
  if err != nil {
    return nil, fmt.Errorf("%+v\n", err)
  }
  return out, nil
}

func (r Resource) getQuery(opts *Options, query []Condition) url.Values {
  v := url.Values{}

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
      search += *cond.Field
    }
    search += cond.Modifiers
    v.Set(search, cond.Value)
  }

  return v
}

//@NotImplemented
func Create() {}
func Delete() {}
func Update() {}

func (r Resource) resourceURI(ref string) string {
  if ref == "" {
    return BASE_PATH + r.child.WAPIObject()
  } else {
    return BASE_PATH + ref
  }
}
