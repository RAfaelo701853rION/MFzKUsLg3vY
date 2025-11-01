// 代码生成时间: 2025-11-01 16:36:00
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// DuplicateFileDetector 结构体封装了文件系统路径和处理器逻辑
type DuplicateFileDetector struct {
    Path string
}

// NewDuplicateFileDetector 创建并返回一个新的DuplicateFileDetector实例
func NewDuplicateFileDetector(path string) *DuplicateFileDetector {
    return &DuplicateFileDetector{
        Path: path,
    }
}

// DetectDuplicates 扫描指定路径下的文件，并检测重复文件
func (d *DuplicateFileDetector) DetectDuplicates() ([]string, error) {
    files, err := os.ReadDir(d.Path)
    if err != nil {
        return nil, err
    }

    fileMap := make(map[string]string)
    duplicates := []string{}

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        filePath := filepath.Join(d.Path, file.Name())
        hash, err := fileHash(filePath)
        if err != nil {
            return nil, err
        }

        if existingPath, exists := fileMap[hash]; exists {
            duplicates = append(duplicates, existingPath, filePath)
        } else {
            fileMap[hash] = filePath
        }
    }

    return duplicates, nil
}

// fileHash 返回文件的内容哈希值
func fileHash(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    var result []byte

    // 使用简单的字符串哈希作为示例
    const multiplier = 37
    var hash int
    buffer := make([]byte, 1024)
    for {
        n, err := file.Read(buffer)
        if n > 0 {
            for i := 0; i < n; i++ {
                hash = hash*multiplier + int(buffer[i])
            }
        }
        if err != nil {
            if err != io.EOF {
                return "", err
            }
            break
        }
    }

    // 将哈希转换为字符串并返回
    return fmt.Sprintf("%d", hash), nil
}

// SetupRoutes 设置Gin路由和中间件
func SetupRoutes() *gin.Engine {
    r := gin.Default()
    r.Use(gin.Recovery()) // 自动恢复中间件，避免程序崩溃

    // 处理检测重复文件的请求
    r.POST("/detect", func(c *gin.Context) {
        var req struct {
            Path string `json:"path" binding:"required"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
            })
            return
        }

        detector := NewDuplicateFileDetector(req.Path)
        duplicates, err := detector.DetectDuplicates()
        if err != nil {
            c.JSON(500, gin.H{
                "error": "Failed to detect duplicates",
                "details": err.Error(),
            })
            return
        }

        if len(duplicates) == 0 {
            c.JSON(200, gin.H{
                "message": "No duplicates found",
            })
        } else {
            c.JSON(200, gin.H{
                "message": "Duplicates found",
                "duplicates": duplicates,
            })
        }
    })

    return r
}

func main() {
    r := SetupRoutes()
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
