package infoblox

// ZoneAuth returns an Infoblox authoritative zone resource.
func (c *Client) ZoneAuth() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "zone_auth",
	}
}

// ZoneAuthObject defines the Network authoritative zone object's fields
type ZoneAuthObject struct {
	Object
	Comment     string `json:"comment,omitempty"`
	Fqdn        string `json:"fqdn,omitempty"`
	NetworkView string `json:"network_view,omitempty"`
	Parent      string `json:"parent,omitempty"`
	View        string `json:"view,omitempty"`
}
