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
  policies {
    name   = "PingOne - Sign On and Password Reset"
    status = "enabled"
    policy_flows {
      flow_id    = "abc123"
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
