package models

type Overlay struct {
	Meta    Internals
	Name    string `json:"name" firestore:"name"`
	Type    string `json:"type" firestore:"type"`
	Content string `json:"content" firestore:"content"`
	X       int    `json:"x" firestore:"x"`
	Y       int    `json:"y" firestore:"y"`
}

func (collection *Collection) NewOverlay(overlayType, content string, x, y int) *Overlay {
	c := &Overlay{
		Meta: collection.Meta.NewInternals("layers"),
		Type: overlayType,
		X:    x,
		Y:    y,
	}
	return c
}
