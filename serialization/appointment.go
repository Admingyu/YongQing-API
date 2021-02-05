package serialization

type AppointmentSer struct {
	ID        int      `json:"id"`
	Name      string   `json:"name" `
	Project   string   `json:"project" `
	Date      string   `json:"date"`
	Location  string   `json:"location"`
	Phone     string   `json:"phone"`
	Status    string   `json:"status"`
	Memo      string   `json:"memo"`
	CreatedAt jsonTime `json:"created_at"`
}
