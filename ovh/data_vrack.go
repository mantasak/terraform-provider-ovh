package ovh

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVrack() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVrackRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Here come all the computed items
			"result": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceVrackRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	escapedServiceName := url.PathEscape(serviceName)
	vrack := &Vrack{}
	err := config.OVHClient.Get(
		fmt.Sprintf(
			"/vrack/%s",
			escapedServiceName,
		),
		&vrack,
	)

	if err != nil {
		return fmt.Errorf("Error calling /vrack/%s:\n\t %q", escapedServiceName, err)
	}

	d.SetId(serviceName)
	d.Set("result", vrack)
	return nil
}
