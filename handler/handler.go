package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "GET opening",
	})
}

func FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "GET opening",
	})
}

func Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "POST opening",
	})
}

func Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "PUT opening",
	})
}

func Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "DELETE opening",
	})
}
