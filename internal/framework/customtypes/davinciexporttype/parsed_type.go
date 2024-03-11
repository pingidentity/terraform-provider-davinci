package davinciexporttype

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringTypable = ParsedType{}
var _ xattr.TypeWithValidate = ParsedType{}

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
	return ParsedValue{}
}

func (t ParsedType) Validate(ctx context.Context, in tftypes.Value, path path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	if in.Type() == nil {
		return diags
	}

	if !in.Type().Is(tftypes.String) {
		err := fmt.Errorf("expected String value, received %T with value: %v", in, in)
		diags.AddAttributeError(
			path,
			"DaVinci Export Parsed Type Validation Error",
			"An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. "+
				"Please report the following to the provider developer:\n\n"+err.Error(),
		)
		return diags
	}

	if !in.IsKnown() || in.IsNull() {
		return diags
	}

	var valueString string

	if err := in.As(&valueString); err != nil {
		diags.AddAttributeError(
			path,
			"DaVinci Export Parsed Type Validation Error",
			"An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. "+
				"Please report the following to the provider developer:\n\n"+err.Error(),
		)

		return diags
	}

	if ok := davinci.ValidFlowsInfoExport([]byte(valueString), davinci.ExportCmpOpts{}); ok {
		diags.AddAttributeError(
			path,
			"Invalid DaVinci Flow Export String Value",
			"A string value was provided that is not valid DaVinci Export JSON string format.  The export should not including subflows as these should be managed separately, as their own independent flows.\n\n"+
				"Please re-export the DaVinci flow without subflows included.\n",
		)

		return diags
	}

	// Validate just the config of the export
	if ok := davinci.ValidFlowExport([]byte(valueString), davinci.ExportCmpOpts{
		IgnoreConfig:              t.IgnoreConfig,
		IgnoreDesignerCues:        t.IgnoreDesignerCues,
		IgnoreEnvironmentMetadata: t.IgnoreEnvironmentMetadata,
		IgnoreUnmappedProperties:  true,
		IgnoreVersionMetadata:     t.IgnoreVersionMetadata,
		IgnoreFlowMetadata:        t.IgnoreFlowMetadata,
	}); !ok {
		diags.AddAttributeError(
			path,
			"Invalid DaVinci Flow Export String Value",
			"A string value was provided that is not valid DaVinci Export JSON string format.\n\n"+
				"Please re-export the DaVinci flow.  If the flow JSON has been correctly exported from the DaVinci environment (and can be re-imported), please report this error to the provider maintainers.\n",
		)

		return diags
	}

	// Warn in case there are AdditionalProperties in the import file (since these aren't cleanly handled in the SDK, while they are preserved on import, there may be unpredictable results in diff calculation)
	if ok := davinci.ValidFlowExport([]byte(valueString), davinci.ExportCmpOpts{
		IgnoreConfig:              true,
		IgnoreDesignerCues:        true,
		IgnoreEnvironmentMetadata: true,
		IgnoreUnmappedProperties:  false,
		IgnoreVersionMetadata:     true,
		IgnoreFlowMetadata:        true,
	}); !t.IgnoreUnmappedProperties && !ok {
		diags.AddAttributeWarning(
			path,
			"DaVinci Export JSON contains unknown properties",
			"The DaVinci Flow Export contains properties that cannot be validated.  These parameters will be preserved on import to the DaVinci service, but there may be unpredictable results in difference calculation.\n",
		)
	}

	return diags
}
