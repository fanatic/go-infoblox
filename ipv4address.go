package infoblox

// https://192.168.2.200/wapidoc/objects/ipv4address.html
func (c *Client) Ipv4address() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "ipv4address",
	}
}

type Ipv4addressObject struct {
	Object
}

func (c *Client) Ipv4addressObject(ref string) *Ipv4addressObject {
	return &Ipv4addressObject{
		Object{
			_ref: ref,
			r:    c.Ipv4address(),
		},
	}
}
