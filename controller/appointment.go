package controller

import (
	"YongQing-API/database"
	"YongQing-API/errors"
	"YongQing-API/middleware"
	"YongQing-API/model"
	"YongQing-API/serialization"
	"log"

	"YongQing-API/schema"

	"github.com/gin-gonic/gin"
)

func RegisterAppointment(rg *gin.RouterGroup) {
	r := rg.Group("/appointment")
	r.GET("", middleware.UserLogin, MyAppointment)
	r.POST("", middleware.UserLogin, MakeAppointment)
}

func MyAppointment(c *gin.Context) {
	UserID, exist := c.Get("userID")
	if !exist {
		log.Println("not exist")
	}
	var data []serialization.AppointmentSer
	err := database.DB.Model(model.Appointment{}).Where("user_id=?", UserID).Scan(&data).Error
	errors.HandleError(err)
	Success(c, data)
}

func MakeAppointment(c *gin.Context) {
	UserID, exist := c.Get("userID")
	if !exist {
		log.Println("not exist")
	}
	var params schema.AppointmentSchema
	c.ShouldBindJSON(&params)
	err := database.DB.Model(model.Appointment{}).Create(&model.Appointment{UserID: UserID.(int), Name: params.Name, Date: params.Date, Location: params.Location, Phone: params.Phone, Memo: params.Memo, Project: params.Project}).Error
	errors.HandleError(err)
	Success(c, nil)
}
