package framework

import (
	"github.com/pingidentity/terraform-provider-davinci/internal/framework/customtypes/davinciexporttype"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

func DaVinciExportTypeToTF(v string, cmpOpts davinci.ExportCmpOpts) davinciexporttype.ParsedValue {
	if v == "" {
		return davinciexporttype.NewParsedNull()
	} else {
		return davinciexporttype.NewParsedValue(v, cmpOpts)
	}
}

func DaVinciExportTypeOkToTF(v *string, ok bool, cmpOpts davinci.ExportCmpOpts) davinciexporttype.ParsedValue {
	if !ok || v == nil {
		return davinciexporttype.NewParsedNull()
	} else {
		return davinciexporttype.NewParsedValue(*v, cmpOpts)
	}
}
