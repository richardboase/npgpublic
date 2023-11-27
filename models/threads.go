package models

import (
	"strings"

	"github.com/google/uuid"
)

func (user *User) NewThread(parent, title string) *Thread {
	t := &Thread{
		Meta:  NewInternals("thread"),
		User:  user.Ref(),
		ID:    uuid.NewString(),
		Title: title,
	}
	t.Meta.Parent = parent
	return t
}

type Thread struct {
	Meta  Internals
	User  UserRef
	Type  string `json:"-" firestore:"-"`
	ID    string `json:"id" firestore:"id"`
	Title string `json:"title" firestore:"title"`
}

func (thread *Thread) Reply(user *User, content string) *ThreadReply {
	r := &ThreadReply{
		Meta:    NewInternals("threadreply"),
		User:    user.Ref(),
		ID:      uuid.NewString(),
		Title:   thread.Title,
		Content: content,
	}
	r.Meta.Parent = thread.ID
	return r
}

type ThreadReply struct {
	Meta    Internals
	User    UserRef
	ID      string `json:"id" firestore:"id"`
	Thread  string `json:"thread" firestore:"thread"`
	Title   string `json:"title" firestore:"title"`
	Content string `json:"content" firestore:"content"`
}

func (reply *ThreadReply) ThreadID() string {
	return strings.Split(reply.ID, "_")[0]
}
