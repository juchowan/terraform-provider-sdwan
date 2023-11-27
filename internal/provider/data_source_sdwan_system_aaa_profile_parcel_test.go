// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanSystemAAAProfileParcel(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSystemAAAProfileParcelConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "authentication_group", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "accounting_group", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "users.0.name", "User1"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "users.0.privilege", "15"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "users.0.public_keys.0.key_string", "AAAAB3NzaC1yc2"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "users.0.public_keys.0.key_type", "ssh-rsa"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.group_name", "RGROUP1"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.vpn", "10"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.source_interface", "GigabitEthernet0"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.servers.0.address", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.servers.0.auth_port", "1812"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.servers.0.acct_port", "1813"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.servers.0.timeout", "5"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.servers.0.retransmit", "3"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.servers.0.secret_key", "cisco123"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.servers.0.key_enum", "7"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "radius_groups.0.servers.0.key_type", "key"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "tacacs_groups.0.group_name", "TGROUP1"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "tacacs_groups.0.vpn", "10"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "tacacs_groups.0.source_interface", "GigabitEthernet0"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "tacacs_groups.0.servers.0.address", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "tacacs_groups.0.servers.0.port", "49"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "tacacs_groups.0.servers.0.timeout", "5"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "tacacs_groups.0.servers.0.secret_key", "cisco123"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "tacacs_groups.0.servers.0.key_enum", "7"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "accounting_rules.0.rule_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "accounting_rules.0.method", "commands"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "accounting_rules.0.level", "15"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "accounting_rules.0.start_stop", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "authorization_console", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "authorization_config_commands", "true"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "authorization_rules.0.rule_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "authorization_rules.0.method", "commands"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "authorization_rules.0.level", "15"),
					resource.TestCheckResourceAttr("data.sdwan_system_aaa_profile_parcel.test", "authorization_rules.0.if_authenticated", "true"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanSystemAAAProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_system_aaa_profile_parcel" "test" {
  name = "TF_TEST"
  description = "Terraform integration test"
  feature_profile_id = sdwan_system_feature_profile.test.id
  authentication_group = true
  accounting_group = true
  server_auth_order = ["local"]
  users = [{
    name = "User1"
    password = "cisco123"
    privilege = "15"
	public_keys = [{
		key_string = "AAAAB3NzaC1yc2"
		key_type = "ssh-rsa"
	}]
  }]
  radius_groups = [{
    group_name = "RGROUP1"
    vpn = 10
    source_interface = "GigabitEthernet0"
	servers = [{
		address = "1.2.3.4"
		auth_port = 1812
		acct_port = 1813
		timeout = 5
		retransmit = 3
		key = "cisco123"
		secret_key = "cisco123"
		key_enum = "7"
		key_type = "key"
	}]
  }]
  tacacs_groups = [{
    group_name = "TGROUP1"
    vpn = 10
    source_interface = "GigabitEthernet0"
	servers = [{
		address = "1.2.3.4"
		port = 49
		timeout = 5
		key = "cisco123"
		secret_key = "cisco123"
		key_enum = "7"
	}]
  }]
  accounting_rules = [{
    rule_id = "1"
    method = "commands"
    level = "15"
    start_stop = true
    group = ["RGROUP1"]
  }]
  authorization_console = true
  authorization_config_commands = true
  authorization_rules = [{
    rule_id = "1"
    method = "commands"
    level = "15"
    group = ["RGROUP1"]
    if_authenticated = true
  }]
}

data "sdwan_system_aaa_profile_parcel" "test" {
  id = sdwan_system_aaa_profile_parcel.test.id
  feature_profile_id = sdwan_system_feature_profile.test.id
}
`
