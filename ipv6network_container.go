package infoblox

// IPv6NetworkContainer returns an Infoblox IPv6 Network Container resource.
func (c *Client) IPv6NetworkContainer() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "ipv6networkcontainer",
	}
}

// IPv6NetworkContainerObject defines the Infoblox Network Container object's fields.
type IPv6NetworkContainerObject struct {
	Object
	Comment          string  `json:"comment,omitempty"`
	Network          string  `json:"network,omitempty"`
	NetworkContainer string  `json:"network_container,omitempty"`
	NetworkView      string  `json:"network_view,omitempty"`
	ExtAttrs         ExtAttr `json:"extattrs,omitempty"`
}
