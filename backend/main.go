package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/match", func(c *gin.Context) {
		var body map[string]string
		c.BindJSON(&body)

		payload, _ := json.Marshal(body)
		resp, _ := http.Post("http://localhost:8000/analyze",
			"application/json",
			bytes.NewBuffer(payload))

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		c.JSON(200, result)
	})

	r.Run(":8181")
}
