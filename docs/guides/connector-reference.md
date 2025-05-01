---
layout: ""
page_title: "Connector Parameter Reference"
description: |-
  The guide describes the connection parameters for all connectors in the DaVinci platform, with examples of how to define within Terraform using the `davinci_connection` resource.
---

# DaVinci Connection Definitions

Below is a list of all available DaVinci Connections available for use in `davinci_connection` resource. 
If the `value` type of a property is not defined it must be inferred.


## 1Kosmos connector

Connector ID (`connector_id` in the resource): `connector1Kosmos`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connector1Kosmos" {
  environment_id = var.pingone_environment_id

  connector_id = "connector1Kosmos"
  name         = "My awesome connector1Kosmos"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## AWS Lambda

Connector ID (`connector_id` in the resource): `connectorAWSLambda`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessKeyId` (string): Access key to connect to AWS Environment. Console display name: "Access Key Id".
* `region` (string): AWS Region where the Lambda function is created. Console display name: "AWS Region".
* `secretAccessKey` (string): Secret Key to access the AWS. Console display name: "AWS Secret Key".


Example:
```terraform
resource "davinci_connection" "connectorAWSLambda" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAWSLambda"
  name         = "My awesome connectorAWSLambda"

  property {
    name  = "accessKeyId"
    type  = "string"
    value = var.connectorawslambda_property_access_key_id
  }

  property {
    name  = "region"
    type  = "string"
    value = "eu-west-1"
  }

  property {
    name  = "secretAccessKey"
    type  = "string"
    value = var.connectorawslambda_property_secret_access_key
  }
}
```


## AWS Login

Connector ID (`connector_id` in the resource): `awsIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (json):  Console display name: "OpenId Parameters".


Example:
```terraform
resource "davinci_connection" "awsIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "awsIdpConnector"
  name         = "My awesome awsIdpConnector"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
```


## AWS Secrets Manager

Connector ID (`connector_id` in the resource): `connectorAmazonAwsSecretsManager`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessKeyId` (string): The AWS Access Key. Console display name: "AWS Access Key".
* `region` (string): The AWS Region. Console display name: "AWS Region".
* `secretAccessKey` (string): The AWS Access Secret. Console display name: "AWS Access Secret".


Example:
```terraform
resource "davinci_connection" "connectorAmazonAwsSecretsManager" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAmazonAwsSecretsManager"
  name         = "My awesome connectorAmazonAwsSecretsManager"

  property {
    name  = "accessKeyId"
    type  = "string"
    value = var.connectoramazonawssecretsmanager_property_access_key_id
  }

  property {
    name  = "region"
    type  = "string"
    value = "eu-west-1"
  }

  property {
    name  = "secretAccessKey"
    type  = "string"
    value = var.connectoramazonawssecretsmanager_property_secret_access_key
  }
}
```


## AbuseIPDB

Connector ID (`connector_id` in the resource): `connectorAbuseipdb`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key gathered from AbuseIPDB tenant. Console display name: "API Key".


Example:
```terraform
resource "davinci_connection" "connectorAbuseipdb" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAbuseipdb"
  name         = "My awesome connectorAbuseipdb"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorabuseipdb_property_api_key
  }
}
```


## ActiveCampaign API

Connector ID (`connector_id` in the resource): `connector-oai-activecampaignapi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authApiKey` (string): The authentication key to the ActiveCampaign API. Console display name: "API Key".
* `authApiVersion` (string): The version of the ActiveCampaign API. Console display name: "API Version".
* `basePath` (string): The base URL for contacting the API. Console display name: "Base Path".


Example:
```terraform
resource "davinci_connection" "connector-oai-activecampaignapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-activecampaignapi"
  name         = "My awesome connector-oai-activecampaignapi"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-activecampaignapi_property_auth_api_key
  }

  property {
    name  = "authApiVersion"
    type  = "string"
    value = var.connector-oai-activecampaignapi_property_auth_api_version
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-activecampaignapi_property_base_path
  }
}
```


## Acuant

Connector ID (`connector_id` in the resource): `connectorAcuant`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorAcuant" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAcuant"
  name         = "My awesome connectorAcuant"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Adobe Marketo

Connector ID (`connector_id` in the resource): `adobemarketoConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): Your Adobe Marketo client ID. Console display name: "Client ID".
* `clientSecret` (string): Your Adobe Marketo client secret. Console display name: "Client Secret".
* `endpoint` (string): The API endpoint for your Adobe Marketo instance, such as "abc123.mktorest.com/rest". Console display name: "API URL".


Example:
```terraform
resource "davinci_connection" "adobemarketoConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "adobemarketoConnector"
  name         = "My awesome adobemarketoConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.adobemarketoconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.adobemarketoconnector_property_client_secret
  }

  property {
    name  = "endpoint"
    type  = "string"
    value = var.adobemarketoconnector_property_endpoint
  }
}
```


## Akamai MFA

Connector ID (`connector_id` in the resource): `akamaiConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "akamaiConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "akamaiConnector"
  name         = "My awesome akamaiConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.akamaiconnector_property_custom_auth
  }
}
```


## Allthenticate

Connector ID (`connector_id` in the resource): `connectorAllthenticate`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorAllthenticate" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAllthenticate"
  name         = "My awesome connectorAllthenticate"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Amazon DynamoDB

Connector ID (`connector_id` in the resource): `connectorAmazonDynamoDB`

Properties (used in the `property` block in the resource as the `name` parameter):

* `awsAccessKey` (string): Your AWS Access Key. Console display name: "AWS Access Key".
* `awsAccessSecret` (string): Access Secret corresponding with Access Key found in Your Security Credentials. Console display name: "AWS Access Secret".
* `awsRegion` (string): The AWS Region you are using the connector for. Console display name: "AWS Region".


Example:
```terraform
resource "davinci_connection" "connectorAmazonDynamoDB" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAmazonDynamoDB"
  name         = "My awesome connectorAmazonDynamoDB"

  property {
    name  = "awsAccessKey"
    type  = "string"
    value = var.connectoramazondynamodb_property_aws_access_key
  }

  property {
    name  = "awsAccessSecret"
    type  = "string"
    value = var.connectoramazondynamodb_property_aws_access_secret
  }

  property {
    name  = "awsRegion"
    type  = "string"
    value = "eu-west-1"
  }
}
```


## Amazon Simple Email Service

Connector ID (`connector_id` in the resource): `amazonSimpleEmailConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `awsAccessKey` (string):  Console display name: "AWS Access Key".
* `awsAccessSecret` (string):  Console display name: "AWS Access Secret".
* `awsRegion` (string):  Console display name: "AWS Region".
* `from` (string): The email address that the message appears to originate from, as registered with your AWS account, such as "support@mycompany.com". Console display name: "From (Default) *".


Example:
```terraform
resource "davinci_connection" "amazonSimpleEmailConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "amazonSimpleEmailConnector"
  name         = "My awesome amazonSimpleEmailConnector"

  property {
    name  = "awsAccessKey"
    type  = "string"
    value = var.amazonsimpleemailconnector_property_aws_access_key
  }

  property {
    name  = "awsAccessSecret"
    type  = "string"
    value = var.amazonsimpleemailconnector_property_aws_access_secret
  }

  property {
    name  = "awsRegion"
    type  = "string"
    value = "eu-west-1"
  }

  property {
    name  = "from"
    type  = "string"
    value = "support@bxretail.org"
  }
}
```


## Annotation

Connector ID (`connector_id` in the resource): `annotationConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "annotationConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "annotationConnector"
  name         = "My awesome annotationConnector"
}
```


## Apple Login

Connector ID (`connector_id` in the resource): `appleConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "appleConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "appleConnector"
  name         = "My awesome appleConnector"

  property {
    name = "customAuth"
    type = "json"
    value = jsonencode({
      "properties" : {
        "providerName" : {
          "displayName" : "Provider Name",
          "preferredControlType" : "textField",
          "value" : "${var.appleconnector_property_provider_name}"
        },
        "skRedirectUri" : {
          "displayName" : "DaVinci Redirect URL",
          "info" : "Your DaVinci redirect URL. This allows an identity provider to redirect the browser back to DaVinci.",
          "preferredControlType" : "textField",
          "disabled" : true,
          "initializeValue" : "SINGULARKEY_REDIRECT_URI",
          "copyToClip" : true
        },
        "iss" : {
          "displayName" : "Issuer",
          "info" : "The issuer registered claim identifies the principal that issued the client secret. Since the client secret was generated for your developer team, use your 10-character Team ID associated with your developer account.",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "${var.appleconnector_property_issuer}"
        },
        "kid" : {
          "displayName" : "Key ID",
          "info" : "A 10-character key identifier generated for the Sign in with Apple private key associated with your developer account.",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "${var.appleconnector_property_key_id}"
        },
        "issuerUrl" : {
          "displayName" : "Issuer URL",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "${var.appleconnector_property_issuer_url}"
        },
        "authorizationEndpoint" : {
          "preferredControlType" : "textField",
          "displayName" : "Authorization Endpoint",
          "required" : true,
          "value" : "${var.appleconnector_property_authorization_endpoint}"
        },
        "tokenEndpoint" : {
          "preferredControlType" : "textField",
          "displayName" : "Token Endpoint",
          "required" : true,
          "value" : "${var.appleconnector_property_token_endpoint}"
        },
        "clientId" : {
          "displayName" : "Client ID",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "${var.appleconnector_property_client_id}"
        },
        "clientSecret" : {
          "displayName" : "Private Key",
          "info" : "Content of your 'Sign in with Apple' private key associated with your developer account.",
          "preferredControlType" : "textArea",
          "secure" : true,
          "required" : true,
          "value" : "${var.appleconnector_property_private_key}"
        },
        "scope" : {
          "displayName" : "Scope",
          "preferredControlType" : "textField",
          "requiredValue" : "email",
          "required" : true,
          "value" : "${var.appleconnector_property_scope}"
        },
        "userConnectorAttributeMapping" : {
          "type" : "object",
          "preferredControlType" : "userConnectorAttributeMapping",
          "newMappingAllowed" : true,
          "title1" : null,
          "title2" : null,
          "sections" : [
            "attributeMapping"
          ],
          "value" : {
            "userPoolConnectionId" : "defaultUserPool",
            "mapping" : {
              "username" : {
                "value1" : "sub"
              },
              "name" : {
                "value1" : "email"
              },
              "email" : {
                "value1" : "email"
              }
            }
          }
        },
        "customAttributes" : {
          "type" : "array",
          "displayName" : "Connector Attributes",
          "preferredControlType" : "tableViewAttributes",
          "info" : "These attributes will be available in User Connector Attribute Mapping.",
          "sections" : [
            "connectorAttributes"
          ],
          "value" : [
            {
              "name" : "sub",
              "description" : "Sub",
              "type" : "string",
              "value" : null,
              "minLength" : "1",
              "maxLength" : "300",
              "required" : true,
              "attributeType" : "sk"
            },
            {
              "name" : "email",
              "description" : "Email",
              "type" : "string",
              "value" : null,
              "minLength" : "1",
              "maxLength" : "250",
              "required" : false,
              "attributeType" : "sk"
            }
          ]
        },
        "disableCreateUser" : {
          "displayName" : "Disable Shadow User Creation",
          "preferredControlType" : "toggleSwitch",
          "value" : false,
          "info" : "A shadow user is implicitly created, unless disabled."
        },
        "returnToUrl" : {
          "displayName" : "Application Return To URL",
          "preferredControlType" : "textField",
          "info" : "When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application."
        }
      }
    })
  }
}
```


## Argyle

Connector ID (`connector_id` in the resource): `argyleConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string):  Console display name: "API Server URL".
* `clientId` (string):  Console display name: "Client ID".
* `clientSecret` (string):  Console display name: "Client Secret".
* `javascriptWebUrl` (string): Argyle loader javascript web URL. Console display name: "Argyle Loader Javascript Web URL".
* `pluginKey` (string):  Console display name: "Plugin Key".


Example:
```terraform
resource "davinci_connection" "argyleConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "argyleConnector"
  name         = "My awesome argyleConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.argyleconnector_property_api_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.argyleconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.argyleconnector_property_client_secret
  }

  property {
    name  = "javascriptWebUrl"
    type  = "string"
    value = var.argyleconnector_property_javascript_web_url
  }

  property {
    name  = "pluginKey"
    type  = "string"
    value = var.argyleconnector_property_plugin_key
  }
}
```


## Asignio

Connector ID (`connector_id` in the resource): `connectorAsignio`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorAsignio" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAsignio"
  name         = "My awesome connectorAsignio"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## AuthID

Connector ID (`connector_id` in the resource): `connectorAuthid`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorAuthid" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAuthid"
  name         = "My awesome connectorAuthid"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## AuthenticID

Connector ID (`connector_id` in the resource): `authenticIdConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accountAccessKey` (string): Your Account Access Key provided by AuthenticID . Console display name: "Account Access Key".
* `androidSDKLicenseKey` (string): License key is whitelisted for specific package name. Console display name: "Android SDK Licence Key".
* `apiUrl` (string): AuthenticID REST API URL for sandbox/production environments. Console display name: "REST API URL".
* `baseUrl` (string): AuthenticID API URL for sandbox/production environments. Console display name: "Base URL".
* `clientCertificate` (string): Your Client Certificate provided by AuthenticID. Console display name: "Client Certificate".
* `clientKey` (string): Your Client Key provided by AuthenticID. Console display name: "Client Key".
* `iOSSDKLicenseKey` (string): License key is whitelisted for specific bundle id. Console display name: "iOS SDK Licence Key".
* `passphrase` (string): Your Certificate Passphrase provided by AuthenticID. Console display name: "Certificate Passphrase".
* `secretToken` (string): Your Secret Token provided by AuthenticID. Console display name: "Secret Token".


Example:
```terraform
resource "davinci_connection" "authenticIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "authenticIdConnector"
  name         = "My awesome authenticIdConnector"

  property {
    name  = "accountAccessKey"
    type  = "string"
    value = var.authenticidconnector_property_account_access_key
  }

  property {
    name  = "androidSDKLicenseKey"
    type  = "string"
    value = var.authenticidconnector_property_android_sdk_license_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.authenticidconnector_property_api_url
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.authenticidconnector_property_base_url
  }

  property {
    name  = "clientCertificate"
    type  = "string"
    value = var.authenticidconnector_property_client_certificate
  }

  property {
    name  = "clientKey"
    type  = "string"
    value = var.authenticidconnector_property_client_key
  }

  property {
    name  = "iOSSDKLicenseKey"
    type  = "string"
    value = var.authenticidconnector_property_ios_sdk_license_key
  }

  property {
    name  = "passphrase"
    type  = "string"
    value = var.authenticidconnector_property_passphrase
  }

  property {
    name  = "secretToken"
    type  = "string"
    value = var.authenticidconnector_property_secret_token
  }
}
```


## Authomize API

Connector ID (`connector_id` in the resource): `connector-oai-authomizeapireference`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authApiKey` (string): Your Authomize API key. Console display name: "API Key".
* `basePath` (string): The base URL for the Authomize API, such as "https://api.authomize.com". Console display name: "Base URL".


Example:
```terraform
resource "davinci_connection" "connector-oai-authomizeapireference" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-authomizeapireference"
  name         = "My awesome connector-oai-authomizeapireference"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-authomizeapireference_property_auth_api_key
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-authomizeapireference_property_base_path
  }
}
```


## Authomize Incident Connector

Connector ID (`connector_id` in the resource): `connectorAuthomize`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): The API Key from the Authomize API Tokens creation page. Console display name: "Authomize API Key".


Example:
```terraform
resource "davinci_connection" "connectorAuthomize" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAuthomize"
  name         = "My awesome connectorAuthomize"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorauthomize_property_api_key
  }
}
```


## Azure AD User Management

Connector ID (`connector_id` in the resource): `azureUserManagementConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseUrl` (string): The Microsoft API URL to target. For a custom value, select Use Custom API URL and enter a value in the Custom API URL field. Console display name: "API URL".
* `customApiUrl` (string): The URL for the Microsoft Graph API, such as "https://graph.microsoft.com/v1.0". Console display name: "Custom API URL".
* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "azureUserManagementConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "azureUserManagementConnector"
  name         = "My awesome azureUserManagementConnector"

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.azureusermanagementconnector_property_base_url
  }

  property {
    name  = "customApiUrl"
    type  = "string"
    value = var.azureusermanagementconnector_property_custom_api_url
  }

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Badge

Connector ID (`connector_id` in the resource): `connectorBadge`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorBadge" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBadge"
  name         = "My awesome connectorBadge"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## BambooHR

Connector ID (`connector_id` in the resource): `bambooConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string):  Console display name: "API Key".
* `baseUrl` (string):  BambooHR Base URL. Console display name: "Base URL".
* `companySubDomain` (string):  Your BambooHR subdomain. Console display name: "Company Sub Domain".
* `flowId` (string): Select ID of the flow to execute when BambooHR sends a webhook. Console display name: "Flow ID".
* `skWebhookUri` (string): Use this url as the Webhook URL in the Third Party Integration's configuration. Console display name: "DaVinci Webhook URL".
* `webhookToken` (string): Create a webhook token and configure it in the bambooHR webhook url. Console display name: "Webhook Token".


Example:
```terraform
resource "davinci_connection" "bambooConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "bambooConnector"
  name         = "My awesome bambooConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.bambooconnector_property_api_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.bambooconnector_property_base_url
  }

  property {
    name  = "companySubDomain"
    type  = "string"
    value = var.bambooconnector_property_company_sub_domain
  }

  property {
    name  = "flowId"
    type  = "string"
    value = var.bambooconnector_property_flow_id
  }

  property {
    name  = "skWebhookUri"
    type  = "string"
    value = var.bambooconnector_property_sk_webhook_uri
  }

  property {
    name  = "webhookToken"
    type  = "string"
    value = var.bambooconnector_property_webhook_token
  }
}
```


## Berbix

Connector ID (`connector_id` in the resource): `connectorBerbix`

Properties (used in the `property` block in the resource as the `name` parameter):

* `domainName` (string): Provide Berbix domain name. Console display name: "Domain Name".
* `path` (string): Provide path of the API. Console display name: "Path".
* `username` (string): Provide your Berbix user name. Console display name: "User Name".


Example:
```terraform
resource "davinci_connection" "connectorBerbix" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBerbix"
  name         = "My awesome connectorBerbix"

  property {
    name  = "domainName"
    type  = "string"
    value = var.connectorberbix_property_domain_name
  }

  property {
    name  = "path"
    type  = "string"
    value = var.connectorberbix_property_path
  }

  property {
    name  = "username"
    type  = "string"
    value = var.connectorberbix_property_username
  }
}
```


## Beyond Identity

Connector ID (`connector_id` in the resource): `connectorBeyondIdentity`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (json):  Console display name: "OpenId Parameters".


Example:
```terraform
resource "davinci_connection" "connectorBeyondIdentity" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBeyondIdentity"
  name         = "My awesome connectorBeyondIdentity"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
```


## BeyondTrust - Password Safe

Connector ID (`connector_id` in the resource): `connectorBTps`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key from your Password Safe environment. Console display name: "API Key".
* `apiUser` (string): API User from your Password Safe environment. Console display name: "API User".
* `domain` (string): Domain of your Password Safe environment. Console display name: "PasswordSafe Hostname".


Example:
```terraform
resource "davinci_connection" "connectorBTps" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBTps"
  name         = "My awesome connectorBTps"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorbtps_property_api_key
  }

  property {
    name  = "apiUser"
    type  = "string"
    value = var.connectorbtps_property_api_user
  }

  property {
    name  = "domain"
    type  = "string"
    value = var.connectorbtps_property_domain
  }
}
```


## BeyondTrust - Privileged Remote Access

Connector ID (`connector_id` in the resource): `connectorBTpra`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientID` (string): PRA API Client ID. Console display name: "Client ID".
* `clientSecret` (string): PRA API Client Secret. Console display name: "Client Secret".
* `praAPIurl` (string): URL of PRA Appliance. Console display name: "PRA Web API Address".


Example:
```terraform
resource "davinci_connection" "connectorBTpra" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBTpra"
  name         = "My awesome connectorBTpra"

  property {
    name  = "clientID"
    type  = "string"
    value = var.connectorbtpra_property_client_i_d
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectorbtpra_property_client_secret
  }

  property {
    name  = "praAPIurl"
    type  = "string"
    value = var.pra_api_url
  }
}
```


## BeyondTrust - Remote Support

Connector ID (`connector_id` in the resource): `connectorBTrs`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientID` (string): RS API Client ID. Console display name: "Client ID".
* `clientSecret` (string): RS API Client Secret. Console display name: "Client Secret".
* `rsAPIurl` (string): URL of RS Appliance. Console display name: "RS Web API Address".


Example:
```terraform
resource "davinci_connection" "connectorBTrs" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBTrs"
  name         = "My awesome connectorBTrs"

  property {
    name  = "clientID"
    type  = "string"
    value = var.connectorbtrs_property_client_i_d
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectorbtrs_property_client_secret
  }

  property {
    name  = "rsAPIurl"
    type  = "string"
    value = var.rs_api_url
  }
}
```


## BioCatch

