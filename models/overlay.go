package models

import "net/http"

type Overlay struct {
	Meta        Internals
	Description string `json:"description" firestore:"description"`
	Type        string `json:"type" firestore:"type"`
	Content     string `json:"content" firestore:"content"`
	X           int    `json:"x" firestore:"x"`
	Y           int    `json:"y" firestore:"y"`
	Font        string `json:"font" firestore:"font"`
}

func (layer *Layer) NewOverlay(description, overlayType, content, font string, x, y int) *Overlay {
	c := &Overlay{
		Meta:        layer.Meta.NewInternals("overlays"),
		Description: description,
		Type:        overlayType,
		Content:     content,
		X:           x,
		Y:           y,
		Font:        font,
	}
	return c
}

func (overlay *Overlay) Validate(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool
	overlay.Description, exists = AssertKeyValue(w, m, "description")
	if !exists {
		return false
	}
	overlay.Type, exists = AssertKeyValue(w, m, "type")
	if !exists {
		return false
	}
	overlay.Content, exists = AssertKeyValue(w, m, "content")
	if !exists {
		return false
	}
	overlay.X, exists = AssertKeyValueInt(w, m, "x")
	if !exists {
		return false
	}
	overlay.X, exists = AssertKeyValueInt(w, m, "x")
	if !exists {
		return false
	}
	overlay.Y, exists = AssertKeyValueInt(w, m, "y")

	return exists
}
