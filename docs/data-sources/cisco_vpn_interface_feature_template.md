---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_cisco_vpn_interface_feature_template Data Source - terraform-provider-sdwan"
subcategory: "Feature Templates"
description: |-
  This data source can read the Cisco VPN Interface feature template.
---

# sdwan_cisco_vpn_interface_feature_template (Data Source)

This data source can read the Cisco VPN Interface feature template.

## Example Usage

```terraform
data "sdwan_cisco_vpn_interface_feature_template" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the feature template

### Read-Only

- `access_lists` (Attributes List) Apply ACL (see [below for nested schema](#nestedatt--access_lists))
- `address` (String) Assign IPv4 address
- `address_variable` (String) Variable name
- `arp_timeout` (Number) Timeout value for dynamically learned ARP entries, <0..2678400> seconds
- `arp_timeout_variable` (String) Variable name
- `auto_bandwidth_detect` (Boolean) Interface auto detect bandwidth
- `auto_bandwidth_detect_variable` (String) Variable name
- `autonegotiate` (Boolean) Link autonegotiation
- `autonegotiate_variable` (String) Variable name
- `bandwidth_downstream` (Number) Interface downstream bandwidth capacity, in kbps
- `bandwidth_downstream_variable` (String) Variable name
- `bandwidth_upstream` (Number) Interface upstream bandwidth capacity, in kbps
- `bandwidth_upstream_variable` (String) Variable name
- `block_non_source_ip` (Boolean) Block packets originating from IP address that is not from this source
- `block_non_source_ip_variable` (String) Variable name
- `core_region` (String) Enable core region
- `core_region_variable` (String) Variable name
- `description` (String) The description of the feature template
- `device_types` (List of String) List of supported device types
- `dhcp` (Boolean) Enable DHCP
- `dhcp_distance` (Number) Set administrative distance for DHCP default route
- `dhcp_distance_variable` (String) Variable name
- `dhcp_variable` (String) Variable name
- `dhcpv6` (Boolean) Enable DHCPv6
- `dhcpv6_variable` (String) Variable name
- `duplex` (String) Duplex mode
- `duplex_variable` (String) Variable name
- `enable_core_region` (Boolean) Enable core region
- `enable_sgt` (Boolean) Enables the interface for CTS SGT authorization and forwarding.
- `gre_tunnel_source_ip` (String) Extend remote TLOC over a GRE tunnel to a local WAN interface
- `gre_tunnel_source_ip_variable` (String) Variable name
- `gre_tunnel_xconnect` (String) Extend remote TLOC over a GRE tunnel to a local WAN interface
- `gre_tunnel_xconnect_variable` (String) Variable name
- `icmp_redirect_disable` (Boolean) Set this option to disable the icmp/icmpv6 redirect packets
- `icmp_redirect_disable_variable` (String) Variable name
- `interface_description` (String) Interface description
- `interface_description_variable` (String) Variable name
- `interface_mtu` (Number) Interface MTU GigabitEthernet0 <1500..1518>, Other GigabitEthernet <1500..9216> in bytes
- `interface_mtu_variable` (String) Variable name
- `interface_name` (String) Interface name: ge0/<0-..> or ge0/<0-..>.vlanid or irb<bridgeid:1-63> or loopback<string> or natpool-<1..31> when present
- `interface_name_variable` (String) Variable name
- `ip_directed_broadcast` (Boolean) IP Directed-Broadcast
- `ip_directed_broadcast_variable` (String) Variable name
- `ip_mtu` (Number) IP MTU for GigabitEthernet main <576..Interface MTU>, GigabitEthernet subinterface <576..9216>, Other Interfaces <576..2000> in bytes
- `ip_mtu_variable` (String) Variable name
- `iperf_server` (String) Iperf server for auto bandwidth detect
- `iperf_server_variable` (String) Variable name
- `ipv4_dhcp_helper` (List of String) List of DHCP IPv4 helper addresses
- `ipv4_dhcp_helper_variable` (String) Variable name
- `ipv4_secondary_addresses` (Attributes List) Assign secondary IP addresses (see [below for nested schema](#nestedatt--ipv4_secondary_addresses))
- `ipv4_vrrps` (Attributes List) Enable VRRP (see [below for nested schema](#nestedatt--ipv4_vrrps))
- `ipv6_access_lists` (Attributes List) Apply IPv6 access list (see [below for nested schema](#nestedatt--ipv6_access_lists))
- `ipv6_address` (String) Assign IPv6 address
- `ipv6_address_variable` (String) Variable name
- `ipv6_dhcp_helpers` (Attributes List) DHCPv6 Helper (see [below for nested schema](#nestedatt--ipv6_dhcp_helpers))
- `ipv6_nat` (Boolean) NAT64 on this interface
- `ipv6_nat_variable` (String) Variable name
- `ipv6_secondary_addresses` (Attributes List) Assign secondary IPv6 addresses (see [below for nested schema](#nestedatt--ipv6_secondary_addresses))
- `ipv6_vrrps` (Attributes List) Enable VRRP (see [below for nested schema](#nestedatt--ipv6_vrrps))
- `load_interval` (Number) Interval for interface load calculation
- `load_interval_variable` (String) Variable name
- `mac_address` (String) Set MAC-layer address
- `mac_address_variable` (String) Variable name
- `media_type` (String) Media type
- `media_type_variable` (String) Variable name
- `name` (String) The name of the feature template
- `nat` (Boolean) Network Address Translation on this interface
- `nat64_interface` (Boolean) NAT64 on this interface
- `nat64_interface_variable` (String) Variable name
- `nat66_interface` (Boolean) NAT66 on this interface
- `nat_inside_source_loopback_interface` (String) Configure NAT Inside Loopback Interface
- `nat_inside_source_loopback_interface_variable` (String) Variable name
- `nat_overload` (Boolean) Enable port translation(PAT)
- `nat_overload_variable` (String) Variable name
- `nat_pool_prefix_length` (Number) Ending IP address of NAT Pool Prefix Length
- `nat_pool_prefix_length_variable` (String) Variable name
- `nat_pool_range_end` (String) Ending IP address of NAT pool range
- `nat_pool_range_end_variable` (String) Variable name
- `nat_pool_range_start` (String) Starting IP address of NAT pool range
- `nat_pool_range_start_variable` (String) Variable name
- `nat_type` (String) NAT type
- `nat_type_variable` (String) Variable name
- `nat_variable` (String) Variable name
- `per_tunnel_qos` (Boolean) Per-tunnel Qos
- `per_tunnel_qos_aggregator` (Boolean) Per-tunnel QoS Aggregator
- `per_tunnel_qos_aggregator_variable` (String) Variable name
- `per_tunnel_qos_variable` (String) Variable name
- `poe` (Boolean) Configure interface as Power-over-Ethernet source
- `poe_variable` (String) Variable name
- `propagate_sgt` (Boolean) Enable/Disable CTS SGT propagation on an interface.
- `qos_adaptive` (Boolean) Adaptive QoS
- `qos_adaptive_bandwidth_downstream` (Number) Adaptive QoS default downstream bandwidth
- `qos_adaptive_bandwidth_downstream_variable` (String) Variable name
- `qos_adaptive_bandwidth_upstream` (Number) Adaptive QoS default upstream bandwidth
- `qos_adaptive_bandwidth_upstream_variable` (String) Variable name
- `qos_adaptive_max_downstream` (Number) Downstream max bandwidth limit
- `qos_adaptive_max_downstream_variable` (String) Variable name
- `qos_adaptive_max_upstream` (Number) Upstream max bandwidth limit
- `qos_adaptive_max_upstream_variable` (String) Variable name
- `qos_adaptive_min_downstream` (Number) Downstream min bandwidth limit
- `qos_adaptive_min_downstream_variable` (String) Variable name
- `qos_adaptive_min_upstream` (Number) Upstream min bandwidth limit
- `qos_adaptive_min_upstream_variable` (String) Variable name
- `qos_adaptive_period` (Number) Periodic timer for adaptive QoS in minutes
- `qos_adaptive_period_variable` (String) Variable name
- `qos_map` (String) Name of QoS map
- `qos_map_variable` (String) Variable name
- `qos_map_vpn` (String) Name of VPN QoS map
- `qos_map_vpn_variable` (String) Variable name
- `rewrite_rule_name` (String) Name of rewrite rule
- `rewrite_rule_name_variable` (String) Variable name
- `secondary_region` (String) Enable secondary region
- `secondary_region_variable` (String) Variable name
- `sgt_enforcement` (Boolean) Enables the interface for CTS SGT authorization and forwarding.
- `sgt_enforcement_sgt` (Number) SGT value between 2 and 65519.
- `sgt_enforcement_sgt_variable` (String) Variable name
- `shaping_rate` (Number) 1ge  interfaces: [0..1000000]kbps; 10ge interfaces: [0..10000000]kbps
- `shaping_rate_variable` (String) Variable name
- `shutdown` (Boolean) Administrative state
- `shutdown_variable` (String) Variable name
- `speed` (String) Set interface speed
- `speed_variable` (String) Variable name
- `static_arps` (Attributes List) Configure static ARP entries (see [below for nested schema](#nestedatt--static_arps))
- `static_nat66_entries` (Attributes List) static NAT (see [below for nested schema](#nestedatt--static_nat66_entries))
- `static_nat_entries` (Attributes List) Configure static NAT entries (see [below for nested schema](#nestedatt--static_nat_entries))
- `static_port_forward_entries` (Attributes List) Configure Port Forward entries (see [below for nested schema](#nestedatt--static_port_forward_entries))
- `static_sgt` (Number) SGT value between 2 and 65519.
- `static_sgt_trusted` (Boolean) Indicates that the interface is trustworthy for CTS.
- `static_sgt_variable` (String) Variable name
- `tcp_mss_adjust` (Number) TCP MSS on SYN packets, in bytes
- `tcp_mss_adjust_variable` (String) Variable name
- `tcp_timeout` (Number) Set NAT TCP session timeout, in minutes
- `tcp_timeout_variable` (String) Variable name
- `template_type` (String) The template type
- `tloc_extension` (String) Extends a local TLOC to a remote node only for vpn 0
- `tloc_extension_variable` (String) Variable name
- `tracker` (List of String) Enable tracker for this interface
- `tracker_variable` (String) Variable name
- `tunnel_bandwidth` (Number) Tunnels Bandwidth Percent
- `tunnel_bandwidth_variable` (String) Variable name
- `tunnel_interface_allow_all` (Boolean) Allow all traffic. Overrides all other allow-service options if allow-service all is set
- `tunnel_interface_allow_all_variable` (String) Variable name
- `tunnel_interface_allow_bgp` (Boolean) Allow/deny BGP
- `tunnel_interface_allow_bgp_variable` (String) Variable name
- `tunnel_interface_allow_dhcp` (Boolean) Allow/Deny DHCP
- `tunnel_interface_allow_dhcp_variable` (String) Variable name
- `tunnel_interface_allow_dns` (Boolean) Allow/Deny DNS
- `tunnel_interface_allow_dns_variable` (String) Variable name
- `tunnel_interface_allow_https` (Boolean) Allow/Deny Https
- `tunnel_interface_allow_https_variable` (String) Variable name
- `tunnel_interface_allow_icmp` (Boolean) Allow/Deny ICMP
- `tunnel_interface_allow_icmp_variable` (String) Variable name
- `tunnel_interface_allow_netconf` (Boolean) Allow/Deny NETCONF
- `tunnel_interface_allow_netconf_variable` (String) Variable name
- `tunnel_interface_allow_ntp` (Boolean) Allow/Deny NTP
- `tunnel_interface_allow_ntp_variable` (String) Variable name
- `tunnel_interface_allow_ospf` (Boolean) Allow/Deny OSPF
- `tunnel_interface_allow_ospf_variable` (String) Variable name
- `tunnel_interface_allow_snmp` (Boolean) Allow/Deny SNMP
- `tunnel_interface_allow_snmp_variable` (String) Variable name
- `tunnel_interface_allow_ssh` (Boolean) Allow/Deny SSH
- `tunnel_interface_allow_ssh_variable` (String) Variable name
- `tunnel_interface_allow_stun` (Boolean) Allow/Deny STUN
- `tunnel_interface_allow_stun_variable` (String) Variable name
- `tunnel_interface_bind_loopback_tunnel` (String) Bind loopback tunnel interface to a physical interface
- `tunnel_interface_bind_loopback_tunnel_variable` (String) Variable name
- `tunnel_interface_border` (Boolean) Set TLOC as border TLOC
- `tunnel_interface_border_variable` (String) Variable name
- `tunnel_interface_carrier` (String) Set carrier for TLOC
- `tunnel_interface_carrier_variable` (String) Variable name
- `tunnel_interface_clear_dont_fragment` (Boolean) Enable clear dont fragment (Currently Only SDWAN Tunnel Interface)
- `tunnel_interface_clear_dont_fragment_variable` (String) Variable name
- `tunnel_interface_color` (String) Set color for TLOC
- `tunnel_interface_color_restrict` (Boolean) Restrict this TLOC behavior
- `tunnel_interface_color_restrict_variable` (String) Variable name
- `tunnel_interface_color_variable` (String) Variable name
- `tunnel_interface_control_connections` (Boolean) Allow Control Connection
- `tunnel_interface_control_connections_variable` (String) Variable name
- `tunnel_interface_encapsulations` (Attributes List) Encapsulation for TLOC (see [below for nested schema](#nestedatt--tunnel_interface_encapsulations))
- `tunnel_interface_exclude_controller_group_list` (List of String) Exclude the following controller groups defined in this list
- `tunnel_interface_exclude_controller_group_list_variable` (String) Variable name
- `tunnel_interface_gre_tunnel_destination_ip` (String) Extend the TLOC to a remote node over GRE tunnel
- `tunnel_interface_gre_tunnel_destination_ip_variable` (String) Variable name
- `tunnel_interface_groups` (List of String) List of groups
- `tunnel_interface_groups_variable` (String) Variable name
- `tunnel_interface_hello_interval` (Number) Set time period of control hello packets <100..600000> milli seconds
- `tunnel_interface_hello_interval_variable` (String) Variable name
- `tunnel_interface_hello_tolerance` (Number) Set tolerance of control hello packets <12..6000> seconds
- `tunnel_interface_hello_tolerance_variable` (String) Variable name
- `tunnel_interface_last_resort_circuit` (Boolean) Set TLOC as last resort
- `tunnel_interface_last_resort_circuit_variable` (String) Variable name
- `tunnel_interface_low_bandwidth_link` (Boolean) Set the interface as a low-bandwidth circuit
- `tunnel_interface_low_bandwidth_link_variable` (String) Variable name
- `tunnel_interface_max_control_connections` (Number) Set the maximum number of control connections for this TLOC
- `tunnel_interface_max_control_connections_variable` (String) Variable name
- `tunnel_interface_nat_refresh_interval` (Number) Set time period of nat refresh packets <1...60> seconds
- `tunnel_interface_nat_refresh_interval_variable` (String) Variable name
- `tunnel_interface_network_broadcast` (Boolean) Accept and respond to network-prefix-directed broadcasts)
- `tunnel_interface_network_broadcast_variable` (String) Variable name
- `tunnel_interface_port_hop` (Boolean) Disallow port hopping on the tunnel interface
- `tunnel_interface_port_hop_variable` (String) Variable name
- `tunnel_interface_propagate_sgt` (Boolean) CTS SGT Propagation configuration
- `tunnel_interface_propagate_sgt_variable` (String) Variable name
- `tunnel_interface_tunnel_tcp_mss` (Number) Tunnel TCP MSS on SYN packets, in bytes
- `tunnel_interface_tunnel_tcp_mss_variable` (String) Variable name
- `tunnel_interface_vbond_as_stun_server` (Boolean) Put this wan interface in STUN mode only
- `tunnel_interface_vbond_as_stun_server_variable` (String) Variable name
- `tunnel_interface_vmanage_connection_preference` (Number) Set interface preference for control connection to vManage <0..8>
- `tunnel_interface_vmanage_connection_preference_variable` (String) Variable name
- `tunnel_qos_mode` (String) Set tunnel QoS mode
- `tunnel_qos_mode_variable` (String) Variable name
- `udp_timeout` (Number) Set NAT UDP session timeout, in minutes
- `udp_timeout_variable` (String) Variable name
- `version` (Number) The version of the feature template

<a id="nestedatt--access_lists"></a>
### Nested Schema for `access_lists`

Read-Only:

- `acl_name` (String) Name of access list
- `acl_name_variable` (String) Variable name
- `direction` (String) Direction
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--ipv4_secondary_addresses"></a>
### Nested Schema for `ipv4_secondary_addresses`

Read-Only:

- `address` (String) IP Address
- `address_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--ipv4_vrrps"></a>
### Nested Schema for `ipv4_vrrps`

