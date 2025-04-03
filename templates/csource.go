package templates

import (
	"text/template"
)

// CSourceTemplate generates a C source file implementation with unique variable names
var CSourceTemplate = template.Must(template.New("csource").Funcs(templateFuncs).Parse(`/**
* {{.Name}}
* {{.Description}}
*/

#include "{{.Name | ToLower}}.h"
#include <string.h>
#include <stdlib.h>

void {{.Name | ToSnakeCase}}_init({{.Name}}_t* p_data) {
    if (p_data == NULL) {
        return;
    }

    {{- range .Items}}
    {{- if .IsArray}}
    memset(p_data->{{.Name}}, {{GetDefaultValueC .Type}}, sizeof(p_data->{{.Name}}));
    {{- else}}
    p_data->{{.Name}} = {{GetDefaultValueC .Type}};
    {{- end}}
    {{- end}}
}

int {{.Name | ToSnakeCase}}_serialize(const {{.Name}}_t* p_data, uint8_t* buffer, size_t buffer_size) {
    if (p_data == NULL || buffer == NULL) {
        return -1;
    }

    // Ensure buffer is large enough
    if (buffer_size < {{CalculateStructSize .}}) {
        return -1;
    }

    size_t offset = 0;
    uint8_t* ptr = buffer;
    size_t item_size = 0;

    {{- range $itemIndex, $item := .Items}}
    {{- if .IsArray}}
    {{- if eq .Type "string"}}
    // String arrays not supported in this simple implementation
    {{- else if NeedsByteSwap .}}
    // Copy {{.Name}} array with byte swapping
    for (int i = 0; i < {{.Length}}; i++) {
        switch ({{.Type}}) {
        case "uint16", "int16":
            {
                uint16_t temp = SWAP_UINT16((uint16_t)p_data->{{.Name}}[i]);
                memcpy(ptr + offset, &temp, sizeof(uint16_t));
                offset += sizeof(uint16_t);
            }
            break;
        case "uint32", "int32":
            {
                uint32_t temp = SWAP_UINT32((uint32_t)p_data->{{.Name}}[i]);
                memcpy(ptr + offset, &temp, sizeof(uint32_t));
                offset += sizeof(uint32_t);
            }
            break;
        case "uint64", "int64":
            {
                uint64_t temp = SWAP_UINT64((uint64_t)p_data->{{.Name}}[i]);
                memcpy(ptr + offset, &temp, sizeof(uint64_t));
                offset += sizeof(uint64_t);
            }
            break;
        case "float":
            {
                float temp = swap_float(p_data->{{.Name}}[i]);
                memcpy(ptr + offset, &temp, sizeof(float));
                offset += sizeof(float);
            }
            break;
        case "double":
            {
                double temp = swap_double(p_data->{{.Name}}[i]);
                memcpy(ptr + offset, &temp, sizeof(double));
                offset += sizeof(double);
            }
            break;
        }
    }
    {{- else}}
    // Direct copy for little-endian or byte types
    item_size = {{GetTypeSizeC .Type}} * {{.Length}};
    memcpy(ptr + offset, p_data->{{.Name}}, item_size);
    offset += item_size;
    {{- end}}
    {{- else}}
    {{- if eq .Type "string"}}
    // For non-array strings, assume null-terminated and copy until null or buffer exhausted
    if (p_data->{{.Name}} != NULL) {
        size_t str_len = strlen(p_data->{{.Name}});
        if (offset + str_len + 1 > buffer_size) {
            return -1;
        }
        memcpy(ptr + offset, p_data->{{.Name}}, str_len + 1);
        offset += str_len + 1;
    } else {
        // Just write a null byte for NULL strings
        *(ptr + offset) = '\0';
        offset++;
    }
    {{- else if NeedsByteSwap .}}
    // Handle byte swapping for multi-byte scalar values
    switch ({{.Type}}) {
    case "uint16", "int16":
        {
            uint16_t temp = SWAP_UINT16((uint16_t)p_data->{{.Name}});
            memcpy(ptr + offset, &temp, sizeof(uint16_t));
            offset += sizeof(uint16_t);
        }
        break;
    case "uint32", "int32":
        {
            uint32_t temp = SWAP_UINT32((uint32_t)p_data->{{.Name}});
            memcpy(ptr + offset, &temp, sizeof(uint32_t));
            offset += sizeof(uint32_t);
        }
        break;
    case "uint64", "int64":
        {
            uint64_t temp = SWAP_UINT64((uint64_t)p_data->{{.Name}});
            memcpy(ptr + offset, &temp, sizeof(uint64_t));
            offset += sizeof(uint64_t);
        }
        break;
    case "float":
        {
            float temp = swap_float(p_data->{{.Name}});
            memcpy(ptr + offset, &temp, sizeof(float));
            offset += sizeof(float);
        }
        break;
    case "double":
        {
            double temp = swap_double(p_data->{{.Name}});
            memcpy(ptr + offset, &temp, sizeof(double));
            offset += sizeof(double);
        }
        break;
    }
    {{- else}}
    // Direct copy for little-endian or byte types
    memcpy(ptr + offset, &p_data->{{.Name}}, {{GetTypeSizeC .Type}});
    offset += {{GetTypeSizeC .Type}};
    {{- end}}
    {{- end}}
    {{- end}}

    return (int)offset;
}

int {{.Name | ToSnakeCase}}_deserialize({{.Name}}_t* p_data, const uint8_t* buffer, size_t buffer_size) {
    if (p_data == NULL || buffer == NULL) {
        return -1;
    }

    // Initialize the structure
    {{.Name | ToSnakeCase}}_init(p_data);

    size_t offset = 0;
    const uint8_t* ptr = buffer;
    size_t item_size = 0;

    {{- range $itemIndex, $item := .Items}}
    {{- if .IsArray}}
    {{- if eq .Type "string"}}
    // String arrays not supported in this simple implementation
    {{- else if NeedsByteSwap .}}
    // Copy array with byte swapping
    for (int i = 0; i < {{.Length}}; i++) {
        if (offset + {{GetTypeSizeC .Type}} > buffer_size) {
            return -1;
        }

        switch ({{.Type}}) {
        case "uint16", "int16":
            {
                uint16_t temp;
                memcpy(&temp, ptr + offset, sizeof(uint16_t));
                p_data->{{.Name}}[i] = SWAP_UINT16(temp);
                offset += sizeof(uint16_t);
            }
            break;
        case "uint32", "int32":
            {
                uint32_t temp;
                memcpy(&temp, ptr + offset, sizeof(uint32_t));
                p_data->{{.Name}}[i] = SWAP_UINT32(temp);
                offset += sizeof(uint32_t);
            }
            break;
        case "uint64", "int64":
            {
                uint64_t temp;
                memcpy(&temp, ptr + offset, sizeof(uint64_t));
                p_data->{{.Name}}[i] = SWAP_UINT64(temp);
                offset += sizeof(uint64_t);
            }
            break;
        case "float":
            {
                float temp;
                memcpy(&temp, ptr + offset, sizeof(float));
                p_data->{{.Name}}[i] = swap_float(temp);
                offset += sizeof(float);
            }
            break;
        case "double":
            {
                double temp;
                memcpy(&temp, ptr + offset, sizeof(double));
                p_data->{{.Name}}[i] = swap_double(temp);
                offset += sizeof(double);
            }
            break;
        }
    }
    {{- else}}
    // Direct copy for little-endian or byte types
    item_size = {{GetTypeSizeC .Type}} * {{.Length}};
    if (offset + item_size > buffer_size) {
        return -1;
    }
    memcpy(p_data->{{.Name}}, ptr + offset, item_size);
    offset += item_size;
    {{- end}}
    {{- else}}
    {{- if eq .Type "string"}}
    // For strings, we need to allocate memory and copy the string
    // This implementation doesn't handle strings - would need memory management
    // Just copy the pointer - assumes the buffer outlives the structure
    p_data->{{.Name}} = (char*)(ptr + offset);
    offset += strlen(p_data->{{.Name}}) + 1;
    {{- else if NeedsByteSwap .}}
    // Handle byte swapping for multi-byte values
    if (offset + {{GetTypeSizeC .Type}} > buffer_size) {
        return -1;
    }

    switch ({{.Type}}) {
    case "uint16", "int16":
        {
            uint16_t temp;
            memcpy(&temp, ptr + offset, sizeof(uint16_t));
            p_data->{{.Name}} = SWAP_UINT16(temp);
            offset += sizeof(uint16_t);
        }
        break;
    case "uint32", "int32":
        {
            uint32_t temp;
            memcpy(&temp, ptr + offset, sizeof(uint32_t));
            p_data->{{.Name}} = SWAP_UINT32(temp);
            offset += sizeof(uint32_t);
        }
        break;
    case "uint64", "int64":
        {
            uint64_t temp;
            memcpy(&temp, ptr + offset, sizeof(uint64_t));
            p_data->{{.Name}} = SWAP_UINT64(temp);
            offset += sizeof(uint64_t);
        }
        break;
    case "float":
        {
            float temp;
            memcpy(&temp, ptr + offset, sizeof(float));
            p_data->{{.Name}} = swap_float(temp);
            offset += sizeof(float);
        }
        break;
    case "double":
        {
            double temp;
            memcpy(&temp, ptr + offset, sizeof(double));
            p_data->{{.Name}} = swap_double(temp);
            offset += sizeof(double);
        }
        break;
    }
    {{- else}}
    // Direct copy for little-endian or byte types
    if (offset + {{GetTypeSizeC .Type}} > buffer_size) {
        return -1;
    }
    memcpy(&p_data->{{.Name}}, ptr + offset, {{GetTypeSizeC .Type}});
    offset += {{GetTypeSizeC .Type}};
    {{- end}}
    {{- end}}
    {{- end}}

    return (int)offset;
}
`))
