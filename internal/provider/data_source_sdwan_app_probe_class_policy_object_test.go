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

func TestAccDataSourceSdwanAppProbeClassPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanAppProbeClassPolicyObjectConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_app_probe_class_policy_object.test", "name", "Example"),
					resource.TestCheckResourceAttr("data.sdwan_app_probe_class_policy_object.test", "forwarding_class", "FC1"),
					resource.TestCheckResourceAttr("data.sdwan_app_probe_class_policy_object.test", "mappings.0.color", "blue"),
					resource.TestCheckResourceAttr("data.sdwan_app_probe_class_policy_object.test", "mappings.0.dscp", "8"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanAppProbeClassPolicyObjectConfig = `

resource "sdwan_app_probe_class_policy_object" "test" {
  name = "Example"
  forwarding_class = "FC1"
  mappings = [{
    color = "blue"
    dscp = 8
  }]
}

data "sdwan_app_probe_class_policy_object" "test" {
  id = sdwan_app_probe_class_policy_object.test.id
}
`
