data "davinci_applications" "all" {
}

output "davinci_applications" {
  value = data.davinci_applications.all.applications
}
