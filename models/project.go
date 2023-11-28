package models

import "github.com/google/uuid"

type Project struct {
	Meta        Internals
	ID          string `json:"id" firestore:"id"`
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
}

func (user *User) NewProject(name, description string) *Project {
	p := &Project{
		Meta:        NewInternals("project"),
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}
	p.Meta.Parent = user.ID
	p.Meta.Moderation.Admins = []string{user.ID}
	return p
}

type ProjectOptions struct {
}
