package infoblox

import "fmt"

// https://192.168.2.200/wapidoc/objects/record.host.html
func (c *Client) RecordHost() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:host",
	}
}

type RecordHostObject struct {
	Object    `json:"-"`
	Ipv4Addrs []HostIpv4Addr `json:"ipv4addrs,omitempty"`
	//Ipv6Addrs string `json:"ipv6addrs,omitempty"`
	Name string `json:"name,omitempty"`
	View string `json:"view,omitempty"`
}

type HostIpv4Addr struct {
	Object           `json:"-"`
	ConfigureForDHCP bool   `json:"configure_for_dhcp,omitempty"`
	Host             string `json:"host,omitempty"`
	Ipv4Addr         string `json:"ipv4addr,omitempty"`
	MAC              string `json:"mac,omitempty"`
}

func (c *Client) RecordHostObject(ref string) *RecordHostObject {
	host := RecordHostObject{}
	host.Object = Object{
		_ref: ref,
		r:    c.RecordHost(),
	}
	return &host
}

func (c *Client) GetRecordHost(ref string) (*RecordHostObject, error) {
	resp, err := c.RecordHostObject(ref).get(nil)
	if err != nil {
		return nil, fmt.Errorf("Could not get created host record: %s", err)
	}

	var out RecordHostObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) FindRecordHost(name string) ([]RecordHostObject, error) {
	field := "name"
	conditions := []Condition{Condition{Field: &field, Value: name}}
	resp, err := c.Ipv4address().find(conditions, nil)
	if err != nil {
		return nil, err
	}

	var out []RecordHostObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
