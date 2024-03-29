---
page_title: "davinci_applications Data Source - terraform-provider-davinci"
subcategory: "Application"
description: |-
  
---

# davinci_applications (Data Source)



## Example Usage

```terraform
data "davinci_applications" "all_applications" {
  environment_id = var.environment_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) The ID of the PingOne environment to retrieve applications from. Must be a valid PingOne resource ID.

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `applications` (Set of Object) A set of applications retrieved from the environment. (see [below for nested schema](#nestedatt--applications))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--applications"></a>
### Nested Schema for `applications`

Read-Only:

- `api_key_enabled` (Boolean)
- `api_keys` (Map of String)
- `created_date` (Number)
- `customer_id` (String)
- `environment_id` (String)
- `id` (String)
- `metadata` (Map of String)
- `name` (String)
- `oauth` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--oauth))
- `policy` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--policy))
- `saml` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--saml))
- `user_pools` (Map of String)
- `user_portal` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--user_portal))

<a id="nestedobjatt--applications--oauth"></a>
### Nested Schema for `applications.oauth`

Read-Only:

- `enabled` (Boolean)
- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--oauth--values))

<a id="nestedobjatt--applications--oauth--values"></a>
### Nested Schema for `applications.oauth.values`

Read-Only:

- `allowed_grants` (Set of String)
- `allowed_scopes` (Set of String)
- `client_secret` (String)
- `enabled` (Boolean)
- `enforce_signed_request_openid` (Boolean)
- `logout_uris` (Set of String)
- `redirect_uris` (Set of String)
- `sp_jwks_openid` (String)
- `sp_jwks_url` (String)



<a id="nestedobjatt--applications--policy"></a>
### Nested Schema for `applications.policy`

Read-Only:

- `created_date` (Number)
- `name` (String)
- `policy_flow` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--policy--policy_flow))
- `policy_id` (String)
- `status` (String)

<a id="nestedobjatt--applications--policy--policy_flow"></a>
### Nested Schema for `applications.policy.policy_flow`

Read-Only:

- `flow_id` (String)
- `success_nodes` (List of String)
- `version_id` (Number)
- `weight` (Number)



<a id="nestedobjatt--applications--saml"></a>
### Nested Schema for `applications.saml`

Read-Only:

- `values` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--saml--values))

<a id="nestedobjatt--applications--saml--values"></a>
### Nested Schema for `applications.saml.values`

Read-Only:

- `audience` (String)
- `enabled` (Boolean)
- `enforce_signed_request` (Boolean)
- `redirect_uri` (String)
- `sp_cert` (String)



<a id="nestedobjatt--applications--user_portal"></a>
### Nested Schema for `applications.user_portal`

Read-Only:

- `add_auth_method_title` (String)
- `cred_page_subtitle` (String)
- `cred_page_title` (String)
- `flow_timeout_seconds` (Number)
- `name_auth_method_title` (String)
- `name_confirm_btn_text` (String)
- `remove_auth_method_title` (String)
- `remove_body_message` (String)
- `remove_cancel_btn_text` (String)
- `remove_confirm_btn_text` (String)
- `remove_message` (String)
- `show_logout_button` (Boolean)
- `show_mfa_button` (Boolean)
- `show_user_info` (Boolean)
- `show_variables` (Boolean)
- `up_title` (String)
- `update_body_message` (String)
- `update_message` (String)