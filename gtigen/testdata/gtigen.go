// Code generated by "gtigen.test.exe -test.paniconexit0 -test.timeout=10m0s"; DO NOT EDIT.

package testdata

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var PersonType = &gti.Type{
	Name: "goki.dev/gti/gtigen/testdata.Person",
	Doc:  "Person represents a person and their attributes.\nThe zero value of a Person is not valid.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{"-type-var", "-instance"}},
		&gti.Directive{Tool: "ki", Directive: "flagtype", Args: []string{"NodeFlags", "-field", "Flag"}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Name", &gti.Field{Name: "Name", Doc: "Name is the name of the person", Directives: gti.Directives{
			&gti.Directive{Tool: "gi", Directive: "toolbar", Args: []string{"-hide"}},
		}}},
		{"Age", &gti.Field{Name: "Age", Doc: "Age is the age of the person", Directives: gti.Directives{
			&gti.Directive{Tool: "gi", Directive: "view", Args: []string{"inline"}},
		}}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{
		{"Introduction", &gti.Method{Name: "Introduction", Doc: "Introduction returns an introduction for the person.\nIt contains the name of the person and their age.", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
			&gti.Directive{Tool: "gi", Directive: "toolbar", Args: []string{"-name", "ShowIntroduction", "-icon", "play", "-show-result", "-confirm"}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"string", &gti.Field{Name: "string", Doc: "", Directives: gti.Directives{}}},
		})}},
	}),
	Instance: &Person{},
}
