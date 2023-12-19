package models

type Overlay struct {
	Meta    Internals
	Type    string `json:"type" firestore:"type"`
	Content string `json:"content" firestore:"content"`
	X       int    `json:"x" firestore:"x"`
	Y       int    `json:"y" firestore:"y"`
}

func (layer *Layer) NewOverlay(overlayType, content string, x, y int) *Overlay {
	c := &Overlay{
		Meta: layer.Meta.NewInternals("overlays"),
		Type: overlayType,
		X:    x,
		Y:    y,
	}
	return c
}
