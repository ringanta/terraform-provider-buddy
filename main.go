package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/ringanta/terraform-provider-buddy/internal/provider"
)

// Run terraform fmt to format Terraform project under the examples folder
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool,
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"
)

func main() {
	opts := &plugin.ServeOpts{ProviderFunc: func() *schema.Provider {
		return provider.New(version)
	}}

	plugin.Serve(opts)
}
