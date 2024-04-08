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
