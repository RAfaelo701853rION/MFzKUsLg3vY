// 代码生成时间: 2025-10-11 02:46:24
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "runtime"
    "time"

    "github.com/gin-gonic/gin"
)

// DiskSpaceManager provides disk space management functionalities.
type DiskSpaceManager struct{
    // Embed a logger to provide logging functionalities.
    Logger *log.Logger
}

// NewDiskSpaceManager initializes a new DiskSpaceManager with a logger.
func NewDiskSpaceManager() *DiskSpaceManager {
    return &DiskSpaceManager{
        Logger: log.New(os.Stdout, "DISK_SPACE_MANAGER: ", log.LstdFlags),
    }
}

// GetDiskUsage gets the disk usage statistics.
// It returns the total, used, free, and used percentage of the disk.
func (dsm *DiskSpaceManager) GetDiskUsage(c *gin.Context) {
    var usage DiskUsage
    var err error

    // Get disk usage information for the current working directory.
    // You may change the path to get usage of a specific directory.
    usage, err = getDiskUsage(".")
    if err != nil {
        dsm.Logger.Printf("Error getting disk usage: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to get disk usage.",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "total":  usage.Total,
        "used":   usage.Used,
        