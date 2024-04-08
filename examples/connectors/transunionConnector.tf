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
