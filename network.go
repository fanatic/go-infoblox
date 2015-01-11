package infoblox

import ()

func (c *Client) Network() *Resource {
  return &Resource{conn: c, child: &Network{}}
}

// https://192.168.2.200/wapidoc/objects/network.html
type Network struct {
}

func (n Network) WAPIObject() string {
  return "network"
}

//Invoke the same-named function on the network resource in WAPI,
//returning an array of available IP addresses.
//You may optionally specify how many IPs you want (num) and which ones to
//exclude from consideration (array of IPv4 addrdess strings).
// func (n Network) NextAvailableIP(num int, exclude []string) {
// if num == 0 {
// num = 1
// }
// body := map[string]string{
// "num":     strconv.Itoa(num),
// "exclude": exclude,
// }
// }

// ##
// # Invoke the same-named function on the network resource in WAPI,
// # returning an array of available IP addresses.
// # You may optionally specify how many IPs you want (num) and which ones to
// # exclude from consideration (array of IPv4 addrdess strings).
// #
// def next_available_ip(num=1, exclude=[])
//   post_body = {
//     num:     num.to_i,
//     exclude: exclude
//   }
//   JSON.parse(connection.post(resource_uri + "?_function=next_available_ip", post_body).body)["ips"]
// end
