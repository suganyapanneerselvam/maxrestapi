package models

import "time"

type Event struct {
	ID       string `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserID      string    `json:"userId"`
}

var events []Event = []Event{}

func (e *Event) Save() {
	events = append(events, *e)
}
func GetAllEvents() []Event {
	return events
}
