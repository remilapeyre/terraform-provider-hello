package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccMessageDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMessageDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.hello_message.test", "message", "hello world"),
				),
			},
			{
				Config: testAccMessageDataSourceConfigTarget,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.hello_message.test", "target", "terraform"),
					resource.TestCheckResourceAttr("data.hello_message.test", "message", "hello terraform"),
				),
			},
		},
	})
}

const testAccMessageDataSourceConfig = `
data "hello_message" "test" {}
`

const testAccMessageDataSourceConfigTarget = `
data "hello_message" "test" {
  target = "terraform"
}
`
