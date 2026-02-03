package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ✅ CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check (optional but good)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Backend running"})
	})

	r.POST("/match", func(c *gin.Context) {
		var body map[string]string
		c.BindJSON(&body)

		payload, _ := json.Marshal(body)

		resp, err := http.Post(
			"http://localhost:8000/analyze",
			"application/json",
			bytes.NewBuffer(payload),
		)
		if err != nil {
			c.JSON(500, gin.H{"error": "AI service unavailable"})
			return
		}

		defer resp.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		c.JSON(200, result)
	})

	// 🔁 Use 8181 (as per your logs)
	r.Run(":8181")
}
