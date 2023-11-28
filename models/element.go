package models

type Element struct {
	Meta    Internals
	Name    string       `json:"name" firestore:"name"`
	Options LayerOptions `json:"options" firestore:"options"`
}

func (layer *Layer) NewElement(name string, options *LayerOptions) *Element {
	c := &Element{
		Meta: layer.Meta.NewInternals("elements"),
		Name: name,
	}
	if options != nil {
		c.Options = *options
	}
	return c
}

type ElementOptions struct {
}
