package models

type Manifest struct {
	Collection *Collection
	Layers     map[string][]*Layer
	Attributes []*Attribute
	Tags       []*Tag
}
