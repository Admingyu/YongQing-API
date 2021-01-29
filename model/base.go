package model

import (
	"time"
)

type BasicModel struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6)"`
}
