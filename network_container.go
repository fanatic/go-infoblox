package infoblox

// NetworkContainer returns an Infoblox Network Container resource.
func (c *Client) NetworkContainer() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "networkcontainer",
	}
}

// NetworkContainerObject defines the Infoblox Network Container object's fields.
type NetworkContainerObject struct {
	Object
	Comment  string  `json:"comment,omitempty"`
	Network  string  `json:"network,omitempty"`
	ExtAttrs ExtAttr `json:"extattrs,omitempty"`
}
