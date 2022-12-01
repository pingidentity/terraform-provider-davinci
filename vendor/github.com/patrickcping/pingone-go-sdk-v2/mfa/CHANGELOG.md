# v0.7.0 (2022-11-06)

* **Feature** Added FIDO Policy API and model [#75](https://github.com/patrickcping/pingone-go-sdk-v2/pull/75)
* **Feature** Added FIDO Custom Device Metadata API and model [#75](https://github.com/patrickcping/pingone-go-sdk-v2/pull/75)

# v0.6.1 (2022-10-10)

* **Bug fix** `lockout` made optional in the `MFASettings` model [#70](https://github.com/patrickcping/pingone-go-sdk-v2/pull/70)

# v0.6.0 (2022-10-09)

* **Bug fix** Corrected the Device policy API [#65](https://github.com/patrickcping/pingone-go-sdk-v2/pull/65)
* **Bug fix** Corrected the Application push credentials API model [#67](https://github.com/patrickcping/pingone-go-sdk-v2/pull/67)

# v0.5.1 (2022-09-18)

* **Enhancement** - Changed model dereferencing strategy for the `CreateMFAPushCredential201Response` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `CreateMFAPushCredentialRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `EntityArrayEmbeddedPushCredentialsInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `UpdateMFAPushCredentialRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Add `environment` attribute block to `MFASettings` model [#50](https://github.com/patrickcping/pingone-go-sdk-v2/pull/50)
* **Enhancement** - Add required attributes to the `MFASettings` model [#50](https://github.com/patrickcping/pingone-go-sdk-v2/pull/50)

# v0.5.0 (2022-09-11)

* **Feature** Added MFA Settings model [#44](https://github.com/patrickcping/pingone-go-sdk-v2/pull/44)
* **Breaking change** `updatedAt` attributes changed to date/time data type [#44](https://github.com/patrickcping/pingone-go-sdk-v2/pull/44)

# v0.4.0 (2022-08-29)

* **Enhancement** Add 404 response processing for all requests [#28](https://github.com/patrickcping/pingone-go-sdk-v2/pull/28)
* **Enhancement** Add missing 400 response processing [#29](https://github.com/patrickcping/pingone-go-sdk-v2/pull/29)

# v0.3.0 (2022-08-05)

* **Feature** Mfa device authentication policies [#13](https://github.com/patrickcping/pingone-go-sdk-v2/pull/13)

# v0.2.0 (2022-07-17)

* **Security** Bump `golang.org/x/net` for `CVE-2021-33194` : https://deps.dev/advisory/osv/GO-2021-0238
* **Feature** Full support for enum values

# v0.1.0 (2022-07-16)

Initial release - rebasing versions to reflect module stability