package infoblox

import (
	"encoding/json"
	"fmt"
)

// RecordCname returns the CNAME record resource
// https://192.168.2.200/wapidoc/objects/record.host.html
func (c *Client) RecordCname() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:cname",
	}
}

// RecordCnameObject defines the CNAME record object's fields
type RecordCnameObject struct {
	Object
	Ref       string `json:"_ref,omitempty"`
	Comment   string `json:"comment,omitempty"`
	Canonical string `json:"canonical,omitempty"`
	Name      string `json:"name,omitempty"`
	Ttl       int    `json:"ttl,omitempty"`
	View      string `json:"view,omitempty"`
}

// RecordCnameObject instantiates an CNAME record object with a WAPI ref
func (c *Client) RecordCnameObject(ref string) *RecordCnameObject {
	cname := RecordCnameObject{}
	cname.Object = Object{
		Ref: ref,
		r:   c.RecordCname(),
	}
	return &cname
}

func (c *Client) FindRecordCname(name string, view string) ([]RecordCnameObject, error) {
	field := "name"
	viewName := "view"
	// conditions := []Condition{Condition{Field: &field, Value: name}}
	conditions := []Condition{
		Condition{
			Field: &field,
			Value: name,
		},
		Condition{
			Field: &viewName,
			Value: view,
		},
	}
	resp, err := c.RecordCname().find(conditions, nil)
	if err != nil {
		return nil, err
	}

	var out []RecordCnameObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetRecordCname fetches an CNAME record from the Infoblox WAPI by its ref
func (c *Client) GetRecordCname(ref string, opts *Options) (*RecordCnameObject, error) {
	resp, err := c.RecordCnameObject(ref).get(opts)
	if err != nil {
		return nil, fmt.Errorf("Could not get created CNAME record: %s", err)
	}
	var out RecordCnameObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) UpdateRecordCname(recordCnameObject RecordCnameObject) (string, error) {
	d, _ := json.Marshal(recordCnameObject)
	existingRecord, err := c.FindRecordCname(recordCnameObject.Name, recordCnameObject.View)
	if err != nil {
		return "", fmt.Errorf("Could not find CNAME record to update: %s", err)
	}
	resp, err := c.RecordCname().UpdateJson(existingRecord[0].Ref, nil, d)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (c *Client) CreateRecordCname(recordCnameObject RecordCnameObject) (string, error) {
	d, _ := json.Marshal(recordCnameObject)
	resp, err := c.RecordCname().CreateJson("record:cname", nil, d)
	if err != nil {
		return "", err
	}
	return resp, nil
}
