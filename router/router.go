package router

import (
	"YongQing-API/controller"

	"github.com/gin-gonic/gin"
)

var GinEngine *gin.Engine

func init() {
	GinEngine = gin.New()
	GinEngine.NoRoute()
	GinEngine.Use(gin.Logger())
	api := GinEngine.Group("/api")
	controller.RegisterAppointment(api)
	controller.RegisterUser(api)
}
