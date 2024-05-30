package davinciexporttype

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringValuable = ParsedValue{}
var _ basetypes.StringValuableWithSemanticEquals = ParsedValue{}
var _ xattr.ValidateableAttribute = ParsedValue{}

type ParsedValue struct {
	basetypes.StringValue
	// ... potentially other fields ...
	davinci.ExportCmpOpts
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
	return ParsedType{
		ExportCmpOpts: v.ExportCmpOpts,
	}
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
	return davinci.Equal(priorFlow, newFlow, davinci.ExportCmpOpts{
		IgnoreConfig:              v.IgnoreConfig,
		IgnoreDesignerCues:        v.IgnoreDesignerCues,
		IgnoreEnvironmentMetadata: v.IgnoreEnvironmentMetadata,
		IgnoreUnmappedProperties:  v.IgnoreUnmappedProperties,
		IgnoreVersionMetadata:     v.IgnoreVersionMetadata,
		IgnoreFlowMetadata:        v.IgnoreFlowMetadata,
	}), diags
}

func (v ParsedValue) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {

	if v.IsNull() || v.IsUnknown() {
		return
	}

	if ok := davinci.ValidFlowsInfoExport([]byte(v.ValueString()), davinci.ExportCmpOpts{}); ok {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid DaVinci Flow Export String Value",
			"A string value was provided that is not valid DaVinci Export JSON string format.  The export should not including subflows as these should be managed separately, as their own independent flows.\n\n"+
				"Please re-export the DaVinci flow without subflows included.\n",
		)

		return
	}

	// Validate just the config of the export
	if ok := davinci.ValidFlowExport([]byte(v.ValueString()), davinci.ExportCmpOpts{
		IgnoreConfig:              v.IgnoreConfig,
		IgnoreDesignerCues:        v.IgnoreDesignerCues,
		IgnoreEnvironmentMetadata: v.IgnoreEnvironmentMetadata,
		IgnoreUnmappedProperties:  true,
		IgnoreVersionMetadata:     v.IgnoreVersionMetadata,
		IgnoreFlowMetadata:        v.IgnoreFlowMetadata,
	}); !ok {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid DaVinci Flow Export String Value",
			"A string value was provided that is not valid DaVinci Export JSON string format.\n\n"+
				"Please re-export the DaVinci flow.  If the flow JSON has been correctly exported from the DaVinci environment (and can be re-imported), please report this error to the provider maintainers.\n",
		)

		return
	}

	// Warn in case there are AdditionalProperties in the import file (since these aren't cleanly handled in the SDK, while they are preserved on import, there may be unpredictable results in diff calculation)
	if ok := davinci.ValidFlowExport([]byte(v.ValueString()), davinci.ExportCmpOpts{
		IgnoreConfig:              true,
		IgnoreDesignerCues:        true,
		IgnoreEnvironmentMetadata: true,
		IgnoreUnmappedProperties:  false,
		IgnoreVersionMetadata:     true,
		IgnoreFlowMetadata:        true,
	}); !v.IgnoreUnmappedProperties && !ok {
		resp.Diagnostics.AddAttributeWarning(
			req.Path,
			"DaVinci Export JSON contains unknown properties",
			"The DaVinci Flow Export contains properties that cannot be validated.  These parameters will be preserved on import to the DaVinci service, but there may be unpredictable results in difference calculation.\n",
		)
	}
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

func NewParsedValue(value string, cmpOpts davinci.ExportCmpOpts) ParsedValue {
	return ParsedValue{
		StringValue:   basetypes.NewStringValue(value),
		ExportCmpOpts: cmpOpts,
	}
}

func NewParsedPointerValue(value *string, cmpOpts davinci.ExportCmpOpts) ParsedValue {
	return ParsedValue{
		StringValue:   basetypes.NewStringPointerValue(value),
		ExportCmpOpts: cmpOpts,
	}
}
