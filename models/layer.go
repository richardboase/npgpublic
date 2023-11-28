package models

type Layer struct {
	Meta    Internals
	ID      string       `json:"id" firestore:"id"`
	Name    string       `json:"name" firestore:"name"`
	Order   int          `json:"order" firestore:"order"`
	Options LayerOptions `json:"options" firestore:"options"`
}

func (collection *Collection) NewLayer(name string, options *LayerOptions) *Layer {
	c := &Layer{
		Meta: collection.Meta.NewInternals("layer"),
		Name: name,
	}
	if options != nil {
		c.Options = *options
	}
	return c
}

type LayerOptions struct {
}
