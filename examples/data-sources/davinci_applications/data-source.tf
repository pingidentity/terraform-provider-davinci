data "davinci_applications" "all" {
  environment_id = var.environment_id
}

output "davinci_applications" {
  value = data.davinci_applications.all.applications
}
