package infoblox

import "fmt"

// This object represents the Infoblox Grid.
// https://ipam.illinois.edu/wapidoc/objects/grid.html
func (c *Client) Grid() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "grid",
	}
}

// GridObject defines the GRID record object's fields
type GridObject struct {
	Object
}

// GridObject instantiates an GRID object with a WAPI ref
func (c *Client) GridObject(ref string) *GridObject {
	obj := GridObject{}
	obj.Object = Object{
		Ref: ref,
		r:   c.Grid(),
	}
	return &obj
}

func (c *Client) GetGrids() ([]GridObject, error) {
	resp, err := c.Grid().find(nil, nil)
	if err != nil {
		return nil, err
	}

	var out []GridObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (grid GridObject) RestartServicesForGrid(member_order string, restart_option string, service_option string) error {
	type RestartOptions struct {
		MemberOrder   string `json:"member_order,omitempty"`
		RestartOption string `json:"restart_option,omitempty"`
		ServiceOption string `json:"service_option,omitempty"`
	}
	restartIfNeeded := &RestartOptions{
		MemberOrder:   member_order,
		RestartOption: restart_option,
		ServiceOption: service_option,
	}
	_, err := grid.FunctionCall("restartservices", restartIfNeeded)
	if err != nil {
		return fmt.Errorf("Error sending request: %v", err)
	}
	return nil
}
