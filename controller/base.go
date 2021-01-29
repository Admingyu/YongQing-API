package controller

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"status": "SUCCESS", "message": nil, "data": data})
}

func Fail(c *gin.Context, msg string) {
	c.JSON(200, gin.H{"status": "SUCCESS", "message": msg, "data": nil})
}
