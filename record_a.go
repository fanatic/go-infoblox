package infoblox

func (c *Client) RecordA() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:a",
	}
}

type RecordAObject struct {
	Object
}

func (c *Client) RecordAObject(ref string) *RecordAObject {
	return &RecordAObject{
		Object{
			Ref: ref,
			r:   c.RecordA(),
		},
	}
}
