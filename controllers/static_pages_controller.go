package controllers

import "github.com/gin-gonic/gin"

type StaticPagesController struct{}

func (con StaticPagesController) Index(context *gin.Context) {
	context.JSON(200, gin.H{"message": "pong"})
	return
}
