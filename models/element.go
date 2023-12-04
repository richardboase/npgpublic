package models

import "github.com/google/uuid"

type Element struct {
	Meta  Internals
	Name  string `json:"name" firestore:"name"`
	Image string `json:"image" firestore:"image"`
	// manifets
	Tags []*Tag `json:",omitempty" firestore:"-"`
}

func (layer *Layer) NewElement(name string) *Element {
	c := &Element{
		Meta:  layer.Meta.NewInternals("elements"),
		Name:  name,
		Image: uuid.NewString(),
	}
	return c
}
