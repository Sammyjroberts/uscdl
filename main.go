package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sammyjroberts/uscdl/templates"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

// Item represents a property within a container
type Item struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	ByteOrder   string `json:"byteOrder"`
	Units       string `json:"units"`
	IsArray     bool   `json:"isArray"`
	Length      int    `json:"length"`
}

// Container represents a struct that contains multiple items
type Container struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Items       []Item `json:"items"`
}

// Config represents the entire configuration
type Config struct {
	Containers []Container `json:"containers"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <config.json> [schema.json]")
	}

	configFile := os.Args[1]
	outputDir := "generated"
	schemaFile := "schema.json"

	if len(os.Args) >= 3 {
		schemaFile = os.Args[2]
	}

	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Read and parse config file
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Validate against JSON Schema
	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile(schemaFile)
	if err != nil {
		log.Fatalf("Failed to compile schema: %v", err)
	}

	var jsonData interface{}
	if err := json.Unmarshal(configData, &jsonData); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	if err := schema.Validate(jsonData); err != nil {
		log.Fatalf("Validation error: %v", err)
	}

	log.Println("Configuration is valid!")

	// Parse JSON into Config struct
	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Generate code for each container
	for _, container := range config.Containers {
		// Convert to template Container type
		tmplContainer := templates.Container{
			Name:        container.Name,
			Description: container.Description,
			Items:       make([]templates.Item, len(container.Items)),
		}

		for i, item := range container.Items {
			tmplContainer.Items[i] = templates.Item{
				Name:        item.Name,
				Type:        item.Type,
				Description: item.Description,
				ByteOrder:   item.ByteOrder,
				Units:       item.Units,
				IsArray:     item.IsArray,
				Length:      item.Length,
			}
		}

		// Generate C header file
		headerFile, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("%s.h", strings.ToLower(container.Name))))
		if err != nil {
			log.Fatalf("Failed to create header file: %v", err)
		}
		defer headerFile.Close()

		if err := templates.CHeaderTemplate.Execute(headerFile, tmplContainer); err != nil {
			log.Fatalf("Failed to render header template: %v", err)
		}
		fmt.Printf("Generated C header file: %s\n", headerFile.Name())

		// Generate C source file
		sourceFile, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("%s.c", strings.ToLower(container.Name))))
		if err != nil {
			log.Fatalf("Failed to create source file: %v", err)
		}
		defer sourceFile.Close()

		if err := templates.CSourceTemplate.Execute(sourceFile, tmplContainer); err != nil {
			log.Fatalf("Failed to render source template: %v", err)
		}
		fmt.Printf("Generated C source file: %s\n", sourceFile.Name())

		// Generate TypeScript file
		tsFile, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("%s.ts", container.Name)))
		if err != nil {
			log.Fatalf("Failed to create TypeScript file: %v", err)
		}
		defer tsFile.Close()

		if err := templates.TypeScriptTemplate.Execute(tsFile, tmplContainer); err != nil {
			log.Fatalf("Failed to render TypeScript template: %v", err)
		}
		fmt.Printf("Generated TypeScript file: %s\n", tsFile.Name())
	}

	fmt.Println("Code generation completed successfully!")
}
