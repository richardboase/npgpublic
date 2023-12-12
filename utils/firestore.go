package utils

import (
	"context"

	"github.com/richardboase/npgpublic/models"
	"github.com/richardboase/npgpublic/sdk/common"
)

func GetDocument(app *common.App, id string, dst interface{}) error {

	i := models.Internal(id)
	path := i.DocPath()

	println("GET DOCUMENT", path)

	doc, err := app.Firestore().Doc(path).Get(context.Background())
	if err != nil {
		return err
	}
	return doc.DataTo(dst)
}