Connector ID (`connector_id` in the resource): `biocatchConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string):  Console display name: "API Server URL".
* `customerId` (string):  Console display name: "Customer ID".
* `javascriptCdnUrl` (string):  Console display name: "Javascript CDN URL".
* `sdkToken` (string):  Console display name: "SDK Token".
* `truthApiKey` (string): Fraudulent/Genuine Session Reporting API Key. Console display name: "Truth-mapping API Key".
* `truthApiUrl` (string): Fraudulent/Genuine Session Reporting. Console display name: "Truth-mapping API URL".


Example:
```terraform
resource "davinci_connection" "biocatchConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "biocatchConnector"
  name         = "My awesome biocatchConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.biocatchconnector_property_api_url
  }

  property {
    name  = "customerId"
    type  = "string"
    value = var.biocatchconnector_property_customer_id
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.biocatchconnector_property_javascript_cdn_url
  }

  property {
    name  = "sdkToken"
    type  = "string"
    value = var.biocatchconnector_property_sdk_token
  }

  property {
    name  = "truthApiKey"
    type  = "string"
    value = var.biocatchconnector_property_truth_api_key
  }

  property {
    name  = "truthApiUrl"
    type  = "string"
    value = var.biocatchconnector_property_truth_api_url
  }
}
```


## Bitbucket Login

Connector ID (`connector_id` in the resource): `bitbucketIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (json):  Console display name: "Oauth2 Parameters".


Example:
```terraform
resource "davinci_connection" "bitbucketIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "bitbucketIdpConnector"
  name         = "My awesome bitbucketIdpConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
```


## CASTLE

Connector ID (`connector_id` in the resource): `castleConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiSecret` (string): Your 32-character Castle API secret, such as “Olc…QBF”. Console display name: "API Secret".


Example:
```terraform
resource "davinci_connection" "castleConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "castleConnector"
  name         = "My awesome castleConnector"

  property {
    name  = "apiSecret"
    type  = "string"
    value = var.castleconnector_property_api_secret
  }
}
```


## CLEAR

Connector ID (`connector_id` in the resource): `connectorClear`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorClear" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorClear"
  name         = "My awesome connectorClear"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.connectorclear_property_custom_auth
  }
}
```


## Challenge

Connector ID (`connector_id` in the resource): `challengeConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "challengeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "challengeConnector"
  name         = "My awesome challengeConnector"
}
```


## Circle Access

Connector ID (`connector_id` in the resource): `connectorCircleAccess`

Properties (used in the `property` block in the resource as the `name` parameter):

* `appKey` (string): App Key. Console display name: "App Key".
* `customAuth` (json):  Console display name: "Custom Parameters".
* `loginUrl` (string): The URL of your Circle Access login. Console display name: "Login Url".
* `readKey` (string): Read Key. Console display name: "Read Key".
* `returnToUrl` (string): When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application. Console display name: "Application Return To URL".
* `writeKey` (string): Write key. Console display name: "Write Key".


Example:
```terraform
resource "davinci_connection" "connectorCircleAccess" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorCircleAccess"
  name         = "My awesome connectorCircleAccess"

  property {
    name  = "appKey"
    type  = "string"
    value = var.connectorcircleaccess_property_app_key
  }

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }

  property {
    name  = "loginUrl"
    type  = "string"
    value = var.connectorcircleaccess_property_login_url
  }

  property {
    name  = "readKey"
    type  = "string"
    value = var.connectorcircleaccess_property_read_key
  }

  property {
    name  = "returnToUrl"
    type  = "string"
    value = var.connectorcircleaccess_property_return_to_url
  }

  property {
    name  = "writeKey"
    type  = "string"
    value = var.connectorcircleaccess_property_write_key
  }
}
```


## Clearbit

Connector ID (`connector_id` in the resource): `connectorClearbit`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Clearbit API Key. Console display name: "API Key".
* `riskApiVersion` (string): Clearbit - Risk API Version. Console display name: "Risk API Version".
* `version` (string): Clearbit - Person API Version. Console display name: "Person API Version".


Example:
```terraform
resource "davinci_connection" "connectorClearbit" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorClearbit"
  name         = "My awesome connectorClearbit"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorclearbit_property_api_key
  }

  property {
    name  = "riskApiVersion"
    type  = "string"
    value = var.connectorclearbit_property_risk_api_version
  }

  property {
    name  = "version"
    type  = "string"
    value = var.connectorclearbit_property_version
  }
}
```


## Cloudflare

Connector ID (`connector_id` in the resource): `connectorCloudflare`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accountId` (string): Cloudflare Account ID. Console display name: "Account ID".
* `apiToken` (string): Cloudflare API Token. Console display name: "API Token".


Example:
```terraform
resource "davinci_connection" "connectorCloudflare" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorCloudflare"
  name         = "My awesome connectorCloudflare"

  property {
    name  = "accountId"
    type  = "string"
    value = var.connectorcloudflare_property_account_id
  }

  property {
    name  = "apiToken"
    type  = "string"
    value = var.connectorcloudflare_property_api_token
  }
}
```


## Code Snippet

Connector ID (`connector_id` in the resource): `codeSnippetConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `code` (string): Follow example for code. Caution: Custom code is for advanced users only. Before using custom code, review the security risks in the DaVinci documentation by searching for "Using custom code safely". Console display name: "Code Snippet".
* `inputSchema` (string): Follow example for JSON schema. Console display name: "Input Schema".
* `outputSchema` (string): Follow example for JSON schema. Console display name: "Output Schema".


Example:
```terraform
resource "davinci_connection" "codeSnippetConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "codeSnippetConnector"
  name         = "My awesome codeSnippetConnector"

  property {
    name  = "code"
    type  = "string"
    value = var.codesnippetconnector_property_code
  }

  property {
    name  = "inputSchema"
    type  = "string"
    value = var.codesnippetconnector_property_input_schema
  }

  property {
    name  = "outputSchema"
    type  = "string"
    value = var.codesnippetconnector_property_output_schema
  }
}
```


## Comply Advantage

Connector ID (`connector_id` in the resource): `complyAdvatangeConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key is the API key that you can retrieve from Comply Advantage Admin Portal. Console display name: "API Key".
* `baseUrl` (string): Comply Advantage API URL for sandbox/production environments. Console display name: "Base URL".


Example:
```terraform
resource "davinci_connection" "complyAdvatangeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "complyAdvatangeConnector"
  name         = "My awesome complyAdvatangeConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.complyadvatangeconnector_property_api_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.complyadvatangeconnector_property_base_url
  }
}
```


## ConnectID

Connector ID (`connector_id` in the resource): `connectIdConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "connectIdConnector"
  name         = "My awesome connectIdConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Cookie

Connector ID (`connector_id` in the resource): `cookieConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `hmacSigningKey` (string): Base64 encoded 256 bit key. Console display name: "HMAC Signing Key".


Example:
```terraform
resource "davinci_connection" "cookieConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "cookieConnector"
  name         = "My awesome cookieConnector"

  property {
    name  = "hmacSigningKey"
    type  = "string"
    value = var.cookieconnector_property_hmac_signing_key
  }
}
```


## Copper API

Connector ID (`connector_id` in the resource): `connector-oai-copperdeveloperapi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `basePath` (string): The base URL for contacting the API. Console display name: "Base Path".
* `contentType` (string): Content type. Console display name: "Content-Type".
* `xPWAccessToken` (string): API Key. Console display name: "X-PW-AccessToken".
* `xPWApplication` (string): Application. Console display name: "X-PW-Application".
* `xPWUserEmail` (string): Email address of token owner. Console display name: "X-PW-UserEmail".


Example:
```terraform
resource "davinci_connection" "connector-oai-copperdeveloperapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-copperdeveloperapi"
  name         = "My awesome connector-oai-copperdeveloperapi"

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-copperdeveloperapi_property_base_path
  }

  property {
    name  = "contentType"
    type  = "string"
    value = var.connector-oai-copperdeveloperapi_property_content_type
  }

  property {
    name  = "xPWAccessToken"
    type  = "string"
    value = var.connector-oai-copperdeveloperapi_property_x_p_w_access_token
  }

  property {
    name  = "xPWApplication"
    type  = "string"
    value = var.connector-oai-copperdeveloperapi_property_x_p_w_application
  }

  property {
    name  = "xPWUserEmail"
    type  = "string"
    value = var.connector-oai-copperdeveloperapi_property_x_p_w_user_email
  }
}
```


## Credova

Connector ID (`connector_id` in the resource): `credovaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseUrl` (string): Base URL for Credova API. Console display name: "Base URL".
* `password` (string): Password for the Credova Developer Portal. Console display name: "Credova Password".
* `username` (string): Username for the Credova Developer Portal. Console display name: "Credova Username".


Example:
```terraform
resource "davinci_connection" "credovaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "credovaConnector"
  name         = "My awesome credovaConnector"

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.credovaconnector_property_base_url
  }

  property {
    name  = "password"
    type  = "string"
    value = var.credovaconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.credovaconnector_property_username
  }
}
```


## CrowdStrike

Connector ID (`connector_id` in the resource): `crowdStrikeConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseURL` (string): The base URL of the CrowdStrike environment. Console display name: "CrowdStrike Base URL".
* `clientId` (string): The Client ID of the application in CrowdStrike. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret provided by CrowdStrike. Console display name: "Client Secret".


Example:
```terraform
resource "davinci_connection" "crowdStrikeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "crowdStrikeConnector"
  name         = "My awesome crowdStrikeConnector"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.crowdstrikeconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.crowdstrikeconnector_property_client_secret
  }
}
```


## Daon IDV

Connector ID (`connector_id` in the resource): `connectorDaonidv`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (json):  Console display name: "OpenId Parameters".


Example:
```terraform
resource "davinci_connection" "connectorDaonidv" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorDaonidv"
  name         = "My awesome connectorDaonidv"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Daon IdentityX

Connector ID (`connector_id` in the resource): `daonConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): The protocol, host and base path to the IdX API. E.g. https://api.identityx-cloud.com/tenant1/IdentityXServices/rest/v1. Console display name: "API Base URL".
* `password` (string): The password of the user to authenticate API calls. Console display name: "Admin Password".
* `username` (string): The userId to authenticate API calls. Console display name: "Admin Username".


Example:
```terraform
resource "davinci_connection" "daonConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "daonConnector"
  name         = "My awesome daonConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.daonconnector_property_api_url
  }

  property {
    name  = "password"
    type  = "string"
    value = var.daonconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.daonconnector_property_username
  }
}
```


## Data Zoo

Connector ID (`connector_id` in the resource): `dataZooConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `password` (string):  Console display name: "Data Zoo Password".
* `username` (string):  Console display name: "Data Zoo Username".


Example:
```terraform
resource "davinci_connection" "dataZooConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "dataZooConnector"
  name         = "My awesome dataZooConnector"

  property {
    name  = "password"
    type  = "string"
    value = var.datazooconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.datazooconnector_property_username
  }
}
```


## Datadog API

Connector ID (`connector_id` in the resource): `connector-oai-datadogapi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authApiKey` (string): The API key for an account that has access to the Datadog API. Console display name: "Authentication API Key".
* `authApplicationKey` (string): The Application key for an account that has access to the Datadog API. Console display name: "Authentication Application Key".
* `basePath` (string): The base URL for contacting the Datadog API, such as "https://api.us3.datadoghq.com". Console display name: "API URL".


Example:
```terraform
resource "davinci_connection" "connector-oai-datadogapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-datadogapi"
  name         = "My awesome connector-oai-datadogapi"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-datadogapi_property_auth_api_key
  }

  property {
    name  = "authApplicationKey"
    type  = "string"
    value = var.connector-oai-datadogapi_property_auth_application_key
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-datadogapi_property_base_path
  }
}
```


## DeBounce

Connector ID (`connector_id` in the resource): `connectorDeBounce`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): A DeBounce API Key is physically a token/code of 13 random alphanumeric characters. If you need to create an API key, please log in to your DeBounce account and then navigate to the API section. Console display name: "API Key".


Example:
```terraform
resource "davinci_connection" "connectorDeBounce" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorDeBounce"
  name         = "My awesome connectorDeBounce"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectordebounce_property_api_key
  }
}
```


## Device Policy

Connector ID (`connector_id` in the resource): `devicePolicyConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "devicePolicyConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "devicePolicyConnector"
  name         = "My awesome devicePolicyConnector"
}
```


## DigiLocker

Connector ID (`connector_id` in the resource): `digilockerConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (json):  Console display name: "Oauth2 Parameters".


Example:
```terraform
resource "davinci_connection" "digilockerConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "digilockerConnector"
  name         = "My awesome digilockerConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Digidentity

Connector ID (`connector_id` in the resource): `digidentityConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (json):  Console display name: "Oauth2 Parameters".


Example:
```terraform
resource "davinci_connection" "digidentityConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "digidentityConnector"
  name         = "My awesome digidentityConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Druva inSync Cloud API

Connector ID (`connector_id` in the resource): `connector-oai-druvainsynccloud`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authClientId` (string): The Client ID of the authenticating application. Console display name: "Client ID".
* `authClientSecret` (string): The Secret Key for the authenticating application. Console display name: "Secret Key".
* `authTokenUrl` (string): The URL used to obtain an access token. Console display name: "Token URL".
* `basePath` (string): The base URL for contacting the API. Console display name: "Base Path".


Example:
```terraform
resource "davinci_connection" "connector-oai-druvainsynccloud" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-druvainsynccloud"
  name         = "My awesome connector-oai-druvainsynccloud"

  property {
    name  = "authClientId"
    type  = "string"
    value = var.connector-oai-druvainsynccloud_property_auth_client_id
  }

  property {
    name  = "authClientSecret"
    type  = "string"
    value = var.connector-oai-druvainsynccloud_property_auth_client_secret
  }

  property {
    name  = "authTokenUrl"
    type  = "string"
    value = var.connector-oai-druvainsynccloud_property_auth_token_url
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-druvainsynccloud_property_base_path
  }
}
```


## Duo

Connector ID (`connector_id` in the resource): `duoConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "duoConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "duoConnector"
  name         = "My awesome duoConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Entrust

Connector ID (`connector_id` in the resource): `entrustConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `applicationId` (string): The application ID for the Identity as a Service application. Console display name: "Application ID".
* `serviceDomain` (string): The domain of the Entrust service. Format is '<customer>.<region>.trustedauth.com'. For example, 'mycompany.us.trustedauth.com'. Console display name: "Service Domain".


Example:
```terraform
resource "davinci_connection" "entrustConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "entrustConnector"
  name         = "My awesome entrustConnector"

  property {
    name  = "applicationId"
    type  = "string"
    value = var.entrustconnector_property_application_id
  }

  property {
    name  = "serviceDomain"
    type  = "string"
    value = var.entrustconnector_property_service_domain
  }
}
```


## Equifax

Connector ID (`connector_id` in the resource): `equifaxConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseUrl` (string): Base URL for Equifax API. Console display name: "Base URL".
* `clientId` (string): When you Create a New App, Equifax will assign a Client ID per environment for the API Product. Console display name: "Client ID".
* `clientSecret` (string): When you Create a New App, Equifax will assign a Client Secret per environment for the API Product. Console display name: "Client Secret".
* `equifaxSoapApiEnvironment` (string): SOAP API WSDL Environment. Console display name: "SOAP API Environment".
* `memberNumber` (string): Unique Identifier of Customer. Please contact Equifax Sales Representative during client onboarding for this value. Console display name: "Member Number".
* `password` (string): Password provided by Equifax for SOAP API. Console display name: "Password for SOAP API".
* `username` (string): Username provided by Equifax for SOAP API. Console display name: "Username for SOAP API".


Example:
```terraform
resource "davinci_connection" "equifaxConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "equifaxConnector"
  name         = "My awesome equifaxConnector"

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.equifaxconnector_property_base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.equifaxconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.equifaxconnector_property_client_secret
  }

  property {
    name  = "equifaxSoapApiEnvironment"
    type  = "string"
    value = var.equifaxconnector_property_equifax_soap_api_environment
  }

  property {
    name  = "memberNumber"
    type  = "string"
    value = var.equifaxconnector_property_member_number
  }

  property {
    name  = "password"
    type  = "string"
    value = var.equifaxconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.equifaxconnector_property_username
  }
}
```


## Error Message

Connector ID (`connector_id` in the resource): `errorConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "errorConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "errorConnector"
  name         = "My awesome errorConnector"
}
```


## Facebook Login

Connector ID (`connector_id` in the resource): `facebookIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (json):  Console display name: "Oauth2 Parameters".


Example:
```terraform
resource "davinci_connection" "facebookIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "facebookIdpConnector"
  name         = "My awesome facebookIdpConnector"

  property {
    name = "oauth2"
    type = "json"
    value = jsonencode({
      "properties" : {
        "providerName" : {
          "type" : "string",
          "displayName" : "Provider Name",
          "preferredControlType" : "textField",
          "value" : "Login with Facebook"
        },
        "skRedirectUri" : {
          "type" : "string",
          "displayName" : "DaVinci Redirect URL",
          "info" : "Enter this in your identity provider configuration to allow it to redirect the browser back to DaVinci. If you use a custom PingOne domain, modify the URL accordingly.",
          "preferredControlType" : "textField",
          "disabled" : true,
          "initializeValue" : "SINGULARKEY_REDIRECT_URI",
          "copyToClip" : true
        },
        "clientId" : {
          "type" : "string",
          "displayName" : "Application ID",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "${var.facebookidpconnector_property_application_id}"
        },
        "clientSecret" : {
          "type" : "string",
          "displayName" : "Client Secret",
          "preferredControlType" : "textField",
          "secure" : true,
          "required" : true,
          "value" : "${var.facebookidpconnector_property_client_secret}"
        },
        "scope" : {
          "type" : "string",
          "displayName" : "Scope",
          "preferredControlType" : "textField",
          "requiredValue" : "email",
          "required" : true,
          "value" : "${var.facebookidpconnector_property_scope}"
        },
        "disableCreateUser" : {
          "displayName" : "Disable Shadow User",
          "preferredControlType" : "toggleSwitch",
          "value" : true,
          "info" : "A shadow user is implicitly created, unless disabled."
        },
        "userConnectorAttributeMapping" : {
          "type" : "object",
          "displayName" : null,
          "preferredControlType" : "userConnectorAttributeMapping",
          "newMappingAllowed" : true,
          "title1" : null,
          "title2" : null,
          "sections" : [
            "attributeMapping"
          ],
          "value" : {
            "userPoolConnectionId" : "defaultUserPool",
            "mapping" : {
              "username" : {
                "value1" : "id"
              },
              "name" : {
                "value1" : "name"
              },
              "email" : {
                "value1" : "email"
              }
            }
          }
        },
        "customAttributes" : {
          "type" : "array",
          "displayName" : "Connector Attributes",
          "preferredControlType" : "tableViewAttributes",
          "info" : "These attributes will be available in User Connector Attribute Mapping.",
          "sections" : [
            "connectorAttributes"
          ],
          "value" : [
            {
              "name" : "id",
              "description" : "ID",
              "type" : "string",
              "value" : null,
              "minLength" : "1",
              "maxLength" : "300",
              "required" : true,
              "attributeType" : "sk"
            },
            {
              "name" : "name",
              "description" : "Display Name",
              "type" : "string",
              "value" : null,
              "minLength" : "1",
              "maxLength" : "250",
              "required" : false,
              "attributeType" : "sk"
            },
            {
              "name" : "email",
              "description" : "Email",
              "type" : "string",
              "value" : null,
              "minLength" : "1",
              "maxLength" : "250",
              "required" : false,
              "attributeType" : "sk"
            }
          ]
        },
        "state" : {
          "displayName" : "Send state with request",
          "value" : true,
          "preferredControlType" : "toggleSwitch",
          "info" : "Send unique state value with every request"
        },
        "returnToUrl" : {
          "displayName" : "Application Return To URL",
          "preferredControlType" : "textField",
          "info" : "When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application.",
          "value" : "${var.facebookidpconnector_property_callback_url}"
        }
      }
    })
  }
}
```


## Fingerprint JS

Connector ID (`connector_id` in the resource): `fingerprintjsConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiToken` (string):  Console display name: "Fingerprint Subscription API Token".
* `javascriptCdnUrl` (string):  Console display name: "Javascript CDN URL".
* `token` (string):  Console display name: "Fingerprint Subscription Browser Token".


Example:
```terraform
resource "davinci_connection" "fingerprintjsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "fingerprintjsConnector"
  name         = "My awesome fingerprintjsConnector"

  property {
    name  = "apiToken"
    type  = "string"
    value = var.fingerprintjsconnector_property_api_token
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.fingerprintjsconnector_property_javascript_cdn_url
  }

  property {
    name  = "token"
    type  = "string"
    value = var.fingerprintjsconnector_property_token
  }
}
```


## Finicity

Connector ID (`connector_id` in the resource): `finicityConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `appKey` (string): Finicity App Key from Finicity Developer Portal. Console display name: "Finicity App Key".
* `baseUrl` (string): Base URL for Finicity API. Console display name: "Base URL".
* `partnerId` (string): The partner id you can obtain from your Finicity developer dashboard. Console display name: "Partner ID".
* `partnerSecret` (string): Partner Secret from Finicity Developer Portal. Console display name: "Partner Secret".


Example:
```terraform
resource "davinci_connection" "finicityConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "finicityConnector"
  name         = "My awesome finicityConnector"

  property {
    name  = "appKey"
    type  = "string"
    value = var.finicityconnector_property_app_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.finicityconnector_property_base_url
  }

  property {
    name  = "partnerId"
    type  = "string"
    value = var.finicityconnector_property_partner_id
  }

  property {
    name  = "partnerSecret"
    type  = "string"
    value = var.finicityconnector_property_partner_secret
  }
}
```


## Flow Analytics

Connector ID (`connector_id` in the resource): `analyticsConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "analyticsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "analyticsConnector"
  name         = "My awesome analyticsConnector"
}
```


## Flow Conductor

Connector ID (`connector_id` in the resource): `flowConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `enforcedSignedToken` (boolean):  Console display name: "Enforce Signed Token".
* `inputSchema` (string): Follow example for JSON schema. Console display name: "Input Schema".
* `pemPublicKey` (string): pem public key. Console display name: "Public Key".


