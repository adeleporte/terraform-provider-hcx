---
page_title: "HCX Lab - Full HCX Connector configuration"
---

This TF file automates the configuration of a HCX lab. It manages creation/update/deletion of :
* Site pairing
* Network profiles (Management, vMotion and Uplink)
* Compute profile
* Service Mesh
* L2 Extension

## Usage example

```hcl

terraform {
  required_providers {
    hcx = {
      versions = ["0.1"]
      source = "vcn.cloud/edu/hcx"
    }
  }
}

provider hcx {
    hcx         = "https://192.168.110.70"
    username    = "administrator@vsphere.local"
    password    = "VMware1!"
}

resource "hcx_site_pairing" "site1" {
    url         = "https://hcx-cloud-01b.corp.local"
    username    = "administrator@vsphere.local"
    password    = "VMware1!"
}

resource "hcx_network_profile" "net_management" {
  vcenter           = hcx_site_pairing.site1.local_vc
  network_name      = "HCX-Management-RegionA01"
  name              = "HCX-Management-RegionA01-profile"
  mtu               = 1500

  start_address     = "192.168.110.151"
  end_address       = "192.168.110.155"

  gateway           = "192.168.110.1"
  prefix_length     = 24
  primary_dns       = "192.168.110.10"
  secondary_dns     = ""
  dns_suffix        = "corp.local"
}



resource "hcx_network_profile" "net_uplink" {
  vcenter           = hcx_site_pairing.site1.local_vc
  network_name      = "HCX-Uplink-RegionA01"
  name              = "HCX-Uplink-RegionA01-profile"
  mtu               = 1600


  start_address     = "192.168.110.156"
  end_address       = "192.168.110.160"


  gateway           = "192.168.110.1"
  prefix_length     = 24
  primary_dns       = "192.168.110.1"
  secondary_dns     = ""
  dns_suffix        = "corp.local"
}

resource "hcx_network_profile" "net_vmotion" {
  vcenter           = hcx_site_pairing.site1.local_vc
  network_name      = "HCX-vMotion-RegionA01"
  name              = "HCX-vMotion-RegionA01-profile"
  mtu               = 1500

  start_address     = "10.10.30.151"
  end_address       = "10.10.30.155"


  gateway           = ""
  prefix_length     = 24
  primary_dns       = ""
  secondary_dns     = ""
  dns_suffix        = ""
}



resource "hcx_compute_profile" "compute_profile_1" {
  name                  = "comp1"
  datacenter            = "RegionA01-ATL"
  cluster               = "RegionA01-COMP01"
  datastore             = "RegionA01-ISCSI01-COMP01"

  management_network    = hcx_network_profile.net_management
  replication_network   = hcx_network_profile.net_management
  uplink_network        = hcx_network_profile.net_uplink
  vmotion_network       = hcx_network_profile.net_vmotion
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

resource "hcx_l2_extension" "l2_extension_1" {
  site_pairing          = hcx_site_pairing.site1
  service_mesh_name     = hcx_service_mesh.service_mesh_1.name
  source_network        = "VM-RegionA01-vDS-COMP"

  destination_t1        = "T1-GW"
  gateway               = "2.2.2.2"
  netmask               = "255.255.255.0"

}


```
