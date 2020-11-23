package main

import (
	"context"

	hcx "github.com/adeleporte/terraform-provider-hcx/hcx"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcevCenter() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcevCenterCreate,
		ReadContext:   resourcevCenterRead,
		UpdateContext: resourcevCenterUpdate,
		DeleteContext: resourcevCenterDelete,

		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcevCenterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*hcx.Client)

	url := d.Get("url").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	body := hcx.InsertvCenterBody{
		Data: hcx.InsertvCenterData{
			Items: []hcx.InsertvCenterDataItem{
				hcx.InsertvCenterDataItem{
					Config: hcx.InsertvCenterDataItemConfig{
						Username: username,
						Password: password,
						URL:      url,
					},
				},
			},
		},
	}

	res, err := hcx.InsertvCenter(client, body)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(res.InsertvCenterData.Items[0].Config.URL)

	return resourcevCenterRead(ctx, d, m)
}

func resourcevCenterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourcevCenterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	return resourcevCenterRead(ctx, d, m)
}

func resourcevCenterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
