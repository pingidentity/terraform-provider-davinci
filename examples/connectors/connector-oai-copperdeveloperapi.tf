resource "davinci_connection" "connector-oai-copperdeveloperapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-copperdeveloperapi"
  name         = "My awesome connector-oai-copperdeveloperapi"

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-copperdeveloperapi_property_base_path
  }

  property {
    name  = "contentType"
    type  = "string"
    value = var.connector-oai-copperdeveloperapi_property_content_type
  }

  property {
    name  = "xPWAccessToken"
    type  = "string"
    value = var.connector-oai-copperdeveloperapi_property_x_p_w_access_token
  }

  property {
    name  = "xPWApplication"
    type  = "string"
    value = var.connector-oai-copperdeveloperapi_property_x_p_w_application
  }

  property {
    name  = "xPWUserEmail"
    type  = "string"
    value = var.connector-oai-copperdeveloperapi_property_x_p_w_user_email
  }
}
