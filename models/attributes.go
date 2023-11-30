package models

import "net/http"

type Attribute struct {
	Meta    Internals
	Name    string           `json:"name" firestore:"name"`
	Options AttributeOptions `json:"options" firestore:"options"`
	Value   int              `json:"value" firestore:"value"`
}

func (collection *Collection) NewAttribute(name string, options AttributeOptions) *Attribute {
	c := &Attribute{
		Meta:    collection.Meta.NewInternals("attributes"),
		Name:    name,
		Options: options,
	}
	return c
}

type AttributeOptions struct {
	Min int
	Max int
}

func (attribute *Attribute) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool

	attribute.Name, exists = AssertKeyValue(w, m, "name")
	if !exists {
		return false
	}

	attribute.Options.Min, exists = AssertKeyValueInt(w, m, "min")
	if !exists {
		return false
	}

	attribute.Options.Max, exists = AssertKeyValueInt(w, m, "max")

	return exists
}
