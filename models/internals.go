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
		ID:       n.ID + "." + class + "-" + uuid.NewString()[:15],
		Class:    class,
		Created:  timestamp,
		Modified: timestamp,
	}
	if len(n.ID) > 0 {
		x.Context.Parent = n.ID
		x.Context.Parents = append(n.Context.Parents, n.ID)
	}
	return x
}

type Internals struct {
	ID         string
	Class      string
	Context    Context
	Moderation Moderation
	Updated    bool
	Created    int64
	Modified   int64
	Stats      map[string]int
}

func (i *Internals) DocPath() string {
	p := strings.Split(string(i.ID[1:]), ".")
	parts := make([][]string, len(p))
	k := ""
	for x, s := range p {
		k += "." + s
		parts[x] = strings.Split(k, ".")

	}
	outs := []string{}
	for _, p := range parts {
		class := strings.Split(p[len(p)-1], "-")[0]
		outs = append(outs, class+"/"+strings.Join(p, "."))
	}
	return strings.Join(outs, "/")
}

func (i *Internals) Firestore(app *common.App) *firestore.DocumentRef {
	return app.Firestore().Doc(i.DocPath())
}

func (i *Internals) FirestoreDoc(app *common.App, ii Internals) *firestore.DocumentRef {
	return app.Firestore().Doc(i.DocPath()).Collection(ii.Class).Doc(ii.ID)
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
