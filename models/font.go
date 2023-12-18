package models

type Font struct {
	Meta Internals
	Name string `json:"name" firestore:"name"`
}

func (project *Project) NewFont(name string) *Font {
	c := &Font{
		Meta: project.Meta.NewInternals("fonts"),
		Name: name,
	}
	return c
}
