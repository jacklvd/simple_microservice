package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	// Category routes
	CategoryRoutes(router)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
