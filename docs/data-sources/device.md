---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_device Data Source - terraform-provider-sdwan"
subcategory: "Inventory"
description: |-
  This data source can read the Device .
---

# sdwan_device (Data Source)

This data source can read the Device .

## Example Usage

```terraform
data "sdwan_device" "example" {
  serial_number = "2AJC9DJ"
  name          = "va-001"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `name` (String) The hostname of a device
- `serial_number` (String) Serial number for device. Could be board or virtual identifier

### Read-Only

- `devices` (Attributes List) List of returned devices (see [below for nested schema](#nestedatt--devices))
- `id` (String) The id of the object

<a id="nestedatt--devices"></a>
### Nested Schema for `devices`

Read-Only:

- `device_id` (String) The device ID as defined in vManage
- `hostname` (String) Hostname for respective device
- `reachability` (String) Reachability of device
- `serial_number` (String) Serial number for device. Could be board or virtual identifier
- `site_id` (String) Site id for respective device
- `state` (String) State for respective device
- `status` (String) Status for respective device
- `uuid` (String) Unique identifier for device