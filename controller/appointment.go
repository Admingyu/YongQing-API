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
	r.PUT("", middleware.UserLogin, CancelAppointment)
	r.GET("/agreement", Agreement)
	r.POST("", middleware.UserLogin, MakeAppointment)
}

func Agreement(c *gin.Context) {
	data := `<div>1、您可预约本年度内的空闲日期。</div>
    <div>2、24小时开放预约，但工作人员回复时间不一定（一般两小时内回复）但不超过当日0点。</div>
    <div>3、预约采取非实名制注册，但请正确填写相关信息，提交后，信息无法更改，注册人为交易双方其中一方。请您如实填写注册信息。</div>`
	Success(c, data)
}

func MyAppointment(c *gin.Context) {
	UserID, exist := c.Get("userID")
	if !exist {
		log.Println("not exist")
	}
	var data []serialization.AppointmentSer
	err := database.DB.Model(model.Appointment{}).Where("user_id=?", UserID).Order("created_at desc,id desc").Scan(&data).Error
	errors.HandleError(err)
	Success(c, data)
}

func MakeAppointment(c *gin.Context) {
	UserID, exist := c.Get("userID")
	if !exist {
		log.Println("not exist")
	}
	var params schema.AppointmentSchema
	err := c.ShouldBindJSON(&params)
	errors.HandleError(err)
	err = database.DB.Model(model.Appointment{}).Create(&model.Appointment{UserID: UserID.(int), Name: params.Name, Date: params.Date, Location: params.Location, Phone: params.Phone, Memo: params.Memo, Project: params.Project}).Error
	errors.HandleError(err)
	Success(c, nil)
}

func CancelAppointment(c *gin.Context) {
	UserID, exist := c.Get("userID")
	if !exist {
		log.Println("not exist")
	}
	var params schema.IDSchema
	err := c.ShouldBindJSON(&params)
	errors.HandleError(err)
	err = database.DB.Model(model.Appointment{}).Where("user_id=? AND id=?", UserID, params.ID).Updates(map[string]interface{}{"status": "CANCELED"}).Error
	errors.HandleError(err)
	Success(c, nil)
}
