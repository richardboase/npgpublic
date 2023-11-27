package models

import "github.com/google/uuid"

type Content struct {
	Type   string `json:"type" firestore:"type"`
	String string `json:"string" firestore:"string"`
}

type Status struct {
	Meta    Internals
	User    UserRef
	ID      string    `json:"id" firestore:"id"`
	Content []Content `json:"content" firestore:"content"`
}

func (user *User) NewStatus(parent string, content ...Content) *Status {
	status := &Status{
		Meta:    NewInternals("status"),
		User:    user.Ref(),
		ID:      uuid.NewString(),
		Content: content,
	}
	status.Meta.Parent = parent
	return status
}
