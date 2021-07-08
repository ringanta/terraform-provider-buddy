package buddy

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider - create a new Buddy provider
func Provider() *schema.Provider {
	return &schema.Provider{}
}
