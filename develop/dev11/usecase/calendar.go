package usecase

import (
	"dev11/customerror"
	"dev11/models"
)

type calendar struct {
	events []models.Event
}

func New() *calendar {
	return &calendar{
		events: []models.Event{},
	}
}

func (c *calendar) CreateEvent(e models.Event) (models.ResultPost, error) {
	e.ID = len(c.events) + 1
	c.events = append(c.events, e)
	return models.ResultPost{Result: "Event created successfully"}, nil
}

func (c *calendar) UpdateEvent(e models.Event) (models.ResultPost, error) {
	for i, event := range c.events {
		if event.ID == e.ID {
			c.events[i] = e
			return models.ResultPost{Result: "Event updated successfully"}, nil
		}
	}
	return models.ResultPost{}, &customerror.CustomBusinessError{Message: "Event not found", Code: 503}
}

func (c *calendar) DeleteEvent(e models.Event) (models.ResultPost, error) {
	for i, event := range c.events {
		if event.ID == e.ID {
			c.events = append(c.events[:i], c.events[i+1:]...)
			return models.ResultPost{Result: "Event deleted successfully"}, nil
		}
	}
	return models.ResultPost{}, &customerror.CustomBusinessError{Message: "Event not found", Code: 503}
}

func (c *calendar) EventsForDay(day string, userID int) (models.ResultGet, error) {
	var events []models.Event
	for _, event := range c.events {
		if event.Day == day && event.UserID == userID {
			events = append(events, event)
		}
	}
	return models.ResultGet{Result: events}, nil
}

func (c *calendar) EventsForWeek(week string, userID int) (models.ResultGet, error) {
	var events []models.Event
	for _, event := range c.events {
		if event.Week == week && event.UserID == userID {
			events = append(events, event)
		}
	}
	return models.ResultGet{Result: events}, nil
}

func (c *calendar) EventsForMonth(month string, userID int) (models.ResultGet, error) {
	var events []models.Event
	for _, event := range c.events {
		if event.Month == month && event.UserID == userID {
			events = append(events, event)
		}
	}
	return models.ResultGet{Result: events}, nil
}
