package models

import "github.com/google/uuid"

func (user *User) NewService() *Service {
	return &Service{
		Meta: NewInternals("service"),
		User: user.Ref(),
		ID:   uuid.NewString(),
	}
}

type Service struct {
	Meta        Internals
	User        UserRef
	ID          string   `json:"id" firestore:"id"`
	Name        string   `json:"name" firestore:"name"`
	Image       string   `json:"image" firestore:"image"`
	Description string   `json:"description" firestore:"description"`
	Duration    string   `json:"duration" firestore:"duration"`
	Concessions bool     `json:"concessions" firestore:"concessions"`
	Price       float64  `json:"price" firestore:"price"`
	Regions     []string `json:"regions" firestore:"regions"`
	Ailments    []string `json:"ailments" firestore:"ailments"`
}
