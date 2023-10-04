---
layout: ""
page_title: "Migrate Application Flow Policies to Individual Resources"
description: |-
  The guide describes how to migrate application flow policies, configured with Terraform, from using being a sub-resource of applications to the new format as individual policies to provide better overall stability of downstream applications
---

# Migrate Application Flow Policies to Individual Resources

The guide describes how to migrate application flow policies, configured with Terraform, from using being a sub-resource of applications to the new format as individual policies to provide better overall stability of downstream applications

This guide applies to users upgrading _from_ DaVinci provider < 0.2.0.

## Background

A defect was discovered in the `policy` field of `resource.davinci_application` that causes policies to be replaced rather than updated when there is a change to be made. The replace behavious causes upstream breaks on pingone_applications that use DaVinci Policies.

## Migration Process

In the DaVinci provider version 0.2.0 the `policy` field has been removed from `resource.davinci_application`. This functionality is now handled in the `davinci_application_flow_policy` resource.

The migration process consists of:

  - Move all instances of `resource.davinci_application.policy` to separate `davinci_application_flow_policy` resources. 
  - Update all references to these new resources accordingly.
  - Before applying the new configuration, import existing managed policies to be part of state.

### Update configuration

The following example shows a before and after of what the relevant configuration may look like for the first two steps: 

**BEFORE**

```terraform
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

resource "pingone_application" "oidc_sdk_sample_app" {
  environment_id = var.pingone_environment_id
  enabled        = true
  name           = "Sample App"
  description    = "A custom sample OIDC application to demonstrate PingOne integration."

  oidc_options {
    type                        = "SINGLE_PAGE_APP"
    grant_types                 = ["AUTHORIZATION_CODE", "IMPLICIT", "REFRESH_TOKEN"]
    response_types              = ["CODE", "TOKEN", "ID_TOKEN"]
    pkce_enforcement            = "S256_REQUIRED"
    token_endpoint_authn_method = "NONE"
    redirect_uris               = var.redirect_uris
    post_logout_redirect_uris   = ["${var.app_url}"]
  }
}

resource "pingone_application_flow_policy_assignment" "login_flow" {
  environment_id = module.environment.environment_id
  application_id = pingone_application.oidc_sdk_sample_app.id
  flow_policy_id = davinci_application.registration_flow_app.policy.* [index(davinci_application.registration_flow_app.policy[*].name, "PingOne - Registration")].policy_id

  priority = 1
}
```

**AFTER:**

```terraform
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
  ## Policy is removed
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
}

## New policy resource is added 
resource "davinci_application_flow_policy" "registration_flow_policy"  {
  environment_id = var.pingone_environment_id
  application_id = davinci_application.registration_flow_app.id
  name   = "PingOne - Registration"
  status = "enabled"
  policy_flow {
    flow_id    = resource.davinci_flow.registration.id
    version_id = -1
    weight     = 100
  }
}

## Remains the same
resource "pingone_application" "oidc_sdk_sample_app" {
  environment_id = var.pingone_environment_id
  enabled        = true
  name           = "Sample App"
  description    = "A custom sample OIDC application to demonstrate PingOne integration."

  oidc_options {
    type                        = "SINGLE_PAGE_APP"
    grant_types                 = ["AUTHORIZATION_CODE", "IMPLICIT", "REFRESH_TOKEN"]
    response_types              = ["CODE", "TOKEN", "ID_TOKEN"]
    pkce_enforcement            = "S256_REQUIRED"
    token_endpoint_authn_method = "NONE"
    redirect_uris               = var.redirect_uris
    post_logout_redirect_uris   = ["${var.app_url}"]
  }
}

resource "pingone_application_flow_policy_assignment" "login_flow" {
  environment_id = module.environment.environment_id
  application_id = pingone_application.oidc_sdk_sample_app.id
  ## Simplified to point to the id of the desired flow policy resource
  flow_policy_id = davinci_application_flow_policy.registration_flow_policy.id

  priority = 1
}
```

### Import Existing Policies

Before applying the new configuration, import any managed resources that were migrated. 
In the example above, this would mean running:

```

```