package models

type Font struct {
	Meta Internals
	Name string `json:"name" firestore:"name"`
}

func (layer *Layer) NewFont(name string) *Font {
	c := &Font{
		Meta: layer.Meta.NewInternals("fonts"),
		Name: name,
	}
	return c
}
