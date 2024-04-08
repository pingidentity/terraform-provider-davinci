resource "davinci_connection" "codeSnippetConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "codeSnippetConnector"
  name         = "My awesome codeSnippetConnector"

  property {
    name  = "code"
    type  = "string"
    value = var.codesnippetconnector_property_code
  }

  property {
    name  = "inputSchema"
    type  = "string"
    value = var.codesnippetconnector_property_input_schema
  }

  property {
    name  = "outputSchema"
    type  = "string"
    value = var.codesnippetconnector_property_output_schema
  }
}
