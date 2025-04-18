package router

import (
	"github.com/Trajy/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func defineRoutes(r *gin.Engine) {
	rv1 := r.Group("/api/v1")
	{
		rv1.GET("/opening", handler.Find)
		rv1.GET("/opening", handler.FindById)
		rv1.POST("/opening", handler.Create)
		rv1.PUT("/opening", handler.Update)
		rv1.DELETE("/opening", handler.Delete)
	}
}
