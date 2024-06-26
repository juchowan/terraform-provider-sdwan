---
subcategory: "Guides"
page_title: "Changelog"
description: |-
    Changelog
---

# Changelog

## 0.3.6

- Fix issue with address family configuration of `sdwan_cisco_bgp_feature_template` resource
- Fix issue with area configuration of `sdwan_cisco_ospf_feature_template` resource
- BREAKING CHANGE: Rename `send_lifetime` attribute of `sdwan_cisco_security_feature_template` resource to `send_lifetime_local`
- BREAKING CHANGE: Rename `accept_lifetime` attribute of `sdwan_cisco_security_feature_template` resource to `accept_lifetime_local`
- Improve error handling when detaching device templates
- Fix templating issue with `sdwan_cisco_vpn_interface_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/165)
- Fix templating issue with `sdwan_cisco_sig_credentials_feature_template` resouce, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/162)
- Enable `global` flag for `sdwan_cisco_sig_credentials_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/163)
- Fix templating issue with `sdwan_cisco_secure_internet_gateway_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/164)
- Fix templating issue with `sdwan_cisco_secure_internet_gateway_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/160)
- BREAKING CHANGE: Rename `aup_enabled` attribute of `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source to `zscaler_aup_enabled`
- BREAKING CHANGE: Rename `aup_block_internet_until_accepted` attribute of `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source to `zscaler_aup_block_internet_until_accepted`
- BREAKING CHANGE: Rename `aup_force_ssl_inspection` attribute of `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source to `zscaler_aup_force_ssl_inspection`
- BREAKING CHANGE: Rename `aup_timeout` attribute of `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source to `zscaler_aup_timeout`
- Fix templating of certain boolean values, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/133), [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/134)
- BREAKING CHANGE: Rename `dot1x_voice_vlan` attribute of `sdwan_switchport_feature_template` resource and data source to `voice_vlan`
- Use `set` type for list attributes with primitive values
- Fix templating issue with `sdwan_cisco_security_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/157)
- Fix templating of endpoint and object trackers with `sdwan_cisco_system_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/159)

## 0.3.5

- Fix import operation of profile parcel resources
- BREAKING CHANGE: Convert `protocol`, `source_port`, `destination_port` attribute of `sdwan_application_aware_routing_definition` to strings to support multiple values
- Add `sdwan_system_omp_profile_parcel` resource and data source
- Add `sdwan_security_policy` resource and data source
- BREAKING CHANGE: Rename `von_configurations` attribute of `sdwan_cedge_aaa_feature_template` resource and data source to `vpn_configurations`
- Add `strict` and `fallbackToBestPath` options to SLA class action parameters of `sdwan_application_aware_routing_policy_definition` resource and data source
- BREAKING CHANGE: Rename `radius_dynamic_author_cts_authorization_list` attribute of `sdwan_cedge_aaa_feature_template` resource and data source to `radius_trustsec_cts_authorization_list`
- Add option to configure `ipv6-unicast` address families to `sdwan_cisco_bgp_feature_template` resource and data source
- Add `stub` and `nssa` options to `sdwan_cisco_ospf_feature_template` resource and data source

## 0.3.4

- Add `sdwan_switchport_feature_template` resource and data source
- Add `sdwan_vpn_interface_svi_feature_template` resource and data source

## 0.3.3

- Fix next hop action and match entry of `sdwan_route_policy_definition` resource and data source

## 0.3.2