Read-Only:

- `group_id` (Number) Group ID
- `group_id_variable` (String) Variable name
- `ip_address` (String) Assign IP Address
- `ip_address_variable` (String) Variable name
- `ipv4_secondary_addresses` (Attributes List) VRRP Secondary IP address (see [below for nested schema](#nestedatt--ipv4_vrrps--ipv4_secondary_addresses))
- `optional` (Boolean) Indicates if list item is considered optional.
- `priority` (Number) Set priority
- `priority_variable` (String) Variable name
- `timer` (Number) Timer interval for successive advertisements, in milliseconds
- `timer_variable` (String) Variable name
- `tloc_preference_change` (Boolean) change TLOC preference
- `tloc_preference_change_value` (Number) Set tloc preference change value
- `tloc_preference_change_value_variable` (String) Variable name
- `track_omp` (Boolean) Track OMP status
- `track_prefix_list` (String) Track Prefix List
- `track_prefix_list_variable` (String) Variable name
- `tracking_objects` (Attributes List) tracking object for VRRP configuration (see [below for nested schema](#nestedatt--ipv4_vrrps--tracking_objects))

<a id="nestedatt--ipv4_vrrps--ipv4_secondary_addresses"></a>
### Nested Schema for `ipv4_vrrps.ipv4_secondary_addresses`

Read-Only:

- `ip_address` (String) VRRP Secondary IP address
- `ip_address_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--ipv4_vrrps--tracking_objects"></a>
### Nested Schema for `ipv4_vrrps.tracking_objects`

Read-Only:

- `decrement_value` (Number) Decrement Value for VRRP priority
- `decrement_value_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `track_action` (String) Track Action
- `track_action_variable` (String) Variable name
- `tracker_id` (Number) Tracker ID
- `tracker_id_variable` (String) Variable name



<a id="nestedatt--ipv6_access_lists"></a>
### Nested Schema for `ipv6_access_lists`

Read-Only:

- `acl_name` (String) Name of access list
- `acl_name_variable` (String) Variable name
- `direction` (String) Direction
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--ipv6_dhcp_helpers"></a>
### Nested Schema for `ipv6_dhcp_helpers`

Read-Only:

- `address` (String) DHCPv6 Helper address
- `address_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `vpn_id` (Number) DHCPv6 Helper VPN
- `vpn_id_variable` (String) Variable name


<a id="nestedatt--ipv6_secondary_addresses"></a>
### Nested Schema for `ipv6_secondary_addresses`

Read-Only:

- `address` (String) IPv6 Address
- `address_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--ipv6_vrrps"></a>
### Nested Schema for `ipv6_vrrps`

Read-Only:

- `group_id` (Number) Group ID
- `group_id_variable` (String) Variable name
- `ipv6_adresses` (Attributes List) IPv6 VRRP (see [below for nested schema](#nestedatt--ipv6_vrrps--ipv6_adresses))
- `optional` (Boolean) Indicates if list item is considered optional.
- `priority` (Number) Set priority
- `priority_variable` (String) Variable name
- `timer` (Number) Timer interval for successive advertisements, in milliseconds
- `timer_variable` (String) Variable name
- `track_omp` (Boolean) Track OMP status
- `track_omp_variable` (String) Variable name
- `track_prefix_list` (String) Track Prefix List
- `track_prefix_list_variable` (String) Variable name

<a id="nestedatt--ipv6_vrrps--ipv6_adresses"></a>
### Nested Schema for `ipv6_vrrps.ipv6_adresses`

Read-Only:

- `ipv6_link_local` (String) Use link-local IPv6 Address
- `ipv6_link_local_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `prefix` (String) Assign Global IPv6 Prefix
- `prefix_variable` (String) Variable name



<a id="nestedatt--static_arps"></a>
### Nested Schema for `static_arps`

Read-Only:

- `ip_address` (String) IP Address
- `ip_address_variable` (String) Variable name
- `mac` (String) MAC address
- `mac_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--static_nat66_entries"></a>
### Nested Schema for `static_nat66_entries`

Read-Only:

- `optional` (Boolean) Indicates if list item is considered optional.
- `source_prefix` (String) Source Prefix
- `source_prefix_variable` (String) Variable name
- `source_vpn_id` (Number) Source VPN ID
- `source_vpn_id_variable` (String) Variable name
- `translated_source_prefix` (String) Translated Source Prefix
- `translated_source_prefix_variable` (String) Variable name


<a id="nestedatt--static_nat_entries"></a>
### Nested Schema for `static_nat_entries`

Read-Only:

- `optional` (Boolean) Indicates if list item is considered optional.
- `source_ip` (String) Source IP address to be translated
- `source_ip_variable` (String) Variable name
- `source_vpn_id` (Number) Configure VPN ID
- `source_vpn_id_variable` (String) Variable name
- `static_nat_direction` (String) Direction of static NAT translation
- `static_nat_direction_variable` (String) Variable name
- `translate_ip` (String) Statically translated source IP address
- `translate_ip_variable` (String) Variable name


<a id="nestedatt--static_port_forward_entries"></a>
### Nested Schema for `static_port_forward_entries`

Read-Only:

- `optional` (Boolean) Indicates if list item is considered optional.
- `protocol` (String) Protocol
- `protocol_variable` (String) Variable name
- `source_ip` (String) Source IP address to be translated
- `source_ip_variable` (String) Variable name
- `source_port` (Number) Source Port
- `source_port_variable` (String) Variable name
- `source_vpn_id` (Number) Configure VPN ID
- `source_vpn_id_variable` (String) Variable name
- `static_nat_direction` (String) Direction of static NAT translation
- `static_nat_direction_variable` (String) Variable name
- `translate_ip` (String) Statically translated source IP address
- `translate_ip_variable` (String) Variable name
- `translate_port` (Number) Translate Port
- `translate_port_variable` (String) Variable name


<a id="nestedatt--tunnel_interface_encapsulations"></a>
### Nested Schema for `tunnel_interface_encapsulations`

Read-Only:

- `encapsulation` (String) Encapsulation
- `optional` (Boolean) Indicates if list item is considered optional.
- `preference` (Number) Set preference for TLOC
- `preference_variable` (String) Variable name
- `weight` (Number) Set weight for TLOC
- `weight_variable` (String) Variable name