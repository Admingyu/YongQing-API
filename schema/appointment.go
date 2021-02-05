package schema

type AppointmentSchema struct {
	Name     string `json:"name" binding:"required,max=50"`
	Date     string `json:"date" binding:"required,max=200"`
	Location string `json:"location" binding:"required,max=200"`
	Phone    string `json:"phone" binding:"required,max=20"`
	Project  string `json:"project" binding:"required,max=50"`
	Memo     string `json:"memo" binding:"max=500"`
}

type IDSchema struct {
	ID int `json:"id" form:"id" binding:"required"`
}
