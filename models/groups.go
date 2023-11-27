package models

import "github.com/google/uuid"

type Group struct {
	Meta        Internals
	Options     GroupOptions
	ID          string    `json:"id" firestore:"id"`
	Region      string    `json:"region" firestore:"region"`
	Admins      []UserRef `json:"admins" firestore:"admins"`
	Moderators  []UserRef `json:"moderators" firestore:"moderators"`
	Name        string    `json:"name" firestore:"name"`
	Description string    `json:"description" firestore:"description"`
	Email       string    `json:"email" firestore:"email"`
	Website     string    `json:"website" firestore:"website"`
	Socials
}

type GroupOptions struct {
	// controls whether anyone can instantly join a group
	Open bool
	// has threads
	Threads bool
	// has meetings
	Meetings bool
	// has events
	Events bool
}

func NewGroup(user *User, region, name, descr string, options GroupOptions) *Group {
	return &Group{
		Meta:        NewInternals("group"),
		ID:          uuid.NewString(),
		Region:      region,
		Admins:      []UserRef{user.Ref()},
		Name:        name,
		Description: descr,
		Options:     options,
	}
}
