---
layout: ""
page_title: "Migrating DaVinci Resources to PingOne Provider"
description: |-
  The guide documents the process of migrating Terraform-managed resources from the legacy DaVinci provider to the new DaVinci resources within the PingOne Terraform provider.
---

# Migrating DaVinci Resources to PingOne Provider

The following guide documents the process of migrating Terraform-managed resources from the legacy DaVinci provider (`pingidentity/davinci`) to the new DaVinci resources within the PingOne Terraform provider (`pingidentity/pingone`).

The goal of this migration is to move configuration managed by the legacy provider to the PingOne provider while minimizing impact to live infrastructure. This involves avoiding deletion or recreation of resources and ensuring that `terraform apply` results in no functional changes during the migration.

## Migration Rationale

Users of the `pingidentity/davinci` provider are recommended to migrate all resources to the `pingidentity/pingone` provider for the following reasons:

### More secure, automation-friendly authentication

- The legacy `pingidentity/davinci` provider relies on human user credentials and browser-based SSO to authenticate.
- The PingOne provider uses PingOne worker applications and standard OAuth client credentials, which is a much better fit for CI/CD pipelines, GitHub Actions, and other non-interactive automation.

### Forward-compatible with PingOne API changes

- The legacy provider is built on an older architecture that does not handle additive changes to the DaVinci APIs well (for example, when new optional properties are added to the flow JSON). In practice, this can lead to situations where a non-breaking change in the platform still causes unexpected diffs or failures when using the old provider.
- The PingOne provider is designed to be forward compatible with these additive changes, reducing the risk of breakage when we evolve the DaVinci APIs.

### First-class support for PingOne DaVinci

- The new DaVinci resources live in the same provider as the rest of PingOne (`pingidentity/pingone`), giving you a single, consistent provider for managing identity services.
- The schema and behavior of the new resources are aligned with our current PingOne v2 APIs and will be the focus of ongoing enhancements.

### Improved migration and day-2 operations

- Using the exporter, you can:
  - Read your live DaVinci configuration from PingOne
  - Generate a version-control-ready Terraform module with proper dependency mapping
  - Produce import blocks and a tfvars file so you can safely bring existing resources under Terraform management
- The recommended migration flow is designed so that your first `terraform plan` and `terraform apply` can be reviewed carefully and, when configured correctly, result in no functional changes to your live flows — only state alignment.

## Prerequisites

Before you begin, ensure you have:

