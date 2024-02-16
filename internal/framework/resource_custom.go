package framework

import "github.com/pingidentity/terraform-provider-davinci/internal/framework/customtypes/davinciexporttype"

func DaVinciExportTypeToTF(v string) davinciexporttype.ParsedValue {
	if v == "" {
		return davinciexporttype.NewParsedNull()
	} else {
		return davinciexporttype.NewParsedValue(v)
	}
}

func DaVinciExportTypeOkToTF(v *string, ok bool) davinciexporttype.ParsedValue {
	if !ok || v == nil {
		return davinciexporttype.NewParsedNull()
	} else {
		return davinciexporttype.NewParsedValue(*v)
	}
}
