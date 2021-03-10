# service_mesh

An HCX Service Mesh is the effective HCX services configuration for a source and destination site. A Service Mesh can be added to a connected Site Pair that has a valid Compute Profile created on both of the sites.

Adding a Service Mesh initiates the deployment of HCX Interconnect virtual appliances on both of the sites. An interconnect Service Mesh is always created at the source site.



## Example Usage

```hcl
resource "hcx_service_mesh" "service_mesh_1" {
  name                            = "sm1"
  site_pairing                    = hcx_site_pairing.site1
  local_compute_profile           = hcx_compute_profile.compute_profile_1.name
  remote_compute_profile          = "Compute-RegionB01"

  app_path_resiliency_enabled     = false
  tcp_flow_conditioning_enabled   = false

  uplink_max_bandwidth            = 10000

  service {
    name                = "INTERCONNECT"
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

}

output "service_mesh_1" {
    value = hcx_service_mesh.service_mesh_1
}

```

## Argument Reference

* `name` - (Required) Name of the service mesh.
* `site_pairing` - (Required) Site pairing used by this service mesh.
* `local_compute_profile` - (Required) Local Compute profile name.
* `remote_compute_profile` - (Required) Remote Compute profile name.
* `app_path_resiliency_enabled` - (Optional) Enable Application Path Resiliency feature. Default is `false`.
* `tcp_flow_conditioning_enabled` - (Optional) Enable TCP flow conditioning feature. Default is `false`.
* `uplink_max_bandwidth` - (Optional) Maximum bandwidth used for uplinks. Default is `10000`.
* `service` - (Required) List of HCX services. (Services selected here must be part of the compute profiles selected).
* `force_delete` - (Optional) Enable/Disable Force Delete of the Service Mesh. Sometimes need when site pairing is not connected anymore.
* `nb_appliances` - (Optional - Default is 1) Nb of NE appliances to deploy (each NE appliance can extend 8 networks)

### Service argument Reference
* `name` - (Required) Name of the HCX service. Value values are: `INTERCONNECT`, `WANOPT`, `VMOTION`, `BULK_MIGRATION`, `RAV`, `NETWORK_EXTENSION`, `DISASTER_RECOVERY`, `SRM`

## Attribute Reference

* `id` - ID of the Service Mesh.
