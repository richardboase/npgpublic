package models

import (
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/richardboase/npgpublic/sdk/common"
)

// NewInternals returns a boilerplate internal object
func (n Internals) NewInternals(class string) Internals {

	timestamp := time.Now().UTC().Unix()

	x := Internals{
		ID:       n.ID + "." + class + "." + uuid.NewString(),
		Class:    class,
		Created:  timestamp,
		Modified: timestamp,
	}
	if len(n.ID) > 0 {
		x.Context.Parent = n.ID
		x.Context.Parents = append(n.Context.Parents, n.Class+"/"+n.ID)
	}
	return x
}

type Internals struct {
	ID    string `json:"id" firestore:"id"`
	Class string
	Context
	Moderation
	Updated    bool
	Searchable bool
	Created    int64
	Modified   int64
	Stats      map[string]int
}

func (i *Internals) Firestore(app *common.App) *firestore.DocumentRef {
	path := strings.Join(strings.Split(i.ID, "."), "/")
	return app.Firestore().Doc(path)
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

type Context struct {
	Parent  string
	Parents []string
	Country string
	Region  string
}

type Moderation struct {
	Admins       []string
	Blocked      bool
	BlockedTime  int64
	BlockedBy    string
	Approved     bool
	ApprovedTime int64
	ApprovedBy   string
}
