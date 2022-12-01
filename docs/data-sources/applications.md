---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "davinci_applications Data Source - terraform-provider-davinci"
subcategory: ""
description: |-
  
---

# davinci_applications (Data Source)



## Example Usage

```terraform
data "davinci_applications" "all" {
}

output "davinci_applications" {
  value = data.davinci_applications.all.applications
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) PingOne environment id

### Read-Only

- `applications` (Set of Object) (see [below for nested schema](#nestedatt--applications))
- `id` (String) The ID of this resource.

<a id="nestedatt--applications"></a>
### Nested Schema for `applications`

Read-Only:

- `api_key_enabled` (Boolean)
- `api_keys` (Map of String)
- `application_id` (String)
- `created_date` (Number)
- `customer_id` (String)
- `environment_id` (String)
- `metadata` (Map of String)
- `name` (String)
- `oauth` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--oauth))
- `policies` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--policies))
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

- `allowed_grants` (List of String)
- `allowed_scopes` (List of String)
- `client_secret` (String)
- `enabled` (Boolean)
- `enforce_signed_request_openid` (Boolean)
- `logout_uris` (List of String)
- `redirect_uris` (List of String)
- `sp_jwks_openid` (String)
- `sp_jwks_url` (String)



<a id="nestedobjatt--applications--policies"></a>
### Nested Schema for `applications.policies`

Read-Only:

- `created_date` (Number)
- `name` (String)
- `policy_flows` (Set of Object) (see [below for nested schema](#nestedobjatt--applications--policies--policy_flows))
- `policy_id` (String)
- `status` (String)

<a id="nestedobjatt--applications--policies--policy_flows"></a>
### Nested Schema for `applications.policies.policy_flows`

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

