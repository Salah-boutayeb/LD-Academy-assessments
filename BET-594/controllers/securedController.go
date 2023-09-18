package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
func TestContoller(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
