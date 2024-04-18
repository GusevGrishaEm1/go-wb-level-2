// Package usecase предоставляет функционал для управления событиями в календаре.
package usecase

import (
	"dev11/customerror"
	"dev11/models"
)

// Calendar представляет собой структуру, содержащую события в календаре.
type Calendar struct {
	events []models.Event
}

// New создает новый экземпляр календаря с пустым срезом событий.
func New() *Calendar {
	return &Calendar{
		events: []models.Event{},
	}
}

// CreateEvent добавляет новое событие в календарь.
func (c *Calendar) CreateEvent(e models.Event) (models.ResultPost, error) {
	e.ID = len(c.events) + 1
	c.events = append(c.events, e)
	return models.ResultPost{Result: "Event created successfully"}, nil
}

// UpdateEvent обновляет существующее событие в календаре.
func (c *Calendar) UpdateEvent(e models.Event) (models.ResultPost, error) {
	for i, event := range c.events {
		if event.ID == e.ID {
			c.events[i] = e
			return models.ResultPost{Result: "Event updated successfully"}, nil
		}
	}
	return models.ResultPost{}, &customerror.CustomBusinessError{Message: "Event not found", Code: 503}
}

// DeleteEvent удаляет событие из календаря.
func (c *Calendar) DeleteEvent(e models.Event) (models.ResultPost, error) {
	for i, event := range c.events {
		if event.ID == e.ID {
			c.events = append(c.events[:i], c.events[i+1:]...)
			return models.ResultPost{Result: "Event deleted successfully"}, nil
		}
	}
	return models.ResultPost{}, &customerror.CustomBusinessError{Message: "Event not found", Code: 503}
}

// EventsForDay извлекает события для определенного дня из календаря.
func (c *Calendar) EventsForDay(day string, userID int) (models.ResultGet, error) {
	var events []models.Event
	for _, event := range c.events {
		if event.Day == day && event.UserID == userID {
			events = append(events, event)
		}
	}
	return models.ResultGet{Result: events}, nil
}

// EventsForWeek извлекает события для определенной недели из календаря.
func (c *Calendar) EventsForWeek(week string, userID int) (models.ResultGet, error) {
	var events []models.Event
	for _, event := range c.events {
		if event.Week == week && event.UserID == userID {
			events = append(events, event)
		}
	}
	return models.ResultGet{Result: events}, nil
}

// EventsForMonth извлекает события для определенного месяца из календаря.
func (c *Calendar) EventsForMonth(month string, userID int) (models.ResultGet, error) {
	var events []models.Event
	for _, event := range c.events {
		if event.Month == month && event.UserID == userID {
			events = append(events, event)
		}
	}
	return models.ResultGet{Result: events}, nil
}
