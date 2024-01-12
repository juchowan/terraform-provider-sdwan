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

func TestAccSdwanSwitchportFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSwitchportFeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "slot", "0"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "sub_slot", "0"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "module_type", "4"),
				),
			},
			{
				Config: testAccSdwanSwitchportFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "slot", "0"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "sub_slot", "0"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "module_type", "4"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.name", "GigabitEthernet0/0/0"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.switchport_mode", "access"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.shutdown", "true"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.speed", "100"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.duplex", "full"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.switchport_access_vlan", "100"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.switchport_trunk_allowed_vlans", "100,200"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.switchport_trunk_native_vlan", "100"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_enable", "true"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_port_control", "auto"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_voice_vlan", "200"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_pae_enable", "true"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_mac_authentication_bypass", "true"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_host_mode", "multi-domain"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_enable_periodic_reauth", "true"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_periodic_reauth_inactivity_timeout", "100"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_periodic_reauth_interval", "60"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_control_direction", "both"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_restricted_vlan", "100"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_guest_vlan", "101"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_critical_vlan", "102"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "interfaces.0.dot1x_enable_criticial_voice_vlan", "true"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "age_out_time", "500"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "static_mac_addresses.0.mac_address", "0000.0000.0000"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "static_mac_addresses.0.if_name", "GigabitEthernet0/0/0"),
					resource.TestCheckResourceAttr("sdwan_switchport_feature_template.test", "static_mac_addresses.0.vlan", "100"),
				),
			},
		},
	})
}

func testAccSdwanSwitchportFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_switchport_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		slot = 0
		sub_slot = 0
		module_type = "4"
	}
	`
}

func testAccSdwanSwitchportFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_switchport_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		slot = 0
		sub_slot = 0
		module_type = "4"
		interfaces = [{
			name = "GigabitEthernet0/0/0"
			switchport_mode = "access"
			shutdown = true
			speed = "100"
			duplex = "full"
			switchport_access_vlan = 100
			switchport_trunk_allowed_vlans = "100,200"
			switchport_trunk_native_vlan = 100
			dot1x_enable = true
			dot1x_port_control = "auto"
			dot1x_authentication_order = ["dot1x"]
			dot1x_voice_vlan = 200
			dot1x_pae_enable = true
			dot1x_mac_authentication_bypass = true
			dot1x_host_mode = "multi-domain"
			dot1x_enable_periodic_reauth = true
			dot1x_periodic_reauth_inactivity_timeout = 100
			dot1x_periodic_reauth_interval = 60
			dot1x_control_direction = "both"
			dot1x_restricted_vlan = 100
			dot1x_guest_vlan = 101
			dot1x_critical_vlan = 102
			dot1x_enable_criticial_voice_vlan = true
		}]
		age_out_time = 500
		static_mac_addresses = [{
			mac_address = "0000.0000.0000"
			if_name = "GigabitEthernet0/0/0"
			vlan = 100
		}]
	}
	`
}