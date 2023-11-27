package models

import "github.com/google/uuid"

type Collection struct {
	Meta    Internals
	ID      string            `json:"id" firestore:"id"`
	Name    string            `json:"name" firestore:"name"`
	Options CollectionOptions `json:"options" firestore:"options"`
}

func (project *Project) NewCollection(name string, options *CollectionOptions) *Collection {
	c := &Collection{
		Meta: NewInternals("collection"),
		ID:   uuid.NewString(),
		Name: name,
	}
	if options != nil {
		c.Options = *options
	}
	c.Meta.Parent = project.ID
	return c
}

type CollectionOptions struct {
	MaxMint int
}
