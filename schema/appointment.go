package schema

type AppointmentSchema struct {
	Name     string `json:"name" binding:"required,max=50"`
	Date     string `json:"date" binding:"required,max=200"`
	Location string `json:"location" binding:"required,max=200"`
	Phone    string `json:"phone" binding:"required,max=20"`
	Memo     string `json:"memo" binding:"max=500"`
}
