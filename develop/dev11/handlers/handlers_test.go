package handlers

import (
	"bytes"
	"dev11/models"
	"dev11/usecase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	calendar := usecase.New()

	event := models.Event{
		UserID:  1,
		Title:   "Test Event",
		Content: "This is a test event",
		Day:     "2022-01-01",
		Week:    "2022-W01",
		Month:   "2022-01",
	}

	// Create a mock request
	reqBody, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRequest(http.MethodPost, "/create_event", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	// Create a controller with the mock calendar
	ctrl := controller{calendar}

	// Call the createEventHandler function
	ctrl.createEventHandler(w, r)

	// Check the response
	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	var result models.ResultPost
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Result != "Event created successfully" {
		t.Errorf("Expected result message %q, got %q", "Event created successfully", result.Result)
	}
}

func createEvent(calendar Calendar, t *testing.T) {
	event := models.Event{
		UserID:  1,
		Title:   "Test Event",
		Content: "This is a test event",
		Day:     "2022-01-01",
		Week:    "2022-W01",
		Month:   "2022-01",
	}
	_, err := calendar.CreateEvent(event)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateEventHandler(t *testing.T) {
	// Create a new instance of your calendar implementation
	calendar := usecase.New()
	createEvent(calendar, t)

	event := models.Event{
		ID:      1,
		UserID:  1,
		Title:   "Test Event",
		Content: "This is a test event",
		Day:     "2022-01-01",
		Week:    "2022-W01",
		Month:   "2022-01",
	}

	// Create a new controller with the calendar instance
	c := controller{calendar}

	reqBody, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request for testing
	req := httptest.NewRequest("POST", "/update_event", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	// Call your updateEventHandler function with the request and response recorder
	c.updateEventHandler(w, req)

	// Add assertions to check the response status code, body, etc.
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
	}

	// Parse the response body
	var result models.ResultPost
	err = json.NewDecoder(w.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Result != "Event updated successfully" {
		t.Errorf("Expected result message %q, got %q", "Event updated successfully", result.Result)
	}

}

func TestDeleteEventHandler(t *testing.T) {
	calendar := usecase.New()
	createEvent(calendar, t)

	event := models.Event{
		ID:      1,
		UserID:  1,
		Title:   "Test Event",
		Content: "This is a test event",
		Day:     "2022-01-01",
		Week:    "2022-W01",
		Month:   "2022-01",
	}

	c := controller{calendar}

	reqBody, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "/delete_event", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	c.deleteEventHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
	}

	var result models.ResultPost
	err = json.NewDecoder(w.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Result != "Event deleted successfully" {
		t.Errorf("Expected result message %q, got %q", "Event deleted successfully", result.Result)
	}
}

func TestEventsForDayHandler(t *testing.T) {
	calendar := usecase.New()
	createEvent(calendar, t)
	c := controller{calendar}

	req, err := http.NewRequest("GET", "/events_for_day?day=2022-01-01&userID=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	c.eventsForDayHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, res.Code)
	}

	var result models.ResultGet
	err = json.Unmarshal(res.Body.Bytes(), &result)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []models.Event{
		{ID: 1, UserID: 1, Title: "Test Event", Content: "This is a test event", Day: "2022-01-01", Week: "2022-W01", Month: "2022-01"},
	}, result.Result)
}

func TestEventsForWeekHandler(t *testing.T) {
	calendar := usecase.New()
	createEvent(calendar, t)
	c := controller{calendar}

	req, err := http.NewRequest("GET", "/events_for_day?week=2022-W01&userID=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	c.eventsForWeekHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, res.Code)
	}

	var result models.ResultGet
	err = json.Unmarshal(res.Body.Bytes(), &result)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []models.Event{
		{ID: 1, UserID: 1, Title: "Test Event", Content: "This is a test event", Day: "2022-01-01", Week: "2022-W01", Month: "2022-01"},
	}, result.Result)
}

func TestEventsForMonthHandler(t *testing.T) {
	calendar := usecase.New()
	createEvent(calendar, t)
	c := controller{calendar}

	req, err := http.NewRequest("GET", "/events_for_day?month=2022-01&userID=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	c.eventsForMonthHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, res.Code)
	}

	var result models.ResultGet
	err = json.Unmarshal(res.Body.Bytes(), &result)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []models.Event{
		{ID: 1, UserID: 1, Title: "Test Event", Content: "This is a test event", Day: "2022-01-01", Week: "2022-W01", Month: "2022-01"},
	}, result.Result)
}
