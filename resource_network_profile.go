package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	hcx "github.com/adeleporte/terraform-provider-hcx/hcx"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkProfileCreate,
		ReadContext:   resourceNetworkProfileRead,
		UpdateContext: resourceNetworkProfileUpdate,
		DeleteContext: resourceNetworkProfileDelete,

		Schema: map[string]*schema.Schema{
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"prefix_length": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"primary_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"secondary_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"dns_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"network_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"start_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"end_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNetworkProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*hcx.Client)

	mtu := d.Get("mtu").(int)
	prefix_length := d.Get("prefix_length").(int)

	name := d.Get("name").(string)
	gateway := d.Get("gateway").(string)
	vcuuid := d.Get("vcenter").(string)
	primary_dns := d.Get("primary_dns").(string)
	secondary_dns := d.Get("secondary_dns").(string)
	dns_suffix := d.Get("dns_suffix").(string)

	network_name := d.Get("network_name").(string)
	network_id, err := hcx.GetNetworkBacking(client, vcuuid, network_name)
	if err != nil {
		return diag.FromErr(err)
	}

	start_address := d.Get("start_address").(string)
	end_address := d.Get("end_address").(string)

	body := hcx.InsertNetworkProfileBody{
		Name:       name,
		Enterprise: "DEFAULT",
		MTU:        mtu,
		Backings: []hcx.Backing{hcx.Backing{
			BackingID:           network_id.EntityID,
			BackingName:         network_name,
			VCenterInstanceUuid: vcuuid,
			Type:                "DistributedVirtualPortgroup",
		},
		},
		IPScopes: []hcx.IPScope{
			hcx.IPScope{
				DnsSuffix:    dns_suffix,
				Gateway:      gateway,
				PrefixLength: prefix_length,
				PrimaryDns:   primary_dns,
				SecondaryDns: secondary_dns,
				NetworkIpRanges: []hcx.NetworkIpRange{hcx.NetworkIpRange{
					StartAddress: start_address,
					EndAddress:   end_address,
				},
				},
			},
		},
		L3TenantManaged: false,
		OwnedBySystem:   true,
	}

	res, err := hcx.InsertNetworkProfile(client, body)

	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for job completion
	for {
		jr, err := hcx.GetJobResult(client, res.Data.JobID)
		if err != nil {
			return diag.FromErr(err)
		}

		if jr.IsDone {
			break
		}
		time.Sleep(5 * time.Second)
	}

	return resourceNetworkProfileRead(ctx, d, m)
}

func resourceNetworkProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*hcx.Client)
	name := d.Get("name").(string)

	id, err := hcx.GetNetworkProfile(client, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return diags
}

func resourceNetworkProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*hcx.Client)

	mtu := d.Get("mtu").(int)
	prefix_length := d.Get("prefix_length").(int)

	name := d.Get("name").(string)
	gateway := d.Get("gateway").(string)
	vcuuid := d.Get("vcenter").(string)
	primary_dns := d.Get("primary_dns").(string)
	secondary_dns := d.Get("secondary_dns").(string)
	dns_suffix := d.Get("dns_suffix").(string)

	network_name := d.Get("network_name").(string)
	network_id, err := hcx.GetNetworkBacking(client, vcuuid, network_name)
	if err != nil {
		return diag.FromErr(err)
	}

	start_address := d.Get("start_address").(string)
	end_address := d.Get("end_address").(string)

	body := hcx.UpdateNetworkProfileBody{
		ObjectId:   d.Id(),
		Name:       name,
		Enterprise: "DEFAULT",
		MTU:        mtu,
		Backings: []hcx.Backing{hcx.Backing{
			BackingID:           network_id.EntityID,
			BackingName:         network_name,
			VCenterInstanceUuid: vcuuid,
			Type:                "DistributedVirtualPortgroup",
		},
		},
		IPScopes: []hcx.IPScope{
			hcx.IPScope{
				DnsSuffix:    dns_suffix,
				Gateway:      gateway,
				PrefixLength: prefix_length,
				PrimaryDns:   primary_dns,
				SecondaryDns: secondary_dns,
				NetworkIpRanges: []hcx.NetworkIpRange{hcx.NetworkIpRange{
					StartAddress: start_address,
					EndAddress:   end_address,
				},
				},
			},
		},
		L3TenantManaged: false,
		OwnedBySystem:   true,
	}

	res, err := hcx.UpdateNetworkProfile(client, body)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	if err != nil {
		return diag.FromErr(errors.New(fmt.Sprintf("%s", buf)))
	}

	// Wait for job completion
	for {
		jr, err := hcx.GetJobResult(client, res.Data.JobID)
		if err != nil {
			return diag.FromErr(err)
		}

		if jr.IsDone {
			break
		}
		time.Sleep(5 * time.Second)
	}

	return resourceNetworkProfileRead(ctx, d, m)
}

func resourceNetworkProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*hcx.Client)
	res, err := hcx.DeleteNetworkProfile(client, d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for job completion
	for {
		jr, err := hcx.GetJobResult(client, res.Data.JobID)
		if err != nil {
			return diag.FromErr(err)
		}

		if jr.IsDone {
			break
		}
		time.Sleep(5 * time.Second)
	}

	return diags
}
