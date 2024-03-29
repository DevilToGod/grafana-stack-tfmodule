package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

/*var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"
)*/

func main() {
	/*var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: grafana.Provider(version)}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/grafana/grafana", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)*/
	var p func() *schema.Provider
	p = grafana.Provider("grafana")
	println(p)
}