Example:
```terraform
resource "davinci_connection" "flowConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "flowConnector"
  name         = "My awesome flowConnector"

  property {
    name  = "enforcedSignedToken"
    type  = "boolean"
    value = var.flowconnector_property_enforced_signed_token
  }

  property {
    name  = "inputSchema"
    type  = "string"
    value = var.flowconnector_property_input_schema
  }

  property {
    name  = "pemPublicKey"
    type  = "string"
    value = var.flowconnector_property_pem_public_key
  }
}
```


## Forter

Connector ID (`connector_id` in the resource): `forterConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiVersion` (string): API Version. Console display name: " Forter API Version".
* `secretKey` (string): Secret Key from Forter tenant. Console display name: "Forter Secret Key".
* `siteId` (string): Site ID from Forter tenant. Console display name: "Forter SiteID".


Example:
```terraform
resource "davinci_connection" "forterConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "forterConnector"
  name         = "My awesome forterConnector"

  property {
    name  = "apiVersion"
    type  = "string"
    value = var.forterconnector_property_api_version
  }

  property {
    name  = "secretKey"
    type  = "string"
    value = var.forterconnector_property_secret_key
  }

  property {
    name  = "siteId"
    type  = "string"
    value = var.forterconnector_property_site_id
  }
}
```


## Freshdesk

Connector ID (`connector_id` in the resource): `connectorFreshdesk`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Make sure that the "APIkey:X" is Base64-encoded before pasting into the text field. Console display name: "Freshdesk API Key".
* `baseURL` (string): The <tenant>.freshdesk.com URL or custom domain. Console display name: "Freshdesk Base URL (or Domain)".
* `version` (string): The current Freshdesk API Version. Console display name: "Freshdesk API Version".


Example:
```terraform
resource "davinci_connection" "connectorFreshdesk" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorFreshdesk"
  name         = "My awesome connectorFreshdesk"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorfreshdesk_property_api_key
  }

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "version"
    type  = "string"
    value = var.connectorfreshdesk_property_version
  }
}
```


## Freshservice

Connector ID (`connector_id` in the resource): `connectorFreshservice`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Your Freshservice API key. Console display name: "API Key".
* `domain` (string): Your Freshservice domain. Example: https://domain.freshservice.com/. Console display name: "Domain".


Example:
```terraform
resource "davinci_connection" "connectorFreshservice" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorFreshservice"
  name         = "My awesome connectorFreshservice"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorfreshservice_property_api_key
  }

  property {
    name  = "domain"
    type  = "string"
    value = var.connectorfreshservice_property_domain
  }
}
```


## Functions

Connector ID (`connector_id` in the resource): `functionsConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "functionsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "functionsConnector"
  name         = "My awesome functionsConnector"
}
```


## GBG

Connector ID (`connector_id` in the resource): `gbgConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `password` (string):  Console display name: "GBG Password".
* `requestUrl` (string):  Console display name: "Request URL".
* `soapAction` (string): SOAP Action is a header required for the soap request. Console display name: "Soap Action URL".
* `username` (string):  Console display name: "GBG Username".


Example:
```terraform
resource "davinci_connection" "gbgConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "gbgConnector"
  name         = "My awesome gbgConnector"

  property {
    name  = "password"
    type  = "string"
    value = var.gbgconnector_property_password
  }

  property {
    name  = "requestUrl"
    type  = "string"
    value = var.gbgconnector_property_request_url
  }

  property {
    name  = "soapAction"
    type  = "string"
    value = var.gbgconnector_property_soap_action
  }

  property {
    name  = "username"
    type  = "string"
    value = var.gbgconnector_property_username
  }
}
```


## GitHub API

Connector ID (`connector_id` in the resource): `connector-oai-github`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiVersion` (string): The GitHub v3 REST API version, such as "2022-11-28". Console display name: "API Version".
* `authBearerToken` (string): The authentication bearer token that has access to GitHub v3 REST API. Console display name: "Authentication Bearer Token".
* `basePath` (string): The base URL for the GitHub API, such as "https://api.github.com". Console display name: "API URL".


Example:
```terraform
resource "davinci_connection" "connector-oai-github" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-github"
  name         = "My awesome connector-oai-github"

  property {
    name  = "apiVersion"
    type  = "string"
    value = var.connector-oai-github_property_api_version
  }

  property {
    name  = "authBearerToken"
    type  = "string"
    value = var.connector-oai-github_property_auth_bearer_token
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-github_property_base_path
  }
}
```


## GitHub Login

Connector ID (`connector_id` in the resource): `githubIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (json):  Console display name: "Oauth2 Parameters".


Example:
```terraform
resource "davinci_connection" "githubIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "githubIdpConnector"
  name         = "My awesome githubIdpConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Google Analytics (Universal Analytics)

Connector ID (`connector_id` in the resource): `connectorGoogleanalyticsUA`

Properties (used in the `property` block in the resource as the `name` parameter):

* `trackingID` (string): The tracking ID / web property ID. The format is UA-XXXX-Y. All collected data is associated by this ID. Console display name: "Tracking ID".
* `version` (string): The Protocol version. The current value is '1'. This will only change when there are changes made that are not backwards compatible. Console display name: "Version".


Example:
```terraform
resource "davinci_connection" "connectorGoogleanalyticsUA" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorGoogleanalyticsUA"
  name         = "My awesome connectorGoogleanalyticsUA"

  property {
    name  = "trackingID"
    type  = "string"
    value = var.tracking_id
  }

  property {
    name  = "version"
    type  = "string"
    value = var.connectorgoogleanalyticsua_property_version
  }
}
```


## Google Chrome Enterprise Device Trust

Connector ID (`connector_id` in the resource): `connectorGoogleChromeEnterprise`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorGoogleChromeEnterprise" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorGoogleChromeEnterprise"
  name         = "My awesome connectorGoogleChromeEnterprise"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Google Login

Connector ID (`connector_id` in the resource): `googleConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (json):  Console display name: "OpenId Parameters".


Example:
```terraform
resource "davinci_connection" "googleConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "googleConnector"
  name         = "My awesome googleConnector"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Google Workspace Admin

Connector ID (`connector_id` in the resource): `googleWorkSpaceAdminConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `iss` (string): The email address associated with the Google Workspace service, such as "google-workspace-admin@xenon-set-123456.iam.gserviceaccount.com". Console display name: "Service Account Email Address".
* `privateKey` (string): The private key associated with the public key that you added to the Google Workspace service. Console display name: "Private Key".
* `sub` (string): The administrator's email address. Console display name: "Admin Email Address".


Example:
```terraform
resource "davinci_connection" "googleWorkSpaceAdminConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "googleWorkSpaceAdminConnector"
  name         = "My awesome googleWorkSpaceAdminConnector"

  property {
    name  = "iss"
    type  = "string"
    value = var.googleworkspaceadminconnector_property_iss
  }

  property {
    name  = "privateKey"
    type  = "string"
    value = var.googleworkspaceadminconnector_property_private_key
  }

  property {
    name  = "sub"
    type  = "string"
    value = var.googleworkspaceadminconnector_property_sub
  }
}
```


## HTTP

Connector ID (`connector_id` in the resource): `httpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `connectionId` (string):  Console display name: "Select an OpenID token management connection for signed HTTP responses.".
* `recaptchaSecretKey` (string): The Secret Key from reCAPTCHA Admin dashboard. Console display name: "reCAPTCHA v2 Secret Key".
* `recaptchaSiteKey` (string): The Site Key from reCAPTCHA Admin dashboard. Console display name: "reCAPTCHA v2 Site Key".
* `whiteList` (string): Enter the hostname for the trusted sites that host your HTML. Note: Ensure that the content hosted on these sites can be trusted and that publishing safeguards are in place to prevent unexpected issues. Console display name: "Trusted Sites".


Example:
```terraform
resource "davinci_connection" "httpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "httpConnector"
  name         = "My awesome httpConnector"

  property {
    name  = "connectionId"
    type  = "string"
    value = var.httpconnector_property_connection_id
  }

  property {
    name  = "recaptchaSecretKey"
    type  = "string"
    value = var.httpconnector_property_recaptcha_secret_key
  }

  property {
    name  = "recaptchaSiteKey"
    type  = "string"
    value = var.httpconnector_property_recaptcha_site_key
  }

  property {
    name  = "whiteList"
    type  = "string"
    value = var.httpconnector_property_white_list
  }
}
```


## HUMAN

Connector ID (`connector_id` in the resource): `connectorHuman`

Properties (used in the `property` block in the resource as the `name` parameter):

* `humanAuthenticationToken` (string): Bearer Token from HUMAN. Console display name: "HUMAN Authentication Token".
* `humanCustomerID` (string): Customer ID from HUMAN. Console display name: "HUMAN Customer ID".
* `humanPolicyName` (string): HUMAN mitigation policy name. Console display name: "HUMAN Policy Name".


Example:
```terraform
resource "davinci_connection" "connectorHuman" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorHuman"
  name         = "My awesome connectorHuman"

  property {
    name  = "humanAuthenticationToken"
    type  = "string"
    value = var.connectorhuman_property_human_authentication_token
  }

  property {
    name  = "humanCustomerID"
    type  = "string"
    value = var.human_customer_id
  }

  property {
    name  = "humanPolicyName"
    type  = "string"
    value = var.connectorhuman_property_human_policy_name
  }
}
```


## HUMAN

Connector ID (`connector_id` in the resource): `humanCompromisedConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `appId` (string): App ID from your HUMAN Tenant. Console display name: "HUMAN App ID".
* `authToken` (string): Auth Token from your HUMAN Tenant. Console display name: "HUMAN Auth Token".


Example:
```terraform
resource "davinci_connection" "humanCompromisedConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "humanCompromisedConnector"
  name         = "My awesome humanCompromisedConnector"

  property {
    name  = "appId"
    type  = "string"
    value = var.humancompromisedconnector_property_app_id
  }

  property {
    name  = "authToken"
    type  = "string"
    value = var.humancompromisedconnector_property_auth_token
  }
}
```


## HYPR

Connector ID (`connector_id` in the resource): `hyprConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "hyprConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "hyprConnector"
  name         = "My awesome hyprConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## HYPR Adapt

Connector ID (`connector_id` in the resource): `connectorHyprAdapt`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessToken` (string): Access Token. Console display name: "HYPR Adapt Access Token".


Example:
```terraform
resource "davinci_connection" "connectorHyprAdapt" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorHyprAdapt"
  name         = "My awesome connectorHyprAdapt"

  property {
    name  = "accessToken"
    type  = "string"
    value = var.connectorhypradapt_property_access_token
  }
}
```


## Have I Been Pwned

Connector ID (`connector_id` in the resource): `haveIBeenPwnedConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string):  Console display name: "Have I Been Pwned API Key".
* `apiUrl` (string):  Console display name: "API Server URL".
* `userAgent` (string):  


Example:
```terraform
resource "davinci_connection" "haveIBeenPwnedConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "haveIBeenPwnedConnector"
  name         = "My awesome haveIBeenPwnedConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.haveibeenpwnedconnector_property_api_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.haveibeenpwnedconnector_property_api_url
  }

  property {
    name  = "userAgent"
    type  = "string"
    value = var.haveibeenpwnedconnector_property_user_agent
  }
}
```


## Hellō Connector

Connector ID (`connector_id` in the resource): `connectorHello`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorHello" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorHello"
  name         = "My awesome connectorHello"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## HubSpot Companies API

Connector ID (`connector_id` in the resource): `connector-oai-hubspotcompanies`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authBearerToken` (string): The authenticating token. Console display name: "Bearer token".
* `basePath` (string): The base URL for contacting the API. Console display name: "Base Path".


Example:
```terraform
resource "davinci_connection" "connector-oai-hubspotcompanies" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-hubspotcompanies"
  name         = "My awesome connector-oai-hubspotcompanies"

  property {
    name  = "authBearerToken"
    type  = "string"
    value = var.connector-oai-hubspotcompanies_property_auth_bearer_token
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-hubspotcompanies_property_base_path
  }
}
```


## Hubspot

Connector ID (`connector_id` in the resource): `connectorHubspot`

Properties (used in the `property` block in the resource as the `name` parameter):

* `bearerToken` (string): Your unique API key. Console display name: "API Key".


Example:
```terraform
resource "davinci_connection" "connectorHubspot" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorHubspot"
  name         = "My awesome connectorHubspot"

  property {
    name  = "bearerToken"
    type  = "string"
    value = var.connectorhubspot_property_bearer_token
  }
}
```


## ID DataWeb

Connector ID (`connector_id` in the resource): `idDatawebConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "idDatawebConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idDatawebConnector"
  name         = "My awesome idDatawebConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## ID R&D

Connector ID (`connector_id` in the resource): `idranddConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string):  Console display name: "API Key".
* `apiUrl` (string):  Console display name: "API Server URL".


Example:
```terraform
resource "davinci_connection" "idranddConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idranddConnector"
  name         = "My awesome idranddConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.idranddconnector_property_api_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.idranddconnector_property_api_url
  }
}
```


## ID.me

Connector ID (`connector_id` in the resource): `idMeConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (json):  Console display name: "Oauth2 Parameters".


Example:
```terraform
resource "davinci_connection" "idMeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idMeConnector"
  name         = "My awesome idMeConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
```


## ID.me - Community Verification

Connector ID (`connector_id` in the resource): `idmecommunityConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (json):  Console display name: "OpenId Parameters".


Example:
```terraform
resource "davinci_connection" "idmecommunityConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idmecommunityConnector"
  name         = "My awesome idmecommunityConnector"

  property {
    name  = "openId"
    type  = "json"
    value = var.idmecommunityconnector_property_open_id
  }
}
```


## ID.me - Identity Verification

Connector ID (`connector_id` in the resource): `connectorIdMeIdentity`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (json):  Console display name: "OpenId Parameters".


Example:
```terraform
resource "davinci_connection" "connectorIdMeIdentity" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIdMeIdentity"
  name         = "My awesome connectorIdMeIdentity"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
```


## IDEMIA

Connector ID (`connector_id` in the resource): `idemiaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apikey` (string):  Console display name: "API Key".
* `baseUrl` (string): Base Url for IDEMIA API. Can be found in the dashboard documents. Console display name: "IDEMIA API base URL".


Example:
```terraform
resource "davinci_connection" "idemiaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idemiaConnector"
  name         = "My awesome idemiaConnector"

  property {
    name  = "apikey"
    type  = "string"
    value = var.idemiaconnector_property_apikey
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.idemiaconnector_property_base_url
  }
}
```


## IDI Data

Connector ID (`connector_id` in the resource): `skPeopleIntelligenceConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authUrl` (string):  Console display name: "Authorization URL".
* `clientId` (string):  Console display name: "Client ID".
* `clientSecret` (string):  Console display name: "Client Secret".
* `dppa` (string):  Console display name: "DPPA".
* `glba` (string):  Console display name: "GLBA".
* `searchUrl` (string):  Console display name: "Search URL".


Example:
```terraform
resource "davinci_connection" "skPeopleIntelligenceConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "skPeopleIntelligenceConnector"
  name         = "My awesome skPeopleIntelligenceConnector"

  property {
    name  = "authUrl"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_auth_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_client_secret
  }

  property {
    name  = "dppa"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_dppa
  }

  property {
    name  = "glba"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_glba
  }

  property {
    name  = "searchUrl"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_search_url
  }
}
```


## IDI coreIDENTITY

Connector ID (`connector_id` in the resource): `connectorIdiVERIFIED`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiSecret` (string): Please enter your API secret that IDI coreIDENTITY has provided you. Console display name: "API Secret".
* `companyKey` (string): Please enter the company key that IDI coreIDENTITY has assigned. Console display name: "Company Key".
* `idiEnv` (string): Please choose which coreIDENTITY environment you would like to query . Console display name: "Environment".
* `siteKey` (string): Please enter your site key that IDI coreIDENTITY has provided you. Console display name: "Site Key".
* `uniqueUrl` (string): Please enter your unique URL that IDI coreIDENTITY has provided you. Console display name: "Unique URL".


Example:
```terraform
resource "davinci_connection" "connectorIdiVERIFIED" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIdiVERIFIED"
  name         = "My awesome connectorIdiVERIFIED"

  property {
    name  = "apiSecret"
    type  = "string"
    value = var.connectoridiverified_property_api_secret
  }

  property {
    name  = "companyKey"
    type  = "string"
    value = var.connectoridiverified_property_company_key
  }

  property {
    name  = "idiEnv"
    type  = "string"
    value = var.connectoridiverified_property_idi_env
  }

  property {
    name  = "siteKey"
    type  = "string"
    value = var.connectoridiverified_property_site_key
  }

  property {
    name  = "uniqueUrl"
    type  = "string"
    value = var.connectoridiverified_property_unique_url
  }
}
```


## IDmelon

Connector ID (`connector_id` in the resource): `connectorIdmelon`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorIdmelon" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIdmelon"
  name         = "My awesome connectorIdmelon"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## IDmission

Connector ID (`connector_id` in the resource): `idmissionConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authDescription` (string):  Console display name: "Authentication Description".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `loginId` (string):  Console display name: "Sign On ID".
* `merchantId` (string):  Console display name: "Merchant ID".
* `password` (string):  Console display name: "Password".
* `productId` (string):  Console display name: "Product ID".
* `productName` (string):  Console display name: "Product Name".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".
* `url` (string):  Console display name: "IDmission Server URL".


Example:
```terraform
resource "davinci_connection" "idmissionConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idmissionConnector"
  name         = "My awesome idmissionConnector"

  property {
    name  = "authDescription"
    type  = "string"
    value = var.idmissionconnector_property_auth_description
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.idmissionconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.idmissionconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.idmissionconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.idmissionconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.idmissionconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.idmissionconnector_property_icon_url_png
  }

  property {
    name  = "loginId"
    type  = "string"
    value = var.idmissionconnector_property_login_id
  }

  property {
    name  = "merchantId"
    type  = "string"
    value = var.idmissionconnector_property_merchant_id
  }

  property {
    name  = "password"
    type  = "string"
    value = var.idmissionconnector_property_password
  }

  property {
    name  = "productId"
    type  = "string"
    value = var.idmissionconnector_property_product_id
  }

  property {
    name  = "productName"
    type  = "string"
    value = var.idmissionconnector_property_product_name
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.idmissionconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.idmissionconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.idmissionconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.idmissionconnector_property_tool_tip
  }

  property {
    name  = "url"
    type  = "string"
    value = var.idmissionconnector_property_url
  }
}
```


## IDmission - OIDC

Connector ID (`connector_id` in the resource): `idmissionOidcConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "idmissionOidcConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idmissionOidcConnector"
  name         = "My awesome idmissionOidcConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.idmissionoidcconnector_property_custom_auth
  }
}
```


## IdRamp

Connector ID (`connector_id` in the resource): `idrampOidcConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "idrampOidcConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idrampOidcConnector"
  name         = "My awesome idrampOidcConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Image

Connector ID (`connector_id` in the resource): `imageConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "imageConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "imageConnector"
  name         = "My awesome imageConnector"
}
```


## Incode

Connector ID (`connector_id` in the resource): `incodeConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "incodeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "incodeConnector"
  name         = "My awesome incodeConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Infinipoint

Connector ID (`connector_id` in the resource): `connectorInfinipoint`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorInfinipoint" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorInfinipoint"
  name         = "My awesome connectorInfinipoint"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Intellicheck

Connector ID (`connector_id` in the resource): `intellicheckConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key from your Intellicheck tenant. Console display name: "API Key".
* `baseUrl` (string): Base URL from your Intellicheck tenant (Including protocol - https://). Console display name: "Base URL".
* `customerId` (string): Customer ID from your Intellicheck tenant. Console display name: "Customer ID".


Example:
```terraform
resource "davinci_connection" "intellicheckConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "intellicheckConnector"
  name         = "My awesome intellicheckConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.intellicheckconnector_property_api_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.intellicheckconnector_property_base_url
  }

  property {
    name  = "customerId"
    type  = "string"
    value = var.intellicheckconnector_property_customer_id
  }
}
```


## Jamf

Connector ID (`connector_id` in the resource): `connectorJamf`

Properties (used in the `property` block in the resource as the `name` parameter):

* `jamfPassword` (string): Enter Password for token. Console display name: "JAMF Password".
* `jamfUsername` (string): Enter Username for token. Console display name: "JAMF Username".
* `serverName` (string): Enter Server Name for Base URL. Console display name: "Server Name".


Example:
```terraform
resource "davinci_connection" "connectorJamf" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorJamf"
  name         = "My awesome connectorJamf"

  property {
    name  = "jamfPassword"
    type  = "string"
    value = var.connectorjamf_property_jamf_password
  }

  property {
    name  = "jamfUsername"
    type  = "string"
    value = var.connectorjamf_property_jamf_username
  }

  property {
    name  = "serverName"
    type  = "string"
    value = var.connectorjamf_property_server_name
  }
}
```


## Jira

Connector ID (`connector_id` in the resource): `jiraConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): You may need to create a token from Jira with your credentials, if you haven't created one. Console display name: "Jira API token".
* `apiUrl` (string): Base URL of the Jira instance. Console display name: "Base Url".
* `email` (string): Email used for your Jira account. Console display name: "Email Address".


