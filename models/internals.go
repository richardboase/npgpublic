package models

import "time"

// NewInternals returns a boilerplate internal object
func NewInternals(class string) Internals {

	timestamp := time.Now().UTC().Unix()

	return Internals{
		Class:    class,
		Created:  timestamp,
		Modified: timestamp,
	}
}

type Internals struct {
	Class      string
	Parent     string
	Country    string
	Region     string
	Moderation struct {
		Blocked      bool
		BlockedTime  int64
		BlockedBy    string
		Approved     bool
		ApprovedTime int64
		ApprovedBy   string
	}
	Updated    bool
	Searchable bool
	Created    int64
	Modified   int64
	Stats      struct {
		Followers int64
		Views     int64
		Likes     int64
		Replies   int64
		Children  int64
	}
}

// Modify updates the timestamp
func (i *Internals) Modify() {
	i.Modified = time.Now().UTC().Unix()
}

// Update sets the metadata to indicate it has updated for a user
func (i *Internals) Update() {
	i.Updated = true
	i.Modify()
}
