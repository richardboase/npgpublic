package models

import "net/http"

type Collection struct {
	Meta    Internals
	Name    string            `json:"name" firestore:"name"`
	Options CollectionOptions `json:"options" firestore:"options"`
}

func (project *Project) NewCollection(name string, options *CollectionOptions) *Collection {
	c := &Collection{
		Meta: project.Meta.NewInternals("collections"),
		Name: name,
	}
	if options != nil {
		c.Options = *options
	}
	return c
}

type CollectionOptions struct {
	MaxMint       int
	ArtworkWidth  int
	ArtworkHeight int
	PrimaryFont   string
	SecondaryFont string
}

func (collection *Collection) ValidateInput(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool

	collection.Name, exists = AssertKeyValue(w, m, "name")
	if !exists {
		return false
	}

	// optional
	collection.Options.PrimaryFont, _ = AssertKeyValue(w, m, "primaryFont")
	collection.Options.SecondaryFont, _ = AssertKeyValue(w, m, "secondaryFont")

	collection.Options.MaxMint, exists = AssertKeyValueInt(w, m, "maxMint")
	if !exists {
		return false
	}

	collection.Options.ArtworkWidth, exists = AssertKeyValueInt(w, m, "artworkWidth")
	if !exists {
		return false
	}

	collection.Options.ArtworkHeight, exists = AssertKeyValueInt(w, m, "artworkHeight")

	return exists
}
