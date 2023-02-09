## 0.1.0

ENHANCEMENTS:

* *_davinci_application: **BREAKING CHANGE** - Change schema block `policies` to `policy`
* *_davinci_application: **BREAKING CHANGE** - Change schema block `policy_flows` to `policy_flow`
resource_davinci_flow: **BREAKING CHANGE** - Change schema block `connections` to `connection_link`
resource_davinci_flow: **BREAKING CHANGE** - Change schema block `subflows` to `subflow_link`
* *_davinci_connection: **BREAKING CHANGE** - Change schema block `properties` to `property ([#42](https://github.com/pingidentity/terraform-provider-pingone/issues/42))
* *_test.go: Updated all tests to use minimal role strategy for main admin user. ([#47](https://github.com/pingidentity/terraform-provider-pingone/issues/47))
* datasource_davinci_connection: Added filter to get connection by name ([#29](https://github.com/pingidentity/terraform-provider-pingone/issues/29))
* resource_davinci_flow: updated connection and subflow dependency fields to `id` and `name`
resource_davinci_connection: updated connection_id and connection_name to `id` and `name`
datasource_davinci_connection: updated connection_id and connection_name to `id` and `name`
datasource_davinci_connections: updated connection_id and connection_name to `id` and `name` ([#33](https://github.com/pingidentity/terraform-provider-pingone/issues/33))

BUG FIXES:

* resource_davinci_application: removed unnecessary application_id property ([#36](https://github.com/pingidentity/terraform-provider-pingone/issues/36))
* go-client: updated client to v0.0.41 for removed omitempty on node data properties ([#32](https://github.com/pingidentity/terraform-provider-pingone/issues/32))
* resource_davinci_application.go: Marked connection property value field at "Sensitive"
data_source_davinci_application.go: Marked connection property value field at "Sensitive" ([#40](https://github.com/pingidentity/terraform-provider-pingone/issues/40))
* resource_davinci_connection.go: Marked connection property value field at "Sensitive" ([#45](https://github.com/pingidentity/terraform-provider-pingone/issues/45))
* sdk: corrected ordering of wait for bootstrap to ensure 5s sleep
resource_flow_test: added test to validate wait for bootstrap
acctest: added models for new resource_flow_test ([#30](https://github.com/pingidentity/terraform-provider-pingone/issues/30))

## 0.0.8

ENHANCEMENTS:

* davinci_flow: closes #14 - Save `davinci_flow.connections` and `davinci_flow.subflows` to terraform state
davinci_flow: Add warning when missing relevant connection dependency ([#14](https://github.com/pingidentity/terraform-provider-pingone/issues/14))

BUG FIXES:

* davinci_flow: `flow_json` is now sensitive
davinci_connection: `properties.value` is now sensitive
davinci_variable: `value` is now sensitive ([#16](https://github.com/pingidentity/terraform-provider-pingone/issues/16))
* davinci_flow: add `Importer` ([#14](https://github.com/pingidentity/terraform-provider-pingone/issues/14))
* sdk: added missed timeout counter increment ([#19](https://github.com/pingidentity/terraform-provider-pingone/issues/19))

## 0.0.6 (Unreleased)

BACKWARDS INCOMPATIBILITIES / NOTES:
