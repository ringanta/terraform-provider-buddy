package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	user_agent string
)

// New - create a new Buddy provider
func New(version string) *schema.Provider {
	user_agent = "terraform-buddy-provider/" + version

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"buddy_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BUDDY_URL", nil),
				Description: "The URL to the Buddy workspace",
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BUDDY_TOKEN", nil),
				Description: "Buddy personal access token",
			},
			"verify_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BUDDY_VERIFY_SSL", true),
				Description: "Whether to verify TLS connection to the Buddy URL",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"buddy_workspace_variable": resourceWorkspaceVariable(),
			"buddy_workspace_member":   resourceWorkspaceMember(),
			"buddy_project_variable":   resourceProjectVariable(),
		},

		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		BuddyURL:  d.Get("buddy_url").(string),
		Token:     d.Get("token").(string),
		VerifySSL: d.Get("verify_ssl").(bool),
	}

	client := newBuddyClient(&config)
	return client, nil
}
