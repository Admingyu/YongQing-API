package serialization

type SlideSer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Link  string `json:"link"`
}

type CaseListSer struct {
	ID       int    `json:"id"`
	Image    string `json:"image"`
	Price    string `json:"price"`
	Desc     string `json:"desc"`
	Location string `json:"location"`
}

type CaseDetail struct {
	ID      int    `json:"id"`
	Image   string `json:"image"`
	Video   string `json:"video"`
	Content string `json:"content"`
}

type CategorySer struct {
	ID   int    `json:"id"`
	Icon string `json:"icon"`
	Text string `json:"text"`
	Page string `json:"page"`
}
