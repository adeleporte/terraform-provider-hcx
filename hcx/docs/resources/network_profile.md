# network_profile

The Network Profile is an abstraction of a Distributed Port group, Standard Port group, or NSX Logical Switch, and the Layer 3 properties of that network. A Network Profile is a subcomponent of a complete Compute Profile.

Create a Network Profile for each network you intend to use with the HCX services. The extension selects these network profiles when creating a Compute Profile and assigned one or more of four Network Profile functions.


## Example Usage

```hcl
resource "hcx_network_profile" "net_management" {
  vcenter       = hcx_site_pairing.site1.local_vc
  network_name  = "HCX-Management-RegionA01"
  name          = "HCX-Management-RegionA01-profile"
  mtu           = 1500

  start_address   = "192.168.110.151"
  end_address     = "192.168.110.155"

  gateway           = "192.168.110.1"
  prefix_length     = 24
  primary_dns       = "192.168.110.10"
  secondary_dns     = ""
  dns_suffix        = "corp.local"
}

output "net_management" {
    value = hcx_network_profile.net_management
}

```

## Argument Reference

* `vcenter` - (Required) Local vCenter Id.
* `network_name` - (Required) Network Name used for this profile.
* `name` - (Required) Name of the network profile.
* `mtu` - (Required) MTU of the network profile.
* `start_address` - (Required) Start address of the IP Pool for this network profile.
* `end_address` - (Required) End address of the IP Pool for this network profile.
* `gateway` - (Optional) Gateway for this network profile.
* `prefix_length` - (Required) Prefix Length for this network profile.
* `primary_dns` - (Optional) Primary DNS for this network profile.
* `secondary_dns` - (Optional) Secondary DNS for this network profile.
* `dns_suffix` - (Optional) DNS suffix for this network profile.

## Attribute Reference

* `id` - ID of the network profile.
