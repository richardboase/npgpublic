package models

import "time"

type Session struct {
	UserID  string
	Expires int64
}

func (user *User) NewSession() *Session {
	return &Session{
		UserID:  user.ID,
		Expires: time.Now().UTC().Unix(),
	}
}