Example:
```terraform
resource "davinci_connection" "jiraConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "jiraConnector"
  name         = "My awesome jiraConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.jiraconnector_property_api_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.jiraconnector_property_api_url
  }

  property {
    name  = "email"
    type  = "string"
    value = var.jiraconnector_property_email
  }
}
```


## Jira Service Desk

Connector ID (`connector_id` in the resource): `connectorJiraServiceDesk`

Properties (used in the `property` block in the resource as the `name` parameter):

* `JIRAServiceDeskAuth` (string): Bearer Authorization Token for JIRA Service Desk. Console display name: "Bearer Authorization Token for JIRA Service Desk".
* `JIRAServiceDeskCreateData` (string): Raw JSON body to create new JIRA service desk request. Example: {   "requestParticipants": ["qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d69abfa3980ce712caae"   ],   "serviceDeskId": "10",   "requestTypeId": "25",   "requestFieldValues": {     "summary": "Request JSD help via REST",     "description": "I need a new *mouse* for my Mac"   } }. Console display name: "Raw JSON for creating new JIRA service desk request".
* `JIRAServiceDeskURL` (string): URL for JIRA Service Desk. Example: your-domain.atlassian.net. Console display name: "JIRA Service Desk URL".
* `JIRAServiceDeskUpdateData` (string): Raw JSON body to update JIRA service desk request. Example: {"id": "1","additionalComment": {"body": "I have fixed the problem."}}. Console display name: "Raw JSON for updating JIRA service desk".
* `method` (string): The HTTP Method. Console display name: "Method".


Example:
```terraform
resource "davinci_connection" "connectorJiraServiceDesk" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorJiraServiceDesk"
  name         = "My awesome connectorJiraServiceDesk"

  property {
    name  = "JIRAServiceDeskAuth"
    type  = "string"
    value = var.jira_service_desk_auth
  }

  property {
    name  = "JIRAServiceDeskCreateData"
    type  = "string"
    value = var.jira_service_desk_create_data
  }

  property {
    name  = "JIRAServiceDeskURL"
    type  = "string"
    value = var.jira_service_desk_url
  }

  property {
    name  = "JIRAServiceDeskUpdateData"
    type  = "string"
    value = var.jira_service_desk_update_data
  }

  property {
    name  = "method"
    type  = "string"
    value = var.connectorjiraservicedesk_property_method
  }
}
```


## Jumio

Connector ID (`connector_id` in the resource): `jumioConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string):  Console display name: "API Key".
* `authDescription` (string):  Console display name: "Authentication Description".
* `authUrl` (string):  Console display name: "Base URL for Authentication".
* `authorizationTokenLifetime` (number): default: 1800 (30 minutes). maximum: 5184000 (60 days). Console display name: "Time Transaction URL Valid (seconds)".
* `baseColor` (string): Must be passed with bgColor. Console display name: "HEX Main Color".
* `bgColor` (string): Must be passed with baseColor. Console display name: "HEX Background Color.".
* `callbackUrl` (string):  Console display name: "Callback URL".
* `clientSecret` (string):  Console display name: "API Secret".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `doNotShowInIframe` (boolean): If this is true, user will be redirected to the verification url and then redirected back when complete. Console display name: "Do not show in iFrame".
* `docVerificationUrl` (string):  Console display name: "Document Verification Url".
* `headerImageUrl` (string): Logo must be: landscape (16:9 or 4:3), min. height of 192 pixels, size 8-64 KB. Console display name: "Custom Header Logo URL".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `locale` (string): Renders content in the specified language. Console display name: "Locale".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```terraform
resource "davinci_connection" "jumioConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "jumioConnector"
  name         = "My awesome jumioConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.jumioconnector_property_api_key
  }

  property {
    name  = "authDescription"
    type  = "string"
    value = var.jumioconnector_property_auth_description
  }

  property {
    name  = "authUrl"
    type  = "string"
    value = var.jumioconnector_property_auth_url
  }

  property {
    name  = "authorizationTokenLifetime"
    type  = "number"
    value = var.jumioconnector_property_authorization_token_lifetime
  }

  property {
    name  = "baseColor"
    type  = "string"
    value = var.jumioconnector_property_base_color
  }

  property {
    name  = "bgColor"
    type  = "string"
    value = var.jumioconnector_property_bg_color
  }

  property {
    name  = "callbackUrl"
    type  = "string"
    value = var.jumioconnector_property_callback_url
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.jumioconnector_property_client_secret
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.jumioconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.jumioconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.jumioconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.jumioconnector_property_details2
  }

  property {
    name  = "doNotShowInIframe"
    type  = "boolean"
    value = var.jumioconnector_property_do_not_show_in_iframe
  }

  property {
    name  = "docVerificationUrl"
    type  = "string"
    value = var.jumioconnector_property_doc_verification_url
  }

  property {
    name  = "headerImageUrl"
    type  = "string"
    value = var.jumioconnector_property_header_image_url
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.jumioconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.jumioconnector_property_icon_url_png
  }

  property {
    name  = "locale"
    type  = "string"
    value = var.jumioconnector_property_locale
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.jumioconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.jumioconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.jumioconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.jumioconnector_property_tool_tip
  }
}
```


## KBA

Connector ID (`connector_id` in the resource): `kbaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authDescription` (string):  Console display name: "Authentication Description".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `formFieldsList` (json):  Console display name: "Fields List".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```terraform
resource "davinci_connection" "kbaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "kbaConnector"
  name         = "My awesome kbaConnector"

  property {
    name  = "authDescription"
    type  = "string"
    value = var.kbaconnector_property_auth_description
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.kbaconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.kbaconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.kbaconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.kbaconnector_property_details2
  }

  property {
    name  = "formFieldsList"
    type  = "json"
    value = var.kbaconnector_property_form_fields_list
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.kbaconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.kbaconnector_property_icon_url_png
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.kbaconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.kbaconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.kbaconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.kbaconnector_property_tool_tip
  }
}
```


## KYXStart

Connector ID (`connector_id` in the resource): `kyxstartConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): KYXStart Client ID. Console display name: "Client ID".
* `clientSecret` (string): KYXStart Client Secret. Console display name: "Client Secret".
* `tenantName` (string): Tenant Name from KYXStart Account. Console display name: "Tenant Name".


Example:
```terraform
resource "davinci_connection" "kyxstartConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "kyxstartConnector"
  name         = "My awesome kyxstartConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.kyxstartconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.kyxstartconnector_property_client_secret
  }

  property {
    name  = "tenantName"
    type  = "string"
    value = var.kyxstartconnector_property_tenant_name
  }
}
```


## Kaizen Secure Voiz

Connector ID (`connector_id` in the resource): `kaizenVoizConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): example: http://<server_root>/ksvvoiceservice/rest/service. Console display name: "API Server URL".
* `applicationName` (string):  Console display name: "Application Name".
* `authDescription` (string):  Console display name: "Authentication Description".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```terraform
resource "davinci_connection" "kaizenVoizConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "kaizenVoizConnector"
  name         = "My awesome kaizenVoizConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.kaizenvoizconnector_property_api_url
  }

  property {
    name  = "applicationName"
    type  = "string"
    value = var.kaizenvoizconnector_property_application_name
  }

  property {
    name  = "authDescription"
    type  = "string"
    value = var.kaizenvoizconnector_property_auth_description
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.kaizenvoizconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.kaizenvoizconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.kaizenvoizconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.kaizenvoizconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.kaizenvoizconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.kaizenvoizconnector_property_icon_url_png
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.kaizenvoizconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.kaizenvoizconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.kaizenvoizconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.kaizenvoizconnector_property_tool_tip
  }
}
```


## Keyless

Connector ID (`connector_id` in the resource): `connectorKeyless`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorKeyless" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorKeyless"
  name         = "My awesome connectorKeyless"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Keyri QR Login

Connector ID (`connector_id` in the resource): `connectorKeyri`

*No properties*


Example:
```terraform
resource "davinci_connection" "connectorKeyri" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorKeyri"
  name         = "My awesome connectorKeyri"
}
```


## LDAP

Connector ID (`connector_id` in the resource): `pingOneLDAPConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne environment ID. Console display name: "Environment ID".
* `gatewayId` (string): Your PingOne LDAP gateway ID. Console display name: "Gateway ID".
* `region` (string): The region in which your PingOne environment exists. Console display name: "Region".


Example:
```terraform
resource "davinci_connection" "pingOneLDAPConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneLDAPConnector"
  name         = "My awesome pingOneLDAPConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingoneldapconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingoneldapconnector_property_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.pingoneldapconnector_property_env_id
  }

  property {
    name  = "gatewayId"
    type  = "string"
    value = var.pingoneldapconnector_property_gateway_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingoneldapconnector_property_region
  }
}
```


## LaunchDarkly API

Connector ID (`connector_id` in the resource): `connector-oai-launchdarklyrestapi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authApiKey` (string): The authentication key to the LaunchDarkly REST API. Console display name: "API Key".
* `basePath` (string): The base URL for contacting the API. Console display name: "Base Path".


Example:
```terraform
resource "davinci_connection" "connector-oai-launchdarklyrestapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-launchdarklyrestapi"
  name         = "My awesome connector-oai-launchdarklyrestapi"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-launchdarklyrestapi_property_auth_api_key
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-launchdarklyrestapi_property_base_path
  }
}
```


## LexisNexis

Connector ID (`connector_id` in the resource): `lexisnexisV2Connector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Your LexisNexis API key, such as “o3x9ywfs26rm1zvl”. Console display name: "API Key".
* `apiUrl` (string): The API URL to target. For a custom value, select Use Custom API URL and enter a value in the Custom API URL field. Console display name: "API URL".
* `orgId` (string): Your LexisNexis organization ID, such as “4en6ll2s”. Console display name: "Organization ID".
* `useCustomApiURL` (string): The API URL to target, such as “https://h.online-metrix.net”. Console display name: "Custom API URL".


Example:
```terraform
resource "davinci_connection" "lexisnexisV2Connector" {
  environment_id = var.pingone_environment_id

  connector_id = "lexisnexisV2Connector"
  name         = "My awesome lexisnexisV2Connector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.lexisnexisv2connector_property_api_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.lexisnexisv2connector_property_api_url
  }

  property {
    name  = "orgId"
    type  = "string"
    value = var.lexisnexisv2connector_property_org_id
  }

  property {
    name  = "useCustomApiURL"
    type  = "string"
    value = var.use_custom_api_url
  }
}
```


## LinkedIn Login

Connector ID (`connector_id` in the resource): `linkedInConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (json):  Console display name: "Oauth2 Parameters".


Example:
```terraform
resource "davinci_connection" "linkedInConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "linkedInConnector"
  name         = "My awesome linkedInConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Location Policy

Connector ID (`connector_id` in the resource): `locationPolicyConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "locationPolicyConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "locationPolicyConnector"
  name         = "My awesome locationPolicyConnector"
}
```


## Mailchimp

Connector ID (`connector_id` in the resource): `connectorMailchimp`

Properties (used in the `property` block in the resource as the `name` parameter):

* `transactionalApiKey` (string): The Transactional API Key is used to send data to the transactional API. Console display name: "Transactional API Key".
* `transactionalApiVersion` (string): Mailchimp - Transactional API Version. Console display name: "Transactional API Version".


Example:
```terraform
resource "davinci_connection" "connectorMailchimp" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorMailchimp"
  name         = "My awesome connectorMailchimp"

  property {
    name  = "transactionalApiKey"
    type  = "string"
    value = var.connectormailchimp_property_transactional_api_key
  }

  property {
    name  = "transactionalApiVersion"
    type  = "string"
    value = var.connectormailchimp_property_transactional_api_version
  }
}
```


## Mailgun

Connector ID (`connector_id` in the resource): `connectorMailgun`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Mailgun API Key. Console display name: "API Key".
* `apiVersion` (string): Mailgun API Version. Console display name: "API Version".
* `mailgunDomain` (string): Name of the desired domain (e.g. mail.mycompany.com). Console display name: "Domain".


Example:
```terraform
resource "davinci_connection" "connectorMailgun" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorMailgun"
  name         = "My awesome connectorMailgun"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectormailgun_property_api_key
  }

  property {
    name  = "apiVersion"
    type  = "string"
    value = var.connectormailgun_property_api_version
  }

  property {
    name  = "mailgunDomain"
    type  = "string"
    value = var.connectormailgun_property_mailgun_domain
  }
}
```


## Mailjet API

Connector ID (`connector_id` in the resource): `connector-oai-mailjetapi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authPassword` (string): API Secret Key. Console display name: "API Secret Key".
* `authUsername` (string): API Key. Console display name: "API Key".
* `basePath` (string): The base URL for contacting the API. Console display name: "Base Path".


Example:
```terraform
resource "davinci_connection" "connector-oai-mailjetapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-mailjetapi"
  name         = "My awesome connector-oai-mailjetapi"

  property {
    name  = "authPassword"
    type  = "string"
    value = var.connector-oai-mailjetapi_property_auth_password
  }

  property {
    name  = "authUsername"
    type  = "string"
    value = var.connector-oai-mailjetapi_property_auth_username
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-mailjetapi_property_base_path
  }
}
```


## Melissa Global Address

Connector ID (`connector_id` in the resource): `melissaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): License Key is the API key that you can retrieve from Melissa Admin Portal. Console display name: "License Key".


Example:
```terraform
resource "davinci_connection" "melissaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "melissaConnector"
  name         = "My awesome melissaConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.melissaconnector_property_api_key
  }
}
```


## Microsoft Dynamics - Customer Insights

Connector ID (`connector_id` in the resource): `microsoftDynamicsCustomerInsightsConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseURL` (string): Base URL. Console display name: "Base URL".
* `clientId` (string): Client ID. Console display name: "Client ID".
* `clientSecret` (string): Client Secret. Console display name: "Client Secret".
* `environmentName` (string): Environment Name. Console display name: "Environment Name".
* `grantType` (string): Grant Type. Console display name: "Grant Type".
* `tenant` (string): Tenant. Console display name: "Tenant".
* `version` (string): Web API Version. Console display name: "Version".


Example:
```terraform
resource "davinci_connection" "microsoftDynamicsCustomerInsightsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "microsoftDynamicsCustomerInsightsConnector"
  name         = "My awesome microsoftDynamicsCustomerInsightsConnector"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_base_u_r_l
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_client_secret
  }

  property {
    name  = "environmentName"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_environment_name
  }

  property {
    name  = "grantType"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_grant_type
  }

  property {
    name  = "tenant"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_tenant
  }

  property {
    name  = "version"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_version
  }
}
```


## Microsoft Edge for Business

Connector ID (`connector_id` in the resource): `connectorMicrosoftEdge`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorMicrosoftEdge" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorMicrosoftEdge"
  name         = "My awesome connectorMicrosoftEdge"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.connectormicrosoftedge_property_custom_auth
  }
}
```


## Microsoft Intune

Connector ID (`connector_id` in the resource): `connectorMicrosoftIntune`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): Client ID. Console display name: "Client ID".
* `clientSecret` (string): Client Secret. Console display name: "Client Secret".
* `grantType` (string): Grant Type. Console display name: "Grant Type".
* `scope` (string): Scope. Console display name: "Scope".
* `tenant` (string): Tenant. Console display name: "Tenant".


Example:
```terraform
resource "davinci_connection" "connectorMicrosoftIntune" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorMicrosoftIntune"
  name         = "My awesome connectorMicrosoftIntune"

  property {
    name  = "clientId"
    type  = "string"
    value = var.connectormicrosoftintune_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectormicrosoftintune_property_client_secret
  }

  property {
    name  = "grantType"
    type  = "string"
    value = var.connectormicrosoftintune_property_grant_type
  }

  property {
    name  = "scope"
    type  = "string"
    value = var.connectormicrosoftintune_property_scope
  }

  property {
    name  = "tenant"
    type  = "string"
    value = var.connectormicrosoftintune_property_tenant
  }
}
```


## Microsoft Login

Connector ID (`connector_id` in the resource): `microsoftIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (json):  Console display name: "OpenId Parameters".


Example:
```terraform
resource "davinci_connection" "microsoftIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "microsoftIdpConnector"
  name         = "My awesome microsoftIdpConnector"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Microsoft Teams

Connector ID (`connector_id` in the resource): `microsoftTeamsConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "microsoftTeamsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "microsoftTeamsConnector"
  name         = "My awesome microsoftTeamsConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## NuData Security

Connector ID (`connector_id` in the resource): `nudataConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "nudataConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "nudataConnector"
  name         = "My awesome nudataConnector"
}
```


## Nuance

Connector ID (`connector_id` in the resource): `nuanceConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authDescription` (string):  Console display name: "Authentication Description".
* `configSetName` (string): The Config Set Name for accessing Nuance API. Console display name: "Config Set Name".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `passphrase1` (string): Passphrase that the user will need to speak for voice sample. Console display name: "Passphrase One".
* `passphrase2` (string): Passphrase that the user will need to speak for voice sample. Console display name: "Passphrase Two".
* `passphrase3` (string): Passphrase that the user will need to speak for voice sample. Console display name: "Passphrase Three".
* `passphrase4` (string): Passphrase that the user will need to speak for voice sample. Console display name: "Passphrase Four".
* `passphrase5` (string): Passphrase that the user will need to speak for voice sample. Console display name: "Passphrase Five".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```terraform
resource "davinci_connection" "nuanceConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "nuanceConnector"
  name         = "My awesome nuanceConnector"

  property {
    name  = "authDescription"
    type  = "string"
    value = var.nuanceconnector_property_auth_description
  }

  property {
    name  = "configSetName"
    type  = "string"
    value = var.nuanceconnector_property_config_set_name
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.nuanceconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.nuanceconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.nuanceconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.nuanceconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.nuanceconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.nuanceconnector_property_icon_url_png
  }

  property {
    name  = "passphrase1"
    type  = "string"
    value = var.nuanceconnector_property_passphrase1
  }

  property {
    name  = "passphrase2"
    type  = "string"
    value = var.nuanceconnector_property_passphrase2
  }

  property {
    name  = "passphrase3"
    type  = "string"
    value = var.nuanceconnector_property_passphrase3
  }

  property {
    name  = "passphrase4"
    type  = "string"
    value = var.nuanceconnector_property_passphrase4
  }

  property {
    name  = "passphrase5"
    type  = "string"
    value = var.nuanceconnector_property_passphrase5
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.nuanceconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.nuanceconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.nuanceconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.nuanceconnector_property_tool_tip
  }
}
```


## OIDC & OAuth IdP

Connector ID (`connector_id` in the resource): `genericConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "genericConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "genericConnector"
  name         = "My awesome genericConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## OPSWAT MetaAccess

Connector ID (`connector_id` in the resource): `connectorOpswat`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientID` (string): Oauth client key for authenticating API calls with MetaAccess. Console display name: "Oauth Client Key".
* `clientSecret` (string): Oauth client secret for authenticating API calls with MetaAccess. Console display name: "Oauth Client Secret".
* `crossDomainApiPort` (string): MetaAccess Cross-Domain API integration port. Console display name: "Cross-Domain API Port".
* `maDomain` (string): MetaAccess domain for your environment. Console display name: "MetaAccess Domain".


Example:
```terraform
resource "davinci_connection" "connectorOpswat" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorOpswat"
  name         = "My awesome connectorOpswat"

  property {
    name  = "clientID"
    type  = "string"
    value = var.connectoropswat_property_client_i_d
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectoropswat_property_client_secret
  }

  property {
    name  = "crossDomainApiPort"
    type  = "string"
    value = var.connectoropswat_property_cross_domain_api_port
  }

  property {
    name  = "maDomain"
    type  = "string"
    value = var.connectoropswat_property_ma_domain
  }
}
```


## OneTrust

