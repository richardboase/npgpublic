package models

type Tag struct {
	Meta  Internals
	Name  string `json:"name" firestore:"name"`
	Color string `json:"color" firestore:"color"`
}

func (collection *Collection) NewTag(name, color string) *Tag {
	c := &Tag{
		Meta:  collection.Meta.NewInternals("tags"),
		Name:  name,
		Color: color,
	}
	return c
}

func (element *Element) NewTag(name, color string) *Tag {
	c := &Tag{
		Meta:  element.Meta.NewInternals("tags"),
		Name:  name,
		Color: color,
	}
	return c
}
