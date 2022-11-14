data "davinci_application" "one" {
  application_id = "<app_id>"
}

output "davinci_app_one_key" {
  value = data.davinci_application.one.api_keys.prod
}
