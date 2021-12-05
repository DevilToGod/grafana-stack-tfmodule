package grafana

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	gapi "../client"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		return strings.TrimSpace(desc)
	}
}

func Provider(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"url": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("GRAFANA_URL", nil),
					Description: "The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.",
				},
				"apikey": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("GRAFANA_API_KEY", nil),
					Description: "API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.",
				},
			},
		}

		p.ConfigureContextFunc = configure(version, p)
		return p
	}
}

type client struct {
	gapi *gapi.Client
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics
		p.UserAgent("grafana-stack-tfmodule", version)

		c := &client{}

		smToken := d.Get("apikey").(string)
		smURL := d.Get("url").(string)
		c.gapi = gapi.NewClient(smURL, smToken, nil)

		return c, diags
	}
}
