## Changelogs

- This is a Terraform project. Whenever a change is made to the schema of a resource or data source, a changelog entry must be created according to the instructions at `contributing/changelog-process.md`
- There should be one changelog entry block per change, in the same file, and the file should be named after the PR number (if the PR number is not known, use `tbc`)
- The changelog entry text should have the resource or data source as the prefix, in markdown backticks. For example "`resource/pingone_flow`: Added new thing."
- The changelog entry should be in past tense, something has been added, changed or removed, and should be correctly punctuated.
- Changelog entries should be created for go module version updates imported directly and only in the go.mod in the root of the project.
- Changelog entires for go module version updates should be `note` types and have the format "bump `{dependency}` {old version} => {new version}"

## Testing
- Every code change should be tested. Only impacted tests should be run. An example command is `TF_ACC=1 go test -v -timeout 1200s -run ^TestAccForm_ github.com/pingidentity/terraform-provider-davinci/internal/service/davinci`.  In this case, the `TestAccForm_` is the prefix or name of the impacted test to run.
- If using 1Password for secrets management, the test command should include `op run --`, as in this example: `TF_ACC=1 op run -- go test -v -timeout 1200s -run ^TestAccForm_ github.com/pingidentity/terraform-provider-davinci/internal/service/davinci`

## Documentation
- If a Terraform schema is changed, use `make generate` to regenerate the Terraform documentation, `make lint` to check linting.
- If any HCL example file is changed, use `make generate` to regenerate the Terraform documentation.

## Project version

- When updating the project version, the version number should be incremented in the makefile.
- The `CHANGELOG.md` file should be updated with a new version header, and "Unreleased" should be used in lieu of a release date.
- If the version number is a minor or major version, the example HCL files in `examples/provider/*.tf` also need to be updated.
