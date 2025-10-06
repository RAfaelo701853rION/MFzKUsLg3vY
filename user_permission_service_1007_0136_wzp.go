// 代码生成时间: 2025-10-07 01:36:20
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "log"
)

// ErrorResponse is a structure that represents an error response
type ErrorResponse struct {
    Error string `json:"error"`
}

// UserPermission represents the user permissions data structure
type UserPermission struct {
    UserID   int    `json:"user_id"`
    RoleName string `json:"role_name"`
}

// NewRouter creates a new Gin router and sets up the routes
func NewRouter() *gin.Engine {
    router := gin.Default()

    // Middleware to handle panics
    router.Use(gin.Recovery())

    // Routes
    router.GET("/permissions", GetPermissions)
    router.POST("/permissions", CreatePermission)

    return router
}

// GetPermissions handles the GET requests to retrieve user permissions
func GetPermissions(c *gin.Context) {
    // Simulate fetching user permissions from a database
    // In a real-world scenario, you would replace this with actual DB calls
    permissions := []UserPermission{
        {UserID: 1, RoleName: "admin"},
        {UserID: 2, RoleName: "user"},
    }

    c.JSON(http.StatusOK, permissions)
}

// CreatePermission handles the POST requests to create a new user permission
func CreatePermission(c *gin.Context) {
    var permission UserPermission
    if err := c.ShouldBindJSON(&permission); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
        return
    }
    // In a real-world scenario, you would save the permission to a database
    // and return a success status with the created permission
    // Simulating a successful save
    c.JSON(http.StatusCreated, permission)
}

// Main function to run the HTTP server
func main() {
    router := NewRouter()
    log.Fatal(router.Run(":8080"))
}
