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

func TestAccDataSourceSdwanSLAClassPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSLAClassPolicyObjectConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "name", "Example"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "jitter", "100"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "latency", "10"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "loss", "1"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "fallback_best_tunnel_criteria", "jitter-loss-latency"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "fallback_best_tunnel_jitter", "100"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "fallback_best_tunnel_latency", "10"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "fallback_best_tunnel_loss", "1"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanSLAClassPolicyObjectConfig = `

resource "sdwan_sla_class_policy_object" "test" {
  name = "Example"
  jitter = 100
  latency = 10
  loss = 1
  fallback_best_tunnel_criteria = "jitter-loss-latency"
  fallback_best_tunnel_jitter = 100
  fallback_best_tunnel_latency = 10
  fallback_best_tunnel_loss = 1
}

data "sdwan_sla_class_policy_object" "test" {
  id = sdwan_sla_class_policy_object.test.id
}
`
