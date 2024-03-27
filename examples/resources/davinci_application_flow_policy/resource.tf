resource "davinci_application" "my_awesome_registration_flow_application" {
  environment_id = var.pingone_environment_id

  name = "My Awesome Registration Application"

  oauth {
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = ["https://auth.pingone.com/0000-0000-000/rp/callback/openid_connect"]
    }
  }
}

resource "davinci_application_flow_policy" "my_awesome_registration_flow_application_policy" {
  environment_id = var.pingone_environment_id
  application_id = davinci_application.my_awesome_registration_flow_application.id

  name   = "PingOne - Registration"
  status = "enabled"

  policy_flow {
    flow_id    = davinci_flow.registration.id
    version_id = -1
    weight     = 100
  }
}
