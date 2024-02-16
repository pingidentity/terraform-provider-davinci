package davinciexporttype

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringValuable = ParsedValue{}
var _ basetypes.StringValuableWithSemanticEquals = ParsedValue{}

type ParsedValue struct {
	basetypes.StringValue
	// ... potentially other fields ...
	ImportFile bool
}

func (v ParsedValue) Equal(o attr.Value) bool {
	other, ok := o.(ParsedValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v ParsedValue) Type(ctx context.Context) attr.Type {
	// ParsedType defined in the schema type section
	return ParsedType{}
}

func (v ParsedValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	// The framework should always pass the correct value type, but always check
	newValue, ok := newValuable.(ParsedValue)

	if !ok {
		diags.AddError(
			"DaVinci Export Parsed Type Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	var priorFlow, newFlow davinci.Flow

	if err := json.Unmarshal([]byte(v.StringValue.ValueString()), &priorFlow); err != nil {
		diags.AddError(
			"DaVinci Export Parsed Type Semantic Equality Check Error",
			"An unexpected error was encountered trying to parse the prior attribute value to perform semantic equals evaluation. This is always an error in the provider. "+
				"Please report the following to the provider developer:\n\n"+err.Error(),
		)
		return false, diags
	}

	if err := json.Unmarshal([]byte(newValue.ValueString()), &newFlow); err != nil {
		diags.AddError(
			"DaVinci Export Parsed Type Semantic Equality Check Error",
			"An unexpected error was encountered trying to parse the new attribute value to perform semantic equals evaluation. This is always an error in the provider. "+
				"Please report the following to the provider developer:\n\n"+err.Error(),
		)
		return false, diags
	}

	// Check whether the flows are equal, ignoring environment metadata and designer UI cues.  Just the flow configuration
	return davinci.FlowEqual(priorFlow, newFlow, davinci.ExportCmpOpts{
		IgnoreConfig:              false,
		IgnoreDesignerCues:        true,
		IgnoreEnvironmentMetadata: !v.ImportFile,
		IgnoreUnmappedProperties:  true,
		IgnoreVersionMetadata:     true,
		IgnoreFlowMetadata:        true,
	}), diags
}

func NewParsedNull() ParsedValue {
	return ParsedValue{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewParsedUnknown() ParsedValue {
	return ParsedValue{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewParsedValue(value string) ParsedValue {
	return ParsedValue{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewParsedPointerValue(value *string) ParsedValue {
	return ParsedValue{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
