package infoblox

// Ipv6address returns an Infoblox Ipv6 address resource
func (c *Client) Ipv6address() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "ipv6address",
	}
}

// Ipv6addressObject defines the Ipv6 address object's fields
type Ipv6addressObject struct {
	Object
	IPAddress   string   `json:"ip_address,omitempty"`
	IsConflict  bool     `json:"is_conflict"`
	Names       []string `json:"names,omitempty"`
	Network     string   `json:"network,omitempty"`
	NetworkView string   `json:"network_view,omitempty"`
	Objects     []string `json:"objects,omitempty"`
	Status      string   `json:"status,omitempty"`
	Types       []string `json:"types,omitempty"`
	Usage       []string `json:"usage,omitempty"`
}
