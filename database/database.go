package database

import (
	. "YongQing-API/config"
	"YongQing-API/errors"
	"YongQing-API/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_ADDR, DB_PORT, DB_NAME)
	conn := mysql.Open(dataSource)
	config := gorm.Config{}
	var err error
	DB, err = gorm.Open(conn, &config)
	errors.HandleError(err)

	DB.AutoMigrate(model.Appointment{}, model.Case{}, model.Category{}, model.Slides{})

}
