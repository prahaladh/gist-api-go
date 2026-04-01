package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Gist matches the GitHub API response structure I am using simply for demonstration. You can expand it with more fields if needed.
type Gist struct {
	URL         string `json:"html_url"`
	Description string `json:"description"`
}

// SetupRouter wraps the API logic so it can be accessed by both main() and tests
func SetupRouter() *gin.Engine {
	// 1. Initialize Gin router with default logging and recovery middleware
	r := gin.Default()

	// 2. Define the route with a named parameter
	r.GET("/:username", func(c *gin.Context) {
		username := c.Param("username")

		// 3. Fetch from GitHub API with a timeout to avoid hanging requests
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get(fmt.Sprintf("https://api.github.com/users/%s/gists", username))

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to reach GitHub"})
			return
		}
		defer resp.Body.Close()

		// Handle non-200 responses (like 404 if user doesn't exist)
		if resp.StatusCode != http.StatusOK {
			c.JSON(resp.StatusCode, gin.H{"error": "GitHub returned an error"})
			return
		}

		// 4. Decode the data
		var gists []Gist
		if err := json.NewDecoder(resp.Body).Decode(&gists); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
			return
		}

		// 5. Gin's built-in JSON responder
		c.JSON(http.StatusOK, gists)
	})

	return r
}

func main() {
	r := SetupRouter()
	// Run the server on port 8080
	r.Run(":8080")
}