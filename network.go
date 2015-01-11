package infoblox

import (
  "fmt"
  "net/url"
  "strconv"
  "strings"
)

// https://192.168.2.200/wapidoc/objects/network.html
func (c *Client) Network() *Resource {
  return &Resource{
    conn:       c,
    wapiObject: "network",
  }
}

type NetworkObject struct {
  Object
}

func (c *Client) NetworkObject(ref string) *NetworkObject {
  return &NetworkObject{
    Object{
      _ref: ref,
      r:    c.Network(),
    },
  }
}

//Invoke the same-named function on the network resource in WAPI,
//returning an array of available IP addresses.
//You may optionally specify how many IPs you want (num) and which ones to
//exclude from consideration (array of IPv4 addrdess strings).
func (n NetworkObject) NextAvailableIP(num int, exclude []string) (map[string]interface{}, error) {
  if num == 0 {
    num = 1
  }

  v := url.Values{}
  if exclude != nil {
    v.Set("exclude", strings.Join(exclude, ","))
  }
  v.Set("num", strconv.Itoa(num))

  out, err := n.FunctionCall("next_available_ip", v)
  if err != nil {
    return nil, fmt.Errorf("Error sending request: %v\n", err)
  }
  return out, nil
}