* Existing Terraform configuration managed by the legacy DaVinci provider (`pingidentity/davinci`)
* The `pingcli-terraformer` command line tool installed. See the [Ping CLI Terraformer plugin repository](https://github.com/pingidentity/pingcli-plugin-terraformer) for installation instructions.
* A PingOne worker application with at least the **DaVinci Admin Read Only** role to read the live configuration
* Terraform 1.5 or later installed

-> You will need a higher role such as "DaVinci Admin" or a custom role with write access to the [flow export endpoint](https://developer.pingidentity.com/pingone-api/davinci/davinci-admin-apis/admin-flow-versions/export-flow-version.html) on the worker application to generate DaVinci Variable dependencies within DaVinci Flows. If the `pingcli-terraformer` tool can't access the endpoint, a warning will be returned and generation of the dependency will be skipped.

## Set up authentication

Configure your worker application credentials for use by the `pingcli-terraformer` tool.

```bash
export PINGCLI_PINGONE_ENVIRONMENT_ID="<your-environment-id>"
export PINGCLI_PINGONE_CLIENT_CREDENTIALS_CLIENT_ID="<your-client-id>"
export PINGCLI_PINGONE_CLIENT_CREDENTIALS_CLIENT_SECRET="<your-client-secret>"
export PINGCLI_PINGONE_REGION_CODE="NA"  # or EU, AP, CA, AU
```

## Export Terraform configuration

Use the `export` command to generate Terraform HCL from your live environment. The export:

- Reads all Terraform-supported resources in the live environment
- Converts API responses to a reusable Terraform module
- Abstracts environment-specific values to variables
- Maps dependencies between resources for deployment ordering

Run the export with the `--include-imports` and `--include-values` flags:

```bash
pingcli-terraformer export \
  --include-imports \
  --include-values
```

### Command flags

| Flag | Description |
|------|-------------|
| `--include-imports` | Generates import blocks for each identified resource, enabling you to bring existing infrastructure under Terraform management |
| `--include-values` | Produces a `terraform.tfvars` file populated with actual values from the live environment |

The export creates a version-control-ready module that can be used standalone or composed into a larger root module.

## Review and prepare generated configuration

Before running `terraform apply`, review the generated configuration and handle secret attributes.

### Configure provider

The export generates a `versions.tf` file that specifies the expected PingOne Terraform provider version. Provider authentication should be inherited from the root module. You can find details in the [Provider Authentication documentation](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs#provider-authentication).

Run `terraform init` to initialize the provider and download dependencies:

```bash
terraform init
```

### Update secret values

PingOne DaVinci has secret attributes (such as client application secrets) that are **not readable using APIs**. You must update these manually before applying:

1. Open the generated `ping-export-terraform.auto.tfvars` file
2. Locate and update all fields marked as `# Secret value - provide manually`

!> All of the variable values in the `ping-export-terraform.auto.tfvars` file could be considered sensitive or environment-specific. This file should not be committed to version control. For automated Terraform deployments, the values to these secrets should be managed through the deployment tool or secrets manager.

### Update configuration references

If your existing Terraform configuration references DaVinci resources (e.g., passing resource IDs to other resources or modules), you need to update these references. Since the export creates a **child module**, you cannot directly reference resources within the module from your root configuration. Instead, you must add outputs to the child module.

#### Add module outputs

-> Future enhancements to the tool plan to optionally populate the `outputs.tf` file with common fields automatically, reducing manual configuration.

The export includes an empty `outputs.tf` file in the generated module (e.g., `ping-export/outputs.tf`). Add outputs for any attributes you need to reference from your root module or other configurations.

**Example output definition:**

```hcl
output "dv_app_api_key" {
  description = "API key value for the DaVinci sample application"
  value       = pingone_davinci_application.pingcli__DaVinci-0020-API-0020-Protect-0020-Sample-0020-Application.api_key.value
  sensitive   = true
}
```

#### Update root module references

Update your root module configuration to reference the new module outputs:

**Before (legacy provider):**

```hcl
resource "example_resource" "demo" {
  api_key = davinci_application.registration_flow_app.api_keys.prod
}
```

**After (using module output):**

```hcl
resource "example_resource" "demo" {
  flow_id = module.ping-export.dv_app_api_key
}
```

#### Finding the correct resource names

To determine the exact resource names for your outputs:

1. Look in the generated module's `.tf` files (e.g., `ping-export/davinci-applications.tf`, `ping-export/davinci-flows.tf`)
2. Find the resource you need to reference
3. Use the full resource reference in your output value

-> Resource names in the export are prefixed with `pingcli__` and use `-0020-` to represent spaces in the original resource names.

#### Validate configuration

After updating all configuration references, validate your Terraform configuration:

```bash
terraform validate
```

This ensures all module outputs are correctly defined and referenced.

## Remove legacy resources from state

Before applying the new configuration, remove the legacy DaVinci provider resources from Terraform state. This prevents Terraform from attempting to destroy the live infrastructure when the legacy resource definitions are replaced. Removing a resource from state does **not** delete the actual resource — it only tells Terraform to stop tracking it.

Use the `terraform state rm` command to remove each legacy resource:

```bash
terraform state rm davinci_connection.my_connector
terraform state rm davinci_application.my_app
terraform state rm davinci_flow.my_flow
# ... continue for all legacy resources
```

After removing all legacy resources, you can also remove the legacy provider configuration and resource definitions from your Terraform files.

## Review and apply the plan

Generate a plan and output it to a file for review:

```bash
terraform plan -no-color > tfplan-migration.txt 2>&1
```

### What to expect in the plan

The size of the plan will correspond to the size of your live infrastructure. The initial `terraform plan` and `terraform apply` can produce the following types of actions.

#### Resources to be imported

The majority of resources will show as items that will be brought under Terraform management without any interaction with the live infrastructure. This indicates that the defined HCL for the resource matches what would be stored in state exactly. For example, `module.ping-export.pingone_davinci_variable.pingcli__myVar_flowInstance will be imported`.

The generated `ping-export-imports.tf` file includes an import block and a commented-out `terraform import` command for each discovered resource. The import commands can be run in a terminal prior to `terraform apply` to minimize the size of the Terraform plan and focus on items that indicate live infrastructure interactions. For an environment with around 100 resources, this sequential import process takes approximately 5 minutes.

#### Configuration updates

Resources that show `will be updated in-place` typically indicate that an API call will be made to the live infrastructure. Some updates might be non-functional, even though an API call is made. Functional and non-functional changes can appear on the same resource, so review all resources with planned changes carefully before proceeding.

Common planned changes include:

- **Default values**: The Terraform provider includes default values for certain attributes to maintain consistency in continuous management. For example, the DaVinci Flow resource schema expects a `default_log_level` of `4`, so this value is injected in the generated HCL even if it's not found on the API read. This change will run an API `PUT` to update the flow and bring the live infrastructure into alignment with the defined configuration.
- **Deploy triggers**: If there is an update to a flow, the current version and published version will be out of sync, so the generated configuration will also look to `deploy` the flow for realignment.
- **Computed values**: Changes marked with `~` and ending in `(known after apply)` represent computed attributes updating to their refreshed state. These are non-functional. A resource whose plan shows only computed updates will not appear as an item to change.
- **Resources with secret values**: The PingOne API doesn't allow reading values of attributes that are considered secrets. For these values, the Terraform provider only looks for mismatches between what is stored in state and what is defined in configuration, rather than refreshing state against live infrastructure. If a `terraform import` command was run prior to this step, an obfuscated value will show in Terraform state. The identified change indicates the Terraform provider will run an API `PUT` to bring the resource into alignment. This is a common case for Connector Instances that include client secrets as values. The `properties` attribute of a Connector Instance is considered sensitive, and child values are grouped together. Ensure that your defined configuration on such resources is exact before running `terraform apply`.

#### Flow deploy resources

- Flow deploy resources will show as "will be created"
- These make API calls to the flow deploy endpoint but don't cause actual changes when the current version equals the deployed version
- Consider these similar to a `terraform import` operation — they bring the deployment status into Terraform state

**Example plan output:**

```text
# module.ping-export.pingone_davinci_flow.my_flow will be updated in-place
  ~ resource "pingone_davinci_flow" "pingcli__OOTB-0020---0020-Account-0020-Recovery-0020-by-0020-Email" {
    ...
      ~ current_version   = 1 -> (known after apply)
      ~ settings          = {
          + log_level = 4
        }
        # (4 unchanged attributes hidden)
    }

  # module.ping-export.pingone_davinci_flow_enable.pingcli__OOTB-0020---0020-Account-0020-Recovery-0020-by-0020-Email will be updated in-place
  ~ resource "pingone_davinci_flow_enable" "pingcli__OOTB-0020---0020-Account-0020-Recovery-0020-by-0020-Email" {
      ~ enabled        = true -> (known after apply)
        id             = "01af583c6b951086992eb3c37aed7af5"
        # (2 unchanged attributes hidden)
    }
```

### Apply the configuration

After you're satisfied with the plan:

```bash
terraform apply
```

After approving the plan, your DaVinci resources are fully migrated to the PingOne Terraform provider.

## Advanced: Manual import and state manipulation

!> Directly manipulating `terraform.tfstate` is generally discouraged. The state file is Terraform's internal record of managed infrastructure, and manual edits can introduce inconsistencies, corrupt state, or cause unexpected resource changes. The recommended approach above uses `terraform plan` and `terraform apply` with import blocks, which lets Terraform manage state transitions safely.

~> If performed carefully in a **lower or non-production environment**, manual state manipulation can be a useful technique to grow confidence in the migration before applying it to production. This approach allows you to pre-populate state with accurate secret values, reducing the number of planned changes on the first apply and giving you a clearer picture of what Terraform will actually modify.

### Run import commands

The generated `ping-export-imports.tf` file includes commented-out `terraform import` commands alongside the import blocks. You can run these commands individually to bring resources into Terraform state before running `terraform apply`:

```bash
terraform import module.ping-export.pingone_davinci_connector_instance.pingcli__Variables "b8093f6b-bc03-4c67-af59-eed648c26628/06922a684039827499bdbdd97f49827b"
terraform import module.ping-export.pingone_davinci_connector_instance.pingcli__Flow-0020-Connector "b8093f6b-bc03-4c67-af59-eed648c26628/2581eb287bb1d9bd29ae9886d675f89f"
# ... continue for all resources
```

For an environment with around 100 resources, this sequential import process takes approximately 5 minutes.

### Update obfuscated secrets in state

After all resources are imported, the state file will contain obfuscated values (`******`) for any attributes the DaVinci API considers secrets (such as connector instance `client_secret` properties or variables with secret-type values). You can replace these obfuscated values with actual secret values directly in `terraform.tfstate`.

1. Open the `terraform.tfstate` file and search for the exact string `******`
2. Replace each occurrence with the corresponding actual secret value
3. Save the state file

**Example — before:**

```json
"properties": "{\"clientId\":{\"value\":\"b8093f6b-abcd-1234-abcd-eed648c26628\"},\"clientSecret\":{\"value\":\"******\"},\"envId\":{\"value\":\"b8093f6b-abcd-1234-abcd-eed648c26628\"},\"region\":{\"value\":\"NA\"}}"
```

**Example — after:**

```json
"properties": "{\"clientId\":{\"value\":\"b8093f6b-abcd-1234-abcd-eed648c26628\"},\"clientSecret\":{\"value\":\"ACTUAL-VALUE-HERE\"},\"envId\":{\"value\":\"b8093f6b-abcd-1234-abcd-eed648c26628\"},\"region\":{\"value\":\"NA\"}}"
```

~> Only replace the exact string `******` — this is DaVinci's obfuscation marker. When resources are **created** by Terraform (not imported), the provider stores the initial value from the declared HCL and only watches for declared changes. If a secret value drifts in the live infrastructure, the drift would not be detected. The `terraform.tfstate` file contains sensitive data. Ensure it is properly secured and never committed to version control.

After updating all obfuscated values, run `terraform plan` to confirm that the state accurately reflects the live infrastructure and the defined configuration.

## Troubleshooting

### Import failures

* Verify that your Terraform provider worker application has the correct **DaVinci Admin** role
* Check that the resource IDs in the generated imports file are correct
* Ensure Terraform 1.5 or later is installed

### Secret value errors

* Confirm that all fields marked `# Secret value - provide manually` in the generated `terraform.tfvars` file have been updated
* Verify that secret values are correctly formatted (no extra spaces or newlines)

### Plan shows unexpected changes

* Review the diff carefully: some resources might receive default values on first apply
* Check whether API responses have changed since the export was generated
* Re-run the export if the environment was modified during the import process

## Additional resources

* [Ping CLI Terraformer repository](https://github.com/pingidentity/pingcli-plugin-terraformer)
* [Terraform import documentation](https://www.terraform.io/docs/cli/import/index.html)
* [PingOne DaVinci documentation](https://docs.pingidentity.com/davinci)
