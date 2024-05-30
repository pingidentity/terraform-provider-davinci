package davinciexporttype

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringTypable = ParsedType{}

type ParsedType struct {
	basetypes.StringType
	// ... potentially other fields ...
	davinci.ExportCmpOpts
}

func (t ParsedType) Equal(o attr.Type) bool {
	other, ok := o.(ParsedType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t ParsedType) String() string {
	return "davinciexporttype.ParsedType"
}

func (t ParsedType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	// ParsedValue defined in the value type section
	value := ParsedValue{
		StringValue:   in,
		ExportCmpOpts: t.ExportCmpOpts,
	}

	return value, nil
}

func (t ParsedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (t ParsedType) ValueType(ctx context.Context) attr.Value {
	// ParsedValue defined in the value type section
	return ParsedValue{
		ExportCmpOpts: t.ExportCmpOpts,
	}
}
