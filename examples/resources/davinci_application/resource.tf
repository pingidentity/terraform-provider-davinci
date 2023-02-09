// example of bootstrapped application
resource "davinci_application" "use_default_flow" {
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
  policy {
    name   = "PingOne - Authentication"
    status = "enabled"
    policy_flow {
      flow_id    = var.davinci_flow_id
      version_id = -1
      weight     = 100
    }
  }
  policy {
    name   = "PingOne - Registration"
    status = "enabled"
    policy_flow {
      flow_id    = resource.davinci_flow.registration.id
      version_id = -1
      weight     = 100
    }
  }
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
}

output "default_app_test_key" {
  value = resource.davinci_application.use_default_flow.api_keys.test
}
