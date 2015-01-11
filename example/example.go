package main

import (
  "fmt"
  "github.com/fanatic/go-infoblox"
  "net/http"
)

func main() {
  ib := infoblox.NewClient("https://192.168.2.200/", "admin", "infoblox")
  out, err := ib.Network().All(nil)
  if err != nil {
    fmt.Printf("%v\n", err)
    return
  }
  fmt.Printf("%v\n\n", out)

  maxResults := 1000
  opts := &infoblox.Options{
    MaxResults:   &maxResults,
    ReturnFields: []string{"network"},
  }

  out, err = ib.Network().All(opts)
  if err != nil {
    fmt.Printf("%v\n", err)
    return
  }
  fmt.Printf("%v\n\n", out)
}

func save() {
  ib := infoblox.NewClient("https://192.168.2.200/", "admin", "infoblox")

  req, err := http.NewRequest("GET", infoblox.BASE_PATH+"network?_return_fields=authority,bootfile,bootserver,comment,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,deny_bootp,disable,email_list,enable_ddns,enable_dhcp_thresholds,enable_email_warnings,enable_ifmap_publishing,enable_snmp_warnings,extensible_attributes,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ipv4addr,lease_scavenge_time,low_water_mark,low_water_mark_reset,members,netmask,network,network_container,network_view,nextserver,options,pxe_lease_time,recycle_leases,update_dns_on_lease_renewal,use_authority,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_ifmap_publishing,use_ignore_dhcp_option_list_request,use_lease_scavenge_time,use_nextserver,use_options,use_recycle_leases,use_update_dns_on_lease_renewal,use_zone_associations,zone_associations", nil)
  if err != nil {
    fmt.Printf("Error creating request: %v\n", err)
    return
  }

  resp, err := ib.SendRequest(req)
  if err != nil {
    fmt.Printf("Error sending request: %v\n", err)
    return
  }

  //fmt.Printf("%v", resp.ReadBody())

  var out []map[string]interface{}
  err = resp.Parse(&out)
  if err != nil {
    fmt.Printf("%+v\n", err)
    return
  }

  fmt.Printf("%v", out)
}
