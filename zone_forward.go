package infoblox

// ZoneForward returns an Infoblox forward zone resource.
func (c *Client) ZoneForward() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "zone_forward",
	}
}

// ZoneForwardObject defines the Network forward zone object's fields
type ZoneForwardObject struct {
	Object
	Comment        string          `json:"comment,omitempty"`
	Address        string          `json:"address,omitempty"`
	FQDN           string          `json:"fqdn,omitempty"`
	Parent         string          `json:"parent,omitempty"`
	View           string          `json:"view,omitempty"`
	ForwardServers []ForwardServer `json:"forward_to,omitempty"`
	ExtAttrs       ExtAttr         `json:"extattrs,omitempty"`
}

// ForwardServer defines ZoneForward forwarding server object.
type ForwardServer struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
}
