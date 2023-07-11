---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_cisco_system_feature_template Resource - terraform-provider-sdwan"
subcategory: "Feature Templates"
description: |-
  This resource can manage a Cisco System feature template.
    - Minimum vManage version: 15.0.0
---

# sdwan_cisco_system_feature_template (Resource)

This resource can manage a Cisco System feature template.
  - Minimum vManage version: `15.0.0`

## Example Usage

```terraform
resource "sdwan_cisco_system_feature_template" "example" {
  name               = "Example"
  description        = "My Example"
  device_types       = ["vedge-C8000V"]
  timezone           = "UTC"
  hostname           = "Router1"
  system_description = "My Description"
  location           = "Building 1"
  latitude           = 40
  longitude          = 50
  geo_fencing        = true
  geo_fencing_range  = 1000
  geo_fencing_sms    = true
  geo_fencing_sms_phone_numbers = [
    {
      number = "+1234567"
    }
  ]
  device_groups         = ["group1"]
  controller_group_list = ["1"]
  system_ip             = "5.5.5.5"
  overlay_id            = 1
  site_id               = 1
  port_offset           = 1
  port_hopping          = true
  control_session_pps   = 300
  track_transport       = true
  track_interface_tag   = 1
  console_baud_rate     = "115200"
  max_omp_sessions      = 5
  multi_tenant          = true
  track_default_gateway = true
  admin_tech_on_failure = true
  idle_timeout          = 100
  trackers = [
    {
      name                        = "tracker1"
      endpoint_ip                 = "5.6.7.8"
      transport_endpoint_ip       = "5.6.7.8"
      transport_endpoint_protocol = "tcp"
      transport_endpoint_port     = 500
      endpoint_dns_name           = "abc.com"
      endpoint_api_url            = "https://1.1.1.1"
      elements                    = ["abc", "def"]
      boolean                     = "or"
      threshold                   = 300
      interval                    = 60
      multiplier                  = 3
      type                        = "interface"
    }
  ]
  object_trackers = [
    {
      object_number = 1
      interface     = "e1"
      sig           = "sig1"
      ip            = "6.6.6.6"
      mask          = "0.0.0.0"
      vpn_id        = 1
      group_tracks_ids = [
        {
          track_id = 1
        }
      ]
      boolean = "and"
    }
  ]
  on_demand_tunnel              = true
  on_demand_tunnel_idle_timeout = 10
  affinity_group_number         = 5
  affinity_group_preference     = [1]
  transport_gateway             = true
  enable_mrf_migration          = "enabled"
  migration_bgp_community       = 100
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the feature template
- `device_types` (List of String) List of supported device types
  - Choices: `vedge-C8000V`, `vedge-C8300-1N1S-4T2X`, `vedge-C8300-1N1S-6T`, `vedge-C8300-2N2S-6T`, `vedge-C8300-2N2S-4T2X`, `vedge-C8500-12X4QC`, `vedge-C8500-12X`, `vedge-C8500-20X6C`, `vedge-C8500L-8S4X`, `vedge-C8200-1N-4T`, `vedge-C8200L-1N-4T`
- `name` (String) The name of the feature template

### Optional

- `admin_tech_on_failure` (Boolean) Collect admin-tech before reboot due to daemon failure
  - Default value: `true`
- `admin_tech_on_failure_variable` (String) Variable name
- `affinity_group_number` (Number) Set the affinity group number for router
  - Range: `1`-`255`
- `affinity_group_number_variable` (String) Variable name
- `affinity_group_preference` (List of String) Set the affinity group preference
- `affinity_group_preference_variable` (String) Variable name
- `console_baud_rate` (String) Set the console baud rate
  - Choices: `1200`, `2400`, `4800`, `9600`, `19200`, `38400`, `57600`, `115200`
  - Default value: `9600`
- `console_baud_rate_variable` (String) Variable name
- `control_session_pps` (Number) Set the policer rate for control sessions
  - Range: `1`-`65535`
  - Default value: `300`
- `control_session_pps_variable` (String) Variable name
- `controller_group_list` (List of String) Configure a list of comma-separated device groups
- `controller_group_list_variable` (String) Variable name
- `device_groups` (List of String) Device groups (Use comma(,) for multiple groups)
- `device_groups_variable` (String) Variable name
- `enable_mrf_migration` (String) Enable migration mode to Multi-Region Fabric
  - Choices: `enabled`, `enabled-from-bgp-core`
- `geo_fencing` (Boolean) Enable Geo fencing
  - Default value: `false`
- `geo_fencing_range` (Number) Set the device’s geo fencing range
  - Range: `100`-`10000`
  - Default value: `100`
- `geo_fencing_range_variable` (String) Variable name
- `geo_fencing_sms` (Boolean) Enable Geo fencing
  - Default value: `false`
- `geo_fencing_sms_phone_numbers` (Attributes List) Set device’s geo fencing SMS phone number (see [below for nested schema](#nestedatt--geo_fencing_sms_phone_numbers))
- `hostname` (String) Set the hostname
- `hostname_variable` (String) Variable name
- `idle_timeout` (Number) Idle CLI timeout in minutes
  - Range: `0`-`300`
- `idle_timeout_variable` (String) Variable name
- `latitude` (Number) Set the device’s physical latitude
  - Range: `-90`-`90`
- `latitude_variable` (String) Variable name
- `location` (String) Set the location of the device
- `location_variable` (String) Variable name
- `longitude` (Number) Set the device’s physical longitude
  - Range: `-180`-`180`
- `longitude_variable` (String) Variable name
- `max_omp_sessions` (Number) Set the maximum number of OMP sessions <1..100> the device can have
  - Range: `1`-`100`
- `max_omp_sessions_variable` (String) Variable name
- `migration_bgp_community` (Number) Set BGP community during migration from BGP-core based network
  - Range: `1`-`4294967295`
- `multi_tenant` (Boolean) Device is multi-tenant
  - Default value: `false`
- `multi_tenant_variable` (String) Variable name
- `object_trackers` (Attributes List) Object Track configuration (see [below for nested schema](#nestedatt--object_trackers))
- `on_demand_tunnel` (Boolean) Enable or disable On-demand Tunnel
  - Default value: `false`
- `on_demand_tunnel_idle_timeout` (Number) Idle CLI timeout in minutes
  - Range: `0`-`300`
- `on_demand_tunnel_idle_timeout_variable` (String) Variable name
- `on_demand_tunnel_variable` (String) Variable name
- `overlay_id` (Number) Set the Overlay ID
  - Range: `1`-`4294967295`
  - Default value: `1`
- `overlay_id_variable` (String) Variable name
- `port_hopping` (Boolean) Enable port hopping
  - Default value: `true`
- `port_hopping_variable` (String) Variable name
- `port_offset` (Number) Set the TLOC port offset when multiple devices are behind a NAT
  - Range: `0`-`19`
  - Default value: `0`
- `port_offset_variable` (String) Variable name
- `region_id` (Number) Set region ID
  - Range: `1`-`63`
- `region_id_variable` (String) Variable name
- `role` (String) Set the role for router
  - Choices: `edge-router`, `border-router`
- `role_variable` (String) Variable name
- `secondary_region_id` (Number) Set secondary region ID
  - Range: `1`-`63`
- `secondary_region_id_variable` (String) Variable name
- `site_id` (Number) Set the site identifier
  - Range: `1`-`4294967295`
- `site_id_variable` (String) Variable name
- `system_description` (String) Set a text description of the device
- `system_description_variable` (String) Variable name
- `system_ip` (String) Set the system IP address
- `system_ip_variable` (String) Variable name
- `timezone` (String) Set the timezone
  - Choices: `Europe/Andorra`, `Asia/Dubai`, `Asia/Kabul`, `America/Antigua`, `America/Anguilla`, `Europe/Tirane`, `Asia/Yerevan`, `Africa/Luanda`, `Antarctica/McMurdo`, `Antarctica/Rothera`, `Antarctica/Palmer`, `Antarctica/Mawson`, `Antarctica/Davis`, `Antarctica/Casey`, `Antarctica/Vostok`, `Antarctica/DumontDUrville`, `Antarctica/Syowa`, `America/Argentina/Buenos_Aires`, `America/Argentina/Cordoba`, `America/Argentina/Salta`, `America/Argentina/Jujuy`, `America/Argentina/Tucuman`, `America/Argentina/Catamarca`, `America/Argentina/La_Rioja`, `America/Argentina/San_Juan`, `America/Argentina/Mendoza`, `America/Argentina/San_Luis`, `America/Argentina/Rio_Gallegos`, `America/Argentina/Ushuaia`, `Pacific/Pago_Pago`, `Europe/Vienna`, `Australia/Lord_Howe`, `Antarctica/Macquarie`, `Australia/Hobart`, `Australia/Currie`, `Australia/Melbourne`, `Australia/Sydney`, `Australia/Broken_Hill`, `Australia/Brisbane`, `Australia/Lindeman`, `Australia/Adelaide`, `Australia/Darwin`, `Australia/Perth`, `Australia/Eucla`, `America/Aruba`, `Europe/Mariehamn`, `Asia/Baku`, `Europe/Sarajevo`, `America/Barbados`, `Asia/Dhaka`, `Europe/Brussels`, `Africa/Ouagadougou`, `Europe/Sofia`, `Asia/Bahrain`, `Africa/Bujumbura`, `Africa/Porto-Novo`, `America/St_Barthelemy`, `Atlantic/Bermuda`, `Asia/Brunei`, `America/La_Paz`, `America/Kralendijk`, `America/Noronha`, `America/Belem`, `America/Fortaleza`, `America/Recife`, `America/Araguaina`, `America/Maceio`, `America/Bahia`, `America/Sao_Paulo`, `America/Campo_Grande`, `America/Cuiaba`, `America/Santarem`, `America/Porto_Velho`, `America/Boa_Vista`, `America/Manaus`, `America/Eirunepe`, `America/Rio_Branco`, `America/Nassau`, `Asia/Thimphu`, `Africa/Gaborone`, `Europe/Minsk`, `America/Belize`, `America/St_Johns`, `America/Halifax`, `America/Glace_Bay`, `America/Moncton`, `America/Goose_Bay`, `America/Blanc-Sablon`, `America/Toronto`, `America/Nipigon`, `America/Thunder_Bay`, `America/Iqaluit`, `America/Pangnirtung`, `America/Resolute`, `America/Atikokan`, `America/Rankin_Inlet`, `America/Winnipeg`, `America/Rainy_River`, `America/Regina`, `America/Swift_Current`, `America/Edmonton`, `America/Cambridge_Bay`, `America/Yellowknife`, `America/Inuvik`, `America/Creston`, `America/Dawson_Creek`, `America/Vancouver`, `America/Whitehorse`, `America/Dawson`, `Indian/Cocos`, `Africa/Kinshasa`, `Africa/Lubumbashi`, `Africa/Bangui`, `Africa/Brazzaville`, `Europe/Zurich`, `Africa/Abidjan`, `Pacific/Rarotonga`, `America/Santiago`, `Pacific/Easter`, `Africa/Douala`, `Asia/Shanghai`, `Asia/Harbin`, `Asia/Chongqing`, `Asia/Urumqi`, `Asia/Kashgar`, `America/Bogota`, `America/Costa_Rica`, `America/Havana`, `Atlantic/Cape_Verde`, `America/Curacao`, `Indian/Christmas`, `Asia/Nicosia`, `Europe/Prague`, `Europe/Berlin`, `Europe/Busingen`, `Africa/Djibouti`, `Europe/Copenhagen`, `America/Dominica`, `America/Santo_Domingo`, `Africa/Algiers`, `America/Guayaquil`, `Pacific/Galapagos`, `Europe/Tallinn`, `Africa/Cairo`, `Africa/El_Aaiun`, `Africa/Asmara`, `Europe/Madrid`, `Africa/Ceuta`, `Atlantic/Canary`, `Africa/Addis_Ababa`, `Europe/Helsinki`, `Pacific/Fiji`, `Atlantic/Stanley`, `Pacific/Chuuk`, `Pacific/Pohnpei`, `Pacific/Kosrae`, `Atlantic/Faroe`, `Europe/Paris`, `Africa/Libreville`, `Europe/London`, `America/Grenada`, `Asia/Tbilisi`, `America/Cayenne`, `Europe/Guernsey`, `Africa/Accra`, `Europe/Gibraltar`, `America/Godthab`, `America/Danmarkshavn`, `America/Scoresbysund`, `America/Thule`, `Africa/Banjul`, `Africa/Conakry`, `America/Guadeloupe`, `Africa/Malabo`, `Europe/Athens`, `Atlantic/South_Georgia`, `America/Guatemala`, `Pacific/Guam`, `Africa/Bissau`, `America/Guyana`, `Asia/Hong_Kong`, `America/Tegucigalpa`, `Europe/Zagreb`, `America/Port-au-Prince`, `Europe/Budapest`, `Asia/Jakarta`, `Asia/Pontianak`, `Asia/Makassar`, `Asia/Jayapura`, `Europe/Dublin`, `Asia/Jerusalem`, `Europe/Isle_of_Man`, `Asia/Kolkata`, `Indian/Chagos`, `Asia/Baghdad`, `Asia/Tehran`, `Atlantic/Reykjavik`, `Europe/Rome`, `Europe/Jersey`, `America/Jamaica`, `Asia/Amman`, `Asia/Tokyo`, `Africa/Nairobi`, `Asia/Bishkek`, `Asia/Phnom_Penh`, `Pacific/Tarawa`, `Pacific/Enderbury`, `Pacific/Kiritimati`, `Indian/Comoro`, `America/St_Kitts`, `Asia/Pyongyang`, `Asia/Seoul`, `Asia/Kuwait`, `America/Cayman`, `Asia/Almaty`, `Asia/Qyzylorda`, `Asia/Aqtobe`, `Asia/Aqtau`, `Asia/Oral`, `Asia/Vientiane`, `Asia/Beirut`, `America/St_Lucia`, `Europe/Vaduz`, `Asia/Colombo`, `Africa/Monrovia`, `Africa/Maseru`, `Europe/Vilnius`, `Europe/Luxembourg`, `Europe/Riga`, `Africa/Tripoli`, `Africa/Casablanca`, `Europe/Monaco`, `Europe/Chisinau`, `Europe/Podgorica`, `America/Marigot`, `Indian/Antananarivo`, `Pacific/Majuro`, `Pacific/Kwajalein`, `Europe/Skopje`, `Africa/Bamako`, `Asia/Rangoon`, `Asia/Ulaanbaatar`, `Asia/Hovd`, `Asia/Choibalsan`, `Asia/Macau`, `Pacific/Saipan`, `America/Martinique`, `Africa/Nouakchott`, `America/Montserrat`, `Europe/Malta`, `Indian/Mauritius`, `Indian/Maldives`, `Africa/Blantyre`, `America/Mexico_City`, `America/Cancun`, `America/Merida`, `America/Monterrey`, `America/Matamoros`, `America/Mazatlan`, `America/Chihuahua`, `America/Ojinaga`, `America/Hermosillo`, `America/Tijuana`, `America/Santa_Isabel`, `America/Bahia_Banderas`, `Asia/Kuala_Lumpur`, `Asia/Kuching`, `Africa/Maputo`, `Africa/Windhoek`, `Pacific/Noumea`, `Africa/Niamey`, `Pacific/Norfolk`, `Africa/Lagos`, `America/Managua`, `Europe/Amsterdam`, `Europe/Oslo`, `Asia/Kathmandu`, `Pacific/Nauru`, `Pacific/Niue`, `Pacific/Auckland`, `Pacific/Chatham`, `Asia/Muscat`, `America/Panama`, `America/Lima`, `Pacific/Tahiti`, `Pacific/Marquesas`, `Pacific/Gambier`, `Pacific/Port_Moresby`, `Asia/Manila`, `Asia/Karachi`, `Europe/Warsaw`, `America/Miquelon`, `Pacific/Pitcairn`, `America/Puerto_Rico`, `Asia/Gaza`, `Asia/Hebron`, `Europe/Lisbon`, `Atlantic/Madeira`, `Atlantic/Azores`, `Pacific/Palau`, `America/Asuncion`, `Asia/Qatar`, `Indian/Reunion`, `Europe/Bucharest`, `Europe/Belgrade`, `Europe/Kaliningrad`, `Europe/Moscow`, `Europe/Volgograd`, `Europe/Samara`, `Asia/Yekaterinburg`, `Asia/Omsk`, `Asia/Novosibirsk`, `Asia/Novokuznetsk`, `Asia/Krasnoyarsk`, `Asia/Irkutsk`, `Asia/Yakutsk`, `Asia/Khandyga`, `Asia/Vladivostok`, `Asia/Sakhalin`, `Asia/Ust-Nera`, `Asia/Magadan`, `Asia/Kamchatka`, `Asia/Anadyr`, `Africa/Kigali`, `Asia/Riyadh`, `Pacific/Guadalcanal`, `Indian/Mahe`, `Africa/Khartoum`, `Europe/Stockholm`, `Asia/Singapore`, `Atlantic/St_Helena`, `Europe/Ljubljana`, `Arctic/Longyearbyen`, `Europe/Bratislava`, `Africa/Freetown`, `Europe/San_Marino`, `Africa/Dakar`, `Africa/Mogadishu`, `America/Paramaribo`, `Africa/Juba`, `Africa/Sao_Tome`, `America/El_Salvador`, `America/Lower_Princes`, `Asia/Damascus`, `Africa/Mbabane`, `America/Grand_Turk`, `Africa/Ndjamena`, `Indian/Kerguelen`, `Africa/Lome`, `Asia/Bangkok`, `Asia/Dushanbe`, `Pacific/Fakaofo`, `Asia/Dili`, `Asia/Ashgabat`, `Africa/Tunis`, `Pacific/Tongatapu`, `Europe/Istanbul`, `America/Port_of_Spain`, `Pacific/Funafuti`, `Asia/Taipei`, `Africa/Dar_es_Salaam`, `Europe/Kiev`, `Europe/Uzhgorod`, `Europe/Zaporozhye`, `Europe/Simferopol`, `Africa/Kampala`, `Pacific/Johnston`, `Pacific/Midway`, `Pacific/Wake`, `America/New_York`, `America/Detroit`, `America/Kentucky/Louisville`, `America/Kentucky/Monticello`, `America/Indiana/Indianapolis`, `America/Indiana/Vincennes`, `America/Indiana/Winamac`, `America/Indiana/Marengo`, `America/Indiana/Petersburg`, `America/Indiana/Vevay`, `America/Chicago`, `America/Indiana/Tell_City`, `America/Indiana/Knox`, `America/Menominee`, `America/North_Dakota/Center`, `America/North_Dakota/New_Salem`, `America/North_Dakota/Beulah`, `America/Denver`, `America/Boise`, `America/Phoenix`, `America/Los_Angeles`, `America/Anchorage`, `America/Juneau`, `America/Sitka`, `America/Yakutat`, `America/Nome`, `America/Adak`, `America/Metlakatla`, `Pacific/Honolulu`, `America/Montevideo`, `Asia/Samarkand`, `Asia/Tashkent`, `Europe/Vatican`, `America/St_Vincent`, `America/Caracas`, `America/Tortola`, `America/St_Thomas`, `Asia/Ho_Chi_Minh`, `Pacific/Efate`, `Pacific/Wallis`, `Pacific/Apia`, `Asia/Aden`, `Indian/Mayotte`, `Africa/Johannesburg`, `Africa/Lusaka`, `Africa/Harare`, `UTC`
  - Default value: `UTC`
- `timezone_variable` (String) Variable name
- `track_default_gateway` (Boolean) Enable or disable default gateway tracking
  - Default value: `true`
- `track_default_gateway_variable` (String) Variable name
- `track_interface_tag` (Number) OMP Tag attached to routes based on interface tracking
  - Range: `1`-`4294967295`
- `track_interface_tag_variable` (String) Variable name
- `track_transport` (Boolean) Configure tracking of transport
  - Default value: `true`
- `track_transport_variable` (String) Variable name
- `trackers` (Attributes List) Tracker configuration (see [below for nested schema](#nestedatt--trackers))
- `transport_gateway` (Boolean) Enable transport gateway
  - Default value: `false`
- `transport_gateway_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the feature template
- `template_type` (String) The template type
- `version` (Number) The version of the feature template

