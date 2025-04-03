package templates

import (
	"strings"
	"text/template"
)

// Helper functions for the template
var templateFuncs = template.FuncMap{
	"ToUpper":             strings.ToUpper,
	"ToLower":             strings.ToLower,
	"ToSnakeCase":         ToSnakeCase,
	"GetCType":            GetCType,
	"GetTypeSizeC":        GetTypeSizeC,
	"GetDefaultValueC":    GetDefaultValueC,
	"CalculateStructSize": CalculateStructSize,
	"NeedsByteSwap":       NeedsByteSwap,
	"GetTSType":           GetTSType,
	"GetDefaultValueTS":   GetDefaultValueTS,
	"sub": func(a, b int) int {
		return a - b
	},
}

// CHeaderTemplate generates a simple C header file with struct definitions
var CHeaderTemplate = template.Must(template.New("cheader").Funcs(templateFuncs).Parse(`/**
* {{.Name}}
* {{.Description}}
*/

#ifndef {{.Name | ToUpper}}_H
#define {{.Name | ToUpper}}_H

#include <stdint.h>
  #include <stdbool.h>

    /**
    * {{.Description}}
    */
    typedef struct {
    {{- range .Items}}
    {{- if .Units}}
    /* {{.Description}} ({{.Units}}) */
    {{- else}}
    /* {{.Description}} */
    {{- end}}
    {{- if .IsArray}}
    {{GetCType .}} {{.Name}}[{{.Length}}];
    {{- else}}
    {{GetCType .}} {{.Name}};
    {{- end}}
    {{- end}}
    } {{.Name}}_t;

    /**
    * Initialize a {{.Name}} structure with default values
    * @param p_data Pointer to the structure to initialize
    */
    void {{.Name | ToSnakeCase}}_init({{.Name}}_t* p_data);

    #endif /* {{.Name | ToUpper}}_H */
    `))
