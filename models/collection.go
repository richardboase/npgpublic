package models

type Collection struct {
	Meta    Internals
	Name    string            `json:"name" firestore:"name"`
	Options CollectionOptions `json:"options" firestore:"options"`
}

func (project *Project) NewCollection(name string, options *CollectionOptions) *Collection {
	c := &Collection{
		Meta: project.Meta.NewInternals("collections"),
		Name: name,
	}
	if options != nil {
		c.Options = *options
	}
	return c
}

type CollectionOptions struct {
	MaxMint int
}
