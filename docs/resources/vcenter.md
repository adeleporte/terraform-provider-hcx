# vcenter

vCenter used by this HCX system.
If resource is created/updates, Application service is restarted.


## Example Usage

```hcl
resource "hcx_vcenter" "vcenter" {
    url         = "https://vcsa-01a.corp.local"
    username    = "administrator@vsphere.local"
    password    = "VMware1!"

    depends_on  = [hcx_activation.activation]
}

```

## Argument Reference

* `url` - (Required) URL of the vCenter.
* `username` - (Required) Username of the vCenter.
* `password` - (Required) Password of the vCenter.


## Attribute Reference

* `id` - UUID of the vcenter.
