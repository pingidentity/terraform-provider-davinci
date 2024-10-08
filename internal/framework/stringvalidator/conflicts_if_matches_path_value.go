package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework/schemavalidator"
)

func ConflictsIfMatchesPathValue(targetValue basetypes.StringValue, expressions ...path.Expression) validator.String {
	return schemavalidator.ConflictsIfMatchesPathValueValidator{
		TargetValue: targetValue,
		Expressions: expressions,
	}
}
