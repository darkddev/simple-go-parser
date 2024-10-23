package main

import (
	"go-parser/routes"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                      // Allow all origins
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"} // Allow GET, POST, OPTIONS methods
	config.AllowHeaders = []string{"Origin", "Content-Type"} // Allow Origin and Content-Type headers
	r.Use(cors.New(config))
	r.MaxMultipartMemory = 512 << 20 // 8 MiB
	// Define a route and its handler
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// Define a POST route and its handler
	r.POST("/amazon", routes.Amazon_PostRequest)
	r.POST("/google", routes.Google_PostRequest)
	r.POST("/walmart", routes.Walmart_PostRequest)

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 16 << 20, // 1 MiB
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
