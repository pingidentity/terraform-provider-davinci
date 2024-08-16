resource "davinci_flow" "my_awesome_main_flow" {
  environment_id = var.environment_id

  name      = "My Awesome Main Flow"
  flow_json = file("./path/to/example-mainflow.json")

  # ... subflow_link and connection_link arguments
}

resource "davinci_variable" "my_awesome_usercontext_variable" {
  environment_id = var.environment_id

  context = "flow"
  flow_id = davinci_flow.my_awesome_main_flow.id

  name        = "usercontext"
  description = "identifies usercontext for functions in flow"
  type        = "string"

  empty_value = true
}