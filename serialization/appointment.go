package serialization

type AppointmentSer struct {
	Name      string   `json:"name" `
	Date      string   `json:"date"`
	Location  string   `json:"location"`
	Phone     string   `json:"phone"`
	Memo      string   `json:"memo"`
	CreatedAt jsonTime `json:"created_at"`
}
