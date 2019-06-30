package infoblox

// IPv6Network returns an Infoblox IPv6 network resource
func (c *Client) IPv6Network() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "ipv6network",
	}
}

// IPv6NetworkObject defines the Infoblox IPv6 Network object's fields
type IPv6NetworkObject struct {
	Object
	Comment          string  `json:"comment,omitempty"`
	Network          string  `json:"network,omitempty"`
	NetworkContainer string  `json:"network_container,omitempty"`
	NetworkView      string  `json:"network_view,omitempty"`
	Netmask          int     `json:"netmask,omitempty"`
	ExtAttrs         ExtAttr `json:"extattrs,omitempty"`
}
