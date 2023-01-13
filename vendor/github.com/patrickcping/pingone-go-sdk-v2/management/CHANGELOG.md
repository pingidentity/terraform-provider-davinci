# v0.14.0 (2023-01-12)

* **Bug fix** - Correct the `TemplateContent` API model [#97](https://github.com/patrickcping/pingone-go-sdk-v2/pull/97)

# v0.13.0 (2023-01-09)

* **Breaking change** Moved `AssignActorRoles` from all `Application` models to just `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Breaking change** Moved `Tags` from all `Application` models to just `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Breaking change** Moved `SupportUnsignedRequestObject` from all `Application` models to just `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Note** bump codegen v6.2.0 => v6.2.1 [#84](https://github.com/patrickcping/pingone-go-sdk-v2/pull/84)
* **Feature** Support for Notifications Settings [#85](https://github.com/patrickcping/pingone-go-sdk-v2/pull/85)
* **Feature** Support for Notifications Policies [#85](https://github.com/patrickcping/pingone-go-sdk-v2/pull/85)
* **Feature** Support for Notifications Templates and Contents [#85](https://github.com/patrickcping/pingone-go-sdk-v2/pull/85)
* **Enhancement** Add support for the WS-Federation application type [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Enhancement** Add support for `HiddenFromAppPortal` property on the `Application` models [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Enhancement** Add support for `AllowWildcardInRedirectUris` property on the `ApplicationOIDC` model [#96](https://github.com/patrickcping/pingone-go-sdk-v2/pull/96)
* **Enhancement** Add support for `InitiateLoginUri` property on the `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Enhancement** Add support for `RefreshTokenRollingGracePeriodDuration` property on the `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Enhancement** Add support for `TargetLinkUri` property on the `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Enhancement** Add support for `HomePageUrl` property on the `ApplicationSAML` model [#96](https://github.com/patrickcping/pingone-go-sdk-v2/pull/96)
* **Enhancement** Add boolean data type support to the Sign On Policy `Equals` Condition [#93](https://github.com/patrickcping/pingone-go-sdk-v2/pull/93)

# v0.12.0 (2022-11-06)

* **Breaking change** Removed the `EnumLicensePackage` enum model [#81](https://github.com/patrickcping/pingone-go-sdk-v2/pull/81)
* **Feature** Support for Branding Settings [#76](https://github.com/patrickcping/pingone-go-sdk-v2/pull/76)
* **Feature** Support for Branding Themes [#76](https://github.com/patrickcping/pingone-go-sdk-v2/pull/76)
* **Feature** Support for Resource Client Secrets [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
* **Enhancement** Add `idToken` and `userInfo` attributes to the `ResourceAttribute` data model [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
* **Enhancement** Add `mappedClaims` attributes to the `ResourceScope` data model [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
* **Enhancement** Add `introspectEndpointAuthMethod` attributes to the `Resource` data model [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
* **Enhancement** Added the `TERMINATED` value to the `EnumLicenseStatus` model [#81](https://github.com/patrickcping/pingone-go-sdk-v2/pull/81)

# v0.11.2 (2022-10-15)

* **Bug fix** - Correct the `Image` API model [#74](https://github.com/patrickcping/pingone-go-sdk-v2/pull/74)

# v0.11.1 (2022-10-10)

* **Enhancement** Add `uriPrefix` (mobile) and `excludedPlatforms` (mobile device integrity check) to the `ApplicationOIDC` data model [#71](https://github.com/patrickcping/pingone-go-sdk-v2/pull/71)

# v0.11.0 (2022-10-09)

* **Feature** Support for Licenses [#64](https://github.com/patrickcping/pingone-go-sdk-v2/pull/64)
* **Feature** Support for Languages [#63](https://github.com/patrickcping/pingone-go-sdk-v2/pull/63)

# v0.10.0 (2022-09-18)

* **Bug fix** - Correct the http endpoint headers object in the `Subscription` model [#52](https://github.com/patrickcping/pingone-go-sdk-v2/pull/52)
* **Bug fix** - Correct the `createGroupNesting` response object [#53](https://github.com/patrickcping/pingone-go-sdk-v2/pull/53)
* **Breaking change** `CreateApplication201Response` changes to `ReadOneApplication200Response` model [#49](https://github.com/patrickcping/pingone-go-sdk-v2/pull/49)
* **Feature** Support for External Link applications [#49](https://github.com/patrickcping/pingone-go-sdk-v2/pull/49)
* **Feature** Support for the PingOne Portal application [#49](https://github.com/patrickcping/pingone-go-sdk-v2/pull/49)
* **Feature** Support for the PingOne Self Service application [#49](https://github.com/patrickcping/pingone-go-sdk-v2/pull/49)
* **Feature** Add the `readOneGroupNesting` API [#53](https://github.com/patrickcping/pingone-go-sdk-v2/pull/53)
* **Feature** Support for Alerts [#54](https://github.com/patrickcping/pingone-go-sdk-v2/pull/54)
* **Enhancement** Add `type` attribute to the `GroupNesting` data model [#53](https://github.com/patrickcping/pingone-go-sdk-v2/pull/53)
* **Enhancement** Support for `kerberos` attribute model in the `Application` data model [#49](https://github.com/patrickcping/pingone-go-sdk-v2/pull/49)
* **Enhancement** - Add `PINGID_WINLOGIN_PASSWORDLESS_AUTHENTICATION` and `PINGID_AUTHENTICATION` to the `EnumSignOnPolicyType` model for workforce based sign on policy actions [#47](https://github.com/patrickcping/pingone-go-sdk-v2/pull/47)
* **Enhancement** - Added `SignOnPolicyActionPingIDWinLoginPasswordless` model to support the PingID Windows Passwordless sign on policy action [#47](https://github.com/patrickcping/pingone-go-sdk-v2/pull/47)
* **Enhancement** - Added `application` attribute to the `ApplicationAttributeMapping` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `CreateApplication201Response` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `CreateApplicationRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `CreateGateway201Response` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `CreateGatewayRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `EntityArrayEmbeddedApplicationsInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `EntityArrayEmbeddedAttributesInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `EntityArrayEmbeddedGatewaysInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `IdentityProvider` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `SignOnPolicyAction` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `SignOnPolicyActionCommonConditionOrOrInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `UpdateApplicationRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `UpdateDomain200Response` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)

# v0.9.0 (2022-09-11)

* **Bug fix** - Correct the `CreateDomain` API response object [#45](https://github.com/patrickcping/pingone-go-sdk-v2/pull/45)
* **Bug fix** - Correct the `UpdateDomain` API response object [#45](https://github.com/patrickcping/pingone-go-sdk-v2/pull/45)
* **Bug fix** - Correct the `CustomDomainCertificate` conflict by adding `CustomDomainCertificateRequest` model [#45](https://github.com/patrickcping/pingone-go-sdk-v2/pull/45)
* **Bug fix** - Correct `Gateway` model attributes [#46](https://github.com/patrickcping/pingone-go-sdk-v2/pull/46)
* **Bug fix** - Correct `GatewayLDAP` model attributes [#46](https://github.com/patrickcping/pingone-go-sdk-v2/pull/46)
* **Bug fix** - Correct the `EnumGatewayPasswordAuthority` values [#46](https://github.com/patrickcping/pingone-go-sdk-v2/pull/46)
* **Feature** Support for Subscriptions data model [#42](https://github.com/patrickcping/pingone-go-sdk-v2/pull/42)
* **Feature** Support for `EmailDomain` and `EmailDomainTrustedEmail` data models [#43](https://github.com/patrickcping/pingone-go-sdk-v2/pull/43)
* **Breaking change** `createdAt` and `updatedAt` attributes changed to date/time data type [#42](https://github.com/patrickcping/pingone-go-sdk-v2/pull/42)
* **Breaking change** `lastUsedAt` attribute on the Gateway Credential model changed to date/time data type [#43](https://github.com/patrickcping/pingone-go-sdk-v2/pull/43)
* **Breaking change** Remove Device Authentication Policy API (moved to MFA module) [#44](https://github.com/patrickcping/pingone-go-sdk-v2/pull/44)

# v0.8.0 (2022-09-04)

* **Bug fix** - Fix `Gateway` response dereferencing for non-`LDAP` types [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)
* **Bug fix** - Fix `ExportCSR` headers, to determine the response type of the CSR [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
* **Bug fix** - Correct the `ImportCSRResponse` API [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
* **Bug fix** - Correct the `GetKey` API when exporting the public certificate [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
* **Bug fix** - Correct the `CreateCertificateFromFile` API [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
* **Enhancement** - Add `Links` to `Gateway` model [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)
* **Enhancement** - Add `X_PEM_FILE` to `EnumCSRExportHeader` (exporting CSR formats) [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
* **Feature** Support for Read All Gateway Credentials [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)
* **Feature** Support for Read One Gateway Credential [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)

# v0.7.0 (2022-09-01)

* **Bug fix** - Correct `ApplicationSAML` model `IdpSigning` attribute [#38](https://github.com/patrickcping/pingone-go-sdk-v2/pull/38)
* **Bug fix** - Correct `ApplicationSAML` model to include read only `Algorithm` attribute [#38](https://github.com/patrickcping/pingone-go-sdk-v2/pull/38)
* **Note** Uplift API version to 2022-08-02 [#32](https://github.com/patrickcping/pingone-go-sdk-v2/pull/32)
* **Note** Documentation formatting change `Certificate` model [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Note** Documentation formatting change `EnumCertificateKeyAlgorithm` [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Note** Documentation formatting change `EnumCertificateKeyStatus` [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Note** Documentation formatting change `EnumCertificateKeySignagureAlgorithm` [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Note** Documentation formatting change `EnumCertificateKeyUsageType` [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Breaking change** Added required fields to the SAML Identity Provider constructor [#31](https://github.com/patrickcping/pingone-go-sdk-v2/pull/31)
* **Breaking change** Changed required attributes on the `Certificate` model [#35](https://github.com/patrickcping/pingone-go-sdk-v2/pull/35)
* **Enhancement** Support for Kerberos Gateway [#32](https://github.com/patrickcping/pingone-go-sdk-v2/pull/32)
* **Enhancement** `EnumCertificateKeyUsageType` includes missing enum values [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Enhancement** `EnumCertificateKeyStatus` includes missing enum values [#37](https://github.com/patrickcping/pingone-go-sdk-v2/pull/37)
* **Enhancement** `EnumCertificateKeySignagureAlgorithm` includes missing enum values [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33), [#36](https://github.com/patrickcping/pingone-go-sdk-v2/pull/36)
* **Enhancement** Better define the Certificate Key update model [#34](https://github.com/patrickcping/pingone-go-sdk-v2/pull/34)

# v0.6.0 (2022-08-29)

* **Bug fix** - Correct `PINGONE_API` from incorrect `PING_ONE_API` in `EnumResourceType` enum [#24](https://github.com/patrickcping/pingone-go-sdk-v2/pull/24)
* **Bug fix** - Correct `EnumIdentityProviderAttributeMappingUpdate` values [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Bug fix** - Fix `IdentityProvider` oneOf decode routine [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Feature** Support for Custom Domains [#26](https://github.com/patrickcping/pingone-go-sdk-v2/pull/26)
* **Feature** Support for Keypairs [#26](https://github.com/patrickcping/pingone-go-sdk-v2/pull/26)
* **Feature** Support for Certificates [#26](https://github.com/patrickcping/pingone-go-sdk-v2/pull/26)
* **Enhancement** Add 404 response processing for all requests [#28](https://github.com/patrickcping/pingone-go-sdk-v2/pull/28)
* **Enhancement** Add missing 400 response processing [#29](https://github.com/patrickcping/pingone-go-sdk-v2/pull/29)
* **Breaking change** API name change for Agreement Languages [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Agreement Revisions [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Agreements [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Application Attribute Mapping [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Application Role Assignment [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Application Secret [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Application Sign on Policy Assignment [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Applications [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Branding Settings [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Branding Themes [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Enable Users [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Gateway Credentials [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Gateway Instances [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Gateway role assignments [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Gateways [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Group Memberships [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Identity Provider Attributes [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Identity Providers [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Language localization [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Languages [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Linked Accounts [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for MFA Pairing Keys [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Notification Settings SMTP [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Notification Settings [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Notification Templates [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Phone Delivery Settings [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Mappings [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Plans [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Revisions [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Rules [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Store Metadata [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Stores [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Resource Attributes [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Resource Scopes [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Resources [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Sessions [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Sign On Policies [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Sign On Policy Actions [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Trusted Email Addresses [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Trusted Email Domains [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User Accounts [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User Agreement Consents [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User ID Verification [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User Passwords [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User Populations [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User Role Assignments [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Users [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)

# v0.5.0 (2022-08-17)

* **Bug fix** - `SignOnPolicyActionProgressiveProfiling` object includes fixed `attributes` attribute, from object type to array type
* **Bug fix** - Restructure `SignOnPolicyAction` model to properly separate policy action attributes from each other
* **Bug fix** - Restructure `SignOnPolicyActionMFA` model to correct `noDevicesMode` attribute (fixed incorrect `noDeviceMode`)
* **Bug fix** - Corrections to `SignOnPolicyActionCommonConditions` model
* **Bug fix** - Moved `confirmIdentityProviderAttributes` boolean to the registration sub model of `SignOnPolicyAction` 

# v0.4.0 (2022-08-10)

* **Enhancement** Added generic `_links` to `Application` model
* **Breaking change** `id` made required in `spVerification.certificates` of the `ApplicationSAML` model
* **Enhancement** Add `mobile.passcodeRefreshDuration` to `ApplicationOIDC` model
* **Breaking change** `accessControl.role.type` made an enum in the `Application` model
* **Enhancement** Added `PING_ONE_DAVINCI` product type

# v0.3.0 (2022-08-05)

* **Enhancement** - *API update 2022-07-18* - `SignOnPolicyAction` object includes new attribute `DeviceAuthenticationPolicy` to replace deprecated `Applications`, `Voice`, `Sms`, `SecurityKey`, `Email`, `BoundBiometrics` and `Authenticator` attributes [#11](https://github.com/patrickcping/pingone-go-sdk-v2/pull/11)
* **Feature** Agreements, Agreement Languages and Agreement Revisions [#14](https://github.com/patrickcping/pingone-go-sdk-v2/pull/14)
* **Feature** Support for Identity Provider models [#15](https://github.com/patrickcping/pingone-go-sdk-v2/pull/15)
* **Bug** fix for `PasswordPolicyMinCharacters` special character issue [#17](https://github.com/patrickcping/pingone-go-sdk-v2/pull/17)

# v0.2.0 (2022-07-17)

* **Security** Bump `golang.org/x/net` for `CVE-2021-33194` : https://deps.dev/advisory/osv/GO-2021-0238
* **Feature** Full support for enum values

# v0.1.0 (2022-07-16)

Initial release - rebasing versions to reflect module stability