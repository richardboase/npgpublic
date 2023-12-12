package models

type Manifest struct {
	Collection *Collection
	Layers     []*Layer
	Attributes []*Attribute
	Tags       []*Tag
}
