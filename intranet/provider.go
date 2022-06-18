package intranet

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kamilturek/intranet"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"session_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTRANET_SESSION_ID", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"intranet_hour_entry": resourceHourEntry(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	sessionId := d.Get("session_id").(string)
	c := intranet.NewClient(sessionId)
	return c, nil
}
