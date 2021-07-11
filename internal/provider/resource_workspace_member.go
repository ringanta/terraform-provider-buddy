package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWorkspaceMember() *schema.Resource {
	return &schema.Resource{
		Description: "`buddy_workspace_member` manages member on a Buddy workspace.\n\n" +
			"A new member can be invited into a workspace using their email address.",

		CreateContext: resourceWorkspaceMemberCreate,
		ReadContext:   resourceWorkspaceMemberRead,
		UpdateContext: resourceWorkspaceMemberUpdate,
		DeleteContext: resourceWorkspaceMemberDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Email address to be invited into Buddy workspace",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Member name",
			},
			"admin": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Flag to indicate whether member has admin right",
			},
		},
	}
}

func resourceWorkspaceMemberCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)

	email := d.Get("email").(string)
	admin := d.Get("admin").(bool)

	member, err := client.CreateWorkspaceMember(email)
	if err != nil {
		return diag.FromErr(err)
	}

	id := strconv.Itoa(member.Id)
	_, err = client.SetAdminRight(id, admin)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)
	return resourceWorkspaceMemberRead(ctx, d, m)
}

func resourceWorkspaceMemberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)
	id := d.Id()

	member, err := client.ReadWorkspaceMember(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("email", member.Email); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", member.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("admin", member.Admin); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceWorkspaceMemberUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)
	id := d.Id()
	admin := d.Get("admin").(bool)

	_, err := client.SetAdminRight(id, admin)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceWorkspaceMemberRead(ctx, d, m)
}

func resourceWorkspaceMemberDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)
	id := d.Id()

	err := client.DeleteWorkspaceMember(id)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
