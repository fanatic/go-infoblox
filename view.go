package infoblox

// View returns an Infoblox View resource.
func (c *Client) View() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "view",
	}
}

// ViewObject defines the Network view object's fields
type ViewObject struct {
	Object
	Comment     string `json:"comment,omitempty"`
	Name        string `json:"name,omitempty"`
	IsDefault   bool   `json:"is_default,omitempty"`
	NetworkView string `json:"network_view,omitempty"`
}
