// 代码生成时间: 2025-09-23 04:37:48
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "gopkg.in/gorilla/schema.v2"
)

// CSVProcessorHandler processes multipart/form-data requests containing CSV files.
func CSVProcessorHandler(c *gin.Context) {
    // Check the existence of file in the request
    if _, err := c.GetRawData(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No file found in the request.",
        })
        return
    }

    // Create a form to parse the request
    form, err := c.MultipartForm()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("Error parsing multipart form: %v", err),
        })
        return
    }

    files := form.File["files"]
    if len(files) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            {
            "error": "No CSV files to process.",
        }})
        return
    }

    // Process each file
    for _, file := range files {
        // Open the file
        src, err := file.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": fmt.Sprintf("Failed to open file: %v", err),
            })
            return
        }
        defer src.Close()

        // Define the destination path for the file
        dst, err := os.Create(filepath.Join("./uploads", file.Filename))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": fmt.Sprintf("Failed to create destination file: %v", err),
            })
            return
        }
        defer dst.Close()

        // Copy the file content from source to destination
        if _, err := io.Copy(dst, src); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": fmt.Sprintf("Failed to copy file content: %v", err),
            })
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "CSV files processed successfully.",
    })
}

func main() {
    r := gin.Default()

    // Register the CSV file processor handler
    r.POST("/process-csv", CSVProcessorHandler)

    // Start the server on port 8080
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
