package infoblox

// https://192.168.2.200/wapidoc/objects/record.host.html
func (c *Client) RecordHost() *Resource {
  return &Resource{
    conn:       c,
    wapiObject: "record:host",
  }
}

type RecordHostObject struct {
  Object
}

func (c *Client) RecordHostObject(ref string) *RecordHostObject {
  return &RecordHostObject{
    Object{
      _ref: ref,
      r:    c.RecordHost(),
    },
  }
}
