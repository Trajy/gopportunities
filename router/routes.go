package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func defineRoutes(r *gin.Engine) {
	rv1 := r.Group("/api/v1")
	{
		rv1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"response": "GET opening",
			})
		})
		rv1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"response": "GET opening",
			})
		})
		rv1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"response": "POST opening",
			})
		})
		rv1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"response": "PUT opening",
			})
		})
		rv1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"response": "DELETE opening",
			})
		})
	}
}
