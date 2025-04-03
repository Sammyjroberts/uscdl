package templates

import (
	"text/template"
)

// TypeScriptTemplate generates a TypeScript interface definition
var TypeScriptTemplate = template.Must(template.New("typescript").Funcs(templateFuncs).Parse(`/**
* {{.Name}}
* {{.Description}}
*/
export interface {{.Name}} {
{{- range .Items}}
{{- if .Units}}
/** {{.Description}} ({{.Units}}) */
{{- else}}
/** {{.Description}} */
{{- end}}
{{.Name}}: {{GetTSType .}};
{{- end}}
}

/**
* Creates a default {{.Name}} object
* @returns A new {{.Name}} with default values
*/
export function create{{.Name}}(): {{.Name}} {
return {
{{- range $i, $item := .Items}}
{{$item.Name}}: {{GetDefaultValueTS $item}}{{if lt $i (sub (len $.Items) 1)}},{{end}}
{{- end}}
};
}

/**
* Serializes a {{.Name}} object to an ArrayBuffer
* @param data The {{.Name}} object to serialize
* @returns An ArrayBuffer containing the serialized data
*/
export function serialize{{.Name}}(data: {{.Name}}): ArrayBuffer {
const buffer = new ArrayBuffer({{CalculateStructSize .}});
const view = new DataView(buffer);
let offset = 0;

{{- range .Items}}
{{- if .IsArray}}
for (let i = 0; i < {{.Length}}; i++) { {{- if eq .Type "uint8" }} view.setUint8(offset, data.{{.Name}}[i]); offset +=1;
  {{- else if eq .Type "uint16" }} view.setUint16(offset, data.{{.Name}}[i], {{if eq .ByteOrder "little"
  }}true{{else}}false{{end}}); offset +=2; {{- else if eq .Type "uint32" }} view.setUint32(offset, data.{{.Name}}[i],
  {{if eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=4; {{- else if eq .Type "int8" }}
  view.setInt8(offset, data.{{.Name}}[i]); offset +=1; {{- else if eq .Type "int16" }} view.setInt16(offset,
  data.{{.Name}}[i], {{if eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=2; {{- else if eq .Type "int32"
  }} view.setInt32(offset, data.{{.Name}}[i], {{if eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=4; {{-
  else if eq .Type "float" }} view.setFloat32(offset, data.{{.Name}}[i], {{if eq .ByteOrder "little"
  }}true{{else}}false{{end}}); offset +=4; {{- else if eq .Type "double" }} view.setFloat64(offset, data.{{.Name}}[i],
  {{if eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=8; {{- else if eq .Type "bool" }}
  view.setUint8(offset, data.{{.Name}}[i] ? 1 : 0); offset +=1; {{- end}} } {{- else}} {{- if eq .Type "uint8" }}
  view.setUint8(offset, data.{{.Name}}); offset +=1; {{- else if eq .Type "uint16" }} view.setUint16(offset,
  data.{{.Name}}, {{if eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=2; {{- else if eq .Type "uint32" }}
  view.setUint32(offset, data.{{.Name}}, {{if eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=4; {{- else
  if eq .Type "int8" }} view.setInt8(offset, data.{{.Name}}); offset +=1; {{- else if eq .Type "int16" }}
  view.setInt16(offset, data.{{.Name}}, {{if eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=2; {{- else if
  eq .Type "int32" }} view.setInt32(offset, data.{{.Name}}, {{if eq .ByteOrder "little" }}true{{else}}false{{end}});
  offset +=4; {{- else if eq .Type "float" }} view.setFloat32(offset, data.{{.Name}}, {{if eq .ByteOrder "little"
  }}true{{else}}false{{end}}); offset +=4; {{- else if eq .Type "double" }} view.setFloat64(offset, data.{{.Name}}, {{if
  eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=8; {{- else if eq .Type "bool" }} view.setUint8(offset,
  data.{{.Name}} ? 1 : 0); offset +=1; {{- else if eq .Type "string" }} // String serialization not fully implemented //
  This is a placeholder for string handling offset +=4; // Allocate space for string length {{- end}} {{- end}} {{-
  end}} return buffer; } /** * Deserializes an ArrayBuffer to a {{.Name}} object * @param buffer The ArrayBuffer
  containing serialized data * @returns A {{.Name}} object with the deserialized data */ export function
  deserialize{{.Name}}(buffer: ArrayBuffer): {{.Name}} { const view=new DataView(buffer); let offset=0; const
  result=create{{.Name}}(); {{- range .Items}} {{- if .IsArray}} const {{.Name}}Array=[]; for (let i=0; i < {{.Length}};
  i++) { {{- if eq .Type "uint8" }} {{.Name}}Array.push(view.getUint8(offset)); offset +=1; {{- else if eq
  .Type "uint16" }} {{.Name}}Array.push(view.getUint16(offset, {{if eq .ByteOrder "little" }}true{{else}}false{{end}}));
  offset +=2; {{- else if eq .Type "uint32" }} {{.Name}}Array.push(view.getUint32(offset, {{if eq .ByteOrder "little"
  }}true{{else}}false{{end}})); offset +=4; {{- else if eq .Type "int8" }} {{.Name}}Array.push(view.getInt8(offset));
  offset +=1; {{- else if eq .Type "int16" }} {{.Name}}Array.push(view.getInt16(offset, {{if eq .ByteOrder "little"
  }}true{{else}}false{{end}})); offset +=2; {{- else if eq .Type "int32" }} {{.Name}}Array.push(view.getInt32(offset,
  {{if eq .ByteOrder "little" }}true{{else}}false{{end}})); offset +=4; {{- else if eq .Type "float" }}
  {{.Name}}Array.push(view.getFloat32(offset, {{if eq .ByteOrder "little" }}true{{else}}false{{end}})); offset +=4; {{-
  else if eq .Type "double" }} {{.Name}}Array.push(view.getFloat64(offset, {{if eq .ByteOrder "little"
  }}true{{else}}false{{end}})); offset +=8; {{- else if eq .Type "bool" }} {{.Name}}Array.push(view.getUint8(offset)
  !==0); offset +=1; {{- end}} } result.{{.Name}}={{.Name}}Array; {{- else}} {{- if eq .Type "uint8" }}
  result.{{.Name}}=view.getUint8(offset); offset +=1; {{- else if eq .Type "uint16" }}
  result.{{.Name}}=view.getUint16(offset, {{if eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=2; {{- else
  if eq .Type "uint32" }} result.{{.Name}}=view.getUint32(offset, {{if eq .ByteOrder "little"
  }}true{{else}}false{{end}}); offset +=4; {{- else if eq .Type "int8" }} result.{{.Name}}=view.getInt8(offset); offset
  +=1; {{- else if eq .Type "int16" }} result.{{.Name}}=view.getInt16(offset, {{if eq .ByteOrder "little"
  }}true{{else}}false{{end}}); offset +=2; {{- else if eq .Type "int32" }} result.{{.Name}}=view.getInt32(offset, {{if
  eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=4; {{- else if eq .Type "float" }}
  result.{{.Name}}=view.getFloat32(offset, {{if eq .ByteOrder "little" }}true{{else}}false{{end}}); offset +=4; {{- else
  if eq .Type "double" }} result.{{.Name}}=view.getFloat64(offset, {{if eq .ByteOrder "little"
  }}true{{else}}false{{end}}); offset +=8; {{- else if eq .Type "bool" }} result.{{.Name}}=view.getUint8(offset) !==0;
  offset +=1; {{- else if eq .Type "string" }} // String deserialization not fully implemented result.{{.Name}}="" ;
  offset +=4; // Skip string length placeholder {{- end}} {{- end}} {{- end}} return result; } `))