Connector ID (`connector_id` in the resource): `oneTrustConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): Your OneTrust application client ID. Console display name: "Client ID".
* `clientSecret` (string): Your OneTrust application client secret. Console display name: "Client Secret".


Example:
```terraform
resource "davinci_connection" "oneTrustConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "oneTrustConnector"
  name         = "My awesome oneTrustConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.onetrustconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.onetrustconnector_property_client_secret
  }
}
```


## Onfido

Connector ID (`connector_id` in the resource): `onfidoConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `androidPackageName` (string): Your Android Application's Package Name. Console display name: "Android Application Package Name".
* `apiKey` (string):  Console display name: "API Key".
* `authDescription` (string):  Console display name: "Authentication Description".
* `baseUrl` (string):  Console display name: "Base URL".
* `connectorName` (string):  Console display name: "Connector Name".
* `customizeSteps` (boolean):  Console display name: "Customize Steps".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iOSBundleId` (string): Your iOS Application's Bundle ID. Console display name: "iOS Application Bundle ID".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `javascriptCSSUrl` (string):  Console display name: "CSS URL".
* `javascriptCdnUrl` (string):  Console display name: "Javascript CDN URL".
* `language` (string):  Console display name: "Language".
* `referenceStepsList` (json):  
* `referrerUrl` (string):  Console display name: "Referrer URL".
* `retrieveReports` (boolean):  Console display name: "Retrieve Reports".
* `shouldCloseOnOverlayClick` (boolean):  Console display name: "Close on Overlay Click".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `stepsList` (boolean): The Proof of Address document capture is currently a BETA feature, and it cannot be used in conjunction with the document and face steps as part of a single SDK flow. Console display name: "ID Verification Steps".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".
* `useLanguage` (boolean):  Console display name: "Customize Language".
* `useModal` (boolean):  Console display name: "Modal".
* `viewDescriptions` (string):  Console display name: "OnFido Description".
* `viewTitle` (string):  Console display name: "OnFido Title".


Example:
```terraform
resource "davinci_connection" "onfidoConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "onfidoConnector"
  name         = "My awesome onfidoConnector"

  property {
    name  = "androidPackageName"
    type  = "string"
    value = var.onfidoconnector_property_android_package_name
  }

  property {
    name  = "apiKey"
    type  = "string"
    value = var.onfidoconnector_property_api_key
  }

  property {
    name  = "authDescription"
    type  = "string"
    value = var.onfidoconnector_property_auth_description
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.onfidoconnector_property_base_url
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.onfidoconnector_property_connector_name
  }

  property {
    name  = "customizeSteps"
    type  = "boolean"
    value = var.onfidoconnector_property_customize_steps
  }

  property {
    name  = "description"
    type  = "string"
    value = var.onfidoconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.onfidoconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.onfidoconnector_property_details2
  }

  property {
    name  = "iOSBundleId"
    type  = "string"
    value = var.onfidoconnector_property_i_o_s_bundle_id
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.onfidoconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.onfidoconnector_property_icon_url_png
  }

  property {
    name  = "javascriptCSSUrl"
    type  = "string"
    value = var.javascript_css_url
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.onfidoconnector_property_javascript_cdn_url
  }

  property {
    name  = "language"
    type  = "string"
    value = var.onfidoconnector_property_language
  }

  property {
    name  = "referenceStepsList"
    type  = "json"
    value = var.onfidoconnector_property_reference_steps_list
  }

  property {
    name  = "referrerUrl"
    type  = "string"
    value = var.onfidoconnector_property_referrer_url
  }

  property {
    name  = "retrieveReports"
    type  = "boolean"
    value = var.onfidoconnector_property_retrieve_reports
  }

  property {
    name  = "shouldCloseOnOverlayClick"
    type  = "boolean"
    value = var.onfidoconnector_property_should_close_on_overlay_click
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.onfidoconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.onfidoconnector_property_show_cred_added_via
  }

  property {
    name  = "stepsList"
    type  = "boolean"
    value = var.onfidoconnector_property_steps_list
  }

  property {
    name  = "title"
    type  = "string"
    value = var.onfidoconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.onfidoconnector_property_tool_tip
  }

  property {
    name  = "useLanguage"
    type  = "boolean"
    value = var.onfidoconnector_property_use_language
  }

  property {
    name  = "useModal"
    type  = "boolean"
    value = var.onfidoconnector_property_use_modal
  }

  property {
    name  = "viewDescriptions"
    type  = "string"
    value = var.onfidoconnector_property_view_descriptions
  }

  property {
    name  = "viewTitle"
    type  = "string"
    value = var.onfidoconnector_property_view_title
  }
}
```


## PaloAlto Prisma Connector

Connector ID (`connector_id` in the resource): `connectorPaloAltoPrisma`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseURL` (string): Prisma Base URL. Console display name: "Prisma Base URL".
* `prismaPassword` (string): Secret Key. Console display name: "Prisma - Secret Key".
* `prismaUsername` (string): Access Key. Console display name: "Prisma - Access Key".


Example:
```terraform
resource "davinci_connection" "connectorPaloAltoPrisma" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorPaloAltoPrisma"
  name         = "My awesome connectorPaloAltoPrisma"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "prismaPassword"
    type  = "string"
    value = var.connectorpaloaltoprisma_property_prisma_password
  }

  property {
    name  = "prismaUsername"
    type  = "string"
    value = var.connectorpaloaltoprisma_property_prisma_username
  }
}
```


## PingAccess Administration

Connector ID (`connector_id` in the resource): `connector-oai-pingaccessadministrativeapi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authPassword` (string): The password for an account that has access to the PingAccess administrative API. Console display name: "Authenticating Password".
* `authUsername` (string): The username for an account that has access to the PingAccess administrative API. Console display name: "Authenticating Username".
* `basePath` (string): The base URL for the PingAccess Administrative API, such as "https://localhost:9000/pa-admin-api/v3". Console display name: "API URL".
* `sslVerification` (string): When enabled, DaVinci verifies the PingAccess SSL certificate and uses encrypted communication. Console display name: "Use SSL Verification".


Example:
```terraform
resource "davinci_connection" "connector-oai-pingaccessadministrativeapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-pingaccessadministrativeapi"
  name         = "My awesome connector-oai-pingaccessadministrativeapi"

  property {
    name  = "authPassword"
    type  = "string"
    value = var.connector-oai-pingaccessadministrativeapi_property_auth_password
  }

  property {
    name  = "authUsername"
    type  = "string"
    value = var.connector-oai-pingaccessadministrativeapi_property_auth_username
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-pingaccessadministrativeapi_property_base_path
  }

  property {
    name  = "sslVerification"
    type  = "string"
    value = var.connector-oai-pingaccessadministrativeapi_property_ssl_verification
  }
}
```


## PingFederate

Connector ID (`connector_id` in the resource): `pingFederateConnectorV2`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (json):  Console display name: "OpenId Parameters".


Example:
```terraform
resource "davinci_connection" "pingFederateConnectorV2" {
  environment_id = var.pingone_environment_id

  connector_id = "pingFederateConnectorV2"
  name         = "My awesome pingFederateConnectorV2"

  property {
    name = "openId"
    type = "json"
    value = jsonencode({
      "properties" : {
        "skRedirectUri" : {
          "type" : "string",
          "displayName" : "Redirect URL",
          "info" : "Enter this in your identity provider configuration to allow it to redirect the browser back to DaVinci. If you use a custom PingOne domain, modify the URL accordingly.",
          "preferredControlType" : "textField",
          "disabled" : true,
          "initializeValue" : "SINGULARKEY_REDIRECT_URI",
          "copyToClip" : true
        },
        "clientId" : {
          "type" : "string",
          "displayName" : "Client ID",
          "placeholder" : "",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "${var.pingfederateconnectorv2_property_client_id}"
        },
        "clientSecret" : {
          "type" : "string",
          "displayName" : "Client Secret",
          "preferredControlType" : "textField",
          "secure" : true,
          "required" : true,
          "value" : "${var.pingfederateconnectorv2_property_client_secret}"
        },
        "scope" : {
          "type" : "string",
          "displayName" : "Scope",
          "preferredControlType" : "textField",
          "requiredValue" : "openid",
          "value" : "${var.pingfederateconnectorv2_property_client_scope}",
          "required" : true
        },
        "issuerUrl" : {
          "type" : "string",
          "displayName" : "Base URL",
          "preferredControlType" : "textField",
          "value" : "${var.pingfederateconnectorv2_property_base_url}",
          "required" : true
        },
        "returnToUrl" : {
          "displayName" : "Application Return To URL",
          "preferredControlType" : "textField",
          "info" : "When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application.",
          "value" : "${var.pingfederateconnectorv2_property_application_callback}"
        }
      }
    })
  }
}
```


## PingFederate Administration

Connector ID (`connector_id` in the resource): `connector-oai-pfadminapi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authPassword` (string): The password for an account that has access to the PingFederate administrative API. Console display name: "Authenticating Password".
* `authUsername` (string): The username for an account that has access to the PingFederate administrative API. Console display name: "Authenticating Username".
* `basePath` (string): The base URL for the PingFederate administrative API, such as "https://8.8.4.4:9999/pf-admin-api/v1". Console display name: "API URL".
* `sslVerification` (string): When enabled, DaVinci verifies the PingFederate SSL certificate and uses encrypted communication. Console display name: "Use SSL Verification".


Example:
```terraform
resource "davinci_connection" "connector-oai-pfadminapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-pfadminapi"
  name         = "My awesome connector-oai-pfadminapi"

  property {
    name  = "authPassword"
    type  = "string"
    value = var.connector-oai-pfadminapi_property_auth_password
  }

  property {
    name  = "authUsername"
    type  = "string"
    value = var.connector-oai-pfadminapi_property_auth_username
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-pfadminapi_property_base_path
  }

  property {
    name  = "sslVerification"
    type  = "string"
    value = var.connector-oai-pfadminapi_property_ssl_verification
  }
}
```


## PingID

Connector ID (`connector_id` in the resource): `pingIdConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "pingIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingIdConnector"
  name         = "My awesome pingIdConnector"

  property {
    name = "customAuth"
    type = "json"
    value = jsonencode({
      "properties" : {
        "pingIdProperties" : {
          "displayName" : "PingID properties file",
          "preferredControlType" : "secureTextArea",
          "hashedVisibility" : true,
          "required" : true,
          "info" : "Paste the contents of the PingID properties file into this field.",
          "value" : "${file(var.pingidconnector_property_pingid_properties_file_path)}"
        },
        "returnToUrl" : {
          "displayName" : "Application Return To URL",
          "preferredControlType" : "textField",
          "info" : "When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application."
        }
      }
    })
  }
}
```


## PingOne

Connector ID (`connector_id` in the resource): `pingOneSSOConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne environment ID. Console display name: "Environment ID".
* `envRegionInfo` (string): If you want to connect with a different PingOne environment, enter the environment and credential information below. Console display name: "The default PingOne environment is configured automatically.".
* `region` (string): The region in which your PingOne environment exists. Console display name: "Region".


Example:
```terraform
resource "davinci_connection" "pingOneSSOConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneSSOConnector"
  name         = "My awesome pingOneSSOConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingone_worker_app_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingone_worker_app_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.pingone_worker_app_environment_id
  }

  property {
    name  = "envRegionInfo"
    type  = "string"
    value = var.pingonessoconnector_property_env_region_info
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingonessoconnector_property_region
  }
}
```


## PingOne Advanced Identity Cloud Access Request

Connector ID (`connector_id` in the resource): `accessRequestConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseURL` (string): The API URL to target. Console display name: "Identity Cloud Base URL".
* `endUserClientId` (string): The Client ID from the end user account. Console display name: "End User Client ID".
* `endUserClientPrivateKey` (string): The Client Private Key from the end user account. Console display name: "End User Client Private Key".
* `realm` (string): The Realm configured in Identity Cloud. Console display name: "Realm".
* `serviceAccountId` (string): The account ID for your Identity Cloud service account. You can find this ID under the account settings of your service account. Console display name: "Service Account ID".
* `serviceAccountPrivateKey` (string): The private key for your Identity Cloud service account. You can find this private key under the account settings of your service account. Console display name: "Service Account Private Key".


Example:
```terraform
resource "davinci_connection" "accessRequestConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "accessRequestConnector"
  name         = "My awesome accessRequestConnector"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.accessrequestconnector_property_base_u_r_l
  }

  property {
    name  = "endUserClientId"
    type  = "string"
    value = var.accessrequestconnector_property_end_user_client_id
  }

  property {
    name  = "endUserClientPrivateKey"
    type  = "string"
    value = var.accessrequestconnector_property_end_user_client_private_key
  }

  property {
    name  = "realm"
    type  = "string"
    value = var.accessrequestconnector_property_realm
  }

  property {
    name  = "serviceAccountId"
    type  = "string"
    value = var.accessrequestconnector_property_service_account_id
  }

  property {
    name  = "serviceAccountPrivateKey"
    type  = "string"
    value = var.accessrequestconnector_property_service_account_private_key
  }
}
```


## PingOne Advanced Identity Cloud Login Connector

Connector ID (`connector_id` in the resource): `pingoneAdvancedIdentityCloudLoginConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (json):  Console display name: "OpenId Parameters".


Example:
```terraform
resource "davinci_connection" "pingoneAdvancedIdentityCloudLoginConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingoneAdvancedIdentityCloudLoginConnector"
  name         = "My awesome pingoneAdvancedIdentityCloudLoginConnector"

  property {
    name  = "openId"
    type  = "json"
    value = var.pingoneadvancedidentitycloudloginconnector_property_open_id
  }
}
```


## PingOne Authentication

Connector ID (`connector_id` in the resource): `pingOneAuthenticationConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "pingOneAuthenticationConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneAuthenticationConnector"
  name         = "My awesome pingOneAuthenticationConnector"
}
```


## PingOne Authorize

Connector ID (`connector_id` in the resource): `pingOneAuthorizeConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of the PingOne worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of the PingOne worker application. Console display name: "Client Secret".
* `endpointURL` (string): The PingOne Authorize decision endpoint or ID to which the connector submits decision requests. Console display name: "Endpoint".


Example:
```terraform
resource "davinci_connection" "pingOneAuthorizeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneAuthorizeConnector"
  name         = "My awesome pingOneAuthorizeConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingoneauthorizeconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingoneauthorizeconnector_property_client_secret
  }

  property {
    name  = "endpointURL"
    type  = "string"
    value = var.endpoint_url
  }
}
```


## PingOne Authorize - API Access Management

Connector ID (`connector_id` in the resource): `pingauthadapter`

*No properties*


Example:
```terraform
resource "davinci_connection" "pingauthadapter" {
  environment_id = var.pingone_environment_id

  connector_id = "pingauthadapter"
  name         = "My awesome pingauthadapter"
}
```


## PingOne Credentials

Connector ID (`connector_id` in the resource): `pingOneCredentialsConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `digitalWalletApplicationId` (string): Identifier (UUID) associated with the credential digital wallet app. Console display name: "Digital Wallet Application ID".
* `envId` (string): Your PingOne Environment ID. Console display name: "Environment ID".
* `region` (string): The region your PingOne environment is in. Console display name: "Region".


Example:
```terraform
resource "davinci_connection" "pingOneCredentialsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneCredentialsConnector"
  name         = "My awesome pingOneCredentialsConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingone_worker_app_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingone_worker_app_client_secret
  }

  property {
    name  = "digitalWalletApplicationId"
    type  = "string"
    value = var.pingonecredentialsconnector_property_digital_wallet_application_id
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.pingone_worker_app_environment_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingonecredentialsconnector_property_region
  }
}
```


## PingOne Forms

Connector ID (`connector_id` in the resource): `pingOneFormsConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "pingOneFormsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneFormsConnector"
  name         = "My awesome pingOneFormsConnector"
}
```


## PingOne MFA

Connector ID (`connector_id` in the resource): `pingOneMfaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne Environment ID. Console display name: "Environment ID".
* `policyId` (string): The ID of your PingOne MFA device authentication policy. Console display name: "Policy ID".
* `region` (string): The region in which your PingOne environment exists. Console display name: "Region".


Example:
```terraform
resource "davinci_connection" "pingOneMfaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneMfaConnector"
  name         = "My awesome pingOneMfaConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingone_worker_app_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingone_worker_app_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.pingone_worker_app_environment_id
  }

  property {
    name  = "policyId"
    type  = "string"
    value = var.pingonemfaconnector_property_policy_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingonemfaconnector_property_region
  }
}
```


## PingOne Notifications

Connector ID (`connector_id` in the resource): `notificationsConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne Environment ID. Console display name: "Environment ID".
* `notificationPolicyId` (string): A unique identifier for the policy. Console display name: "Notification Policy ID".
* `region` (string): The region in which your PingOne environment exists. Console display name: "Region".


Example:
```terraform
resource "davinci_connection" "notificationsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "notificationsConnector"
  name         = "My awesome notificationsConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.notificationsconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.notificationsconnector_property_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.notificationsconnector_property_env_id
  }

  property {
    name  = "notificationPolicyId"
    type  = "string"
    value = var.notificationsconnector_property_notification_policy_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.notificationsconnector_property_region
  }
}
```


## PingOne Protect

Connector ID (`connector_id` in the resource): `pingOneRiskConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The id for your Application found in Ping's Dashboard. Console display name: "Client ID".
* `clientSecret` (string): Client Secret from your App in Ping's Dashboard. Console display name: "Client Secret".
* `envId` (string): Your Environment ID provided by Ping. Console display name: "Environment ID".
* `region` (string): The region your PingOne environment is in. Console display name: "Region".


Example:
```terraform
resource "davinci_connection" "pingOneRiskConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneRiskConnector"
  name         = "My awesome pingOneRiskConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingone_worker_app_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingone_worker_app_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.pingone_worker_app_environment_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingoneriskconnector_property_region
  }
}
```


## PingOne RADIUS Gateway

Connector ID (`connector_id` in the resource): `pingOneIntegrationsConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "pingOneIntegrationsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneIntegrationsConnector"
  name         = "My awesome pingOneIntegrationsConnector"
}
```


## PingOne Scope Consent

Connector ID (`connector_id` in the resource): `pingOneScopeConsentConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne Environment ID. Console display name: "Environment ID".
* `region` (string): The region in which your PingOne environment exists. Console display name: "Region".


Example:
```terraform
resource "davinci_connection" "pingOneScopeConsentConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneScopeConsentConnector"
  name         = "My awesome pingOneScopeConsentConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingone_worker_app_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingone_worker_app_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.pingone_worker_app_environment_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingonescopeconsentconnector_property_region
  }
}
```


## PingOne Verify

Connector ID (`connector_id` in the resource): `pingOneVerifyConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): The Client ID of your PingOne Worker application. Console display name: "Client ID".
* `clientSecret` (string): The Client Secret of your PingOne Worker application. Console display name: "Client Secret".
* `envId` (string): Your PingOne Environment ID. Console display name: "Environment ID".
* `region` (string): The region your PingOne environment is in. Console display name: "Region".


Example:
```terraform
resource "davinci_connection" "pingOneVerifyConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneVerifyConnector"
  name         = "My awesome pingOneVerifyConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingone_worker_app_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingone_worker_app_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.pingone_worker_app_environment_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingoneverifyconnector_property_region
  }
}
```


## Private ID

Connector ID (`connector_id` in the resource): `privateidConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "privateidConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "privateidConnector"
  name         = "My awesome privateidConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.privateidconnector_property_custom_auth
  }
}
```


## Prove

Connector ID (`connector_id` in the resource): `payfoneConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `appClientId` (string):  Console display name: "App Client ID".
* `baseUrl` (string):  Console display name: "Prove Base URL".
* `clientId` (string):  Console display name: "Client ID".
* `password` (string):  Console display name: "Password".
* `simulatorMode` (boolean):  Console display name: "Simulator Mode?".
* `simulatorPhoneNumber` (string):  Console display name: "Simulator Phone Number".
* `skCallbackBaseUrl` (string): Use this url as the callback base URL. Console display name: "Callback Base URL".
* `username` (string):  Console display name: "Username".


Example:
```terraform
resource "davinci_connection" "payfoneConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "payfoneConnector"
  name         = "My awesome payfoneConnector"

  property {
    name  = "appClientId"
    type  = "string"
    value = var.payfoneconnector_property_app_client_id
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.payfoneconnector_property_base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.payfoneconnector_property_client_id
  }

  property {
    name  = "password"
    type  = "string"
    value = var.payfoneconnector_property_password
  }

  property {
    name  = "simulatorMode"
    type  = "boolean"
    value = var.payfoneconnector_property_simulator_mode
  }

  property {
    name  = "simulatorPhoneNumber"
    type  = "string"
    value = var.payfoneconnector_property_simulator_phone_number
  }

  property {
    name  = "skCallbackBaseUrl"
    type  = "string"
    value = var.payfoneconnector_property_sk_callback_base_url
  }

  property {
    name  = "username"
    type  = "string"
    value = var.payfoneconnector_property_username
  }
}
```


## Prove International

Connector ID (`connector_id` in the resource): `proveConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseUrl` (string):  Console display name: "Prove Base URL".
* `clientId` (string):  Console display name: "Prove Client ID".
* `grantType` (string):  Console display name: "Prove Grant Type".
* `password` (string):  Console display name: "Prove Password".
* `username` (string):  Console display name: "Prove Username".


Example:
```terraform
resource "davinci_connection" "proveConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "proveConnector"
  name         = "My awesome proveConnector"

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.proveconnector_property_base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.proveconnector_property_client_id
  }

  property {
    name  = "grantType"
    type  = "string"
    value = var.proveconnector_property_grant_type
  }

  property {
    name  = "password"
    type  = "string"
    value = var.proveconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.proveconnector_property_username
  }
}
```


## RSA

Connector ID (`connector_id` in the resource): `rsaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessId` (string): RSA Access ID from Administration API key file. Console display name: "Access ID".
* `accessKey` (string): RSA Access Key from Administration API key file. Console display name: "Access Key".
* `baseUrl` (string): Base URL for RSA API that is provided in Administration API key file. Console display name: "Base URL".


Example:
```terraform
resource "davinci_connection" "rsaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "rsaConnector"
  name         = "My awesome rsaConnector"

  property {
    name  = "accessId"
    type  = "string"
    value = var.rsaconnector_property_access_id
  }

  property {
    name  = "accessKey"
    type  = "string"
    value = var.rsaconnector_property_access_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.rsaconnector_property_base_url
  }
}
```


## ReadID by Inverid

