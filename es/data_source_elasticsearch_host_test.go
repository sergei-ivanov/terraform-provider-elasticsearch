package es

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccElasticsearchDataSourceHost_basic(t *testing.T) {
	var providers []*schema.Provider
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories(&providers),
		Steps: []resource.TestStep{
			{
				Config: testAccElasticsearchDataSourceHost,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.elasticsearch_host.test", "id"),
				),
			},
		},
	})
}

var testAccElasticsearchDataSourceHost = `
data "elasticsearch_host" "test" {
  active = true
}
`
