// 代码生成时间: 2025-11-03 13:01:58
package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// BatchRenameRequest 定义批量重命名的请求结构
type BatchRenameRequest struct {
    SourceDir string `json:"sourceDir"` // 源目录
    RenameMap map[string]string `json:"renameMap"` // 重命名映射
}

// BatchRenameResponse 定义批量重命名的响应结构
type BatchRenameResponse struct {
    SuccessCount int `json:"successCount"` // 成功重命名的文件数量
    FailedCount  int `json:"failedCount"`  // 重命名失败的文件数量
}

func main() {
    router := gin.Default()

    router.POST("/rename", func(c *gin.Context) {
        var req BatchRenameRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        resp := BatchRenameResponse{SuccessCount: 0, FailedCount: 0}

        for oldName, newName := range req.RenameMap {
            oldPath := filepath.Join(req.SourceDir, oldName)
            newPath := filepath.Join(req.SourceDir, newName)
            if err := os.Rename(oldPath, newPath); err != nil {
                resp.FailedCount++
                c.Error(err)
            } else {
                resp.SuccessCount++
            }
        }

        c.JSON(http.StatusOK, resp)
    })

    router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

// 请注意，这个示例代码没有处理文件系统权限错误，也没有处理文件不存在的情况。
// 在实际应用中，应该添加更多的错误处理和验证逻辑。