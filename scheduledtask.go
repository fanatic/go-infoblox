package infoblox

// https://192.168.2.200/wapidoc/objects/scheduledtask.html
func (c *Client) ScheduledTask() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "scheduledtask",
	}
}

type ScheduledTaskObject struct {
	Object
}

func (c *Client) ScheduledTaskObject(ref string) *ScheduledTaskObject {
	return &ScheduledTaskObject{
		Object{
			_ref: ref,
			r:    c.ScheduledTask(),
		},
	}
}
