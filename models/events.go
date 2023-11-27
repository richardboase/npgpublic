package models

import (
	"log"

	"github.com/google/uuid"
)

type EventType string

const (
	Retreat  EventType = "retreat"
	Festival EventType = "festival"
)

// NewEvent constructs the event object
func (group *Group) NewEvent(name, description string, start int64, options EventOptions) (event *Event) {
	event = &Event{
		Meta:        NewInternals("event"),
		Options:     options,
		Group:       group.ID,
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		Time:        start,
	}
	log.Println(*event)
	return
}

type Event struct {
	Meta        Internals
	Options     EventOptions
	Online      OnlineEvent
	Group       string  `json:"group" firestore:"group"`
	ID          string  `json:"id" firestore:"id"`
	Name        string  `json:"name" firestore:"name"`
	Description string  `json:"description" firestore:"description"`
	Location    string  `json:"location" firestore:"location"`
	Lat         float64 `json:"lat" firestore:"lat"`
	Lng         float64 `json:"lng" firestore:"lng"`
	Time        int64   `json:"time" firestore:"time"`
	// hours
	Duration float64 `json:"duration" firestore:"duration"`
}

type EventOptions struct {
	// Event type
	Type string
	// is visible to general search
	Public bool
	// costs no money
	Free bool
	//
	EventType string
	// for accessibility
	DisabledAccess bool
}

type OnlineEvent struct {
	OnlineReadme    string
	MeetingID       string
	MeetingPassword string
	Link            string
	Phone           string
}

type EventChatMessage struct {
	Meta    Internals
	Event   string
	Message string
}

func (event *Event) NewMessage(msg string) *EventChatMessage {
	return &EventChatMessage{
		Meta:    NewInternals("eventchatmessage"),
		Event:   event.ID,
		Message: msg,
	}
}
