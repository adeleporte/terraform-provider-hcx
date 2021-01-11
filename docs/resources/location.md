# location

Select the nearest major city to where the HCX system is geographically located. HCX sites are represented visually in the Dashboard.


## Example Usage

```hcl
resource "hcx_location" "location" {
    city        = "Paris"
    country     = "France"
    province    = "Ile-de-France"
    latitude    = 48.86669293
    longitude   = 2.333335326
}

```

## Argument Reference

* `city` - (Optional) City of this HCX Site.
* `country` - (Optional) Country of this HCX Site.
* `province` - (Optional) Province of this HCX Site.
* `latitude` - (Optional) Latitude of this HCX Site.
* `longitude` - (Optional) Longitude of this HCX Site.

