package infoblox

// https://192.168.2.200/wapidoc/objects/record.host.html
func (c *Client) RecordCname() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:cname",
	}
}

type RecordCnameObject struct {
	Object
}

func (c *Client) RecordCnameObject(ref string) *RecordCnameObject {
	return &RecordCnameObject{
		Object{
			Ref: ref,
			r:   c.RecordCname(),
		},
	}
}
