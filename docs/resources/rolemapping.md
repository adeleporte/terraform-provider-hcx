# rolemapping

Assign the HCX Roles to the vCenter User Groups that are allowed to perform HCX operations.



## Example Usage

```hcl
resource "hcx_rolemapping" "rolemapping" {
    sso = hcx_sso.sso.id

    admin {
      user_group = "vsphere.local\\Administrators"
    }

    admin {
      user_group = "corp.local\\Administrators"
    }

    enterprise {
      user_group = "corp.local\\Administrators"
    }
}

```

## Argument Reference

* `sso` - (Required) ID of the SSO Lookup Service.
* `admin` - (Optional) Group List for Admin users.
* `enterpise` - (Optional) Group List for Enterprise users.

### Admin & Enterprise Argument Reference
* `user_group` - (Optional) Group name.

