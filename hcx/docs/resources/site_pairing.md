# site_pairing

A Site Pair establishes the connection needed for management, authentication, and orchestration of HCX services across a source and destination environment.

In HCX Connector to HCX Cloud deployments, the HCX Connector is deployed at the legacy or source vSphere environment. The HCX Connector creates a unidirectional site pairing to an HCX Cloud system. In this type of site pairing, all HCX Service Mesh connections, Migration and Network Extension operations, including reverse migrations, are always initiated from the HCX Connector at the source.


## Example Usage

```hcl
resource "hcx_site_pairing" "site1" {
    url         = "https://hcx-cloud-01b.corp.local"
    username    = "administrator@vsphere.local"
    password    = "VMware1!"
}

output "hcx_site_pairing_site1" {
    value = hcx_site_pairing.site1
}

```

## Argument Reference

* `url` - (Required) URL of the remote cloud.
* `username` - (Required) Username used for remote cloud authentication.
* `password` - (Required) Password used for remote cloud authentication.


## Attribute Reference

* `id` - ID of the site pairing.
* `local_vc` - ID of the local vCenter.
* `local_endpoint_id` - Endpoint ID of the local HCX site.
* `local_name` - Endpoint Name of the local HCX site.
* `remote_name` - Endpoint Name of the remote HCX site.
* `remote_endpoint_type` - Endpoint Type of the remote HCX site.
* `remote_resource_id` - Resource ID of the remote HCX site.
* `remote_resource_name` - Resource Name of the remote HCX site.
* `remote_resource_type` - Resource Type of the remote HCX site.
