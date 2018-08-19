package infoblox

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// https://102.168.2.200/wapidoc/objects/record.a.html
func (c *Client) GlobalSearch() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "search",
	}
}

type SearchObject struct {
	Object
	Comment  string `json:"comment,omitempty"`
	Ipv4Addr string `json:"ipv4addr,omitempty"`
	Name     string `json:"name,omitempty"`
	Ttl      int    `json:"ttl,omitempty"`
	View     string `json:"view,omitempty"`
}

func (c *Client) SearchObject(ref string) *RecordAObject {
	a := RecordAObject{}
	a.Object = Object{
		Ref: ref,
		r:   c.RecordA(),
	}
	return &a
}

func (c *Client) Search(name string, objtype string) (string, error) {
	field := "search_string"
	conditions := []Condition{
		Condition{
			Field:     &field,
			Modifiers: "~",
			Value:     name,
		},
	}
	if objtype != "" {
		objtypestring := "objtype"
		conditions = append(conditions, Condition{
			Field: &objtypestring,
			Value: objtype,
		})

	}
	resp, respErr := c.GlobalSearch().find(conditions, nil)
	if respErr != nil {
		return "", respErr
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 400 {
		var err APIErrorResponse
		if err2 := json.Unmarshal(body, &err); err2 != nil {
			return "", err2
		}
		return "", errors.New(err.Text)
	}

	return string(body[:]), nil
}
