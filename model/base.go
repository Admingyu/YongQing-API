package model

import (
	"fmt"
	"time"
)

type BasicModel struct {
	ID        int      `json:"id" gorm:"primary_key"`
	CreatedAt jsonTime `json:"created_at" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)"`
	UpdatedAt jsonTime `json:"updated_at" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6)"`
}

type jsonTime time.Time

//实现它的json序列化方法
func (this jsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type jsonDate time.Time

//实现它的json序列化方法
func (this jsonDate) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02"))
	return []byte(stamp), nil
}