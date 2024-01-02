package models

import "net/http"

type Layer struct {
	Meta  Internals
	Name  string `json:"name" firestore:"name"`
	Type  string `json:"type" firestore:"type"`
	Order int    `json:"order" firestore:"order"`
	// this is for building the manifest
	Layers     []*Layer
	Elements   []*Element
	Attributes *Attribute
}

func (collection *Collection) NewLayer(name, layerType string, order int) *Layer {
	c := &Layer{
		Meta:  collection.Meta.NewInternals("layers"),
		Name:  name,
		Type:  layerType,
		Order: order,
	}
	return c
}

func (layer *Layer) Validate(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool
	layer.Name, exists = AssertKeyValue(w, m, "name")
	if !exists {
		return false
	}
	layer.Type, exists = AssertKeyValue(w, m, "type")

	return exists
}
