package router

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()
	defineRoutes(r)
	r.Run()
}
