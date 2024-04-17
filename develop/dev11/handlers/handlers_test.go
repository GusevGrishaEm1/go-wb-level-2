package handlers

import (
	"bytes"
	"dev11/models"
	"dev11/usecase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
	var response models.ResultPost
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	// Add more assertions to check the response data
	// For example, you can check if the returned event matches the expected event
}

func TestDeleteEventHandler(t *testing.T) {
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

	// Create a new HTTP request for testing
	req := httptest.NewRequest("POST", "/delete_event", nil)
	w := httptest.NewRecorder()

	// Call your deleteEventHandler function with the request and response recorder
	c.deleteEventHandler(w, req)

	// Add assertions to check the response status code, body, etc.
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
	}

	// Parse the response body
	var response models.ResultPost
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	// Add more assertions to check the response data
	// For example, you can check if the returned event matches the expected event
}
