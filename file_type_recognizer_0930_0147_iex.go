// 代码生成时间: 2025-09-30 01:47:21
package main

import (
    "fmt"
    "mime"
    "net/http"

    // Importing Gin framework
    "github.com/gin-gonic/gin"
)

// FileTypeRecognizer is a Gin.HandlerFunc that recognizes the file type of the uploaded file.
func FileTypeRecognizer(c *gin.Context) {
    // Check if a file was uploaded
    if _, header, err := c.Request.FormFile("file"); err != nil {
        // Handle error if file is not uploaded
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No file uploaded.",
        })
        return
    } else {
        // Get the MIME type of the file
        contentType := header.Header.Get("Content-Type")
        if mimeType, _, err := mime.ParseMediaType(contentType); err == nil {
            // Return the MIME type of the file
            c.JSON(http.StatusOK, gin.H{
                "filetype": mimeType,
            })
        } else {
            // Handle error if the MIME type cannot be determined
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Unable to determine file type.",
            })
            return
        }
    }
}

// main function to run the Gin web server
func main() {
    r := gin.Default()

    // Register the handler for the file type recognition
    r.POST("/recognize", FileTypeRecognizer)

    // Start the server on port 8080
    r.Run(":8080")
}
