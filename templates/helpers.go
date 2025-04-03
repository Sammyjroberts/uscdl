package templates

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

// Item represents a property within a container
type Item struct {
	Name        string
	Type        string
	Description string
	ByteOrder   string
	Units       string
	IsArray     bool
	Length      int
}

// Container represents a struct that contains multiple items
type Container struct {
	Name        string
	Description string
	Items       []Item
}

// CTypeMapping maps JSON types to C types
var CTypeMapping = map[string]string{
	"uint8":  "uint8_t",
	"uint16": "uint16_t",
	"uint32": "uint32_t",
	"uint64": "uint64_t",
	"int8":   "int8_t",
	"int16":  "int16_t",
	"int32":  "int32_t",
	"int64":  "int64_t",
	"float":  "float",
	"double": "double",
	"bool":   "bool",
	"string": "char*",
}

// TSTypeMapping maps JSON types to TypeScript types
var TSTypeMapping = map[string]string{
	"uint8":  "number",
	"uint16": "number",
	"uint32": "number",
	"uint64": "number",
	"int8":   "number",
	"int16":  "number",
	"int32":  "number",
	"int64":  "number",
	"float":  "number",
	"double": "number",
	"bool":   "boolean",
	"string": "string",
}

// GetCType returns the C type for a given item
func GetCType(item Item) string {
	cType, ok := CTypeMapping[item.Type]
	if !ok {
		return "void"
	}
	return cType
}

// GetTSType returns the TypeScript type for a given item
func GetTSType(item Item) string {
	tsType, ok := TSTypeMapping[item.Type]
	if !ok {
		return "any"
	}

	if item.IsArray {
		return fmt.Sprintf("%s[]", tsType)
	}

	return tsType
}

// GetDefaultValueC returns the default value for a C type
func GetDefaultValueC(itemType string) string {
	switch itemType {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
		return "0"
	case "float", "double":
		return "0.0"
	case "bool":
		return "false"
	case "string":
		return "NULL"
	default:
		return "0"
	}
}

// GetDefaultValueTS returns the default value for a TypeScript type and item
func GetDefaultValueTS(item Item) string {
	if item.IsArray {
		switch item.Type {
		case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float", "double":
			return fmt.Sprintf("Array(%d).fill(0)", item.Length)
		case "bool":
			return fmt.Sprintf("Array(%d).fill(false)", item.Length)
		case "string":
			return fmt.Sprintf("Array(%d).fill('')", item.Length)
		default:
			return "[]"
		}
	}

	switch item.Type {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float", "double":
		return "0"
	case "bool":
		return "false"
	case "string":
		return "''"
	default:
		return "null"
	}
}

// GetTypeSizeC returns the size in bytes of a C type
func GetTypeSizeC(itemType string) string {
	switch itemType {
	case "uint8", "int8", "bool":
		return "1"
	case "uint16", "int16":
		return "2"
	case "uint32", "int32", "float":
		return "4"
	case "uint64", "int64", "double":
		return "8"
	case "string":
		return "sizeof(char*)"
	default:
		return "1"
	}
}

// NeedsByteSwap determines if a type needs byte swapping based on endianness
func NeedsByteSwap(item Item) bool {
	if item.ByteOrder == "big" {
		return true
	}
	return false
}

// CalculateStructSize calculates the approximate size of a struct in bytes
func CalculateStructSize(container Container) string {
	size := 0
	for _, item := range container.Items {
		var itemSize int
		switch item.Type {
		case "uint8", "int8", "bool":
			itemSize = 1
		case "uint16", "int16":
			itemSize = 2
		case "uint32", "int32", "float":
			itemSize = 4
		case "uint64", "int64", "double":
			itemSize = 8
		case "string":
			itemSize = 4 // Pointer size on 32-bit systems
		}

		if item.IsArray {
			itemSize *= item.Length
		}

		size += itemSize
	}
	return fmt.Sprintf("%d", size)
}

// ByteOrderFunctionsNeeded checks if any items in the container need byte swapping
func ByteOrderFunctionsNeeded(container Container) bool {
	for _, item := range container.Items {
		if item.ByteOrder == "big" &&
			(item.Type == "uint16" || item.Type == "uint32" || item.Type == "uint64" ||
				item.Type == "int16" || item.Type == "int32" || item.Type == "int64" ||
				item.Type == "float" || item.Type == "double") {
			return true
		}
	}
	return false
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	return strcase.ToSnake(s)
}

// ToCamelCase converts a string to CamelCase
func ToCamelCase(s string) string {
	words := strings.Split(s, "_")
	for i := range words {
		if i == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = strings.Title(strings.ToLower(words[i]))
		}
	}
	return strings.Join(words, "")
}