Connector ID (`connector_id` in the resource): `inveridConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `getApiKey` (string): Viewer API Key provided to you by Inverid. Console display name: "ReadID Viewer API Key".
* `host` (string): Hostname provided to you by Inverid. Console display name: "ReadID Hostname".
* `postApiKey` (string): Submitter API Key provided to you by Inverid. Console display name: "ReadID Submitter API Key".
* `skWebhookUri` (string): Use this url as the Webhook URL in the Third Party Integration's configuration. Console display name: "Redirect Webhook URI".
* `timeToLive` (string): Specify the duration (in minutes) a users session should stay active. Value must be between 30 and 72000. Console display name: "Time to live for ReadySession".


Example:
```terraform
resource "davinci_connection" "inveridConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "inveridConnector"
  name         = "My awesome inveridConnector"

  property {
    name  = "getApiKey"
    type  = "string"
    value = var.inveridconnector_property_get_api_key
  }

  property {
    name  = "host"
    type  = "string"
    value = var.inveridconnector_property_host
  }

  property {
    name  = "postApiKey"
    type  = "string"
    value = var.inveridconnector_property_post_api_key
  }

  property {
    name  = "skWebhookUri"
    type  = "string"
    value = var.inveridconnector_property_sk_webhook_uri
  }

  property {
    name  = "timeToLive"
    type  = "string"
    value = var.inveridconnector_property_time_to_live
  }
}
```


## SAML

Connector ID (`connector_id` in the resource): `samlConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "samlConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "samlConnector"
  name         = "My awesome samlConnector"
}
```


## SAML IdP

Connector ID (`connector_id` in the resource): `samlIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `saml` (json):  Console display name: "SAML Parameters".


Example:
```terraform
resource "davinci_connection" "samlIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "samlIdpConnector"
  name         = "My awesome samlIdpConnector"

  property {
    name  = "saml"
    type  = "json"
    value = jsonencode({})
  }
}
```


## SAP Identity API

Connector ID (`connector_id` in the resource): `connector-oai-sapidentityapis`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authApiKey` (string): The authentication key to the SAP Identity APIs. Console display name: "API Key".
* `basePath` (string): The base URL for contacting the API. Console display name: "Base Path".


Example:
```terraform
resource "davinci_connection" "connector-oai-sapidentityapis" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-sapidentityapis"
  name         = "My awesome connector-oai-sapidentityapis"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-sapidentityapis_property_auth_api_key
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-sapidentityapis_property_base_path
  }
}
```


## SEON

Connector ID (`connector_id` in the resource): `seonConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseURL` (string): The API URL to target. Console display name: "API Base URL".
* `licenseKey` (string): Your SEON license key. For help, see the SEON REST API documentation. Console display name: "License Key".


Example:
```terraform
resource "davinci_connection" "seonConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "seonConnector"
  name         = "My awesome seonConnector"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "licenseKey"
    type  = "string"
    value = var.seonconnector_property_license_key
  }
}
```


## SMTP Client

Connector ID (`connector_id` in the resource): `smtpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `hostname` (string): Example: smtp-relay.gmail.com. Console display name: "SMTP Server/Host".
* `name` (string): Optional hostname of the client, used for identifying to the server, defaults to hostname of the machine. Console display name: "Client Name".
* `password` (string):  Console display name: "Password".
* `port` (number): Example: 25. Console display name: "SMTP Port".
* `secureFlag` (boolean):  Console display name: "Secure Flag?".
* `username` (string):  Console display name: "Username".


Example:
```terraform
resource "davinci_connection" "smtpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "smtpConnector"
  name         = "My awesome smtpConnector"

  property {
    name  = "hostname"
    type  = "string"
    value = var.smtpconnector_property_hostname
  }

  property {
    name  = "name"
    type  = "string"
    value = var.smtpconnector_property_name
  }

  property {
    name  = "password"
    type  = "string"
    value = var.smtpconnector_property_password
  }

  property {
    name  = "port"
    type  = "number"
    value = var.smtpconnector_property_port
  }

  property {
    name  = "secureFlag"
    type  = "boolean"
    value = var.smtpconnector_property_secure_flag
  }

  property {
    name  = "username"
    type  = "string"
    value = var.smtpconnector_property_username
  }
}
```


## SailPoint IdentityNow

Connector ID (`connector_id` in the resource): `connectorIdentityNow`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientId` (string): Client Id for your client found in IdentityNow's Dashboard. Console display name: "Client ID".
* `clientSecret` (string): Client Secret from your client in IdentityNow's Dashboard. Console display name: "Client Secret".
* `tenant` (string): The org name is displayed within the Org Details section of the dashboard. Console display name: "IdentityNow Tenant".


Example:
```terraform
resource "davinci_connection" "connectorIdentityNow" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIdentityNow"
  name         = "My awesome connectorIdentityNow"

  property {
    name  = "clientId"
    type  = "string"
    value = var.connectoridentitynow_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectoridentitynow_property_client_secret
  }

  property {
    name  = "tenant"
    type  = "string"
    value = var.connectoridentitynow_property_tenant
  }
}
```


## Salesforce

Connector ID (`connector_id` in the resource): `salesforceConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `adminUsername` (string): The username of your Salesforce administrator account. Console display name: "Username".
* `consumerKey` (string): The consumer key shown on your Salesforce connected app. Console display name: "Consumer Key".
* `domainName` (string): Your Salesforce domain name, such as "mycompany-dev-ed". Console display name: "Domain Name".
* `environment` (string): If the environment you specify in the Domain Name field is part of a sandbox organization, select Sandbox. Otherwise, select Production. Console display name: "Environment".
* `privateKey` (string): The private key that corresponds to the X.509 certificate you added to your Salesforce connected app. Console display name: "Private Key".


Example:
```terraform
resource "davinci_connection" "salesforceConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "salesforceConnector"
  name         = "My awesome salesforceConnector"

  property {
    name  = "adminUsername"
    type  = "string"
    value = var.salesforceconnector_property_admin_username
  }

  property {
    name  = "consumerKey"
    type  = "string"
    value = var.salesforceconnector_property_consumer_key
  }

  property {
    name  = "domainName"
    type  = "string"
    value = var.salesforceconnector_property_domain_name
  }

  property {
    name  = "environment"
    type  = "string"
    value = var.salesforceconnector_property_environment
  }

  property {
    name  = "privateKey"
    type  = "string"
    value = var.salesforceconnector_property_private_key
  }
}
```


## Salesforce Marketing Cloud (BETA)

Connector ID (`connector_id` in the resource): `connectorSalesforceMarketingCloud`

Properties (used in the `property` block in the resource as the `name` parameter):

* `SalesforceMarketingCloudURL` (string): URL for Salesforce Marketing Cloud. Example: https://YOUR_SUBDOMAIN.rest.marketingcloudapis.com. Console display name: "Salesforce Marketing Cloud URL".
* `accountId` (string): Account identifier, or MID, of the target business unit. Use to switch between business units. If you don’t specify account_id, the returned access token is in the context of the business unit that created the integration. Console display name: "Account ID".
* `clientId` (string): Client ID issued when you create the API integration in Installed Packages. Console display name: "Client ID".
* `clientSecret` (string): Client secret issued when you create the API integration in Installed Packages. Console display name: "Client Secret".
* `scope` (string): Space-separated list of data-access permissions for your application. Console display name: "Scope".


Example:
```terraform
resource "davinci_connection" "connectorSalesforceMarketingCloud" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSalesforceMarketingCloud"
  name         = "My awesome connectorSalesforceMarketingCloud"

  property {
    name  = "SalesforceMarketingCloudURL"
    type  = "string"
    value = var.salesforce_marketing_cloud_url
  }

  property {
    name  = "accountId"
    type  = "string"
    value = var.connectorsalesforcemarketingcloud_property_account_id
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.connectorsalesforcemarketingcloud_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectorsalesforcemarketingcloud_property_client_secret
  }

  property {
    name  = "scope"
    type  = "string"
    value = var.connectorsalesforcemarketingcloud_property_scope
  }
}
```


## Saviynt Connector Flows

Connector ID (`connector_id` in the resource): `connectorSaviyntFlow`

Properties (used in the `property` block in the resource as the `name` parameter):

* `domainName` (string): Provide your Saviynt domain name. Console display name: "Saviynt Domain Name".
* `path` (string): Provide your Saviynt path name. Console display name: "Saviynt Path Name".
* `saviyntPassword` (string): Provide your Saviynt password. Console display name: "Saviynt Password".
* `saviyntUserName` (string): Provide your Saviynt user name. Console display name: "Saviynt User Name".


Example:
```terraform
resource "davinci_connection" "connectorSaviyntFlow" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSaviyntFlow"
  name         = "My awesome connectorSaviyntFlow"

  property {
    name  = "domainName"
    type  = "string"
    value = var.connectorsaviyntflow_property_domain_name
  }

  property {
    name  = "path"
    type  = "string"
    value = var.connectorsaviyntflow_property_path
  }

  property {
    name  = "saviyntPassword"
    type  = "string"
    value = var.connectorsaviyntflow_property_saviynt_password
  }

  property {
    name  = "saviyntUserName"
    type  = "string"
    value = var.connectorsaviyntflow_property_saviynt_user_name
  }
}
```


## Screen

Connector ID (`connector_id` in the resource): `screenConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "screenConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "screenConnector"
  name         = "My awesome screenConnector"
}
```


## SecurID

Connector ID (`connector_id` in the resource): `securIdConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): The URL of your SecurID authentication API, such as "https://company.auth.securid.com". Console display name: "SecurID Authentication API REST URL".
* `clientKey` (string): Your SecurID authentication client key, such as "vowc450ahs6nry66vok0pvaizwnfr43ewsqcm7tz". Console display name: "Client Key".


Example:
```terraform
resource "davinci_connection" "securIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "securIdConnector"
  name         = "My awesome securIdConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.securidconnector_property_api_url
  }

  property {
    name  = "clientKey"
    type  = "string"
    value = var.securidconnector_property_client_key
  }
}
```


## Securonix

Connector ID (`connector_id` in the resource): `connectorSecuronix`

Properties (used in the `property` block in the resource as the `name` parameter):

* `domainName` (string): Domain Name. Console display name: "Domain Name".
* `token` (string): Token for authentication. Console display name: "Token".


Example:
```terraform
resource "davinci_connection" "connectorSecuronix" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSecuronix"
  name         = "My awesome connectorSecuronix"

  property {
    name  = "domainName"
    type  = "string"
    value = var.connectorsecuronix_property_domain_name
  }

  property {
    name  = "token"
    type  = "string"
    value = var.connectorsecuronix_property_token
  }
}
```


## Segment

Connector ID (`connector_id` in the resource): `connectorSegment`

Properties (used in the `property` block in the resource as the `name` parameter):

* `version` (string): Segment - HTTP Tracking API Version. Console display name: "HTTP Tracking API Version".
* `writeKey` (string): The Write Key is used to send data to a specific workplace. Console display name: "Write Key".


Example:
```terraform
resource "davinci_connection" "connectorSegment" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSegment"
  name         = "My awesome connectorSegment"

  property {
    name  = "version"
    type  = "string"
    value = var.connectorsegment_property_version
  }

  property {
    name  = "writeKey"
    type  = "string"
    value = var.connectorsegment_property_write_key
  }
}
```


## SentiLink

Connector ID (`connector_id` in the resource): `sentilinkConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `account` (string): Account ID of SentiLink. Console display name: "Account ID".
* `apiUrl` (string):  Console display name: "API URL".
* `javascriptCdnUrl` (string):  Console display name: "Javascript CDN URL".
* `token` (string): Token ID for SentiLink account. Console display name: "Token ID".


Example:
```terraform
resource "davinci_connection" "sentilinkConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "sentilinkConnector"
  name         = "My awesome sentilinkConnector"

  property {
    name  = "account"
    type  = "string"
    value = var.sentilinkconnector_property_account
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.sentilinkconnector_property_api_url
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.sentilinkconnector_property_javascript_cdn_url
  }

  property {
    name  = "token"
    type  = "string"
    value = var.sentilinkconnector_property_token
  }
}
```


## ServiceNow

Connector ID (`connector_id` in the resource): `servicenowConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `adminUsername` (string): Your ServiceNow administrator username. Console display name: "Username".
* `apiUrl` (string): The API URL to target, such as "https://mycompany.service-now.com". Console display name: "API URL".
* `password` (string): Your ServiceNow administrator password. Console display name: "Password".


Example:
```terraform
resource "davinci_connection" "servicenowConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "servicenowConnector"
  name         = "My awesome servicenowConnector"

  property {
    name  = "adminUsername"
    type  = "string"
    value = var.servicenowconnector_property_admin_username
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.servicenowconnector_property_api_url
  }

  property {
    name  = "password"
    type  = "string"
    value = var.servicenowconnector_property_password
  }
}
```


## Shopify Connector

Connector ID (`connector_id` in the resource): `connectorShopify`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessToken` (string): Your store's unique Admin API Access Token that goes into the X-Shopify-Access-Token property. Required scopes when generating Admin API Access Token: 'read_customers' and 'write_customers'. Note any Custom Shopify API calls you intend to use with this connector via Make Custom API Call capability, will have to be added as well. Console display name: "Admin API Access Token".
* `apiVersion` (string): The Shopify version name ( ex. 2022-04 ). Console display name: "API Version Name".
* `multipassSecret` (string): Shopify Multipass Secret. Console display name: "Multipass Secret".
* `multipassStoreDomain` (string): Shopify Multipass Store Domain (yourstorename.myshopify.com). Console display name: "Multipass Store Domain".
* `yourStoreName` (string): The name of your store as Shopify identifies you ( first text that comes after HTTPS:// ). Console display name: "Store Name".


Example:
```terraform
resource "davinci_connection" "connectorShopify" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorShopify"
  name         = "My awesome connectorShopify"

  property {
    name  = "accessToken"
    type  = "string"
    value = var.connectorshopify_property_access_token
  }

  property {
    name  = "apiVersion"
    type  = "string"
    value = var.connectorshopify_property_api_version
  }

  property {
    name  = "multipassSecret"
    type  = "string"
    value = var.connectorshopify_property_multipass_secret
  }

  property {
    name  = "multipassStoreDomain"
    type  = "string"
    value = var.connectorshopify_property_multipass_store_domain
  }

  property {
    name  = "yourStoreName"
    type  = "string"
    value = var.connectorshopify_property_your_store_name
  }
}
```


## Sift

Connector ID (`connector_id` in the resource): `siftConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key from Sift Tenant. Console display name: "Sift API Key".


Example:
```terraform
resource "davinci_connection" "siftConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "siftConnector"
  name         = "My awesome siftConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.siftconnector_property_api_key
  }
}
```


## Signicat

Connector ID (`connector_id` in the resource): `connectorSignicat`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorSignicat" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSignicat"
  name         = "My awesome connectorSignicat"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Silverfort

Connector ID (`connector_id` in the resource): `silverfortConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Silverfort Risk API Key. Console display name: "Risk API Key".
* `appUserSecret` (string): Silverfort App User Secret. Console display name: "App User Secret".
* `consoleApi` (string): Silverfort App User ID. Console display name: "App User ID".


Example:
```terraform
resource "davinci_connection" "silverfortConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "silverfortConnector"
  name         = "My awesome silverfortConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.silverfortconnector_property_api_key
  }

  property {
    name  = "appUserSecret"
    type  = "string"
    value = var.silverfortconnector_property_app_user_secret
  }

  property {
    name  = "consoleApi"
    type  = "string"
    value = var.silverfortconnector_property_console_api
  }
}
```


## Sinch

Connector ID (`connector_id` in the resource): `sinchConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `acceptLanguage` (string): Language of SMS sent, if using Sinch provided templates will be chosen based on Accept-Language header. Examples include, but are not limited to pl-PL, no-NO, en-US. Console display name: "Language".
* `applicationKey` (string): Verification Application Key from your Sinch Account. Console display name: "Sinch Application Key".
* `secretKey` (string): Verification Secret Key from your Sinch Account. Console display name: "Sinch Secret Key".


Example:
```terraform
resource "davinci_connection" "sinchConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "sinchConnector"
  name         = "My awesome sinchConnector"

  property {
    name  = "acceptLanguage"
    type  = "string"
    value = var.sinchconnector_property_accept_language
  }

  property {
    name  = "applicationKey"
    type  = "string"
    value = var.sinchconnector_property_application_key
  }

  property {
    name  = "secretKey"
    type  = "string"
    value = var.sinchconnector_property_secret_key
  }
}
```


## Singpass Login

Connector ID (`connector_id` in the resource): `singpassLoginConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "singpassLoginConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "singpassLoginConnector"
  name         = "My awesome singpassLoginConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Slack Login

Connector ID (`connector_id` in the resource): `slackConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (json):  Console display name: "Oauth2 Parameters".


Example:
```terraform
resource "davinci_connection" "slackConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "slackConnector"
  name         = "My awesome slackConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Smarty Address Validator

Connector ID (`connector_id` in the resource): `connectorSmarty`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authId` (string): Smarty Authentication ID (Found on 'API Keys' tab in Smarty tenant). Console display name: "Auth ID".
* `authToken` (string): Smarty Authentication Token (Found on 'API Keys' tab in Smarty tenant). Console display name: "Auth Token".
* `license` (string): Smarty License Value (Found on 'Subscriptions' tab in Smarty tenant). Console display name: "License".


Example:
```terraform
resource "davinci_connection" "connectorSmarty" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSmarty"
  name         = "My awesome connectorSmarty"

  property {
    name  = "authId"
    type  = "string"
    value = var.connectorsmarty_property_auth_id
  }

  property {
    name  = "authToken"
    type  = "string"
    value = var.connectorsmarty_property_auth_token
  }

  property {
    name  = "license"
    type  = "string"
    value = var.connectorsmarty_property_license
  }
}
```


## Socure

Connector ID (`connector_id` in the resource): `socureConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): ID+ Key is the API key that you can retrieve from Socure Admin Portal. Console display name: "ID+ Key".
* `baseUrl` (string): The Socure API URL to target. For a custom value, select Use Custom API URL and enter a value in the Custom API URL field. Console display name: "API URL".
* `customApiUrl` (string): The URL for the Socure API, such as "https://example.socure.com". Console display name: "Custom API URL".
* `sdkKey` (string): SDK Key that you can retrieve from Socure Admin Portal. Console display name: "SDK Key".
* `skWebhookUri` (string): Use this url as the Webhook URL in the Third Party Integration's configuration. Console display name: "Webhook URL".


Example:
```terraform
resource "davinci_connection" "socureConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "socureConnector"
  name         = "My awesome socureConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.socureconnector_property_api_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.socureconnector_property_base_url
  }

  property {
    name  = "customApiUrl"
    type  = "string"
    value = var.socureconnector_property_custom_api_url
  }

  property {
    name  = "sdkKey"
    type  = "string"
    value = var.socureconnector_property_sdk_key
  }

  property {
    name  = "skWebhookUri"
    type  = "string"
    value = var.socureconnector_property_sk_webhook_uri
  }
}
```


## Splunk

Connector ID (`connector_id` in the resource): `splunkConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): The Base API URL for Splunk. Console display name: "Base URL".
* `port` (number): API Server Port. Console display name: "Port".
* `token` (string): Splunk Token to make API requests. Console display name: "Token".


Example:
```terraform
resource "davinci_connection" "splunkConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "splunkConnector"
  name         = "My awesome splunkConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.splunkconnector_property_api_url
  }

  property {
    name  = "port"
    type  = "number"
    value = var.splunkconnector_property_port
  }

  property {
    name  = "token"
    type  = "string"
    value = var.splunkconnector_property_token
  }
}
```


## Spotify

Connector ID (`connector_id` in the resource): `connectorSpotify`

Properties (used in the `property` block in the resource as the `name` parameter):

* `oauth2` (json):  Console display name: "Oauth2 Parameters".


Example:
```terraform
resource "davinci_connection" "connectorSpotify" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSpotify"
  name         = "My awesome connectorSpotify"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
```


## SpyCloud Enterprise Protection

Connector ID (`connector_id` in the resource): `connectorSpycloud`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Contact SpyCloud to acquire an Employee ATO Prevention API Key that will work with DaVinci. Console display name: "SpyCloud Employee ATO Prevention API Key".


Example:
```terraform
resource "davinci_connection" "connectorSpycloud" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSpycloud"
  name         = "My awesome connectorSpycloud"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorspycloud_property_api_key
  }
}
```


## String

Connector ID (`connector_id` in the resource): `stringsConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "stringsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "stringsConnector"
  name         = "My awesome stringsConnector"
}
```


## TMT Analysis

Connector ID (`connector_id` in the resource): `tmtConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key for TMT Analysis. Console display name: "API Key".
* `apiSecret` (string): API Secret for TMT Analysis. Console display name: "API Secret".
* `apiUrl` (string): The Base API URL for TMT Analysis. Console display name: "Base URL".


Example:
```terraform
resource "davinci_connection" "tmtConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "tmtConnector"
  name         = "My awesome tmtConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.tmtconnector_property_api_key
  }

  property {
    name  = "apiSecret"
    type  = "string"
    value = var.tmtconnector_property_api_secret
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.tmtconnector_property_api_url
  }
}
```


## Tableau

Connector ID (`connector_id` in the resource): `connectorTableau`

Properties (used in the `property` block in the resource as the `name` parameter):

