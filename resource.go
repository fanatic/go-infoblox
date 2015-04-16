package infoblox

import (
  "bytes"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "net/url"
  "strconv"
  "strings"
)

//Resource represents a WAPI object type
type Resource struct {
  conn       *Client
  wapiObject string
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
func (r Resource) All(opts *Options) ([]map[string]interface{}, error) {
  return r.Find([]Condition{}, opts)
}

// Find resources with query parameters. Conditions are combined with AND
// logic.  When a field is a list of extensible attribute that can have multiple
// values, the condition is true if any value in the list matches.
func (r Resource) Find(query []Condition, opts *Options) ([]map[string]interface{}, error) {
  q := r.getQuery(opts, query, url.Values{})

  log.Printf("GET %s\n", r.resourceURI()+"?"+q.Encode())
  req, err := http.NewRequest("GET", r.resourceURI()+"?"+q.Encode(), nil)
  if err != nil {
    return nil, fmt.Errorf("Error creating request: %v\n", err)
  }

  resp, err := r.conn.SendRequest(req)
  if err != nil {
    return nil, fmt.Errorf("Error sending request: %v\n", err)
  }

  //fmt.Printf("%v", resp.ReadBody())

  var out []map[string]interface{}
  err = resp.Parse(&out)
  if err != nil {
    return nil, fmt.Errorf("%+v\n", err)
  }
  return out, nil
}

func (r Resource) Create(data url.Values, opts *Options, body interface{}) (string, error) {
  q := r.getQuery(opts, []Condition{}, data)
  q.Set("_return_fields", "") //Force object response
  log.Printf("POST %s\n", r.resourceURI()+"?"+q.Encode())

  var req *http.Request
  var err error
  if body == nil {
    // Send URL-encoded data in the request body
    req, err = http.NewRequest("POST", r.resourceURI(), strings.NewReader(q.Encode()))
    if err != nil {
      return "", fmt.Errorf("Error creating request: %v\n", err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  } else {
    // Put url-encoded data in the URL and send the body parameter as a JSON body.
    bodyJSON, err := json.Marshal(body)
    if err != nil {
      return "", fmt.Errorf("Error creating request: %v\n", err)
    }
    log.Printf("POST body: %s\n", bodyJSON)
    req, err = http.NewRequest("POST", r.resourceURI()+"?"+q.Encode(), bytes.NewReader(bodyJSON))
    if err != nil {
      return "", fmt.Errorf("Error creating request: %v\n", err)
    }
    req.Header.Set("Content-Type", "application/json")
  }

  resp, err := r.conn.SendRequest(req)
  if err != nil {
    return "", fmt.Errorf("Error sending request: %v\n", err)
  }

  //fmt.Printf("%v", resp.ReadBody())

  var out map[string]string
  err = resp.Parse(&out)
  if err != nil {
    return "", fmt.Errorf("%+v\n", err)
  }
  return out["_ref"], nil
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

func (r Resource) resourceURI() string {
  return BASE_PATH + r.wapiObject
}
