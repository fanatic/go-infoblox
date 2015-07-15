package infoblox

func (c *Client) RecordPtr() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:ptr",
	}
}

type RecordPtrObject struct {
	Object
}

func (c *Client) RecordPtrObject(ref string) *RecordPtrObject {
	return &RecordPtrObject{
		Object{
			Ref: ref,
			r:   c.RecordPtr(),
		},
	}
}
