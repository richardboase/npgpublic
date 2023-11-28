package models

type Attribute struct {
	Meta    Internals
	Name    string           `json:"name" firestore:"name"`
	Options AttributeOptions `json:"options" firestore:"options"`
}

func (collection *Collection) NewAttribute(name string, options *AttributeOptions) *Attribute {
	c := &Attribute{
		Meta: collection.Meta.NewInternals("attributes"),
		Name: name,
	}
	if options != nil {
		c.Options = *options
	}
	return c
}

type AttributeOptions struct {
}
