package infoblox

import "fmt"

func (c *Client) RecordPtr() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:ptr",
	}
}

type RecordPtrObject struct {
	Object
	Comment  string `json:"comment,omitempty"`
	Ipv4Addr string `json:"ipv4addr,omitempty"`
	Ipv6Addr string `json:"ipv6addr,omitempty"`
	Name     string `json:"name,omitempty"`
	PtrDname string `json:"ptrdname,omitempty"`
	Ttl      int    `json:"ttl,omitempty"`
	View     string `json:"view,omitempty"`
}

func (c *Client) RecordPtrObject(ref string) *RecordPtrObject {
	ptr := RecordPtrObject{}
	ptr.Object = Object{
		Ref: ref,
		r:   c.RecordPtr(),
	}
	return &ptr
}

func (c *Client) GetRecordPtr(ref string, opts *Options) (*RecordPtrObject, error) {
	resp, err := c.RecordPtrObject(ref).get(opts)
	if err != nil {
		return nil, fmt.Errorf("Could not get created PTR record: %s", err)
	}
	var out RecordPtrObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) FindRecordPtr(name string) ([]RecordPtrObject, error) {
	field := "name"
	conditions := []Condition{Condition{Field: &field, Value: name}}
	resp, err := c.RecordPtr().find(conditions, nil)
	if err != nil {
		return nil, err
	}

	var out []RecordPtrObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
