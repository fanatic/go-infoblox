// Implements an Infoblox DNS/DHCP appliance client library in Go
package infoblox

import (
  "crypto/tls"
  "fmt"
  "log"
  "net/http"
  "net/url"
  "os"
  "strings"
)

var (
  WAPI_VERSION = "1.4.1"
  BASE_PATH    = "/wapi/v" + WAPI_VERSION + "/"
  DEBUG        = false
)

// Implements a Infoblox WAPI client.
// https://192.168.2.200/wapidoc/#transport-and-authentication
type Client struct {
  Host       string
  Password   string
  Username   string
  HttpClient *http.Client
}

// Creates a new Infoblox client with the supplied user/pass configuration.
// Supports the use of HTTP proxies through the $HTTP_PROXY env var.
// For example:
//     export HTTP_PROXY=http://localhost:8888
//
// When using a proxy, disable TLS certificate verification with the following:
//     export TLS_INSECURE=1
func NewClient(host, username, password string) *Client {
  var (
    req, _    = http.NewRequest("GET", host, nil)
    proxy, _  = http.ProxyFromEnvironment(req)
    transport *http.Transport
    tlsconfig *tls.Config
  )
  tlsconfig = &tls.Config{
    InsecureSkipVerify: os.Getenv("TLS_INSECURE") != "",
  }
  if tlsconfig.InsecureSkipVerify {
    log.Printf("WARNING: SSL cert verification  disabled\n")
  }
  transport = &http.Transport{
    TLSClientConfig: tlsconfig,
  }
  if proxy != nil {
    transport.Proxy = http.ProxyURL(proxy)
  }
  return &Client{
    Host: host,
    HttpClient: &http.Client{
      Transport: transport,
    },
    Username: username,
    Password: password,
  }
}

// Sends a HTTP request through this instance's HTTP client.
func (c *Client) SendRequest(req *http.Request) (resp *APIResponse, err error) {
  u := req.URL.String()
  if !strings.HasPrefix(u, "http") {
    u = fmt.Sprintf("%v%v", c.Host, u)
    req.URL, err = url.Parse(u)
    if err != nil {
      return
    }
  }
  req.SetBasicAuth(c.Username, c.Password)
  var r *http.Response
  r, err = c.HttpClient.Do(req)
  resp = (*APIResponse)(r)
  return
}