<a id="nestedatt--geo_fencing_sms_phone_numbers"></a>
### Nested Schema for `geo_fencing_sms_phone_numbers`

Optional:

- `number` (String) Mobile number, ex: +1231234414
- `number_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--object_trackers"></a>
### Nested Schema for `object_trackers`

Optional:

- `boolean` (String) Type of grouping to be performed for tracker group
  - Choices: `and`, `or`
- `boolean_variable` (String) Variable name
- `group_tracks_ids` (Attributes List) Tracks id in group configuration (see [below for nested schema](#nestedatt--object_trackers--group_tracks_ids))
- `interface` (String) interface name
- `interface_variable` (String) Variable name
- `ip` (String) IP address of route
- `ip_variable` (String) Variable name
- `mask` (String) Route Ip Mask
  - Default value: `0.0.0.0`
- `mask_variable` (String) Variable name
- `object_number` (Number) Object tracker ID
  - Range: `1`-`1000`
- `object_number_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `sig` (String) service sig
- `sig_variable` (String) Variable name
- `vpn_id` (Number) VPN
  - Range: `0`-`65527`
  - Default value: `0`

<a id="nestedatt--object_trackers--group_tracks_ids"></a>
### Nested Schema for `object_trackers.group_tracks_ids`

Optional:

- `optional` (Boolean) Indicates if list item is considered optional.
- `track_id` (Number) Track id
  - Range: `1`-`1000`