- BREAKING CHANGE: Rename `source_port` attribute to `source_ports` of `sdwan_ipv4_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `destination_port` attribute to `destination_ports` of `sdwan_ipv4_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `source_port` attribute to `source_ports` of `sdwan_ipv6_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `destination_port` attribute to `destination_ports` of `sdwan_ipv6_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `source_port` attribute to `source_ports` of `sdwan_ipv4_device_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `source_port` attribute to `source_ports` of `sdwan_ipv6_device_acl_policy_definition` resource and data source
- Add `traffic_class` attribute to `sdwan_ipv6_acl_policy_definition` resource and data source
- Fix next hop action entry of `sdwan_route_policy_definition` resource and data source

## 0.3.1

- Fix idempotency issue with DNS redirect configuration of `sdwan_traffic_data_policy_definition` resource
- Add option to select both `gre` and `ipsec` encaps for local TLOC set action of `sdwan_traffic_data_policy_definition` resource
- Fix issue with `exportTo` action entry of `sdwan_custom_control_topology_policy_definition` resource
- Fix issue with service TLOC lists of `sdwan_traffic_data_policy_definition` resource

## 0.3.0

- Add `sdwan_system_feature_profile` resource and data source
- Add `sdwan_service_feature_profile` resource and data source
- Add `sdwan_transport_feature_profile` resource and data source
- Add `sdwan_cli_feature_profile` resource and data source
- Add `sdwan_system_aaa_profile_parcel` resource and data source
- Add `sdwan_system_banner_profile_parcel` resource and data source
- Add `sdwan_system_bfd_profile_parcel` resource and data source
- Fix issue with `sdwan_centralized_policy` not accepting all valid parameters
- Fix issue with `sdwan_custom_control_topology_policy_definition` not accepting all possible values
- BREAKING CHANGE: Convert `protocol`, `source_port`, `destination_port` attributes of `sdwan_traffic_data_policy_definition` to strings to support multiple values
- Fix issue with `sdwan_traffic_data_policy_definition` not supporting local and restrict options
- Add `sdwan_dns_security_policy_definition` resource and data source
- Add `sdwan_system_global_profile_parcel` resource and data source
- BREAKING CHANGE: Rename `sdwan_acl_policy_definition` resource and data source to `sdwan_ipv4_acl_policy_definition` and update attributes
- BREAKING CHANGE: Rename `sdwan_device_acl_policy_definition` resource and data source to `sdwan_ipv4_device_acl_policy_definition` and update attributes
- Add `sdwan_ipv6_acl_policy_definition` resource and data source
- Add `sdwan_ipv6_device_acl_policy_definition` resource and data source

## 0.2.11

- Fix idempotency issue with various resource and boolean values
- Remove redundant `per_tunnel_qos` and `per_tunnel_qos_aggregator` attributes from `sdwan_cisco_vpn_interface_feature_template` resource and data source
- Add `base_action` attribute to `sdwan_custom_control_topology_policy_definition` resource and data source
- Add `base_action` attribute to `sdwan_traffic_data_policy_definition` resource and data source
- Allow full path variables in `sdwan_attach_feature_device_template` resource `variables` section
- Always include empty lists in payloads to fix various GUI compatibility issues
- Skip `sdwan_attach_feature_device_template` template detachment when device template is no longer attached

## 0.2.10

- Fix issue with `sdwan_cisco_ospf_feature_template` resource when not configuring route policies
- Fix issue with `sdwan_cisco_snmp_feature_template` resource when configuring SNMPv2
- Add `sdwan_device` data source
- Add `sdwan_vedge_inventory` data source
- Fix idempotency issue with `sdwan_cisco_ospf_feature_template` resource when enabling `default_information_originate`

## 0.2.9

- Add `sdwan_advanced_malware_protection_policy_definition` resource and data source
- BREAKING CHANGE: Rename `group` attribute of `sdwan_cedge_aaa_feature_template` resource and data source to `groups` and fix type
- Use type `Set` for `device_types` attributes of feature template resources and data sources
- Add `sdwan_tls_ssl_decryption_policy_definition` resource and data source
- Add `sdwan_advanced_inspection_profile_policy_definition` resource and data source

## 0.2.8

- Correctly populate system-level variables when attaching existing (factory default) feature templates
- Add `sdwan_intrusion_prevention_policy_definition` resource and data source
- Add `sdwan_tls_ssl_profile_policy_definition` resource and data source
- BREAKING CHANGE: Rename `timers_spf_initial_holf` attribute of `sdwan_cisco_ospf_feature_template` resource and data source to `timers_spf_initial_hold`
- Fix issue with setting an `omp_tag` match condition using a `sdwan_route_policy_definition` resource

## 0.2.7

- Add option to feature template data sources to identify templates by its name

## 0.2.6

- Fix deployment of attached device templates without feature template changes

## 0.2.5

- Add `sdwan_local_application_list_policy_object` resource and data source
- Add `sdwan_domain_list_policy_object` resource and data source
- Add `sdwan_ips_signature_list_policy_object` resource and data source
- Add `sdwan_allow_url_list_policy_object` resource and data source
- Add `sdwan_block_url_list_policy_object` resource and data source
- Add `sdwan_zone_list_policy_object` resource and data source
- Add `sdwan_port_list_policy_object` resource and data source
- Add `sdwan_protocol_list_policy_object` resource and data source
- Add `sdwan_geo_location_list_policy_object` resource and data source
- Add `sdwan_data_fqdn_prefix_list_policy_object` resource and data source
- Add `sdwan_object_group_policy_definition` resource and data source
- Add `sdwan_rule_set_policy_definition` resource and data source

## 0.2.4

- BREAKING CHANGE: Rename `ipv6_adresses` attribute of `sdwan_cisco_vpn_interface_feature_template` resource and data source to `ipv6_addresses`

## 0.2.3

- Add `sdwan_centralized_policy` resource and data source
- Add `sdwan_activate_centralized_policy` resource

## 0.2.2

- Add `sdwan_application_aware_routing_policy_definition` resource and data source
- Add `sdwan_cflowd_policy_definition` resource and data source
- Add `sdwan_traffic_data_policy_definition` resource and data source
- Add `set_parameters` attribute to `sdwan_acl_policy_definition` resource and data source
- BREAKING CHANGE: Remove `dscp` and `next_hop` attributes of `sdwan_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `layer2cos` attribute of `sdwan_rewrite_rule_policy_definition` resource and data source to `layer2_cos`
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_class_map_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_app_probe_class_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_mirror_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_policer_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_preferred_color_group_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_sla_class_policy_object` resource and data source
- Add `app_probe_class_version` attribute to `sdwan_sla_class_policy_object` resource and data source

## 0.2.1

- Add `sdwan_cisco_ospf_feature_template` resource and data source
- Add `sdwan_cisco_vpn_interface_ipsec_feature_template` resource and data source
- Add `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source
- Add `sdwan_hub_and_spoke_topology_policy_definition` resource and data source
- Add `sdwan_mesh_topology_policy_definition` resource and data source
- Add `sdwan_custom_control_topology_policy_definition` resource and data source
- Add `sdwan_vpn_membership_policy_definition` resource and data source

## 0.2.0

- BREAKING CHANGE: Completely revamped the provider based on `github.com/netascode/terraform-provider-sdwan` codebase, replacing all existing resources and data sources

## 0.1.0 (July 23, 2021)

- Initial Release

