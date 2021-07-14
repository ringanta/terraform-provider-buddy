package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProjectMember() *schema.Resource {
	return &schema.Resource{
		Description: "`buddy_project_member` manages member on a Buddy project.\n\n" +
			"Member is granted access to a project using their ID and permission sets ID.",

		CreateContext: resourceProjectMemberCreate,
		ReadContext:   resourceProjectMemberRead,
		UpdateContext: resourceProjectMemberUpdate,
		DeleteContext: resourceProjectMemberDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"project_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Project name",
			},
			"member_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Member ID",
			},
			"permission_set_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "ID of permission set that will be granted to the user",
			},
		},
	}
}

func resourceProjectMemberCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)

	projectName := d.Get("project_name").(string)
	memberId := d.Get("member_id").(string)
	permissionSetId := d.Get("permission_set_id").(int)
	variable := buddyRequestProjectMember{
		Id: memberId,
		PermissionSet: buddyId{
			Id: permissionSetId,
		},
	}

	_, err := client.CreateProjectMember(projectName, variable)
	if err != nil {
		return diag.FromErr(err)
	}

	id := fmt.Sprintf("%v:%v", projectName, memberId)

	d.SetId(id)
	return resourceProjectMemberRead(ctx, d, m)
}

func resourceProjectMemberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)
	ids := strings.Split(d.Id(), ":")

	member, err := client.ReadProjectMember(ids[0], ids[1])
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("project_name", ids[0]); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("member_id", strconv.Itoa(member.Id)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("permission_set_id", member.PermissionSet.Id); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceProjectMemberUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)
	ids := strings.Split(d.Id(), ":")
	permissionSetId := d.Get("permission_set_id").(int)
	variable := buddyRequestPermissionSet{
		PermissionSet: buddyId{
			Id: permissionSetId,
		},
	}

	_, err := client.UpdateProjectMember(ids[0], ids[1], variable)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceProjectMemberRead(ctx, d, m)
}

func resourceProjectMemberDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(buddyClient)
	ids := strings.Split(d.Id(), ":")

	err := client.DeleteProjectMember(ids[0], ids[1])
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
