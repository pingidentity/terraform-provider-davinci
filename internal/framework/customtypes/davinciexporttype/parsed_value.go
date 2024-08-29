package davinciexporttype

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
		IgnoreFlowVariables:       v.IgnoreFlowVariables,
		NodeOpts:                  v.NodeOpts,
	}), diags
}

func (v ParsedValue) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {

	if v.IsNull() || v.IsUnknown() {
		return
	}

	var flows davinci.Flows

	// should really use the actual validator
	err := davinci.Unmarshal([]byte(v.ValueString()), &flows, davinci.ExportCmpOpts{})
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid DaVinci Flow Export String Value",
			"A string value was provided that is not valid DaVinci Export JSON for this provider.\n\n"+
				v.parseValidationErrorMessage(err)+"\n",
		)
		return
	}

	if len(flows.Flow) > 0 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid DaVinci Flow Export String Value",
			"A string value was provided that is not valid DaVinci Export JSON string format.  The export should not including subflows as these should be managed separately, as their own independent flows.\n\n"+
				"Please re-export the DaVinci flow without subflows included.\n",
		)

		return
	}

	// Validate just the config of the export
	err = davinci.ValidFlowExport([]byte(v.ValueString()), davinci.ExportCmpOpts{
		IgnoreConfig:              v.IgnoreConfig,
		IgnoreDesignerCues:        v.IgnoreDesignerCues,
		IgnoreEnvironmentMetadata: v.IgnoreEnvironmentMetadata,
		IgnoreUnmappedProperties:  true,
		IgnoreVersionMetadata:     v.IgnoreVersionMetadata,
		IgnoreFlowMetadata:        v.IgnoreFlowMetadata,
		IgnoreFlowVariables:       v.IgnoreFlowVariables,
		NodeOpts:                  v.NodeOpts,
	})

	if err != nil {
		tflog.Debug(ctx, "Invalid DaVinci Flow Export String Value", map[string]interface{}{
			"error": err,
		})

		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid DaVinci Flow Export String Value",
			"A string value was provided that is not valid DaVinci Export JSON for this provider.\n\n"+
				v.parseValidationErrorMessage(err)+"\n",
		)

		return
	}

	// Warn in case there are AdditionalProperties in the import file (since these aren't cleanly handled in the SDK, while they are preserved on import, there may be unpredictable results in diff calculation)
	err = davinci.ValidFlowExport([]byte(v.ValueString()), davinci.ExportCmpOpts{
		IgnoreConfig:              true,
		IgnoreDesignerCues:        true,
		IgnoreEnvironmentMetadata: true,
		IgnoreUnmappedProperties:  false,
		IgnoreVersionMetadata:     true,
		IgnoreFlowMetadata:        true,
		IgnoreFlowVariables:       true,
	})

	if !v.IgnoreUnmappedProperties && err != nil {
		tflog.Debug(ctx, "Invalid DaVinci Flow Export String Value", map[string]interface{}{
			"error": err,
		})

		resp.Diagnostics.AddAttributeWarning(
			req.Path,
			"DaVinci Export JSON contains unknown properties",
			v.parseValidationErrorMessage(err)+"\n",
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

func (v ParsedValue) parseValidationErrorMessage(err error) string {
	var equatesEmptyError *davinci.EquatesEmptyTypeError
	var missingRequiredFlowFieldsError *davinci.MissingRequiredFlowFieldsTypeError
	var unknownAdditionalFieldsError *davinci.UnknownAdditionalFieldsTypeError
	var minFlowDefsError *davinci.MinFlowDefinitionsExceededTypeError
	var maxFlowDefsError *davinci.MaxFlowDefinitionsExceededTypeError
	switch {
	case errors.Is(err, davinci.ErrInvalidJson):
		return "The DaVinci Flow Export JSON is not valid JSON.  Please re-export the DaVinci flow."
	case errors.Is(err, davinci.ErrEmptyFlow):
		return "The DaVinci Flow Export JSON is empty.  Please re-export the DaVinci flow."
	case errors.Is(err, davinci.ErrNoFlowDefinition):
		return "No flow definition found in the DaVinci Flow Export JSON.  Expecting exactly one flow definition.  Please re-export the DaVinci flow."
	case errors.Is(err, davinci.ErrMissingSaveVariableValues):
		return "Save flow variable nodes are present but are missing variable values in the DaVinci Flow Export JSON.  Please re-export the DaVinci flow ensuring that variable values are included."
	case errors.As(err, &equatesEmptyError):
		return "The DaVinci Flow Export JSON has been evaluated to be empty according to plan diff criteria.  Please re-export the DaVinci flow."
	case errors.As(err, &missingRequiredFlowFieldsError):
		return "The DaVinci Flow Export JSON has been evaluated to be missing required fields.  Please re-export the DaVinci flow."
	case errors.As(err, &unknownAdditionalFieldsError):
		return "The DaVinci Flow Export contains unknown properties that cannot be validated.  These parameters will be preserved on import to the DaVinci service, but there may be unpredictable results in difference calculation."
	case errors.As(err, &minFlowDefsError):
		return fmt.Sprintf("There are not enough flows exported in the flow group.  Expecting a minimum of %d", minFlowDefsError.Min)
	case errors.As(err, &maxFlowDefsError):
		return fmt.Sprintf("There are too many flows exported in the flow group.  Expecting a maximum of %d", maxFlowDefsError.Max)
	default:
		return fmt.Sprintf("An unexpected error was encountered while validating the DaVinci Export JSON string: %s.  This is always an error in the provider and should be reported to the provider maintainers. ", err)
	}
}
