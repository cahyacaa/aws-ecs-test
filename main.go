package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("ENV Failed to initialized")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/testing/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hi %s!, Congratulations You Did It Works!!", name),
		})
	})
	r.GET("/greeting/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %s! Changes At %v changes", name, time.Now().Format(time.RFC3339)),
		})
	})
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
