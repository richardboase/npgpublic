package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Meeting struct {
	Meta      Internals
	Options   GroupOptions
	ID        string  `json:"id" firestore:"id"`
	Title     string  `json:"title" firestore:"title"`
	Address   string  `json:"address" firestore:"address"`
	Cost      float64 `json:"cost" firestore:"cost"`
	Date      string  `json:"date" firestore:"date"`
	Time      string  `json:"time" firestore:"time"`
	Timestamp int64   `json:"timestamp" firestore:"timestamp"`
	// number of minutes
	Duration int `json:"duration" firestore:"duration"`
}

func (group *Group) NewMeeting(title, address string, cost float64, date, theTime string, duration int) (*Meeting, error) {

	t, err := time.Parse("2006-01-02T15:04:05.000Z", date)
	if err != nil {
		return nil, err
	}

	hs := strings.Split(theTime, ":")
	hours, err := strconv.Atoi(hs[0])
	if err != nil {
		return nil, err
	}
	minutes, err := strconv.Atoi(hs[1])
	if err != nil {
		return nil, err
	}
	for x := 0; x < hours; x++ {
		t.Add(time.Hour)
	}
	for x := 0; x < minutes; x++ {
		t.Add(time.Minute)
	}
	meeting := &Meeting{
		Meta:    NewInternals("meeting"),
		ID:      uuid.NewString(),
		Title:   title,
		Address: address,
		Cost:    cost,
		Date:    date,
		Time:    theTime,
		// use this for db order-by queries
		Timestamp: t.UTC().Unix(),
		Duration:  duration,
	}
	meeting.Meta.Parent = group.ID
	return meeting, nil
}

type MeetingComment struct {
	Meta    Internals
	User    UserRef
	ID      string `json:"id" firestore:"id"`
	Content string `json:"content" firestore:"content"`
}

func (meeting *Meeting) NewComment(user *User, content string) *MeetingComment {
	comment := &MeetingComment{
		Meta:    NewInternals("meetingcomment"),
		User:    user.Ref(),
		ID:      uuid.NewString(),
		Content: content,
	}
	comment.Meta.Parent = meeting.ID
	return comment
}
