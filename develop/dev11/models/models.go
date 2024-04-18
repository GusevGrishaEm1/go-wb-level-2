package models

// Event представляет событие.
type Event struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Day     string `json:"day"`
	Week    string `json:"week"`
	Month   string `json:"month"`
}

// ResultPost представляет результат постинга.
type ResultPost struct {
	Result string `json:"result"`
}

// ResultGet представляет результат получения событий.
type ResultGet struct {
	Result []Event `json:"result"`
}

// Error представляет ошибку.
type Error struct {
	Error string `json:"error"`
}

// ConfigServer представляет конфигурацию сервера.
type ConfigServer struct {
	Addr string
}
