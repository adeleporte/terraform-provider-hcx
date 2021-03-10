# compute_profile

A Compute Profile contains the compute, storage, and network settings that HCX uses on this site to deploy the Interconnect-dedicated virtual appliances when a Service Mesh is added.

Create a Compute Profile in the Multi-Site Service Mesh interface in both the source and the destination HCX environments using the planned configuration options for each site, respectively.



## Example Usage

```hcl
resource "hcx_compute_profile" "compute_profile_1" {
  name                  = "comp1"
  datacenter            = "RegionA01-ATL"
  cluster               = "RegionA01-COMP01"
  datastore             = "RegionA01-ISCSI01-COMP01"

  management_network    = hcx_network_profile.net_management.id
  replication_network   = hcx_network_profile.net_management.id
  uplink_network        = hcx_network_profile.net_uplink.id
  vmotion_network       = hcx_network_profile.net_vmotion.id
  dvs                   = "RegionA01-vDS-COMP"

  service {
    name                = "INTERCONNECT"
  }

  service {
    name                = "WANOPT"
  }

  service {
    name                = "VMOTION"
  }

  service {
    name                = "BULK_MIGRATION"
  }

  service {
    name                = "RAV"
  }

  service {
    name                = "NETWORK_EXTENSION"
  }

  service {
    name                = "DISASTER_RECOVERY"
  }

  service {
    name                = "SRM"
  }

}

output "compute_profile_1" {
    value = hcx_compute_profile.compute_profile_1
}

```

## Argument Reference

* `name` - (Required) Name of the compute profile.
* `datacenter` - (Required) Datacenter where HCX Services will be available.
* `cluster` - (Required) Cluster used for HCX appliances deployment.
* `datastore` - (Required) Datastore used for HCX appliances deployment.
* `management_network` - (Required) Management network profile. (ID)
* `replication_network` - (Required) Replication network profile. (ID)
* `vmotion_network` - (Required) vMotion network profile. (ID)
* `uplink_network` - (Required) Uplink network profile. (ID)
* `dvs` - (Required) DVS used for L2 extension.
* `service` - (Required) List of HCX services.

### Service argument Reference
* `name` - (Required) Name of the HCX service. Value values are: `INTERCONNECT`, `WANOPT`, `VMOTION`, `BULK_MIGRATION`, `RAV`, `NETWORK_EXTENSION`, `DISASTER_RECOVERY`, `SRM`

## Attribute Reference

* `id` - ID of the compute profile.
