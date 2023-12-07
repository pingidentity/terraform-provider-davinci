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
