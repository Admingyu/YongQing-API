package model

type Appointment struct {
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Date     string `json:"name" gorm:"type:varchar(500)`
	Location string `json:"name" gorm:"type:varchat(200)"`
	Phone    string `json:"phone" gorm:"type:varchar(20)"`
	Memo     string `json:"memo" gorm:"type:varchar(500)"`
}
