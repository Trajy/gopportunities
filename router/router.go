package router

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()
	defineRoutes(r)
	r.Run()
}

func defineRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})
}
