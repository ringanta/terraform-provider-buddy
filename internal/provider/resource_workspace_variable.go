package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWorkspaceVariable() *schema.Resource {
	return &schema.Resource{
		Description: "`buddy_workspace_variable` manages variable under the workspace scope.\n\n" +
			"Variable under the workspace scoped is accessible by all projects under the workspace. " +
			"Use this variable to distributed the same variable that is used by multiple projects.",

		CreateContext: resourceWorkspaceVariableCreate,
		ReadContext:   resourceWorkpaceVariableRead,
		UpdateContext: resourceWorkspaceVariableUpdate,
		DeleteContext: resourceWorkpaceVariableDelete,
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
				Description: "Hash of the encrypted variable value",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "VAR",
				Description: "Variable type. Currently only support VAR",
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
				Description: "Flag to decide whether variable encrypted",
			},
			"ssh_key": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag to decide whether the variable is an SSH key",
			},
		},
	}
}

func resourceWorkspaceVariableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)

	key := d.Get("key").(string)
	value := d.Get("value").(string)
	varType := d.Get("type").(string)
	description := d.Get("description").(string)
	settable := d.Get("settable").(bool)
	encrypted := d.Get("encrypted").(bool)
	variable := buddyRequestWorkspaceVariable{
		Key:         key,
		Value:       value,
		Type:        varType,
		Description: description,
		Settable:    settable,
		Encrypted:   encrypted,
	}

	globalVar, err := client.CreateWorkspaceVariable(variable)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(globalVar.Id))
	return resourceWorkpaceVariableRead(ctx, d, m)
}

func resourceWorkpaceVariableRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)

	id := d.Id()
	data, err := client.ReadWorkspaceVariable(id)
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

	if err := d.Set("value_hash", data.Value); err != nil {
		return diag.FromErr(err)
	}

	if !data.Encrypted {
		if err := d.Set("value", data.Value); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceWorkspaceVariableUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)

	id := d.Id()
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	varType := d.Get("type").(string)
	description := d.Get("description").(string)
	settable := d.Get("settable").(bool)
	encrypted := d.Get("encrypted").(bool)
	variable := buddyRequestWorkspaceVariable{
		Key:         key,
		Value:       value,
		Type:        varType,
		Description: description,
		Settable:    settable,
		Encrypted:   encrypted,
	}

	_, err := client.UpdateWorkspaceVariable(id, variable)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceWorkpaceVariableRead(ctx, d, m)
}

func resourceWorkpaceVariableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)

	id := d.Id()
	err := client.DeleteVariable(id)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
