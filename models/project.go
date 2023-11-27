package models

import "github.com/google/uuid"

type Project struct {
	Meta        Internals
	ID          string `json:"id" firestore:"id"`
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
}

func NewProject(name, description string) *Project {
	return &Project{
		Meta:        NewInternals("topic"),
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}
}
