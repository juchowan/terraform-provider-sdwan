resource "sdwan_app_probe_class_policy_object" "example" {
  name             = "Example"
  forwarding_class = "FC1"
  mappings = [
    {
      color = "blue"
      dscp  = 8
    }
  ]
}
