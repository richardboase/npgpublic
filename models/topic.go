package models

import "github.com/google/uuid"

type Topic struct {
	Meta        Internals
	ID          string `json:"id" firestore:"id"`
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
}

func NewTopic(name, description string) *Topic {
	return &Topic{
		Meta:        NewInternals("topic"),
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}
}

type TopicUpdate struct {
	Meta   Internals
	User   UserRef
	ID     string       `json:"id" firestore:"id"`
	Update string       `json:"update" firestore:"update"`
	Media  []*MediaFile `json:"media" firestore:"media"`
}

func (user *User) NewTopicUpdate(update string, mediaFiles ...*MediaFile) *TopicUpdate {
	return &TopicUpdate{
		Meta:   NewInternals("topicupdate"),
		User:   user.Ref(),
		ID:     uuid.NewString(),
		Update: update,
		Media:  mediaFiles,
	}
}

type TopicQuestion struct {
	Meta     Internals
	User     UserRef
	ID       string `json:"id" firestore:"id"`
	Question string `json:"question" firestore:"question"`
}

func (user *User) NewTopicQuestion(question string) *TopicQuestion {
	return &TopicQuestion{
		Meta:     NewInternals("topicquestion"),
		User:     user.Ref(),
		ID:       uuid.NewString(),
		Question: question,
	}
}
