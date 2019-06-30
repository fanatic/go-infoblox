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
	Comment          string  `json:"comment,omitempty"`
	Network          string  `json:"network,omitempty"`
	NetworkContainer string  `json:"network_container,omitempty"`
	NetworkView      string  `json:"network_view,omitempty"`
	ExtAttrs         ExtAttr `json:"extattrs,omitempty"`
}
