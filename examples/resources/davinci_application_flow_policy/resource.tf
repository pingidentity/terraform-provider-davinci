// example of bootstrapped application
resource "davinci_application" "registration_flow_app" {
  name           = "PingOne SSO Connection"
  environment_id = var.pingone_environment_id
  oauth {
    enabled = true
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = ["https://auth.pingone.com/0000-0000-000/rp/callback/openid_connect"]
    }
  }
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
}

resource "davinci_application_flow_policy" "registration_flow_policy" {
  environment_id = var.pingone_environment_id
  application_id = davinci_application.registration_flow_app.id
  name           = "PingOne - Registration"
  status         = "enabled"
  policy_flow {
    flow_id    = resource.davinci_flow.registration.id
    version_id = -1
    weight     = 100
  }
}
