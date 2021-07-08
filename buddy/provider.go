package buddy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider - create a new Buddy provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"buddy_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BUDDY_URL", nil),
				Description: "The URL to the Buddy workspace",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Default:     schema.EnvDefaultFunc("BUDDY_TOKEN", nil),
				Description: "Buddy personal access token",
			},
			"verify_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BUDDY_VERIFY_SSL", true),
				Description: "Whether to verify TLS connection to the Buddy URL",
			},
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
