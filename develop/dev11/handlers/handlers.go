// Package handlers содержит обработчики HTTP-запросов.
package handlers

import (
	"dev11/customerror"
	"dev11/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Calendar представляет интерфейс для работы с календарем.
type Calendar interface {
	// CreateEvent создает событие в календаре.
	CreateEvent(e models.Event) (models.ResultPost, error)
	// UpdateEvent обновляет событие в календаре.
	UpdateEvent(e models.Event) (models.ResultPost, error)
	// DeleteEvent удаляет событие из календаря.
	DeleteEvent(e models.Event) (models.ResultPost, error)
	// EventsForDay возвращает события для указанного дня и пользователя.
	EventsForDay(day string, userID int) (models.ResultGet, error)
	// EventsForWeek возвращает события для указанной недели и пользователя.
	EventsForWeek(week string, userID int) (models.ResultGet, error)
	// EventsForMonth возвращает события для указанного месяца и пользователя.
	EventsForMonth(month string, userID int) (models.ResultGet, error)
}

// StartServer запускает HTTP-сервер.
//
// Параметры:
//   - calendar: интерфейс для работы с календарем.
//   - config: конфигурация сервера.
func StartServer(calendar Calendar, config models.ConfigServer) {
	controller := controller{calendar}

	r := chi.NewMux()

	r.Use(logMiddleware)

	r.Post("/create_event", controller.createEventHandler)
	r.Post("/update_event", controller.updateEventHandler)
	r.Post("/delete_event", controller.deleteEventHandler)
	r.Get("/events_for_day", controller.eventsForDayHandler)
	r.Get("/events_for_week", controller.eventsForWeekHandler)
	r.Get("/events_for_month", controller.eventsForMonthHandler)

	log.Fatal(http.ListenAndServe(config.Addr, r))
}

type controller struct {
	Calendar
}

func (c controller) createEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.CreateEvent(event)
	if err != nil {
		customErr := &customerror.CustomBusinessError{}
		if errors.As(err, &customErr) {
			errorResponse := models.Error{Error: customErr.Error()}
			errJSON, _ := json.Marshal(errorResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(customErr.Code)
			w.Write(errJSON)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c controller) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.UpdateEvent(event)
	if err != nil {
		customErr := &customerror.CustomBusinessError{}
		if errors.As(err, &customErr) {
			errorResponse := models.Error{Error: customErr.Error()}
			errJSON, _ := json.Marshal(errorResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(customErr.Code)
			w.Write(errJSON)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c controller) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.DeleteEvent(event)
	if err != nil {
		customErr := &customerror.CustomBusinessError{}
		if errors.As(err, &customErr) {
			errorResponse := models.Error{Error: customErr.Error()}
			errJSON, _ := json.Marshal(errorResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(customErr.Code)
			w.Write(errJSON)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c controller) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Query().Get("day")

	if day == "" {
		http.Error(w, "Missing 'day' query parameter", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.EventsForDay(day, userID)
	if err != nil {
		customErr := &customerror.CustomBusinessError{}
		if errors.As(err, &customErr) {
			errorResponse := models.Error{Error: customErr.Error()}
			errJSON, _ := json.Marshal(errorResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(customErr.Code)
			w.Write(errJSON)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c controller) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	week := r.URL.Query().Get("week")

	if week == "" {
		http.Error(w, "Missing 'week' query parameter", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.EventsForWeek(week, userID)
	if err != nil {
		customErr := &customerror.CustomBusinessError{}
		if errors.As(err, &customErr) {
			errorResponse := models.Error{Error: customErr.Error()}
			errJSON, _ := json.Marshal(errorResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(customErr.Code)
			w.Write(errJSON)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c controller) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	month := r.URL.Query().Get("month")

	if month == "" {
		http.Error(w, "Missing 'month' query parameter", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.EventsForMonth(month, userID)
	if err != nil {
		customErr := &customerror.CustomBusinessError{}
		if errors.As(err, &customErr) {
			errorResponse := models.Error{Error: customErr.Error()}
			errJSON, _ := json.Marshal(errorResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(customErr.Code)
			w.Write(errJSON)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func logMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}
