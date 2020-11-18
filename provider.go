package main

import (
	"context"

	hcx "github.com/adeleporte/terraform-provider-hcx/hcx"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	//	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hcx": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HCX_URL", nil),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("HCX_USER", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("HCX_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hcx_site_pairing":    resourceSitePairing(),
			"hcx_network_profile": resourceNetworkProfile(),
			"hcx_compute_profile": resourceComputeProfile(),
			"hcx_service_mesh":    resourceServiceMesh(),
			"hcx_l2_extension":    resourceL2Extension(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"hcx_network_backing": dataSourceNetworkBacking(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	hcxurl := d.Get("hcx").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	if (hcxurl != "") && (username != "") && (password != "") {
		c, err := hcx.NewClient(&hcxurl, &username, &password)
		//c := &http.Client{Timeout: 10 * time.Second}

		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	return nil, diag.Errorf("Missing credentials")
}