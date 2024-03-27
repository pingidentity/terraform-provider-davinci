data "davinci_application" "by_app_id" {
  environment_id = var.environment_id
  application_id = var.application_id
}

output "davinci_app_one_key" {
  value = data.davinci_application.one.api_keys.prod
}
