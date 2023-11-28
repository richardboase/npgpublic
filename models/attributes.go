package models

type Attribute struct {
	Meta Internals
	Name string `json:"name" firestore:"name"`
	Min  int    `json:"min" firestore:"min"`
	Max  int    `json:"max" firestore:"max"`
}

func (collection *Collection) NewAttribute(name string, min, max int) *Attribute {
	c := &Attribute{
		Meta: collection.Meta.NewInternals("attributes"),
		Name: name,
		Min:  min,
		Max:  max,
	}
	return c
}

type AttributeOptions struct {
}