- `track_id_variable` (String) Variable name



<a id="nestedatt--trackers"></a>
### Nested Schema for `trackers`

Optional:

- `boolean` (String) Type of grouping to be performed for tracker group
  - Choices: `or`, `and`
  - Default value: `or`
- `boolean_variable` (String) Variable name
- `elements` (List of String) Tracker member names separated by space
- `elements_variable` (String) Variable name
- `endpoint_api_url` (String) API url of endpoint
- `endpoint_api_url_variable` (String) Variable name
- `endpoint_dns_name` (String) DNS name of endpoint
- `endpoint_dns_name_variable` (String) Variable name
- `endpoint_ip` (String) IP address of endpoint
- `endpoint_ip_variable` (String) Variable name
- `interval` (Number) Probe interval <10..600> seconds
  - Range: `20`-`600`
  - Default value: `60`
- `interval_variable` (String) Variable name
- `multiplier` (Number) Probe failure multiplier <1..10> failed attempts
  - Range: `1`-`10`
  - Default value: `3`
- `multiplier_variable` (String) Variable name
- `name` (String) Tracker name
- `name_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `threshold` (Number) Probe Timeout threshold <100..1000> milliseconds
  - Range: `100`-`1000`
  - Default value: `300`
- `threshold_variable` (String) Variable name
- `transport_endpoint_ip` (String) IP address of endpoint
- `transport_endpoint_ip_variable` (String) Variable name
- `transport_endpoint_port` (Number) TCP/UDP port pf endpoint
  - Range: `1`-`65535`
- `transport_endpoint_port_variable` (String) Variable name
- `transport_endpoint_protocol` (String) transport protocol: TCP/UDP
  - Choices: `tcp`, `udp`
- `transport_endpoint_protocol_variable` (String) Variable name
- `type` (String) Default(Interface)
  - Choices: `interface`, `static-route`
  - Default value: `interface`
- `type_variable` (String) Variable name

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_cisco_system_feature_template.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```