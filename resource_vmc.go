package main

import (
	"context"
	"time"

	hcx "github.com/adeleporte/terraform-provider-hcx/hcx"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVmc() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVmcCreate,
		ReadContext:   resourceVmcRead,
		UpdateContext: resourceVmcUpdate,
		DeleteContext: resourceVmcDelete,

		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("VMC_API_TOKEN", nil),
			},
			"sddc_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVmcCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//client := m.(*hcx.Client)

	token := d.Get("token").(string)
	sddc_name := d.Get("sddc_name").(string)

	// Authenticate with VMware Cloud Services
	access_token, err := hcx.VmcAuthenticate(token)
	if err != nil {
		return diag.FromErr(err)
	}

	hcx_auth, err := hcx.HcxCloudAuthenticate(access_token)
	if err != nil {
		return diag.FromErr(err)
	}

	sddc, err := hcx.GetSddc(hcx_auth, sddc_name)
	if err != nil {
		return diag.FromErr(err)
	}

	// Check if already activated
	if sddc.DeploymentStatus == "ACTIVE" {
		return diag.Errorf("Already activated")
	}

	// Activate HCX
	_, err = hcx.ActivateHcxOnSDDC(hcx_auth, sddc.ID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for task to be completed
	for {
		sddc, err := hcx.GetSddc(hcx_auth, sddc_name)
		if err != nil {
			return diag.FromErr(err)
		}

		if sddc.DeploymentStatus == "ACTIVE" {
			break
		}

		if sddc.DeploymentStatus == "ACTIVATION_FAILED" {
			return diag.Errorf("Activation failed")
		}

		time.Sleep(10 * time.Second)
	}

	return resourceVmcRead(ctx, d, m)
}

func resourceVmcRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	token := d.Get("token").(string)
	sddc_name := d.Get("sddc_name").(string)

	// Authenticate with VMware Cloud Services
	access_token, err := hcx.VmcAuthenticate(token)
	if err != nil {
		return diag.FromErr(err)
	}

	hcx_auth, err := hcx.HcxCloudAuthenticate(access_token)
	if err != nil {
		return diag.FromErr(err)
	}

	sddc, err := hcx.GetSddc(hcx_auth, sddc_name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sddc.ID)
	d.Set("cloud_url", sddc.CloudURL)
	d.Set("cloud_name", sddc.CloudName)
	d.Set("cloud_type", sddc.CloudType)

	return diags
}

func resourceVmcUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	return resourceVmcRead(ctx, d, m)
}

func resourceVmcDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	token := d.Get("token").(string)
	sddc_name := d.Get("sddc_name").(string)

	// Authenticate with VMware Cloud Services
	access_token, err := hcx.VmcAuthenticate(token)
	if err != nil {
		return diag.FromErr(err)
	}

	hcx_auth, err := hcx.HcxCloudAuthenticate(access_token)
	if err != nil {
		return diag.FromErr(err)
	}

	sddc, err := hcx.GetSddc(hcx_auth, sddc_name)
	if err != nil {
		return diag.FromErr(err)
	}

	// Deactivate HCX
	_, err = hcx.DeactivateHcxOnSDDC(hcx_auth, sddc.ID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for task to be completed
	for {
		sddc, err := hcx.GetSddc(hcx_auth, sddc_name)
		if err != nil {
			return diag.FromErr(err)
		}

		if sddc.DeploymentStatus == "DE-ACTIVATED" {
			break
		}

		if sddc.DeploymentStatus == "DEACTIVATION_FAILED" {
			return diag.Errorf("Deactivation failed")
		}

		time.Sleep(10 * time.Second)
	}

	return diags
}
