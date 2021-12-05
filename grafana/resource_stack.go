package grafana

import (
	"context"

	gapi "../client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type InstanceName struct {
	ID    int64
	Email string
	Role  string
}

func ResourceStack() *schema.Resource {
	return &schema.Resource{

		Description: `
* [Official documentation](https://grafana.com/docs/grafana/latest/manage-users/server-admin/server-admin-manage-orgs/)
* [HTTP API](https://grafana.com/docs/grafana/latest/http_api/org/)
`,

		CreateContext: CreateStack,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The display name for the new stack",
			},
			"slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subdomain that the Grafana instance will be available at",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Choose a region for your stack. For example, you can specify the United States: us, or Europe: eu",
			},
		},
	}
}

func CreateStack(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	slug := d.Get("slug").(string)
	region := d.Get("region").(string)
	c := m.(*client).gapi

	request := gapi.CreateStackRequest{ctx}

	response, err := c.CreateStack(request)
	if err != nil {
		return diag.FromErr(err)
	}
	println(response.URL)

	return diag.Diagnostics{}
}
