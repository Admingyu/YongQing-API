package controller

import (
	"YongQing-API/database"
	"YongQing-API/errors"
	"YongQing-API/model"
	"YongQing-API/schema"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterUser(rg *gin.RouterGroup) {
	r := rg.Group("/user")
	r.POST("", UserOpen)
}

func UserOpen(c *gin.Context) {
	var params schema.UserCodeSchema
	err := c.ShouldBindQuery(&params)
	errors.HandleError(err)
	// openID, sessionKey := utils.CodeToJSession(params.Code)
	openID, sessionKey := "oOERN5LYuF31nmjRok-ijTAHgEko", ""
	log.Println(openID, sessionKey)
	if openID == "" {
		Fail(c, "INVALID_CODE")
		return
	}
	err = database.DB.FirstOrCreate(&model.User{OpenID: openID}, &model.User{OpenID: openID}).Error
	errors.HandleError(err)
	Success(c, map[string]interface{}{"openID": openID})
}
