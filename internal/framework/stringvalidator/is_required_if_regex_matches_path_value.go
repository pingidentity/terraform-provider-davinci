package stringvalidator

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework/schemavalidator"
)

// IsRequiredIfRegexMatchesPathValue validates if the provided regex matches
// the value at the provided path expression(s). If matched, the current argument is required.
//
// If a list of expressions is provided, all expressions are checked until a match is found,
// or the list of expressions is exhausted.
func IsRequiredIfRegexMatchesPathValue(regexp *regexp.Regexp, message string, expressions ...path.Expression) validator.String {
	return schemavalidator.IsRequiredIfRegexMatchesPathValueValidator{
		Regexp:      regexp,
		Message:     message,
		Expressions: expressions,
	}
}