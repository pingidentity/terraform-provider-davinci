## 0.4.0 (Unreleased)

NOTES:

* Bump `github.com/golangci/golangci-lint` from 1.59.0 => 1.59.1 ([#345](https://github.com/pingidentity/terraform-provider-davinci/issues/345))
* Bump `github.com/golangci/golangci-lint` from 1.59.1 => 1.60.1 ([#350](https://github.com/pingidentity/terraform-provider-davinci/issues/350))
* Bump `github.com/hashicorp/go-getter` from 1.7.4 => 1.7.5 ([#345](https://github.com/pingidentity/terraform-provider-davinci/issues/345))
* Bump `github.com/hashicorp/terraform-plugin-docs` from 0.19.3 => 0.19.4 ([#345](https://github.com/pingidentity/terraform-provider-davinci/issues/345))
* Bump `github.com/hashicorp/terraform-plugin-framework` from 1.8.0 => 1.11.0 ([#345](https://github.com/pingidentity/terraform-provider-davinci/issues/345))
* Bump `github.com/patrickcping/pingone-go-sdk-v2/management` from 0.39.0 => 0.43.0 ([#345](https://github.com/pingidentity/terraform-provider-davinci/issues/345))
* Bump `github.com/patrickcping/pingone-go-sdk-v2` from 0.11.9 => 0.12.3 ([#345](https://github.com/pingidentity/terraform-provider-davinci/issues/345))
* Bump `github.com/samir-gandhi/davinci-client-go` from 0.3.0 => 0.4.0 ([#350](https://github.com/pingidentity/terraform-provider-davinci/issues/350))
* bump go version from 1.22.3 => 1.22.5 ([#345](https://github.com/pingidentity/terraform-provider-davinci/issues/345))

BUG FIXES:

* `resource/davinci_flow`: Resolve warnings that state that DaVinci JSON files contain unknown properties (August 2024). ([#350](https://github.com/pingidentity/terraform-provider-davinci/issues/350))

## 0.3.3 (4 June 2024)

NOTES:

* `resource/davinci_flow`: Replace deprecated flow attribute validator. ([#318](https://github.com/pingidentity/terraform-provider-davinci/issues/318))
* bump `github.com/bflad/tfproviderlint` from 0.29.0 => 0.30.0 ([#326](https://github.com/pingidentity/terraform-provider-davinci/issues/326))
* bump `github.com/golangci/golangci-lint` 1.58.2 => 1.59.0 ([#318](https://github.com/pingidentity/terraform-provider-davinci/issues/318))
* bump `github.com/hashicorp/terraform-plugin-docs` from 0.19.2 => 0.19.3 ([#326](https://github.com/pingidentity/terraform-provider-davinci/issues/326))
* bump `github.com/hashicorp/terraform-plugin-framework` 1.7.0 => 1.8.0 ([#318](https://github.com/pingidentity/terraform-provider-davinci/issues/318))
* bump `github.com/hashicorp/terraform-plugin-mux` 0.15.0 => 0.16.0 ([#318](https://github.com/pingidentity/terraform-provider-davinci/issues/318))
* bump `github.com/hashicorp/terraform-plugin-sdk/v2` from 2.33.0 => 2.34.0 ([#326](https://github.com/pingidentity/terraform-provider-davinci/issues/326))
* bump `github.com/patrickcping/pingone-go-sdk-v2/management` 0.38.0 => 0.39.0 ([#318](https://github.com/pingidentity/terraform-provider-davinci/issues/318))
* bump `github.com/patrickcping/pingone-go-sdk-v2` from 0.11.8 => 0.11.9 ([#326](https://github.com/pingidentity/terraform-provider-davinci/issues/326))
* bump `github.com/terraform-linters/tflint` from 0.50.3 => 0.51.1 ([#326](https://github.com/pingidentity/terraform-provider-davinci/issues/326))
* bump go version from 1.21.1 => 1.22.3 ([#326](https://github.com/pingidentity/terraform-provider-davinci/issues/326))

ENHANCEMENTS:

* `resource/davinci_flow`: Added `flow_variables.value` to allow the variable's default value to be updated. ([#325](https://github.com/pingidentity/terraform-provider-davinci/issues/325))

BUG FIXES:

* `resource/davinci_flow`: Fix issue whereby descriptions are not updated. ([#319](https://github.com/pingidentity/terraform-provider-davinci/issues/319))
* `resource/davinci_flow`: Fix issue whereby flow variables cannot be updated, leading to error. ([#325](https://github.com/pingidentity/terraform-provider-davinci/issues/325))
* `resource/davinci_flow`: Fix panic crash when flow instance and/or company variables are included in a flow export. ([#316](https://github.com/pingidentity/terraform-provider-davinci/issues/316))
* `resource/davinci_flow`: Fixed `flow_variables.type` so that it refers to the data type of the variable (as is the original intention), rather than the type of the variable object. ([#325](https://github.com/pingidentity/terraform-provider-davinci/issues/325))
* `resource/davinci_flow`: Where a description is not provided in the Terraform schema, the description from the flow export will be applied as a fallback. ([#319](https://github.com/pingidentity/terraform-provider-davinci/issues/319))

## 0.3.2 (28 May 2024)

NOTES:

* `resource/davinci_flow`: `flow_variables` now includes a plan derived from the flow import JSON. ([#308](https://github.com/pingidentity/terraform-provider-davinci/issues/308))
* bump `github.com/golangci/golangci-lint` 1.55.2 => 1.58.2 ([#307](https://github.com/pingidentity/terraform-provider-davinci/issues/307))
* bump `github.com/hashicorp/go-getter` 1.7.2 => 1.7.4 ([#307](https://github.com/pingidentity/terraform-provider-davinci/issues/307))
* bump `github.com/hashicorp/terraform-plugin-docs` 0.18.0 => 0.19.2 ([#307](https://github.com/pingidentity/terraform-provider-davinci/issues/307))
* bump `github.com/hashicorp/terraform-plugin-go` 0.22.1 => 0.23.0 ([#307](https://github.com/pingidentity/terraform-provider-davinci/issues/307))
* bump `github.com/samir-gandhi/dvgenerate` 0.0.10 => 0.0.11 ([#307](https://github.com/pingidentity/terraform-provider-davinci/issues/307))
* bump `golang.org/x/net` 0.22.0 => 0.25.0 ([#307](https://github.com/pingidentity/terraform-provider-davinci/issues/307))

BUG FIXES:

* `resource/davinci_flow`: Fixed "Error parsing `flow_json`" error when the `flow_json` string is unknown during plan. ([#306](https://github.com/pingidentity/terraform-provider-davinci/issues/306))

## 0.3.1 (12 April 2024)

NOTES:

* bump `github.com/hashicorp/terraform-plugin-framework` 1.5.0 => 1.7.0 ([#286](https://github.com/pingidentity/terraform-provider-davinci/issues/286))
* bump `github.com/hashicorp/terraform-plugin-go` 0.21.0 => 0.22.1 ([#286](https://github.com/pingidentity/terraform-provider-davinci/issues/286))
* bump `github.com/hashicorp/terraform-plugin-mux` 0.14.0 => 0.15.0 ([#286](https://github.com/pingidentity/terraform-provider-davinci/issues/286))
* bump `github.com/hashicorp/terraform-plugin-sdk/v2` 2.31.0 => 2.33.0 ([#286](https://github.com/pingidentity/terraform-provider-davinci/issues/286))
* bump `github.com/katbyte/terrafmt` 0.5.2 => 0.5.3 ([#286](https://github.com/pingidentity/terraform-provider-davinci/issues/286))
* bump `github.com/samir-gandhi/davinci-client-go` 0.2.0 => 0.3.0 ([#288](https://github.com/pingidentity/terraform-provider-davinci/issues/288))

BUG FIXES:

* `resource/davinci_flow`: Fixed inconsistent state errors when importing a flow with a non-default log setting. ([#288](https://github.com/pingidentity/terraform-provider-davinci/issues/288))
* `resource/davinci_flow`: Fixed panic error when the flow JSON contains a flow conductor node which isn't using a subflow capability. ([#285](https://github.com/pingidentity/terraform-provider-davinci/issues/285))

## 0.3.0 (10 April 2024)

BREAKING CHANGES:

* `resource/davinci_flow`: The `flow_json` parameter does not now represent the flow after import.  It now only represents the flow JSON payload from the source system.  When needing to use the resulting flow after import in the target environment, `flow_export_json` should be used. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))

NOTES:

* Added plugin mux factory and plugin framework (v6 protocol) provider to facilitate migration from SDKv2 (v5 protocol). ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* Added regex validation for resource/data-source parameters that require platform IDs. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* Code optimisations to remove unnecessary SDK retry logic. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* Index documentation updated to reflect latest best practices for admin role requirements and creation of PingOne environments that do not have demo/bootstrapped configuration automatically applied. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* Provider updated thoughout to support environments that have been created without demo/bootstrapped configuration automatically applied.  Creation of environments without demo/bootstrapped configuration is best practice going forward. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* Upgrade GO to `v1.21`. ([#245](https://github.com/pingidentity/terraform-provider-davinci/issues/245))
* `data-source/davinci_application`: Deprecated the `user_portal` and `saml` block parameters as they are no longer used.  The parameters will be removed in the next major release. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_application`: Un-deprecate the `application_id` parameter, and deprecated the `id` parameter for the purpose of fetching an application by it's ID.  The `application_id` should be used going forward. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_applications`: Address code scanning advisories. ([#245](https://github.com/pingidentity/terraform-provider-davinci/issues/245))
* `data-source/davinci_applications`: Corrected documentation category. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_applications`: Deprecated the `user_portal` and `saml` block parameters as they are no longer used.  The parameters will be removed in the next major release. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_connections`: Deprecated the `connections.company_id` parameter as it is a duplicate of the `environment_id` parameter. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_application_flow_policy`: Address code scanning advisories. ([#245](https://github.com/pingidentity/terraform-provider-davinci/issues/245))
* `resource/davinci_application_flow_policy`: Fix import documentation example. ([#248](https://github.com/pingidentity/terraform-provider-davinci/issues/248))
* `resource/davinci_application`: Deprecated the `user_portal` and `saml` parameters as they are no longer used.  The parameters will be removed in the next major version release. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_connection`: Moved connector reference to separate guide. ([#279](https://github.com/pingidentity/terraform-provider-davinci/issues/279))
* `resource/davinci_connection`: Updated connector reference documentation. ([#245](https://github.com/pingidentity/terraform-provider-davinci/issues/245))
* `resource/davinci_connection`: When the `property.type` parameter is not set for a property object, the default is now set to `string`. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: Address code scanning advisories. ([#247](https://github.com/pingidentity/terraform-provider-davinci/issues/247))
* `resource/davinci_flow`: Migrated to plugin framework. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: The `deploy` parameter is now deprecated.  Deployment on import and update is now implicit.  This parameter will be removed in the next major release. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: The `flow_variables` computed attribute has changed data type.  Previously the attribute was a block type.  Going forward, the attribute is nested set type.  There are no changes expected to HCL to use the new data type. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: Updated warning messages when subflows and connectors are left unmapped.  Going forward, all subflows and connections in a flow should be mapped using the `connection_link` and `subflow_link` parameters. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* bump `github.com/cloudflare/circl` 1.3.3 => 1.3.7 ([#278](https://github.com/pingidentity/terraform-provider-davinci/issues/278))
* bump `github.com/go-git/go-git/v5` v5.9.0 => v5.11.0 ([#242](https://github.com/pingidentity/terraform-provider-davinci/issues/242))
* bump `github.com/hashicorp/terraform-plugin-docs` 0.16.0 => 0.18.0 ([#278](https://github.com/pingidentity/terraform-provider-davinci/issues/278))
* bump `github.com/hashicorp/terraform-plugin-sdk/v2` v2.30.0 => v2.31.0 ([#246](https://github.com/pingidentity/terraform-provider-davinci/issues/246))
* bump `github.com/patrickcping/pingone-go-sdk-v2/management` 0.34.0 => 0.36.0 ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* bump `github.com/patrickcping/pingone-go-sdk-v2/management` 0.36.0 => 0.38.0 ([#278](https://github.com/pingidentity/terraform-provider-davinci/issues/278))
* bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.32.0 => v0.34.0 ([#246](https://github.com/pingidentity/terraform-provider-davinci/issues/246))
* bump `github.com/patrickcping/pingone-go-sdk-v2` 0.11.5 => 0.11.8 ([#278](https://github.com/pingidentity/terraform-provider-davinci/issues/278))
* bump `github.com/samir-gandhi/davinci-client-go` 0.0.55 => 0.1.0 ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* bump `github.com/samir-gandhi/davinci-client-go` 0.1.0 => 0.2.0 ([#279](https://github.com/pingidentity/terraform-provider-davinci/issues/279))
* bump `github.com/terraform-linters/tflint` 0.48.0 => 0.50.3 ([#278](https://github.com/pingidentity/terraform-provider-davinci/issues/278))
* bump `golang.org/x/crypto` v0.16.0 => v0.17.0 ([#241](https://github.com/pingidentity/terraform-provider-davinci/issues/241))
* bump `google.golang.org/protobuf` 1.31.0 => 1.33.0 ([#278](https://github.com/pingidentity/terraform-provider-davinci/issues/278))

ENHANCEMENTS:

* Added ability to append custom text information to the default User Agent. ([#249](https://github.com/pingidentity/terraform-provider-davinci/issues/249))
* `data-source/davinci_applications`: Added a configurable timeout parameter and updated the default timeout to `20m` due to eventual consistency considerations on environment creation. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_connection`: Added a configurable timeout parameter and updated the default timeout to `20m` due to eventual consistency considerations on environment creation. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_connection`: Added a new `connection_id` parameter and deprecated the `id` parameter for the purpose of fetching a connection by it's ID.  The `connection_id` should be used going forward. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_connections`: Added a configurable timeout parameter and updated the default timeout to `20m` due to eventual consistency considerations on environment creation. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_application_flow_policy`: Added `policy_flow.allowed_ip_list` to be able to configure an allowed IP list for policy flows. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_connection`: Added ability to configure complex properties as JSON format. ([#279](https://github.com/pingidentity/terraform-provider-davinci/issues/279))
* `resource/davinci_flow`: Added the `connection_link.replace_import_connection_id` parameter to be able to replace the name of connectors with the specified connection ID in a flow import. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: Added the `description` parameter to be able to override the description of the flow on import. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: Added the `flow_configuration_json` and `flow_export_json` attributes to the resource.  `flow_configuration_json` is used to compute configuration changes and drift, while `flow_export_json` is used as a record of the resulting flow once imported to the target environment. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: Added the `subflow_link.replace_import_subflow_id` parameter to be able to replace the name of subflows with the specified subflow ID in a flow import. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: The `name` parameter is now an optional field, to be able to override the name of the flow on import. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))

BUG FIXES:

* Fixed bug where resources may be created in the incorrect PingOne environment when multiple DaVinci environments are being configured in the same `apply` routine. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* Fixed intermittent `Unable to retrieve access_token within 60s for environment` error. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_applications`: Fixed bug where not all applications are returned for an environment due to eventual consistency considerations on environment creation. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_applications`: Fixed issue where not all connections are returned due to eventual consistency considerations on environment creation. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_connection`: Fixed `unable to identify value type, only string or boolean is currently supported` error when reading a connection that has an integer property value. ([#276](https://github.com/pingidentity/terraform-provider-davinci/issues/276))
* `data-source/davinci_connection`: Fixed issue where a connection isn't returned due to eventual consistency considerations on environment creation. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `data-source/davinci_connections`: Fixed `unable to identify value type, only string or boolean is currently supported` error when reading connections that have an integer property value. ([#276](https://github.com/pingidentity/terraform-provider-davinci/issues/276))
* `data-source/davinci_connections`: Fixed the `connections.customer_id` attribute not being stored into state. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_application_flow_policy`: Fixed error when specifying a flow policy without `policy_flow.name` or `policy_flow.version` parameters.  The `policy_flow.name` and `policy_flow.version` are now required fields. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_application_flow_policy`: Fixed error when specifying a flow policy without any `policy_flow` blocks.  The `policy_flow` block is now a required field. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_application_flow_policy`: Fixed inability to set `policy_flow.success_nodes`. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_application_flow_policy`: Fixed inability to update the environment ID and application ID once created (resource now requires replacement). ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_application`: Fixed inability to set `api_key_enabled`, `oauth.enabled` and `oauth.values.enabled` to `false`. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_application`: Fixed panic crash when defining an application with an `oauth` block, but without `oauth.values`. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_connection`: Fixed `Error retrieving connectors / 7005` error in certain conditions. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_connection`: Fixed `unable to identify value type, only string or boolean is currently supported` error when importing a connection that has an integer property value. ([#276](https://github.com/pingidentity/terraform-provider-davinci/issues/276))
* `resource/davinci_connection`: Fixed inability to import the `defaultUserPool` User Pool connector to Terraform state. ([#279](https://github.com/pingidentity/terraform-provider-davinci/issues/279))
* `resource/davinci_connection`: Fixed inability to update the environment ID, connection name and ID once created (resource now requires replacement). ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_connection`: Fixed inconsistent plan when specifying a property data type. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: Corrected flow drift calculation errors, causing plan inconsistency. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: Corrected panic errors on flow drift calculation and connection/subflow re-mapping. ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_flow`: Fixed inability to update the environment ID after initial configuration (resource requires replacement). ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))
* `resource/davinci_variable`: Fixed inability to update the environment ID once created (resource now requires replacement). ([#250](https://github.com/pingidentity/terraform-provider-davinci/issues/250))

## 0.2.1

BUG FIXES:

* `resource/davinci_application`: add policy block back with deprecation notice ([#236](https://github.com/pingidentity/terraform-provider-davinci/issues/236))

## 0.2.0

NOTES:

* Added `terraform import` documentation to all applicable resources. ([#195](https://github.com/pingidentity/terraform-provider-davinci/issues/195))

ENHANCEMENTS:

* `resource/application_flow_policy`: separate application policies to have their own lifecycle and avoid affecting upstream apps. ([#214](https://github.com/pingidentity/terraform-provider-davinci/issues/214))

BUG FIXES:

* `internal/sdk`: capture additional retryable errors on flow sign-in ([#128](https://github.com/pingidentity/terraform-provider-davinci/issues/128))
* `resource/davinci_application`: Fixed error when attempting to import resource state. ([#196](https://github.com/pingidentity/terraform-provider-davinci/issues/196))
* `resource/davinci_connection`: Fixed error when attempting to import resource state. ([#196](https://github.com/pingidentity/terraform-provider-davinci/issues/196))
* `resource/davinci_flow`: Fixed error when attempting to import resource state. ([#196](https://github.com/pingidentity/terraform-provider-davinci/issues/196))
* `resource/davinci_flow`: additional error handling for misconfigured flows ([#230](https://github.com/pingidentity/terraform-provider-davinci/issues/230))
* `resource/davinci_variable`: Fixed error when attempting to import resource state. ([#196](https://github.com/pingidentity/terraform-provider-davinci/issues/196))

## 0.1.13

BUG FIXES:

* `datasource_connection`: update datasource to account for empty value ([#204](https://github.com/pingidentity/terraform-provider-davinci/issues/204))
* `resource_connection`: handling for resource removed without terraform. ([#191](https://github.com/pingidentity/terraform-provider-davinci/issues/191))

## 0.1.12

ENHANCEMENTS:

* `provider`: Support for custom User Agent API request header ([#182](https://github.com/pingidentity/terraform-provider-davinci/issues/182))

BUG FIXES:

* `resource_application`: Add full CRUD logic to application policies ([#183](https://github.com/pingidentity/terraform-provider-davinci/issues/183))

## 0.1.11

BUG FIXES:

* `resource_connection`: creation and update of connection directly updates state. Read avoids obfuscated values. ([#168](https://github.com/pingidentity/terraform-provider-davinci/issues/168))

DOCUMENTATION:

* `provider`: Add example of new PingOne role usage: "DaVinci Admin" ([#165](https://github.com/pingidentity/terraform-provider-davinci/issues/165))

## 0.1.10

ENHANCEMENTS:

* `provider`: Support host_url as provider input parameter
`davinci-client-go`: Update to v0.0.53 for host_url support ([#151](https://github.com/pingidentity/terraform-provider-davinci/issues/151))

BUG FIXES:

* `internal_sweep_client`: Update model to match new pingone client schema. ([#163](https://github.com/pingidentity/terraform-provider-davinci/issues/163))

## 0.1.9

ENHANCEMENTS:

* `provider`: added handling to optionally accept davinci access token in place of username/password ([#124](https://github.com/pingidentity/terraform-provider-davinci/issues/124))

## 0.1.8

BUG FIXES:

* `davinci_api_client`: update to v0.0.52. http default timeout raised to 300s ([#131](https://github.com/pingidentity/terraform-provider-davinci/issues/131))

## 0.1.7

BUG FIXES:

* `davinci_flow`: Diff handling for flow_json.settings.logLevel ([#125](https://github.com/pingidentity/terraform-provider-davinci/issues/125))

DOCUMENTATION:

* `davinci_connection`: added plugin to generate connection detail docs

## 0.1.6

NOTES:

* Updated the index document to refer to detailed getting started guide at `terraform.pingidentity.com`. ([#104](https://github.com/pingidentity/terraform-provider-davinci/issues/104))
* `GNUMakefile`: updated release validation checks and preparation for multi-region tests. ([#85](https://github.com/pingidentity/terraform-provider-davinci/issues/85))

BUG FIXES:

* `davinci_flow`: switch for subflowVersionId types ([#106](https://github.com/pingidentity/terraform-provider-davinci/issues/106))
* `go_client`: v0.0.48 corrected logic for empty nodes of type CONNECTION ([#98](https://github.com/pingidentity/terraform-provider-davinci/issues/98))

## 0.1.5

BUG FIXES:

* `go_client`: fixed flow variable loop logic
`davinci_flow`: correct variable diff function ([#95](https://github.com/pingidentity/terraform-provider-davinci/issues/95))

## 0.1.4

BUG FIXES:

* `go_client`: Updated Models for flow and variable
`davinci_flow`: Added computed flow_variable section ([#93](https://github.com/pingidentity/terraform-provider-davinci/issues/93))

## 0.1.3

NOTES:

* Added attribute schema to the documentation index. ([#70](https://github.com/pingidentity/terraform-provider-davinci/issues/70))
* Updated index documentation examples. ([#70](https://github.com/pingidentity/terraform-provider-davinci/issues/70))
* `davinci_connection`: Adjusted example HCL. ([#70](https://github.com/pingidentity/terraform-provider-davinci/issues/70))

ENHANCEMENTS:

* `davinci-client-go`: Updated application create function for p1 session flow policies ([#69](https://github.com/pingidentity/terraform-provider-davinci/issues/69))

BUG FIXES:

* `davinci_application`: Updated read funtion to unset id if application is not found ([#73](https://github.com/pingidentity/terraform-provider-davinci/issues/73))

## 0.1.1

BUG FIXES:

* data_source_application.go: added `id` field. added deprecation notice for `application_id`
data_source_application_test.go: organized tests to get by `application_id` and get by `id`
data_source_connnection.go: updated to use correct `d.SetId()` function
resource_application.go: updated to use correct `d.SetId()` function
resource_connection.go: updated to use correct `d.SetId()` function
resource_flow.go: removed `id` schema element. updated to use correct `d.SetId()` function ([#63](https://github.com/pingidentity/terraform-provider-davinci/issues/63))

## 0.1.0

ENHANCEMENTS:

* data_source_application.go: added `id` field. added deprecation notice for `application_id`
data_source_application_test.go: organized tests to get by `application_id` and get by `id`
data_source_connnection.go: updated to use correct `d.SetId()` function
resource_application.go: updated to use correct `d.SetId()` function
resource_connection.go: updated to use correct `d.SetId()` function
resource_flow.go: removed `id` schema element. updated to use correct `d.SetId()` function ([#63](https://github.com/pingidentity/terraform-provider-davinci/issues/63))

## 0.1.0

ENHANCEMENTS:

* *_davinci_application: **BREAKING CHANGE** - Change schema block `policies` to `policy`
* *_davinci_application: **BREAKING CHANGE** - Change schema block `policy_flows` to `policy_flow`
resource_davinci_flow: **BREAKING CHANGE** - Change schema block `connections` to `connection_link`
resource_davinci_flow: **BREAKING CHANGE** - Change schema block `subflows` to `subflow_link`
* *_davinci_connection: **BREAKING CHANGE** - Change schema block `properties` to `property ([#42](https://github.com/pingidentity/terraform-provider-davinci/issues/42))
* *_test.go: Updated all tests to use minimal role strategy for main admin user. ([#47](https://github.com/pingidentity/terraform-provider-davinci/issues/47))
* datasource_davinci_connection: Added filter to get connection by name ([#29](https://github.com/pingidentity/terraform-provider-davinci/issues/29))
* resource_davinci_flow: updated connection and subflow dependency fields to `id` and `name`
resource_davinci_connection: updated connection_id and connection_name to `id` and `name`
datasource_davinci_connection: updated connection_id and connection_name to `id` and `name`
datasource_davinci_connections: updated connection_id and connection_name to `id` and `name` ([#33](https://github.com/pingidentity/terraform-provider-davinci/issues/33))

BUG FIXES:

* resource_davinci_application: removed unnecessary application_id property ([#36](https://github.com/pingidentity/terraform-provider-davinci/issues/36))
* go-client: updated client to v0.0.41 for removed omitempty on node data properties ([#32](https://github.com/pingidentity/terraform-provider-davinci/issues/32))
* resource_davinci_application.go: Marked connection property value field at "Sensitive"
data_source_davinci_application.go: Marked connection property value field at "Sensitive" ([#40](https://github.com/pingidentity/terraform-provider-davinci/issues/40))
* resource_davinci_connection.go: Marked connection property value field at "Sensitive" ([#45](https://github.com/pingidentity/terraform-provider-davinci/issues/45))
* sdk: corrected ordering of wait for bootstrap to ensure 5s sleep
resource_flow_test: added test to validate wait for bootstrap
acctest: added models for new resource_flow_test ([#30](https://github.com/pingidentity/terraform-provider-davinci/issues/30))

## 0.0.8

ENHANCEMENTS:

* davinci_flow: closes #14 - Save `davinci_flow.connections` and `davinci_flow.subflows` to terraform state
davinci_flow: Add warning when missing relevant connection dependency ([#14](https://github.com/pingidentity/terraform-provider-davinci/issues/14))

BUG FIXES:

* davinci_flow: `flow_json` is now sensitive
davinci_connection: `properties.value` is now sensitive
davinci_variable: `value` is now sensitive ([#16](https://github.com/pingidentity/terraform-provider-davinci/issues/16))
* davinci_flow: add `Importer` ([#14](https://github.com/pingidentity/terraform-provider-davinci/issues/14))
* sdk: added missed timeout counter increment ([#19](https://github.com/pingidentity/terraform-provider-davinci/issues/19))

## 0.0.6 (Unreleased)

BACKWARDS INCOMPATIBILITIES / NOTES:
