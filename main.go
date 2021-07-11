package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/ringanta/terraform-provider-buddy/buddy"
)

// Run terraform fmt to format Terraform project under the examples folder
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool,
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return buddy.Provider()
		},
	})
}
