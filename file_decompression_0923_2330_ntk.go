// 代码生成时间: 2025-09-23 23:30:52
package main

import (
    "compress/zip"
    "net/http"
    "os"
    "path"
    "strings"

    "github.com/gin-gonic/gin"
)

// DecompressHandler is the Gin handler function for decompressing files.
func DecompressHandler(c *gin.Context) {
    // Check if the file is in the request
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{'error': 'No file part found in the request'})
        return
    }

    // Open the file
    src, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{'error': 'Unable to open the file'})
        return
    }
    defer src.Close()

    // Create a directory to hold the extracted files
    destDir := path.Join(os.TempDir(), "decompressed")
    if _, err := os.Stat(destDir); os.IsNotExist(err) {
        err = os.MkdirAll(destDir, 0755)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{'error': 'Unable to create destination directory'})
            return
        }
    }

    // Decompress the file
    err = unzip(src, destDir)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{'error': 'Decompression failed'})
        return
    }

    // Return a success response
    c.JSON(http.StatusOK, gin.H{'message': 'File decompressed successfully'})
}

// unzip is a utility function to decompress zip files.
func unzip(src *os.File, dest string) error {
    reader, err := zip.OpenReader(src.Name())
    if err != nil {
        return err
    }
    defer reader.Close()

    for _, file := range reader.File {
        path := path.Join(dest, file.Name)
        if file.FileInfo().IsDir() {
            os.MkdirAll(path, 0755)
            continue
        }

        // Create the directory structure
        if _, err := os.Stat(path[0:len(path)-len(path.Base())]); os.IsNotExist(err) {
            os.MkdirAll(path[0:len(path)-len(path.Base())], 0755)
        }

        // Open the file to write
        outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
        if err != nil {
            return err
        }
        defer outFile.Close()

        // Copy the contents to the new file
        rc, err := file.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        _, err = copyData(outFile, rc)
        if err != nil {
            return err
        }
    }
    return nil
}

// copyData copies data from src to dst until either EOF is reached on src or an error occurs.
func copyData(dst *os.File, src io.Reader) (int64, error) {
    return io.Copy(dst, src)
}

func main() {
    r := gin.Default()

    // Define the route for file decompression
    r.POST("/decompress", DecompressHandler)

    // Start the server
    r.Run() // listening and serving on 0.0.0.0:8080
}
