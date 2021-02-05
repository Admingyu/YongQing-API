package controller

import (
	"YongQing-API/database"
	"YongQing-API/errors"
	"YongQing-API/middleware"
	"YongQing-API/model"
	"YongQing-API/schema"
	"YongQing-API/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterUser(rg *gin.RouterGroup) {
	r := rg.Group("/user")
	r.GET("/code", UserOpen)
	r.POST("/info", middleware.UserLogin, UserInfo)
}

func UserOpen(c *gin.Context) {
	var params schema.UserCodeSchema
	err := c.ShouldBindQuery(&params)
	errors.HandleError(err)
	// openID, sessionKey := utils.CodeToJSession(params.Code)
	openID, sessionKey := utils.CodeToJSession(params.Code)
	log.Println(openID, sessionKey)
	if openID == "" {
		Fail(c, "INVALID_CODE")
		return
	}
	err = database.DB.FirstOrCreate(&model.User{OpenID: openID}, &model.User{OpenID: openID}).Error
	errors.HandleError(err)
	Success(c, map[string]interface{}{"openID": openID})
}

func UserInfo(c *gin.Context) {
	var params schema.UserInfoSchema
	err := c.ShouldBindJSON(&params)
	errors.HandleError(err)
	userID, exist := c.Get("userID")
	if !exist {
		log.Println("not exist")
	}
	userUp := model.User{
		NickName:  params.NickName,
		AvatarUrl: params.AvatarUrl,
		Gender:    params.Gender,
		Province:  params.Province,
		City:      params.City,
		Language:  params.Language,
		Country:   params.Country,
	}
	err = database.DB.Model(&model.User{}).Where("id=?", userID).Updates(&userUp).Error
	errors.HandleError(err)

	Success(c, nil)
}