* `addFlowPermissionsRequestBody` (string): Add Flow Permissions Request Body in XML Format. Example: <tsRequest><task><flowRun><flow id="flow-id"/><flowRunSpec><flowParameterSpecs><flowParameterSpec parameterId="parameter-id" overrideValue= "overrideValue"/><flowParameterSpecs><flowRunSpec></flowRun></task></tsRequest>. Console display name: "Add Flow Permissions Request Body in XML Format.".
* `addUsertoSiteRequestBody` (string): Add User to Site Request Body in XML Format. Example: <tsRequest><user name="user-name" siteRole="site-role" authSetting="auth-setting" /></tsRequest>. Console display name: "Add User to Site Request Body in XML Format.".
* `apiVersion` (string): The version of the API to use, such as 3.16. Console display name: "api-version".
* `authId` (string): The Tableau-Auth sent along with every request. Console display name: "auth-ID".
* `createScheduleBody` (string): This should contain the entire XML. Eg: <tsRequest><schedule name="schedule-name"priority="schedule-priority"type="schedule-type"frequency="schedule-frequency"executionOrder="schedule-execution-order"><frequencyDetails start="start-time" end="end-time"><intervals><interval interval-expression /></intervals></frequencyDetails></schedule></tsRequest>. Console display name: "XML file format to be used for creating schedule".
* `datasourceId` (string): The ID of the flow. Console display name: "datasource-id".
* `flowId` (string): The flow-id value for the flow you want to add permissions to. Console display name: "flow-id".
* `groupId` (string): The ID of the group. Console display name: "group-id".
* `jobId` (string): The ID of the job. Console display name: "job-id".
* `scheduleId` (string): The ID of the schedule that you are associating with the data source. Console display name: "schedule-id".
* `serverUrl` (string): The tableau server URL Example: https://www.tableau.com:8030. Console display name: "server-url".
* `siteId` (string): The ID of the site that contains the view. Console display name: "site-id".
* `taskId` (string): The ID of the extract refresh task. Console display name: "task-id".
* `updateScheduleRequestBody` (string): This should contain the entire XML. Eg: <tsRequest><schedule name="hourly-schedule-1" priority="50" type="Extract" frequency="Hourly" executionOrder="Parallel"><frequencyDetails start="18:30:00" end="23:00:00"><intervals><interval hours="2" /></intervals></frequencyDetails></schedule></tsRequest>. Console display name: "XML file format to be used for updating schedule".
* `updateUserRequestBody` (string): Update User Request Body in XML Format. <tsRequest><user fullName="new-full-name" email="new-email" password="new-password" siteRole="new-site-role" authSetting="new-auth-setting" /></tsRequest>. Console display name: "Update User Request Body in XML Format.".
* `userId` (string): The ID of the user to get/give information for. Console display name: "user-id".
* `workbookId` (string): The ID of the workbook to add to the schedule. Console display name: "workbook-id".


Example:
```terraform
resource "davinci_connection" "connectorTableau" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorTableau"
  name         = "My awesome connectorTableau"

  property {
    name  = "addFlowPermissionsRequestBody"
    type  = "string"
    value = var.connectortableau_property_add_flow_permissions_request_body
  }

  property {
    name  = "addUsertoSiteRequestBody"
    type  = "string"
    value = var.connectortableau_property_add_userto_site_request_body
  }

  property {
    name  = "apiVersion"
    type  = "string"
    value = var.connectortableau_property_api_version
  }

  property {
    name  = "authId"
    type  = "string"
    value = var.connectortableau_property_auth_id
  }

  property {
    name  = "createScheduleBody"
    type  = "string"
    value = var.connectortableau_property_create_schedule_body
  }

  property {
    name  = "datasourceId"
    type  = "string"
    value = var.connectortableau_property_datasource_id
  }

  property {
    name  = "flowId"
    type  = "string"
    value = var.connectortableau_property_flow_id
  }

  property {
    name  = "groupId"
    type  = "string"
    value = var.connectortableau_property_group_id
  }

  property {
    name  = "jobId"
    type  = "string"
    value = var.connectortableau_property_job_id
  }

  property {
    name  = "scheduleId"
    type  = "string"
    value = var.connectortableau_property_schedule_id
  }

  property {
    name  = "serverUrl"
    type  = "string"
    value = var.connectortableau_property_server_url
  }

  property {
    name  = "siteId"
    type  = "string"
    value = var.connectortableau_property_site_id
  }

  property {
    name  = "taskId"
    type  = "string"
    value = var.connectortableau_property_task_id
  }

  property {
    name  = "updateScheduleRequestBody"
    type  = "string"
    value = var.connectortableau_property_update_schedule_request_body
  }

  property {
    name  = "updateUserRequestBody"
    type  = "string"
    value = var.connectortableau_property_update_user_request_body
  }

  property {
    name  = "userId"
    type  = "string"
    value = var.connectortableau_property_user_id
  }

  property {
    name  = "workbookId"
    type  = "string"
    value = var.connectortableau_property_workbook_id
  }
}
```


## Talend Identities Management API

Connector ID (`connector_id` in the resource): `connector-oai-talendim`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authBearerToken` (string): The authenticating token. Console display name: "Bearer Token".
* `basePath` (string): The base URL for contacting the API. Console display name: "Base Path".


Example:
```terraform
resource "davinci_connection" "connector-oai-talendim" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-talendim"
  name         = "My awesome connector-oai-talendim"

  property {
    name  = "authBearerToken"
    type  = "string"
    value = var.connector-oai-talendim_property_auth_bearer_token
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-talendim_property_base_path
  }
}
```


## Talend SCIM API

Connector ID (`connector_id` in the resource): `connector-oai-talendscim`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authBearerToken` (string): The authenticating token. Console display name: "Bearer Token".
* `basePath` (string): The base URL for contacting the API. Console display name: "Base Path".


Example:
```terraform
resource "davinci_connection" "connector-oai-talendscim" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-talendscim"
  name         = "My awesome connector-oai-talendscim"

  property {
    name  = "authBearerToken"
    type  = "string"
    value = var.connector-oai-talendscim_property_auth_bearer_token
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-talendscim_property_base_path
  }
}
```


## Teleport

Connector ID (`connector_id` in the resource): `nodeConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "nodeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "nodeConnector"
  name         = "My awesome nodeConnector"
}
```


## Telesign

Connector ID (`connector_id` in the resource): `telesignConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authDescription` (string):  Console display name: "Authentication Description".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `password` (string):  Console display name: "Password".
* `providerName` (string):  Console display name: "Provider Name".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".
* `username` (string):  Console display name: "Username".


Example:
```terraform
resource "davinci_connection" "telesignConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "telesignConnector"
  name         = "My awesome telesignConnector"

  property {
    name  = "authDescription"
    type  = "string"
    value = var.telesignconnector_property_auth_description
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.telesignconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.telesignconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.telesignconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.telesignconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.telesignconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.telesignconnector_property_icon_url_png
  }

  property {
    name  = "password"
    type  = "string"
    value = var.telesignconnector_property_password
  }

  property {
    name  = "providerName"
    type  = "string"
    value = var.telesignconnector_property_provider_name
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.telesignconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.telesignconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.telesignconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.telesignconnector_property_tool_tip
  }

  property {
    name  = "username"
    type  = "string"
    value = var.telesignconnector_property_username
  }
}
```


## Token Management

Connector ID (`connector_id` in the resource): `skOpenIdConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "skOpenIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "skOpenIdConnector"
  name         = "My awesome skOpenIdConnector"
}
```


## TransUnion TLOxp

Connector ID (`connector_id` in the resource): `tutloxpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): The URL for your TransUnion API. Unnecessary to change unless you're testing against a demo tenant. Console display name: "API URL".
* `dppaCode` (string): The DPPA code that determines the level of data access in the API. Console display name: "DPPA Purpose Code".
* `glbCode` (string): The GLB code that determines the level of data access in the API. Console display name: "GLB Purpose Code".
* `password` (string): The password for your API User. Console display name: "Password".
* `username` (string): The username for your API user. Console display name: "Username".


Example:
```terraform
resource "davinci_connection" "tutloxpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "tutloxpConnector"
  name         = "My awesome tutloxpConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.tutloxpconnector_property_api_url
  }

  property {
    name  = "dppaCode"
    type  = "string"
    value = var.tutloxpconnector_property_dppa_code
  }

  property {
    name  = "glbCode"
    type  = "string"
    value = var.tutloxpconnector_property_glb_code
  }

  property {
    name  = "password"
    type  = "string"
    value = var.tutloxpconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.tutloxpconnector_property_username
  }
}
```


## TransUnion TruValidate

Connector ID (`connector_id` in the resource): `transunionConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string): The Base API URL for TransUnion. Console display name: "Base URL".
* `docVerificationPassword` (string): Password for Document Verification, provided by TransUnion. Console display name: "Password".
* `docVerificationPublicKey` (string): Public Key for Document Verification, provided by TransUnion. Console display name: "Public Key".
* `docVerificationSecret` (string): Secret for Document Verification, provided by TransUnion. Console display name: "Secret".
* `docVerificationSiteId` (string): Site ID for Document Verification, provided by TransUnion. Console display name: "Site ID".
* `docVerificationUsername` (string): Username for Document Verification, provided by TransUnion. Console display name: "Username".
* `idVerificationPassword` (string): Password for ID Verification, provided by TransUnion. Console display name: "Password".
* `idVerificationPublicKey` (string): Public Key for ID Verification, provided by TransUnion. Console display name: "Public Key".
* `idVerificationSecret` (string): Secret for ID Verification, provided by TransUnion. Console display name: "Secret".
* `idVerificationSiteId` (string): Site ID for ID Verification, provided by TransUnion. Console display name: "Site ID".
* `idVerificationUsername` (string): Username for ID Verification, provided by TransUnion. Console display name: "Username".
* `kbaPassword` (string): Password for KBA, provided by TransUnion. Console display name: "Password".
* `kbaPublicKey` (string): Public Key for KBA, provided by TransUnion. Console display name: "Public Key".
* `kbaSecret` (string): Secret for KBA, provided by TransUnion. Console display name: "Secret".
* `kbaSiteId` (string): Site ID for KBA, provided by TransUnion. Console display name: "Site ID".
* `kbaUsername` (string): Username for KBA, provided by TransUnion. Console display name: "Username".
* `otpPassword` (string): Password for otp Verification, provided by TransUnion. Console display name: "Password".
* `otpPublicKey` (string): Public Key for otp Verification, provided by TransUnion. Console display name: "Public Key".
* `otpSecret` (string): Secret for otp Verification, provided by TransUnion. Console display name: "Secret".
* `otpSiteId` (string): Site ID for otp Verification, provided by TransUnion. Console display name: "Site ID".
* `otpUsername` (string): Username for otp Verification, provided by TransUnion. Console display name: "Username".


Example:
```terraform
resource "davinci_connection" "transunionConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "transunionConnector"
  name         = "My awesome transunionConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.transunionconnector_property_api_url
  }

  property {
    name  = "docVerificationPassword"
    type  = "string"
    value = var.transunionconnector_property_doc_verification_password
  }

  property {
    name  = "docVerificationPublicKey"
    type  = "string"
    value = var.transunionconnector_property_doc_verification_public_key
  }

  property {
    name  = "docVerificationSecret"
    type  = "string"
    value = var.transunionconnector_property_doc_verification_secret
  }

  property {
    name  = "docVerificationSiteId"
    type  = "string"
    value = var.transunionconnector_property_doc_verification_site_id
  }

  property {
    name  = "docVerificationUsername"
    type  = "string"
    value = var.transunionconnector_property_doc_verification_username
  }

  property {
    name  = "idVerificationPassword"
    type  = "string"
    value = var.transunionconnector_property_id_verification_password
  }

  property {
    name  = "idVerificationPublicKey"
    type  = "string"
    value = var.transunionconnector_property_id_verification_public_key
  }

  property {
    name  = "idVerificationSecret"
    type  = "string"
    value = var.transunionconnector_property_id_verification_secret
  }

  property {
    name  = "idVerificationSiteId"
    type  = "string"
    value = var.transunionconnector_property_id_verification_site_id
  }

  property {
    name  = "idVerificationUsername"
    type  = "string"
    value = var.transunionconnector_property_id_verification_username
  }

  property {
    name  = "kbaPassword"
    type  = "string"
    value = var.transunionconnector_property_kba_password
  }

  property {
    name  = "kbaPublicKey"
    type  = "string"
    value = var.transunionconnector_property_kba_public_key
  }

  property {
    name  = "kbaSecret"
    type  = "string"
    value = var.transunionconnector_property_kba_secret
  }

  property {
    name  = "kbaSiteId"
    type  = "string"
    value = var.transunionconnector_property_kba_site_id
  }

  property {
    name  = "kbaUsername"
    type  = "string"
    value = var.transunionconnector_property_kba_username
  }

  property {
    name  = "otpPassword"
    type  = "string"
    value = var.transunionconnector_property_otp_password
  }

  property {
    name  = "otpPublicKey"
    type  = "string"
    value = var.transunionconnector_property_otp_public_key
  }

  property {
    name  = "otpSecret"
    type  = "string"
    value = var.transunionconnector_property_otp_secret
  }

  property {
    name  = "otpSiteId"
    type  = "string"
    value = var.transunionconnector_property_otp_site_id
  }

  property {
    name  = "otpUsername"
    type  = "string"
    value = var.transunionconnector_property_otp_username
  }
}
```


## Treasure Data

Connector ID (`connector_id` in the resource): `treasureDataConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Treasure Data API Key. Console display name: "API Key".


Example:
```terraform
resource "davinci_connection" "treasureDataConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "treasureDataConnector"
  name         = "My awesome treasureDataConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.treasuredataconnector_property_api_key
  }
}
```


## Trulioo

Connector ID (`connector_id` in the resource): `connectorTrulioo`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientID` (string): Trulioo Client ID. Console display name: "Client ID".
* `clientSecret` (string): Trulioo Client Secret. Console display name: "Client Secret".


Example:
```terraform
resource "davinci_connection" "connectorTrulioo" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorTrulioo"
  name         = "My awesome connectorTrulioo"

  property {
    name  = "clientID"
    type  = "string"
    value = var.connectortrulioo_property_client_i_d
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectortrulioo_property_client_secret
  }
}
```


## Twilio

Connector ID (`connector_id` in the resource): `twilioConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accountSid` (string):  Console display name: "Account Sid".
* `authDescription` (string):  Console display name: "Authentication Description".
* `authMessageTemplate` (string):  Console display name: "Text Message Template (Authentication)".
* `authToken` (string):  Console display name: "Auth Token".
* `connectorName` (string):  Console display name: "Connector Name".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `registerMessageTemplate` (string):  Console display name: "Text Message Template (Registration)".
* `senderPhoneNumber` (string):  Console display name: "Sender Phone Number".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```terraform
resource "davinci_connection" "twilioConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "twilioConnector"
  name         = "My awesome twilioConnector"

  property {
    name  = "accountSid"
    type  = "string"
    value = var.twilioconnector_property_account_sid
  }

  property {
    name  = "authDescription"
    type  = "string"
    value = var.twilioconnector_property_auth_description
  }

  property {
    name  = "authMessageTemplate"
    type  = "string"
    value = var.twilioconnector_property_auth_message_template
  }

  property {
    name  = "authToken"
    type  = "string"
    value = var.twilioconnector_property_auth_token
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.twilioconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.twilioconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.twilioconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.twilioconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.twilioconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.twilioconnector_property_icon_url_png
  }

  property {
    name  = "registerMessageTemplate"
    type  = "string"
    value = var.twilioconnector_property_register_message_template
  }

  property {
    name  = "senderPhoneNumber"
    type  = "string"
    value = var.twilioconnector_property_sender_phone_number
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.twilioconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.twilioconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.twilioconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.twilioconnector_property_tool_tip
  }
}
```


## TypingDNA

Connector ID (`connector_id` in the resource): `typingdnaConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "typingdnaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "typingdnaConnector"
  name         = "My awesome typingdnaConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.typingdnaconnector_property_custom_auth
  }
}
```


## UnifyID

Connector ID (`connector_id` in the resource): `unifyIdConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accountId` (string):  Console display name: "Account ID".
* `apiKey` (string):  Console display name: "API Key".
* `connectorName` (string):  Console display name: "Connector Name".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `sdkToken` (string):  Console display name: "SDK Token".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `toolTip` (string):  Console display name: "Tooltip".


Example:
```terraform
resource "davinci_connection" "unifyIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "unifyIdConnector"
  name         = "My awesome unifyIdConnector"

  property {
    name  = "accountId"
    type  = "string"
    value = var.unifyidconnector_property_account_id
  }

  property {
    name  = "apiKey"
    type  = "string"
    value = var.unifyidconnector_property_api_key
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.unifyidconnector_property_connector_name
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.unifyidconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.unifyidconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.unifyidconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.unifyidconnector_property_icon_url_png
  }

  property {
    name  = "sdkToken"
    type  = "string"
    value = var.unifyidconnector_property_sdk_token
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.unifyidconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.unifyidconnector_property_show_cred_added_via
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.unifyidconnector_property_tool_tip
  }
}
```


## User Policy

Connector ID (`connector_id` in the resource): `userPolicyConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `passwordExpiryInDays` (number): Choose 0 for never expire. Console display name: "Expires in the specified number of days".
* `passwordExpiryNotification` (boolean):  Console display name: "Notify user before password expires".
* `passwordLengthMax` (number):  Console display name: "Maximum Password Length".
* `passwordLengthMin` (number):  Console display name: "Minimum Password Length".
* `passwordLockoutAttempts` (number):  Console display name: "Number of failed login attempts before account is locked".
* `passwordPreviousXPasswords` (number): Choose 0 if any previous passwords are allowed. This is not recommended. Console display name: "Number of unique user passwords associated with a user".
* `passwordRequireLowercase` (boolean): Should the password contain lowercase characters?. Console display name: "Require Lowercase Characters".
* `passwordRequireNumbers` (boolean): Should the password contain numbers?. Console display name: "Require Numbers".
* `passwordRequireSpecial` (boolean): Should the password contain special character?. Console display name: "Require Special Characters".
* `passwordRequireUppercase` (boolean): Should the password contain uppercase characters?. Console display name: "Require Uppercase Characters".
* `passwordSpacesOk` (boolean): Are spaces allowed in the password?. Console display name: "Spaces Accepted".
* `passwordsEnabled` (boolean):  Console display name: "Passwords Feature Enabled?".
* `temporaryPasswordExpiryInDays` (number): If an administrator sets a temporary password, choose how long before it expires. Console display name: "Temporary password expires in the specified number of days".


Example:
```terraform
resource "davinci_connection" "userPolicyConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "userPolicyConnector"
  name         = "My awesome userPolicyConnector"

  property {
    name  = "passwordExpiryInDays"
    type  = "number"
    value = var.userpolicyconnector_property_password_expiry_in_days
  }

  property {
    name  = "passwordExpiryNotification"
    type  = "boolean"
    value = var.userpolicyconnector_property_password_expiry_notification
  }

  property {
    name  = "passwordLengthMax"
    type  = "number"
    value = var.userpolicyconnector_property_password_length_max
  }

  property {
    name  = "passwordLengthMin"
    type  = "number"
    value = var.userpolicyconnector_property_password_length_min
  }

  property {
    name  = "passwordLockoutAttempts"
    type  = "number"
    value = var.userpolicyconnector_property_password_lockout_attempts
  }

  property {
    name  = "passwordPreviousXPasswords"
    type  = "number"
    value = var.userpolicyconnector_property_password_previous_x_passwords
  }

  property {
    name  = "passwordRequireLowercase"
    type  = "boolean"
    value = var.userpolicyconnector_property_password_require_lowercase
  }

  property {
    name  = "passwordRequireNumbers"
    type  = "boolean"
    value = var.userpolicyconnector_property_password_require_numbers
  }

  property {
    name  = "passwordRequireSpecial"
    type  = "boolean"
    value = var.userpolicyconnector_property_password_require_special
  }

  property {
    name  = "passwordRequireUppercase"
    type  = "boolean"
    value = var.userpolicyconnector_property_password_require_uppercase
  }

  property {
    name  = "passwordSpacesOk"
    type  = "boolean"
    value = var.userpolicyconnector_property_password_spaces_ok
  }

  property {
    name  = "passwordsEnabled"
    type  = "boolean"
    value = var.userpolicyconnector_property_passwords_enabled
  }

  property {
    name  = "temporaryPasswordExpiryInDays"
    type  = "number"
    value = var.userpolicyconnector_property_temporary_password_expiry_in_days
  }
}
```


## User Pool

Connector ID (`connector_id` in the resource): `skUserPool`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAttributes` (json):  


Example:
```terraform
resource "davinci_connection" "skUserPool" {
  environment_id = var.pingone_environment_id

  connector_id = "skUserPool"
  name         = "My awesome skUserPool"

  property {
    name = "customAttributes"
    type = "json"
    value = jsonencode({
      "type" : "array",
      "preferredControlType" : "tableViewAttributes",
      "sections" : [
        "connectorAttributes"
      ],
      "value" : [
        {
          "name" : "username",
          "description" : "Username",
          "type" : "string",
          "value" : null,
          "minLength" : "1",
          "maxLength" : "300",
          "required" : true,
          "attributeType" : "sk"
        },
        {
          "name" : "firstName",
          "description" : "First Name",
          "type" : "string",
          "value" : null,
          "minLength" : "1",
          "maxLength" : "100",
          "required" : false,
          "attributeType" : "sk"
        },
        {
          "name" : "lastName",
          "description" : "Last Name",
          "type" : "string",
          "value" : null,
          "minLength" : "1",
          "maxLength" : "100",
          "required" : false,
          "attributeType" : "sk"
        },
        {
          "name" : "name",
          "description" : "Display Name",
          "type" : "string",
          "value" : null,
          "minLength" : "1",
          "maxLength" : "250",
          "required" : false,
          "attributeType" : "sk"
        },
        {
          "name" : "email",
          "description" : "Email",
          "type" : "string",
          "value" : null,
          "minLength" : "1",
          "maxLength" : "250",
          "required" : false,
          "attributeType" : "sk"
        }
      ]
    })
  }
}
```


