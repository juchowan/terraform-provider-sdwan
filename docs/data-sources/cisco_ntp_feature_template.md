---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_cisco_ntp_feature_template Data Source - terraform-provider-sdwan"
subcategory: "Feature Templates"
description: |-
  This data source can read the Cisco NTP feature template.
---

# sdwan_cisco_ntp_feature_template (Data Source)

This data source can read the Cisco NTP feature template.

## Example Usage

```terraform
data "sdwan_cisco_ntp_feature_template" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the feature template

### Read-Only

- `authentication_keys` (Attributes List) Set MD5 authentication key (see [below for nested schema](#nestedatt--authentication_keys))
- `description` (String) The description of the feature template
- `device_types` (List of String) List of supported device types
- `master` (Boolean) Configure device as NTP master
- `master_source_interface` (String) Set interface for NTP Master
- `master_source_interface_variable` (String) Variable name
- `master_stratum` (Number) Master Stratum <1..15>
- `master_stratum_variable` (String) Variable name
- `master_variable` (String) Variable name
- `name` (String) The name of the feature template
- `servers` (Attributes List) Configure NTP servers (see [below for nested schema](#nestedatt--servers))
- `template_type` (String) The template type
- `trusted_keys` (List of String) Designate authentication key as trustworthy
- `trusted_keys_variable` (String) Variable name
- `version` (Number) The version of the feature template

<a id="nestedatt--authentication_keys"></a>
### Nested Schema for `authentication_keys`

Read-Only:

- `id` (Number) MD5 authentication key ID
- `id_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `value` (String) Enter cleartext or AES-encrypted MD5 authentication key
- `value_variable` (String) Variable name


<a id="nestedatt--servers"></a>
### Nested Schema for `servers`

Read-Only:

- `authentication_key_id` (Number) Set authentication key for the server
- `authentication_key_id_variable` (String) Variable name
- `hostname_ip` (String) Set hostname or IP address of server
- `hostname_ip_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `prefer` (Boolean) Prefer this NTP server
- `prefer_variable` (String) Variable name
- `source_interface` (String) Set interface to use to reach NTP server
- `source_interface_variable` (String) Variable name
- `version` (Number) Set NTP version
- `version_variable` (String) Variable name
- `vpn_id` (Number) Set VPN in which NTP server is located
- `vpn_id_variable` (String) Variable name