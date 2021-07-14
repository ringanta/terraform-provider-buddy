package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWorkspaceMember() *schema.Resource {
	return &schema.Resource{
		Description: "`buddy_workspace_member` get information about workspace member",

		ReadContext: dataSourceWorkspaceMemberRead,

		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Email address of the member",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Member name",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Member ID",
			},
			"title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Member title",
			},
		},
	}
}

func dataSourceWorkspaceMemberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)
	email := d.Get("email").(string)

	member, err := client.GetUser(email)
	if err != nil {
		return diag.FromErr(err)
	}

	if member.Url == "" && member.Id == 0 {
		return diag.FromErr(fmt.Errorf("User not found: " + email))
	}

	if err := d.Set("email", member.Email); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", member.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("title", member.Title); err != nil {
		return diag.FromErr(err)
	}

	id := strconv.Itoa(member.Id)
	if err := d.Set("id", id); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)

	return nil
}
