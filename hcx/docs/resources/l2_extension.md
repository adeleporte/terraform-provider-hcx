# l2_extension

You can bridge local network segments between HCX-enabled data centers with HCX Network Extension.

With VMware HCX Network Extension (HCX-NE), a High-Performance (4–6 Gbps) service, you can extend the Virtual Machine networks to a VMware HCX-enabled remote site. Virtual Machines that are migrated or created on the extended segment at the remote site are Layer 2 next to virtual machines placed on the origin network. Using Network Extension a remote site's resources can be quickly consumed. With Network Extension , the default gateway for the extended network only exists at the source site. Traffic from virtual machines (on remote extended networks) that must be routed returns to the source site gateway.



## Example Usage

```hcl
resource "hcx_l2_extension" "l2_extension_1" {
  site_pairing                    = hcx_site_pairing.site1
  service_mesh_name               = hcx_service_mesh.service_mesh_1.name
  source_network                  = "VM-RegionA01-vDS-COMP"

  destination_t1                  = "T1-GW"
  gateway                         = "2.2.2.2"
  netmask                         = "255.255.255.0"

}

output "l2_extension_1" {
    value = hcx_l2_extension.l2_extension_1
}

```

## Argument Reference

* `site_pairing` - (Required) Site pairing used by this service mesh.
* `service_mesh_name` - (Required) Name of the Service Mesh to be used for this L2 extension.
* `source_network` - (Required) Source Network. Must be a dvpg which is vlan tagged.
* `destination_t1` - (Required) Name of the T1 NSX-T router at destination.
* `gateway` - (Required) Gateway address to configure on the T1. Should be equal to the existing default gateway at source site.
* `netmask` - (Required) Netmask

## Attribute Reference

* `id` - ID of the L2 extension.
