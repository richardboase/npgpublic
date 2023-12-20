package models

type Layer struct {
	Meta  Internals
	Name  string `json:"name" firestore:"name"`
	Type  string `json:"type" firestore:"type"`
	Order int    `json:"order" firestore:"order"`
	// this is for building the manifest
	Overlays   []*Overlay
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
