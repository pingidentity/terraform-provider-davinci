# Go API client for mfa

The PingOne Platform API covering the PingOne MFA service

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2021-10-17
- Package version: 0.7.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import mfa "github.com/patrickcping/pingone-go-sdk-v2/mfa"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), mfa.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), mfa.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), mfa.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), mfa.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.pingone.eu*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationsApplicationMFAPushCredentialsApi* | [**CreateMFAPushCredential**](docs/ApplicationsApplicationMFAPushCredentialsApi.md#createmfapushcredential) | **Post** /v1/environments/{environmentID}/applications/{applicationID}/pushCredentials | CREATE MFA Push Credential
*ApplicationsApplicationMFAPushCredentialsApi* | [**DeleteMFAPushCredential**](docs/ApplicationsApplicationMFAPushCredentialsApi.md#deletemfapushcredential) | **Delete** /v1/environments/{environmentID}/applications/{applicationID}/pushCredentials/{pushCredentialID} | DELETE MFA Push Credential
*ApplicationsApplicationMFAPushCredentialsApi* | [**ReadAllMFAPushCredentials**](docs/ApplicationsApplicationMFAPushCredentialsApi.md#readallmfapushcredentials) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/pushCredentials | READ All MFA Push Credentials
*ApplicationsApplicationMFAPushCredentialsApi* | [**ReadOneMFAPushCredential**](docs/ApplicationsApplicationMFAPushCredentialsApi.md#readonemfapushcredential) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/pushCredentials/{pushCredentialID} | READ One MFA Push Credential
*ApplicationsApplicationMFAPushCredentialsApi* | [**UpdateMFAPushCredential**](docs/ApplicationsApplicationMFAPushCredentialsApi.md#updatemfapushcredential) | **Put** /v1/environments/{environmentID}/applications/{applicationID}/pushCredentials/{pushCredentialID} | UPDATE MFA Push Credential
*DeviceAuthenticationPolicyApi* | [**CreateDeviceAuthenticationPolicies**](docs/DeviceAuthenticationPolicyApi.md#createdeviceauthenticationpolicies) | **Post** /v1/environments/{environmentID}/deviceAuthenticationPolicies | CREATE Device Authentication Policy
*DeviceAuthenticationPolicyApi* | [**DeleteDeviceAuthenticationPolicy**](docs/DeviceAuthenticationPolicyApi.md#deletedeviceauthenticationpolicy) | **Delete** /v1/environments/{environmentID}/deviceAuthenticationPolicies/{deviceAuthenticationPolicyID} | DELETE Device Authentication Policy
*DeviceAuthenticationPolicyApi* | [**ReadDeviceAuthenticationPolicies**](docs/DeviceAuthenticationPolicyApi.md#readdeviceauthenticationpolicies) | **Get** /v1/environments/{environmentID}/deviceAuthenticationPolicies | READ Device Authentication Policies
*DeviceAuthenticationPolicyApi* | [**ReadOneDeviceAuthenticationPolicy**](docs/DeviceAuthenticationPolicyApi.md#readonedeviceauthenticationpolicy) | **Get** /v1/environments/{environmentID}/deviceAuthenticationPolicies/{deviceAuthenticationPolicyID} | READ One Device Authentication Policy
*DeviceAuthenticationPolicyApi* | [**UpdateDeviceAuthenticationPolicy**](docs/DeviceAuthenticationPolicyApi.md#updatedeviceauthenticationpolicy) | **Put** /v1/environments/{environmentID}/deviceAuthenticationPolicies/{deviceAuthenticationPolicyID} | UPDATE Device Authentication Policy
*FIDODeviceApi* | [**CreateFidoDevice**](docs/FIDODeviceApi.md#createfidodevice) | **Post** /v1/environments/{environmentID}/fidoDevicesMetadata | CREATE FIDO Device
*FIDODeviceApi* | [**DeleteFidoDevice**](docs/FIDODeviceApi.md#deletefidodevice) | **Delete** /v1/environments/{environmentID}/fidoDevicesMetadata/{fidoDeviceID} | DELETE FIDO Device
*FIDODeviceApi* | [**ReadFidoDevices**](docs/FIDODeviceApi.md#readfidodevices) | **Get** /v1/environments/{environmentID}/fidoDevicesMetadata | READ All FIDO Devices
*FIDODeviceApi* | [**ReadOneFidoDevice**](docs/FIDODeviceApi.md#readonefidodevice) | **Get** /v1/environments/{environmentID}/fidoDevicesMetadata/{fidoDeviceID} | READ One FIDO Device
*FIDOPolicyApi* | [**CreateFidoPolicy**](docs/FIDOPolicyApi.md#createfidopolicy) | **Post** /v1/environments/{environmentID}/fidoPolicies | CREATE FIDO Policy
*FIDOPolicyApi* | [**DeleteFidoPolicy**](docs/FIDOPolicyApi.md#deletefidopolicy) | **Delete** /v1/environments/{environmentID}/fidoPolicies/{fidoPolicyID} | DELETE FIDO Policy
*FIDOPolicyApi* | [**ReadFidoPolicies**](docs/FIDOPolicyApi.md#readfidopolicies) | **Get** /v1/environments/{environmentID}/fidoPolicies | READ FIDO Policies
*FIDOPolicyApi* | [**ReadOneFidoPolicy**](docs/FIDOPolicyApi.md#readonefidopolicy) | **Get** /v1/environments/{environmentID}/fidoPolicies/{fidoPolicyID} | READ One FIDO Policy
*FIDOPolicyApi* | [**UpdateFIDOPolicy**](docs/FIDOPolicyApi.md#updatefidopolicy) | **Put** /v1/environments/{environmentID}/fidoPolicies/{fidoPolicyID} | UPDATE FIDO Policy
*MFASettingsApi* | [**ReadMFASettings**](docs/MFASettingsApi.md#readmfasettings) | **Get** /v1/environments/{environmentID}/mfaSettings | READ MFA Settings
*MFASettingsApi* | [**ResetMFASettings**](docs/MFASettingsApi.md#resetmfasettings) | **Delete** /v1/environments/{environmentID}/mfaSettings | RESET MFA Settings
*MFASettingsApi* | [**UpdateMFASettings**](docs/MFASettingsApi.md#updatemfasettings) | **Put** /v1/environments/{environmentID}/mfaSettings | UPDATE MFA Settings
*UsersEnableUsersMFAApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledGet**](docs/UsersEnableUsersMFAApi.md#v1environmentsenvironmentidusersuseridmfaenabledget) | **Get** /v1/environments/{environmentID}/users/{userID}/mfaEnabled | READ User MFA Enabled
*UsersEnableUsersMFAApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledPut**](docs/UsersEnableUsersMFAApi.md#v1environmentsenvironmentidusersuseridmfaenabledput) | **Put** /v1/environments/{environmentID}/users/{userID}/mfaEnabled | UPDATE User MFA Enabled
*UsersMFADevicesApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDelete**](docs/UsersMFADevicesApi.md#v1environmentsenvironmentidusersuseriddevicesdelete) | **Delete** /v1/environments/{environmentID}/users/{userID}/devices | DELETE Device Order
*UsersMFADevicesApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete**](docs/UsersMFADevicesApi.md#v1environmentsenvironmentidusersuseriddevicesdeviceiddelete) | **Delete** /v1/environments/{environmentID}/users/{userID}/devices/{deviceID} | DELETE MFA User Device
*UsersMFADevicesApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet**](docs/UsersMFADevicesApi.md#v1environmentsenvironmentidusersuseriddevicesdeviceidget) | **Get** /v1/environments/{environmentID}/users/{userID}/devices/{deviceID} | READ One MFA User Device
*UsersMFADevicesApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut**](docs/UsersMFADevicesApi.md#v1environmentsenvironmentidusersuseriddevicesdeviceidlogsput) | **Put** /v1/environments/{environmentID}/users/{userID}/devices/{deviceID}/logs | SEND MFA Device Logs
*UsersMFADevicesApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut**](docs/UsersMFADevicesApi.md#v1environmentsenvironmentidusersuseriddevicesdeviceidnicknameput) | **Put** /v1/environments/{environmentID}/users/{userID}/devices/{deviceID}/nickname | UPDATE Device Nickname
*UsersMFADevicesApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost**](docs/UsersMFADevicesApi.md#v1environmentsenvironmentidusersuseriddevicesdeviceidpost) | **Post** /v1/environments/{environmentID}/users/{userID}/devices/{deviceID} | ACTIVATE MFA User Device
*UsersMFADevicesApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDDevicesGet**](docs/UsersMFADevicesApi.md#v1environmentsenvironmentidusersuseriddevicesget) | **Get** /v1/environments/{environmentID}/users/{userID}/devices | READ All MFA User Devices
*UsersMFADevicesApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDDevicesPost**](docs/UsersMFADevicesApi.md#v1environmentsenvironmentidusersuseriddevicespost) | **Post** /v1/environments/{environmentID}/users/{userID}/devices | SET Device Order
*UsersMFAPairingKeysApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDDelete**](docs/UsersMFAPairingKeysApi.md#v1environmentsenvironmentidusersuseridpairingkeyspairingkeyiddelete) | **Delete** /v1/environments/{environmentID}/users/{userID}/pairingKeys/{pairingKeyID} | DELETE MFA Pairing Key
*UsersMFAPairingKeysApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDGet**](docs/UsersMFAPairingKeysApi.md#v1environmentsenvironmentidusersuseridpairingkeyspairingkeyidget) | **Get** /v1/environments/{environmentID}/users/{userID}/pairingKeys/{pairingKeyID} | READ One MFA Pairing Key
*UsersMFAPairingKeysApi* | [**V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPost**](docs/UsersMFAPairingKeysApi.md#v1environmentsenvironmentidusersuseridpairingkeyspost) | **Post** /v1/environments/{environmentID}/users/{userID}/pairingKeys | CREATE MFA Pairing Key


## Documentation For Models

 - [CreateMFAPushCredentialRequest](docs/CreateMFAPushCredentialRequest.md)
 - [DeviceAuthenticationPolicy](docs/DeviceAuthenticationPolicy.md)
 - [DeviceAuthenticationPolicyFIDODevice](docs/DeviceAuthenticationPolicyFIDODevice.md)
 - [DeviceAuthenticationPolicyMobile](docs/DeviceAuthenticationPolicyMobile.md)
 - [DeviceAuthenticationPolicyMobileApplicationsInner](docs/DeviceAuthenticationPolicyMobileApplicationsInner.md)
 - [DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment](docs/DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment.md)
 - [DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization](docs/DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization.md)
 - [DeviceAuthenticationPolicyMobileApplicationsInnerOtp](docs/DeviceAuthenticationPolicyMobileApplicationsInnerOtp.md)
 - [DeviceAuthenticationPolicyMobileApplicationsInnerPush](docs/DeviceAuthenticationPolicyMobileApplicationsInnerPush.md)
 - [DeviceAuthenticationPolicyMobileOtp](docs/DeviceAuthenticationPolicyMobileOtp.md)
 - [DeviceAuthenticationPolicyMobileOtpWindow](docs/DeviceAuthenticationPolicyMobileOtpWindow.md)
 - [DeviceAuthenticationPolicyMobileOtpWindowStepSize](docs/DeviceAuthenticationPolicyMobileOtpWindowStepSize.md)
 - [DeviceAuthenticationPolicyOfflineDevice](docs/DeviceAuthenticationPolicyOfflineDevice.md)
 - [DeviceAuthenticationPolicyOfflineDeviceOtp](docs/DeviceAuthenticationPolicyOfflineDeviceOtp.md)
 - [DeviceAuthenticationPolicyOfflineDeviceOtpFailure](docs/DeviceAuthenticationPolicyOfflineDeviceOtpFailure.md)
 - [DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown](docs/DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown.md)
 - [DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime](docs/DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime.md)
 - [DeviceAuthenticationPolicyTotp](docs/DeviceAuthenticationPolicyTotp.md)
 - [DeviceAuthenticationPolicyTotpOtp](docs/DeviceAuthenticationPolicyTotpOtp.md)
 - [EntityArray](docs/EntityArray.md)
 - [EntityArrayEmbedded](docs/EntityArrayEmbedded.md)
 - [EntityArrayEmbeddedPushCredentialsInner](docs/EntityArrayEmbeddedPushCredentialsInner.md)
 - [EnumFIDOAttestationRequirements](docs/EnumFIDOAttestationRequirements.md)
 - [EnumFIDOResidentKeyRequirement](docs/EnumFIDOResidentKeyRequirement.md)
 - [EnumMFADevicePolicyMobileExtraVerification](docs/EnumMFADevicePolicyMobileExtraVerification.md)
 - [EnumMFADevicePolicyMobileIntegrityDetection](docs/EnumMFADevicePolicyMobileIntegrityDetection.md)
 - [EnumMFAPushCredentialAttrType](docs/EnumMFAPushCredentialAttrType.md)
 - [EnumMFASettingsDeviceSelection](docs/EnumMFASettingsDeviceSelection.md)
 - [EnumMFASettingsPairingKeyFormat](docs/EnumMFASettingsPairingKeyFormat.md)
 - [EnumTimeUnit](docs/EnumTimeUnit.md)
 - [FIDOPolicy](docs/FIDOPolicy.md)
 - [FIDOPolicyAllowedAuthenticatorsInner](docs/FIDOPolicyAllowedAuthenticatorsInner.md)
 - [MFAPushCredential](docs/MFAPushCredential.md)
 - [MFAPushCredentialAPNS](docs/MFAPushCredentialAPNS.md)
 - [MFAPushCredentialAPNSAllOf](docs/MFAPushCredentialAPNSAllOf.md)
 - [MFAPushCredentialResponse](docs/MFAPushCredentialResponse.md)
 - [MFASettings](docs/MFASettings.md)
 - [MFASettingsAuthentication](docs/MFASettingsAuthentication.md)
 - [MFASettingsLockout](docs/MFASettingsLockout.md)
 - [MFASettingsPairing](docs/MFASettingsPairing.md)
 - [ObjectEnvironment](docs/ObjectEnvironment.md)
 - [P1Error](docs/P1Error.md)
 - [P1ErrorDetailsInner](docs/P1ErrorDetailsInner.md)
 - [P1ErrorDetailsInnerInnerError](docs/P1ErrorDetailsInnerInnerError.md)
 - [UpdateMFAPushCredentialRequest](docs/UpdateMFAPushCredentialRequest.md)


## Documentation For Authorization



### bearer

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


