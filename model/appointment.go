package model

type Appointment struct {
	BasicModel
	User   User `gorm:"foreignKey:UserID"`
	UserID int  `json:"user_id" gorm:"comment:'用户id'"`

	Name     string `json:"name" gorm:"type:varchar(100);comment:'称呼'"`
	Date     string `json:"date" gorm:"type:varchar(500);comment:'日期'"`
	Location string `json:"location" gorm:"type:varchar(200);comment:'地址'"`
	Phone    string `json:"phone" gorm:"type:varchar(20);comment:'电话'"`
	Memo     string `json:"memo" gorm:"type:varchar(500);comment:'备注'"`
	Status   string `json:"status" gorm:"type:enum('SUBMITTED', 'CANCELED', 'SETTLED');default:'SUBMITTED';comment:'状态'"`
}
