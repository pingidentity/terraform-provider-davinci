resource "davinci_connection" "userPolicyConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "userPolicyConnector"
  name         = "My awesome userPolicyConnector"

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
}
