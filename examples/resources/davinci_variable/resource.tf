resource "davinci_variable" "my_awesome_region_variable" {
  environment_id = var.environment_id

  name        = "region"
  context     = "company"
  description = "identifies region for functions in flow"
  value       = "northamerica"
  type        = "string"
}
