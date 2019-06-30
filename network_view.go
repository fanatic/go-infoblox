package infoblox

// NetworkView returns an Infoblox network view resource.
func (c *Client) NetworkView() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "networkview",
	}
}

// NetworkViewObject defines the Network view object's fields
type NetworkViewObject struct {
	Object
	Name      string `json:"name,omitempty"`
	IsDefault bool   `json:"is_default,omitempty"`
}
