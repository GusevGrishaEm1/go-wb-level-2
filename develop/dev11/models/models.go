package models

type Event struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Day     string `json:"day"`
	Week    string `json:"week"`
	Month   string `json:"month"`
}

type ResultPost struct {
	Result string `json:"result"`
}

type ResultGet struct {
	Result []Event `json:"result"`
}

type Error struct {
	Error string `json:"error"`
}

type ConfigServer struct {
	Addr string
}
