package controller

import (
	"YongQing-API/database"
	"YongQing-API/errors"
	"YongQing-API/middleware"
	"YongQing-API/model"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterAppointment(rg *gin.RouterGroup) {
	r := rg.Group("/appointment")
	r.GET("", middleware.UserLogin, MyAppointment)
}

func MyAppointment(c *gin.Context) {
	UserID, exist := c.Get("userID")
	if !exist {
		log.Println("not exist")
	}
	var data []model.Appointment
	err := database.DB.Model(model.Appointment{}).Where("user_id=?", UserID).Scan(&data).Error
	errors.HandleError(err)
	Success(c, data)
}
