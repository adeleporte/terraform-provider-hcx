# vmc

This resource manages the HCX activation / desactivation at the VMC side. When HCX is activated, it is also configured with appropriate network and compute profiles. Make sure that HCX appliances are reachable from the HCX connector for other resources to work (Firewall configuration).



## Example Usage

```hcl

resource "hcx_vmc" "vmc_nico" {  
    sddc_name   = "nvibert-VELOCLOUD"
}

resource "hcx_site_pairing" "vmc" {
    url         = hcx_vmc.vmc_nico.cloud_url
    username    = "cloudadmin@vmc.local"
    password    = var.vmc_vcenter_password
}



```

## Argument Reference

* `sddc_name` - (Optional) Name of the SDDC. If not specified, sddc_id must be set.
* `sddc_id` - (Optional) ID of the SDDC. If not specified, sddc_name must be set.
* `token` - (Required) VMware Cloud Service API Token. Generated from the VMware Cloud Services Console / My account / API Tokens. Environment variable VMC_API_TOKEN can be used to avoid setting the token in the code.



## Attribute Reference

* `id` - ID of the SDDC.
* `cloud_url` - URL of HCX Cloud. Use this attribute for the site pairing configuration.
* `cloud_type` - Type of cloud. Should be nsp for VMC.
* `cloud_name` - Name of the HCX Cloud.
