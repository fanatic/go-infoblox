package infoblox

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "strings"
)

//Resource represents a WAPI object
type Object struct {
  _ref string
  r    *Resource
}

func (o Object) Get(opts *Options) (map[string]interface{}, error) {
  q := o.r.getQuery(opts, []Condition{}, url.Values{})

  log.Printf("GET %s\n", o.objectURI()+"?"+q.Encode())
  req, err := http.NewRequest("GET", o.objectURI()+"?"+q.Encode(), nil)
  if err != nil {
    return nil, fmt.Errorf("Error creating request: %v\n", err)
  }

  resp, err := o.r.conn.SendRequest(req)
  if err != nil {
    return nil, fmt.Errorf("Error sending request: %v\n", err)
  }

  //fmt.Printf("%v", resp.ReadBody())

  var out map[string]interface{}
  err = resp.Parse(&out)
  if err != nil {
    return nil, fmt.Errorf("%+v\n", err)
  }
  return out, nil
}

func (o Object) Update(data url.Values, opts *Options) (string, error) {
  q := o.r.getQuery(opts, []Condition{}, data)
  q.Set("_return_fields", "") //Force object response

  log.Printf("PUT %s\n", o.objectURI()+"?"+q.Encode())
  req, err := http.NewRequest("PUT", o.objectURI(), strings.NewReader(q.Encode()))
  if err != nil {
    return "", fmt.Errorf("Error creating request: %v\n", err)
  }
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  resp, err := o.r.conn.SendRequest(req)
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

func (o Object) Delete(opts *Options) error {
  q := o.r.getQuery(opts, []Condition{}, url.Values{})

  log.Printf("DELETE %s\n", o.objectURI()+"?"+q.Encode())
  req, err := http.NewRequest("DELETE", o.objectURI()+"?"+q.Encode(), nil)
  if err != nil {
    return fmt.Errorf("Error creating request: %v\n", err)
  }

  resp, err := o.r.conn.SendRequest(req)
  if err != nil {
    return fmt.Errorf("Error sending request: %v\n", err)
  }

  //fmt.Printf("%v", resp.ReadBody())

  var out interface{}
  err = resp.Parse(&out)
  if err != nil {
    return fmt.Errorf("%+v\n", err)
  }
  return nil
}

func (o Object) FunctionCall(functionName string, inputFields url.Values) (map[string]interface{}, error) {
  inputFields.Set("_function", functionName)

  log.Printf("POST %s\n", o.objectURI()+"?"+inputFields.Encode())
  req, err := http.NewRequest("POST", o.objectURI(), strings.NewReader(inputFields.Encode()))
  if err != nil {
    return nil, fmt.Errorf("Error creating request: %v\n", err)
  }
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  resp, err := o.r.conn.SendRequest(req)
  if err != nil {
    return nil, fmt.Errorf("Error sending request: %v\n", err)
  }

  //fmt.Printf("%v", resp.ReadBody())

  var out map[string]interface{}
  err = resp.Parse(&out)
  if err != nil {
    return nil, fmt.Errorf("%+v\n", err)
  }
  return out, nil
}

func (o Object) objectURI() string {
  return BASE_PATH + o._ref
}
