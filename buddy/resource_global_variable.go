package buddy

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGlobalVariableCreate,
		ReadContext:   resourceGlobalVariableRead,
		UpdateContext: resourceGlobalVariableUpdate,
		DeleteContext: resourceGlobalVariableDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Variable name",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Variable value",
				Sensitive:   true,
			},
			"value_hash": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Hash of the encrypted variable",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "VAR",
				Description: "Variable type. Can be VAR, SSH_KEY, or FILE",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Variable description",
			},
			"settable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Flag to decide whether the variable is settable by pipeline run",
			},
			"encrypted": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Flag to check whether variable encrypted",
			},
			"ssh_key": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag to check whether the variable is SSH key",
			},
		},
	}
}

func resourceGlobalVariableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)

	key := d.Get("key").(string)
	value := d.Get("value").(string)
	varType := d.Get("type").(string)
	description := d.Get("description").(string)
	settable := d.Get("settable").(bool)
	encrypted := d.Get("encrypted").(bool)

	globalVar, err := client.CreateGlobalVariable(key, value, varType, description, settable, encrypted)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(globalVar.Id))
	return resourceGlobalVariableRead(ctx, d, m)
}

func resourceGlobalVariableRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)

	id := d.Id()
	data, err := client.ReadGlobalVariable(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("key", data.Key); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("ssh_key", data.SSHKey); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("settable", data.Settable); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("encrypted", data.Encrypted); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", data.Description); err != nil {
		return diag.FromErr(err)
	}

	if data.Encrypted {
		if err := d.Set("value_hash", data.Value); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("value", data.Value); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGlobalVariableUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)

	id := d.Id()
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	varType := d.Get("type").(string)
	description := d.Get("description").(string)
	settable := d.Get("settable").(bool)
	encrypted := d.Get("encrypted").(bool)

	_, err := client.UpdateGlobalVariable(id, key, value, varType, description, settable, encrypted)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGlobalVariableRead(ctx, d, m)
}

func resourceGlobalVariableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)

	id := d.Id()
	err := client.DeleteGlobalVariable(id)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
