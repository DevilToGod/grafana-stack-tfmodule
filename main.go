package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"
)

type CreateStackRequest struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Region string `json:"region"`
}

type CreateStackResponse struct {
	// ID field only returned after Grafana v7.
	URL       string `json:"url"`
	ClusterID string `json:"clusterId"`
}

func main() {

	plugin.Serve(&plugin.ServeOpts{ProviderFunc: cmdb.Provider})
}
