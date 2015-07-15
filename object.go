package infoblox

import (
	"fmt"
	"net/url"
)

//Resource represents a WAPI object
type Object struct {
	Ref string `json:"_ref"`
	r   *Resource
}

func (o Object) Get(opts *Options) (map[string]interface{}, error) {
	resp, err := o.get(opts)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = resp.Parse(&out)
	if err != nil {
		return nil, fmt.Errorf("%+v\n", err)
	}

	return out, nil
}

func (o Object) get(opts *Options) (*APIResponse, error) {
	q := o.r.getQuery(opts, []Condition{}, url.Values{})

	resp, err := o.r.conn.SendRequest("GET", o.objectURI()+"?"+q.Encode(), "", nil)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %v\n", err)
	}
	return resp, nil
}

func (o Object) Update(data url.Values, opts *Options) (string, error) {
	q := o.r.getQuery(opts, []Condition{}, data)
	q.Set("_return_fields", "") //Force object response

	resp, err := o.r.conn.SendRequest("PUT", o.objectURI(), q.Encode(), map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
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

	resp, err := o.r.conn.SendRequest("DELETE", o.objectURI()+"?"+q.Encode(), "", nil)
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

	resp, err := o.r.conn.SendRequest("POST", o.objectURI(), inputFields.Encode(), map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
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
	return BASE_PATH + o.Ref
}
