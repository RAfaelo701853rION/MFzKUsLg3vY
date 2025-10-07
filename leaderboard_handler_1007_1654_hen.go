// 代码生成时间: 2025-10-07 16:54:48
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "encoding/json"
    "fmt"
)

// Leaderboard represents the structure of the leaderboard data
type Leaderboard struct {
    // Id is a unique identifier for each entry
    Id int `json:"id"`
    // Name is the name of the player
    Name string `json:"name"`
    // Score is the player's score
    Score int `json:"score"`
}

// leaderboardData represents the data store for the leaderboard
var leaderboardData = []Leaderboard{
    {Id: 1, Name: "Alice", Score: 250},
    {Id: 2, Name: "Bob", Score: 180},
    {Id: 3, Name: "Charlie", Score: 300},
    // Add more entries as needed
}

// GetLeaderboard handles GET requests to retrieve the leaderboard
func GetLeaderboard(c *gin.Context) {
    // Return the leaderboard data in JSON format
    c.JSON(http.StatusOK, leaderboardData)
}

// AddEntry handles POST requests to add a new entry to the leaderboard
func AddEntry(c *gin.Context) {
    var newEntry Leaderboard
    if err := c.ShouldBindJSON(&newEntry); err != nil {
        // Handle binding error
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("Invalid input: %s", err.Error()),
        })
        return
    }
    // Add the new entry to the leaderboard
    leaderboardData = append(leaderboardData, newEntry)
    c.JSON(http.StatusCreated, gin.H{
        "message": "Entry added successfully",
    })
}

// main is the entry point of the application
func main() {
    r := gin.Default()
    // Use Gin middleware to log requests
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // Define the routes
    r.GET("/leaderboard", GetLeaderboard)
    r.POST("/leaderboard", AddEntry)

    // Start the server on port 8080
    r.Run(":8080")
}
