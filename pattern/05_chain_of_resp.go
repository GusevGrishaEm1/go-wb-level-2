package pattern

import (
	"fmt"
	"net/http"
)

// Handler interface
type Handler interface {
	SetNext(handler Handler)
	Handle(req *http.Request)
}

// BaseHandler struct
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// ConcreteHandler struct
type ConcreteHandler struct {
	BaseHandler
	name string
	path string
}

func (h *ConcreteHandler) Handle(req *http.Request) {
	if req.URL.Path == h.path {
		fmt.Printf("%s обработал HTTP запрос для %sn", h.name, h.path)
	} else if h.next != nil {
		fmt.Printf("%s передал HTTP запрос дальшеn", h.name)
		h.next.Handle(req)
	} else {
		fmt.Printf("Нет обработчика для HTTP запроса %sn", req.URL.Path)
	}
}

func test_05() {
	// Создаем обработчики запросов
	homeHandler := &ConcreteHandler{name: "Домашняя страница", path: "/"}
	aboutHandler := &ConcreteHandler{name: "О нас", path: "/about"}
	contactHandler := &ConcreteHandler{name: "Контакты", path: "/contact"}

	// Строим цепочку обработчиков
	homeHandler.SetNext(aboutHandler)
	aboutHandler.SetNext(contactHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeHandler.Handle(r)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		homeHandler.Handle(r)
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		homeHandler.Handle(r)
	})

	fmt.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", nil)
}
