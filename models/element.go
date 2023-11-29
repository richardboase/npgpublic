package models

type Element struct {
	Meta  Internals
	Name  string `json:"name" firestore:"name"`
	Image string `json:"image" firestore:"image"`
}

func (layer *Layer) NewElement(name string, options *LayerOptions) *Element {
	c := &Element{
		Meta: layer.Meta.NewInternals("elements"),
		Name: name,
	}
	return c
}
