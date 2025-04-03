package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) // Enable CORS for frontend access

	// Define routes
	e.GET("/schema.json", serveFile("../schema.json"))
	e.GET("/adcs.json", serveFile("../adcs.json"))
	e.GET("/eps.json", serveFile("../eps.json"))

	// Add route to get all generated files
	e.GET("/generated", getGeneratedFiles)
	e.GET("/generated/:filename", serveGeneratedFile)
	// Add API endpoint to generate code (future extension)
	e.POST("/generate", generateCode)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}

// serveFile returns a handler that serves the specified file
func serveFile(filename string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Try opening the file from the current directory
		data, err := os.ReadFile(filename)
		if err != nil {
			return c.String(http.StatusNotFound, "File not found")
		}

		return c.Blob(http.StatusOK, "application/json", data)
	}
}

// generateCode is a placeholder for future code generation functionality
func generateCode(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "Not implemented yet",
	})
}

// getGeneratedFiles returns information about all files in the generated directory
func getGeneratedFiles(c echo.Context) error {
	files, err := os.ReadDir("../generated")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to read generated directory: " + err.Error(),
		})
	}

	type FileInfo struct {
		Name string `json:"name"`
		Type string `json:"type"` // "c", "h", or "ts"
		Size int64  `json:"size"`
	}

	var fileList []FileInfo
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		fileType := ""
		switch ext {
		case ".c":
			fileType = "c"
		case ".h":
			fileType = "h"
		case ".ts":
			fileType = "ts"
		default:
			continue // Skip files with unsupported extensions
		}

		fileList = append(fileList, FileInfo{
			Name: file.Name(),
			Type: fileType,
			Size: file.Size(),
		})
	}

	return c.JSON(http.StatusOK, fileList)
}

// serveGeneratedFile serves a specific file from the generated directory
func serveGeneratedFile(c echo.Context) error {
	filename := c.Param("filename")

	// Basic security check - don't allow path traversal
	if strings.Contains(filename, "..") {
		return c.String(http.StatusBadRequest, "Invalid filename")
	}

	filepath := filepath.Join("../generated", filename)
	data, err := os.ReadFile(filepath)
	if err != nil {
		return c.String(http.StatusNotFound, "File not found")
	}

	contentType := "text/plain"
	if strings.HasSuffix(filename, ".json") {
		contentType = "application/json"
	}

	return c.Blob(http.StatusOK, contentType, data)
}
