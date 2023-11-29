package models

type Tag struct {
	Meta            Internals
	Name            string `json:"name" firestore:"name"`
	BackgroundColor string `json:"backgroundColor" firestore:"backgroundColor"`
	TextColor       string `json:"textColor" firestore:"textColor"`
}

func (collection *Collection) NewTag(name, backgroundColor, textColor string) *Tag {
	c := &Tag{
		Meta:            collection.Meta.NewInternals("tags"),
		Name:            name,
		BackgroundColor: backgroundColor,
		TextColor:       textColor,
	}
	return c
}

func (element *Element) NewTag(name, backgroundColor, textColor string) *Tag {
	c := &Tag{
		Meta:            element.Meta.NewInternals("tags"),
		Name:            name,
		BackgroundColor: backgroundColor,
		TextColor:       textColor,
	}
	return c
}
