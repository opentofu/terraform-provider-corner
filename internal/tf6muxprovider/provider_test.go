// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf6muxprovider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccResourceUser(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"tf6muxprovider": func() (tfprotov6.ProviderServer, error) {
				provider, err := New()

				return provider(), err
			},
		},
		Steps: []resource.TestStep{
			{
				Config: configResourceUserBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tf6muxprovider_user1.example", "age", "123"),
					resource.TestCheckResourceAttr("tf6muxprovider_user1.example", "email", "example1@example.com"),
					resource.TestCheckResourceAttr("tf6muxprovider_user1.example", "language", "en"),
					resource.TestCheckResourceAttr("tf6muxprovider_user1.example", "name", "Example Name 1"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "age", "234"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "email", "example2@example.com"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "language", "en"),
					resource.TestCheckResourceAttr("tf6muxprovider_user2.example", "name", "Example Name 2"),
				),
			},
		},
	})
}

const configResourceUserBasic = `
resource "tf6muxprovider_user1" "example" {
  age   = 123
  email = "example1@example.com"
  name  = "Example Name 1"
}

resource "tf6muxprovider_user2" "example" {
  age   = 234
  email = "example2@example.com"
  name  = "Example Name 2"
}
`

func TestAccFunctionString(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"tf6muxprovider": func() (tfprotov6.ProviderServer, error) {
				provider, err := New()

				return provider(), err
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `
				output "test1" {
					value = provider::tf6muxprovider::string1("str1")
				}

				output "test2" {
					value = provider::tf6muxprovider::string2("str2")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test1", knownvalue.StringExact("str1")),
					statecheck.ExpectKnownOutputValue("test2", knownvalue.StringExact("str2")),
				},
			},
		},
	})
}
