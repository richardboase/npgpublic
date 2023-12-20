package models

type Manifest struct {
	Collection *Collection
	Fonts      map[string]string
	Layers     map[string][]*Layer
	Attributes []*Attribute
	Tags       []*Tag
}