## ValidSoft

Connector ID (`connector_id` in the resource): `connectorValidsoft`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorValidsoft" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorValidsoft"
  name         = "My awesome connectorValidsoft"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Variable

Connector ID (`connector_id` in the resource): `variablesConnector`

*No properties*


Example:
```terraform
resource "davinci_connection" "variablesConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "variablesConnector"
  name         = "My awesome variablesConnector"
}
```


## Venafi Account Service API

Connector ID (`connector_id` in the resource): `connector-oai-venafi`

Properties (used in the `property` block in the resource as the `name` parameter):

* `authApiKey` (string): The authentication key to the Venafi as a Service API for Account Service Operations. Console display name: "API Key".
* `basePath` (string): The base URL for contacting the API. Console display name: "Base Path".


Example:
```terraform
resource "davinci_connection" "connector-oai-venafi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-venafi"
  name         = "My awesome connector-oai-venafi"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-venafi_property_auth_api_key
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-venafi_property_base_path
  }
}
```


## Vericlouds

Connector ID (`connector_id` in the resource): `connectorVericlouds`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiSecret` (string): The API secret assigned by VeriClouds to the customer. The secret is also used for decrypting sensitive data such as leaked passwords. It is important to never share the secret with any 3rd party. Console display name: "apiSecret".
* `apikey` (string): The API key assigned by VeriClouds to the customer. Console display name: "apiKey".


Example:
```terraform
resource "davinci_connection" "connectorVericlouds" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorVericlouds"
  name         = "My awesome connectorVericlouds"

  property {
    name  = "apiSecret"
    type  = "string"
    value = var.connectorvericlouds_property_api_secret
  }

  property {
    name  = "apikey"
    type  = "string"
    value = var.connectorvericlouds_property_apikey
  }
}
```


## Veriff

Connector ID (`connector_id` in the resource): `veriffConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `access_token` (string): The API Key provided by Veriff, such as "323aa031-b4af-4e12-b354-de0da91a2ab0". Console display name: "API Key".
* `baseUrl` (string): The API URL to target, such as “https://stationapi.veriff.com/”. Console display name: "Base URL".
* `password` (string): The Share Secret Key from Veriff to create HMAC signature, such as "20bf4sf0-fbg7-488c-b4f1-d9594lf707bk". Console display name: "Shared Secret Key".


Example:
```terraform
resource "davinci_connection" "veriffConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "veriffConnector"
  name         = "My awesome veriffConnector"

  property {
    name  = "access_token"
    type  = "string"
    value = var.veriffconnector_property_access_token
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.veriffconnector_property_base_url
  }

  property {
    name  = "password"
    type  = "string"
    value = var.veriffconnector_property_password
  }
}
```


## Verosint

Connector ID (`connector_id` in the resource): `connector443id`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): This is the API key from your Verosint account. Remember, Your API KEY is like a serial number for your policy. If you want to utilize more than one policy, you can generate another API KEY and tailor that to a custom policy. Console display name: "API Key".


Example:
```terraform
resource "davinci_connection" "connector443id" {
  environment_id = var.pingone_environment_id

  connector_id = "connector443id"
  name         = "My awesome connector443id"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connector443id_property_api_key
  }
}
```


## Vidos

Connector ID (`connector_id` in the resource): `mailchainConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Enter your Vidos API Key obtained from the Vidos Dashboard with appropriate resolver or verifier permissions (visit https://dashboard.vidos.id/iam/api-keys). Console display name: "Vidos API Key".
* `version` (string): The verification API specification version. Console display name: "Verifier Version".


Example:
```terraform
resource "davinci_connection" "mailchainConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "mailchainConnector"
  name         = "My awesome mailchainConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.mailchainconnector_property_api_key
  }

  property {
    name  = "version"
    type  = "string"
    value = var.mailchainconnector_property_version
  }
}
```


## Webhook

Connector ID (`connector_id` in the resource): `webhookConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `urls` (string): POST requests will be made to these registered url as selected later. Console display name: "Register URLs".


Example:
```terraform
resource "davinci_connection" "webhookConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "webhookConnector"
  name         = "My awesome webhookConnector"

  property {
    name  = "urls"
    type  = "string"
    value = var.webhookconnector_property_urls
  }
}
```


## WhatsApp for Business

Connector ID (`connector_id` in the resource): `connectorWhatsAppBusiness`

Properties (used in the `property` block in the resource as the `name` parameter):

* `accessToken` (string): WhatsApp Access Token. Console display name: "Access Token".
* `appSecret` (string): WhatsApp App Secret for the application, it is used to verify the webhook signatures. Console display name: "App Secret".
* `skWebhookUri` (string): Use this url as the Webhook URL in the Third Party Integration's configuration. Console display name: "Redirect Webhook URI".
* `verifyToken` (string): Meta webhook verify token. Console display name: "Webhook Verify Token".
* `version` (string): WhatsApp Graph API Version. Console display name: "Version".


Example:
```terraform
resource "davinci_connection" "connectorWhatsAppBusiness" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorWhatsAppBusiness"
  name         = "My awesome connectorWhatsAppBusiness"

  property {
    name  = "accessToken"
    type  = "string"
    value = var.connectorwhatsappbusiness_property_access_token
  }

  property {
    name  = "appSecret"
    type  = "string"
    value = var.connectorwhatsappbusiness_property_app_secret
  }

  property {
    name  = "skWebhookUri"
    type  = "string"
    value = var.connectorwhatsappbusiness_property_sk_webhook_uri
  }

  property {
    name  = "verifyToken"
    type  = "string"
    value = var.connectorwhatsappbusiness_property_verify_token
  }

  property {
    name  = "version"
    type  = "string"
    value = var.connectorwhatsappbusiness_property_version
  }
}
```


## WinMagic

Connector ID (`connector_id` in the resource): `connectorWinmagic`

Properties (used in the `property` block in the resource as the `name` parameter):

* `openId` (json):  Console display name: "OpenId Parameters".


Example:
```terraform
resource "davinci_connection" "connectorWinmagic" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorWinmagic"
  name         = "My awesome connectorWinmagic"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
```


## WireWheel

Connector ID (`connector_id` in the resource): `wireWheelConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `baseURL` (string): The base API URL of the WireWheel environment. Console display name: "WireWheel Base API URL".
* `clientId` (string): Client ID from WireWheel Channel settings. Console display name: "Client ID".
* `clientSecret` (string): Client Secret from WireWheel Channel settings. Console display name: "Client Secret".
* `issuerId` (string): Issuer URL from WireWheel Channel settings. Console display name: "Issuer URL".


Example:
```terraform
resource "davinci_connection" "wireWheelConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "wireWheelConnector"
  name         = "My awesome wireWheelConnector"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.wirewheelconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.wirewheelconnector_property_client_secret
  }

  property {
    name  = "issuerId"
    type  = "string"
    value = var.wirewheelconnector_property_issuer_id
  }
}
```


## X Login

Connector ID (`connector_id` in the resource): `twitterIdpConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "twitterIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "twitterIdpConnector"
  name         = "My awesome twitterIdpConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Yoti

Connector ID (`connector_id` in the resource): `yotiConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "yotiConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "yotiConnector"
  name         = "My awesome yotiConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## Zendesk

Connector ID (`connector_id` in the resource): `connectorZendesk`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiToken` (string): An Active Zendesk API Token (admin center->Apps&Integrations->Zendesk API). Console display name: "Zendesk API Token".
* `emailUsername` (string): Email used as 'username' for your Zendesk account. Console display name: "Email of User (username)".
* `subdomain` (string): Your Zendesk subdomain (ex. {subdomain}.zendesk.com/api/v2/...). Console display name: "Subdomain".


Example:
```terraform
resource "davinci_connection" "connectorZendesk" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorZendesk"
  name         = "My awesome connectorZendesk"

  property {
    name  = "apiToken"
    type  = "string"
    value = var.connectorzendesk_property_api_token
  }

  property {
    name  = "emailUsername"
    type  = "string"
    value = var.connectorzendesk_property_email_username
  }

  property {
    name  = "subdomain"
    type  = "string"
    value = var.connectorzendesk_property_subdomain
  }
}
```


## Zoop.one

Connector ID (`connector_id` in the resource): `zoopConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `agencyId` (string):  Console display name: "Zoop Agency ID".
* `apiKey` (string):  Console display name: "Zoop API Key".
* `apiUrl` (string):  Console display name: "Zoop API URL".


Example:
```terraform
resource "davinci_connection" "zoopConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "zoopConnector"
  name         = "My awesome zoopConnector"

  property {
    name  = "agencyId"
    type  = "string"
    value = var.zoopconnector_property_agency_id
  }

  property {
    name  = "apiKey"
    type  = "string"
    value = var.zoopconnector_property_api_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.zoopconnector_property_api_url
  }
}
```


## Zscaler ZIA

Connector ID (`connector_id` in the resource): `connectorZscaler`

Properties (used in the `property` block in the resource as the `name` parameter):

* `basePath` (string): basePath. Console display name: "Base Path".
* `baseURL` (string): baseURL. Console display name: "Base URL".
* `zscalerAPIkey` (string): Zscaler APIkey. Console display name: "Zscaler APIkey".
* `zscalerPassword` (string): Zscaler Domain Password. Console display name: "Zscaler Password".
* `zscalerUsername` (string): Zscaler Domain Username. Console display name: "Zscaler Username".


Example:
```terraform
resource "davinci_connection" "connectorZscaler" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorZscaler"
  name         = "My awesome connectorZscaler"

  property {
    name  = "basePath"
    type  = "string"
    value = var.connectorzscaler_property_base_path
  }

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "zscalerAPIkey"
    type  = "string"
    value = var.zscaler_api_key
  }

  property {
    name  = "zscalerPassword"
    type  = "string"
    value = var.connectorzscaler_property_zscaler_password
  }

  property {
    name  = "zscalerUsername"
    type  = "string"
    value = var.connectorzscaler_property_zscaler_username
  }
}
```


## iProov

Connector ID (`connector_id` in the resource): `iproovConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `allowLandscape` (boolean):  Console display name: "Allow Landscape".
* `apiKey` (string):  Console display name: "API Key".
* `authDescription` (string):  Console display name: "Authentication Description".
* `baseUrl` (string):  Console display name: "Base URL".
* `color1` (string): Ex. #000000. Console display name: "Loading Tint Color".
* `color2` (string): Ex. #000000. Console display name: "Not Ready Tint Color".
* `color3` (string): Ex. #000000. Console display name: "Ready Tint Color".
* `color4` (string): Ex. #000000. Console display name: "Liveness Tint Color".
* `connectorName` (string):  Console display name: "Connector Name".
* `customTitle` (string): Specify a custom title to be shown. Defaults to show an iProov-generated message. Set to empty string "" to hide the message entirely.  Console display name: "Custom Title".
* `description` (string):  Console display name: "Description".
* `details1` (string):  Console display name: "Credentials Details 1".
* `details2` (string):  Console display name: "Credentials Details 2".
* `enableCameraSelector` (boolean):  Console display name: "Enable Camera Selector".
* `iconUrl` (string):  Console display name: "Icon URL".
* `iconUrlPng` (string):  Console display name: "Icon URL in PNG".
* `javascriptCSSUrl` (string):  Console display name: "CSS URL".
* `javascriptCdnUrl` (string):  Console display name: "Javascript CDN URL".
* `kioskMode` (boolean):  Console display name: "Kiosk Mode".
* `logo` (string): You can use a custom logo by simply passing a relative link, absolute path or data URI to your logo. If you do not want a logo to show pass the logo attribute as null. Console display name: "Logo".
* `password` (string):  Console display name: "Password".
* `secret` (string):  Console display name: "Secret".
* `showCountdown` (boolean):  Console display name: "Show Countdown".
* `showCredAddedOn` (boolean):  Console display name: "Show Credentials Added On?".
* `showCredAddedVia` (boolean):  Console display name: "Show Credentials Added through ?".
* `startScreenTitle` (string):  Console display name: "Start Screen Title".
* `title` (string):  Console display name: "Title".
* `toolTip` (string):  Console display name: "Tooltip".
* `username` (string):  Console display name: "Username".


Example:
```terraform
resource "davinci_connection" "iproovConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "iproovConnector"
  name         = "My awesome iproovConnector"

  property {
    name  = "allowLandscape"
    type  = "boolean"
    value = var.iproovconnector_property_allow_landscape
  }

  property {
    name  = "apiKey"
    type  = "string"
    value = var.iproovconnector_property_api_key
  }

  property {
    name  = "authDescription"
    type  = "string"
    value = var.iproovconnector_property_auth_description
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.iproovconnector_property_base_url
  }

  property {
    name  = "color1"
    type  = "string"
    value = var.iproovconnector_property_color1
  }

  property {
    name  = "color2"
    type  = "string"
    value = var.iproovconnector_property_color2
  }

  property {
    name  = "color3"
    type  = "string"
    value = var.iproovconnector_property_color3
  }

  property {
    name  = "color4"
    type  = "string"
    value = var.iproovconnector_property_color4
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.iproovconnector_property_connector_name
  }

  property {
    name  = "customTitle"
    type  = "string"
    value = var.iproovconnector_property_custom_title
  }

  property {
    name  = "description"
    type  = "string"
    value = var.iproovconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.iproovconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.iproovconnector_property_details2
  }

  property {
    name  = "enableCameraSelector"
    type  = "boolean"
    value = var.iproovconnector_property_enable_camera_selector
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.iproovconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.iproovconnector_property_icon_url_png
  }

  property {
    name  = "javascriptCSSUrl"
    type  = "string"
    value = var.javascript_css_url
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.iproovconnector_property_javascript_cdn_url
  }

  property {
    name  = "kioskMode"
    type  = "boolean"
    value = var.iproovconnector_property_kiosk_mode
  }

  property {
    name  = "logo"
    type  = "string"
    value = var.iproovconnector_property_logo
  }

  property {
    name  = "password"
    type  = "string"
    value = var.iproovconnector_property_password
  }

  property {
    name  = "secret"
    type  = "string"
    value = var.iproovconnector_property_secret
  }

  property {
    name  = "showCountdown"
    type  = "boolean"
    value = var.iproovconnector_property_show_countdown
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.iproovconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.iproovconnector_property_show_cred_added_via
  }

  property {
    name  = "startScreenTitle"
    type  = "string"
    value = var.iproovconnector_property_start_screen_title
  }

  property {
    name  = "title"
    type  = "string"
    value = var.iproovconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.iproovconnector_property_tool_tip
  }

  property {
    name  = "username"
    type  = "string"
    value = var.iproovconnector_property_username
  }
}
```


## iProov API

Connector ID (`connector_id` in the resource): `iproovV2Connector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Your iProov Service Provider API key. This can be obtained from your iPortal account. Please contact support@iproov.com for more information. Console display name: "iProov API Key".
* `secret` (string): Your iProov Service Provider Secret. This can be obtained from your iPortal account. Please contact support@iproov.com for more information. Console display name: "iProov Secret".
* `tenant` (string): The iProov tenant URL (do not include https://). This can be obtained from your iPortal account. Please contact support@iproov.com for more information. Console display name: "iProov Tenant".


Example:
```terraform
resource "davinci_connection" "iproovV2Connector" {
  environment_id = var.pingone_environment_id

  connector_id = "iproovV2Connector"
  name         = "My awesome iproovV2Connector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.iproovv2connector_property_api_key
  }

  property {
    name  = "secret"
    type  = "string"
    value = var.iproovv2connector_property_secret
  }

  property {
    name  = "tenant"
    type  = "string"
    value = var.iproovv2connector_property_tenant
  }
}
```


## iProov OIDC

Connector ID (`connector_id` in the resource): `connectorSvipe`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorSvipe" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSvipe"
  name         = "My awesome connectorSvipe"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```


## iovation

Connector ID (`connector_id` in the resource): `iovationConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiUrl` (string):  Console display name: "API Server URL".
* `javascriptCdnUrl` (string): iovation loader javascript CDN. Console display name: "iovation loader Javascript CDN URL".
* `subKey` (string): This will be an iovation assigned value that tracks requests from your site. This is primarily used for debugging and troubleshooting purposes. Console display name: "Sub Key".
* `subscriberAccount` (string):  Console display name: "Subscriber Account".
* `subscriberId` (string):  Console display name: "Subscriber ID".
* `subscriberPasscode` (string):  Console display name: "Subscriber Passcode".
* `version` (string): This is the version of the script to load. The value should either correspond to a specific version you wish to use, or one of the following aliases to get the latest version of the code: general5 - the latest stable version of the javascript, early5 - the latest available version of the javascript. Console display name: "Version".


Example:
```terraform
resource "davinci_connection" "iovationConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "iovationConnector"
  name         = "My awesome iovationConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.iovationconnector_property_api_url
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.iovationconnector_property_javascript_cdn_url
  }

  property {
    name  = "subKey"
    type  = "string"
    value = var.iovationconnector_property_sub_key
  }

  property {
    name  = "subscriberAccount"
    type  = "string"
    value = var.iovationconnector_property_subscriber_account
  }

  property {
    name  = "subscriberId"
    type  = "string"
    value = var.iovationconnector_property_subscriber_id
  }

  property {
    name  = "subscriberPasscode"
    type  = "string"
    value = var.iovationconnector_property_subscriber_passcode
  }

  property {
    name  = "version"
    type  = "string"
    value = var.iovationconnector_property_version
  }
}
```


## ipgeolocation.io

Connector ID (`connector_id` in the resource): `connectorIPGeolocationio`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): Developer subscription API key. Console display name: "API key".


Example:
```terraform
resource "davinci_connection" "connectorIPGeolocationio" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIPGeolocationio"
  name         = "My awesome connectorIPGeolocationio"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectoripgeolocationio_property_api_key
  }
}
```


## ipregistry

Connector ID (`connector_id` in the resource): `connectorIPregistry`

Properties (used in the `property` block in the resource as the `name` parameter):

* `apiKey` (string): API Key used to authenticate to the ipregistry.co API. Console display name: "API Key".


Example:
```terraform
resource "davinci_connection" "connectorIPregistry" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIPregistry"
  name         = "My awesome connectorIPregistry"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectoripregistry_property_api_key
  }
}
```


## ipstack

Connector ID (`connector_id` in the resource): `connectorIPStack`

Properties (used in the `property` block in the resource as the `name` parameter):

* `allowInsecureIPStackConnection` (string): The Free IPStack Subscription Plan does not support HTTPS connections. For more information refer to https://ipstack.com/plan. Console display name: "Allow Insecure ipstack Connection?".
* `apiKey` (string): The ipstack API key to use the service. Console display name: "API key".


Example:
```terraform
resource "davinci_connection" "connectorIPStack" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIPStack"
  name         = "My awesome connectorIPStack"

  property {
    name  = "allowInsecureIPStackConnection"
    type  = "string"
    value = var.allow_insecure_ip_stack_connection
  }

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectoripstack_property_api_key
  }
}
```


## mParticle

Connector ID (`connector_id` in the resource): `mparticleConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `clientID` (string): Client ID from mParticle tenant. Console display name: "Client ID".
* `clientSecret` (string): Client Secret from mParticle tenant. Console display name: "Client Secret".
* `pod` (string): Pod from mParticle tenant. Only required for 'Upload an event batch' capability. Console display name: "Pod".


Example:
```terraform
resource "davinci_connection" "mparticleConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "mparticleConnector"
  name         = "My awesome mparticleConnector"

  property {
    name  = "clientID"
    type  = "string"
    value = var.mparticleconnector_property_client_i_d
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.mparticleconnector_property_client_secret
  }

  property {
    name  = "pod"
    type  = "string"
    value = var.mparticleconnector_property_pod
  }
}
```


## neoEYED

Connector ID (`connector_id` in the resource): `neoeyedConnector`

Properties (used in the `property` block in the resource as the `name` parameter):

* `appKey` (string): Unique key for the application. Console display name: "Application Key".
* `javascriptCdnUrl` (string): URL of javascript CDN of neoEYED. Console display name: "Javascript CDN URL".


Example:
```terraform
resource "davinci_connection" "neoeyedConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "neoeyedConnector"
  name         = "My awesome neoeyedConnector"

  property {
    name  = "appKey"
    type  = "string"
    value = var.neoeyedconnector_property_app_key
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.neoeyedconnector_property_javascript_cdn_url
  }
}
```


## randomuser.me

Connector ID (`connector_id` in the resource): `connectorRandomUserMe`

*No properties*


Example:
```terraform
resource "davinci_connection" "connectorRandomUserMe" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorRandomUserMe"
  name         = "My awesome connectorRandomUserMe"
}
```


## tru.ID

Connector ID (`connector_id` in the resource): `connectorTruid`

Properties (used in the `property` block in the resource as the `name` parameter):

* `customAuth` (json):  Console display name: "Custom Parameters".


Example:
```terraform
resource "davinci_connection" "connectorTruid" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorTruid"
  name         = "My awesome connectorTruid"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
```

