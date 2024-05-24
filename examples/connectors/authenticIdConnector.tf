resource "davinci_connection" "authenticIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "authenticIdConnector"
  name         = "My awesome authenticIdConnector"

  property {
    name  = "accountAccessKey"
    type  = "string"
    value = var.authenticidconnector_property_account_access_key
  }

  property {
    name  = "androidSDKLicenseKey"
    type  = "string"
    value = var.authenticidconnector_property_android_sdk_license_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.authenticidconnector_property_api_url
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.authenticidconnector_property_base_url
  }

  property {
    name  = "clientCertificate"
    type  = "string"
    value = var.authenticidconnector_property_client_certificate
  }

  property {
    name  = "clientKey"
    type  = "string"
    value = var.authenticidconnector_property_client_key
  }

  property {
    name  = "iOSSDKLicenseKey"
    type  = "string"
    value = var.authenticidconnector_property_ios_sdk_license_key
  }

  property {
    name  = "passphrase"
    type  = "string"
    value = var.authenticidconnector_property_passphrase
  }

  property {
    name  = "secretToken"
    type  = "string"
    value = var.authenticidconnector_property_secret_token
  }
}
