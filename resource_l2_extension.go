package main

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	hcx "github.com/adeleporte/terraform-provider-hcx/hcx"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceL2Extension() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceL2ExtensionCreate,
		ReadContext:   resourceL2ExtensionRead,
		UpdateContext: resourceL2ExtensionUpdate,
		DeleteContext: resourceL2ExtensionDelete,

		Schema: map[string]*schema.Schema{
			"site_pairing": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
			},
			"service_mesh_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"destination_t1": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceL2ExtensionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*hcx.Client)

	site_pairing := d.Get("site_pairing").(map[string]interface{})
	vcGuid := site_pairing["local_vc"].(string)

	//service_mesh := d.Get("service_mesh").(map[string]interface{})
	source_network := d.Get("source_network").(string)
	destination_t1 := d.Get("destination_t1").(string)
	gateway := d.Get("gateway").(string)
	netmask := d.Get("netmask").(string)

	destination_endpoint_id := site_pairing["id"].(string)
	destination_endpoint_name := site_pairing["remote_name"].(string)
	destination_endpoint_type := site_pairing["remote_endpoint_type"].(string)

	destination_resource_id := site_pairing["remote_resource_id"].(string)
	destination_resource_name := site_pairing["remote_resource_name"].(string)
	destination_resource_type := site_pairing["remote_resource_type"].(string)

	dvpg, err := hcx.GetNetworkBacking(client, vcGuid, source_network)
	if err != nil {
		return diag.FromErr(err)
	}

	appliance, err := hcx.GetAppliance(client, site_pairing["local_endpoint_id"].(string))
	if err != nil {
		return diag.FromErr(err)
	}

	body := hcx.InsertL2ExtensionBody{
		Gateway: gateway,
		Netmask: netmask,
		DestinationNetwork: hcx.DestinationNetwork{
			GatewayId: destination_t1,
		},
		Dns: []string{},
		Features: hcx.Features{
			EgressOptimization: false,
		},
		SourceAppliance: hcx.SourceAppliance{
			ApplianceId: appliance.ApplianceId,
		},
		SourceNetwork: hcx.SourceNetwork{
			NetworkId:   dvpg.EntityID,
			NetworkName: dvpg.Name,
			NetworkType: "dvpg",
		},
		VcGuid: vcGuid,
		Destination: hcx.Destination{
			EndpointId:   destination_endpoint_id,
			EndpointName: destination_endpoint_name,
			EndpointType: destination_endpoint_type,
			ResourceId:   destination_resource_id,
			ResourceName: destination_resource_name,
			ResourceType: destination_resource_type,
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	res2, err := hcx.InsertL2Extension(client, body)

	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for job completion
	for {
		jr, err := hcx.GetJobResult(client, res2.ID)
		if err != nil {
			return diag.FromErr(err)
		}

		if jr.IsDone {
			break
		}
		time.Sleep(5 * time.Second)
	}

	// Get L2 Extension ID
	res3, err := hcx.GetL2Extensions(client, dvpg.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(res3.StretchId)

	return resourceL2ExtensionRead(ctx, d, m)

}

func resourceL2ExtensionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceL2ExtensionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	return resourceL2ExtensionRead(ctx, d, m)
}

func resourceL2ExtensionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*hcx.Client)

	res, err := hcx.DeleteL2Extension(client, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for job completion
	for {
		jr, err := hcx.GetJobResult(client, res.ID)
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