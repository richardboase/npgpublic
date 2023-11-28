package utils

import (
	"context"
	"strings"

	"github.com/richardboase/npgpublic/sdk/common"
)

func GetDocument(app *common.App, id string, dst interface{}) error {

	path := FirestorePath(id)
	println("GET DOCUMENT", path)

	doc, err := app.Firestore().Doc(path).Get(context.Background())
	if err != nil {
		return err
	}
	return doc.DataTo(dst)
}

func FirestorePath(id string) string {
	p := strings.Split(string(id[1:]), ".")
	println(len(p))

	parts := make([][]string, len(p))

	k := ""
	for x, s := range p {
		k += "." + s
		println("K", s)
		parts[x] = strings.Split(k, ".")
		println(strings.Join(parts[x], "."))

	}

	outs := []string{}
	for _, p := range parts {

		class := strings.Split(p[len(p)-1], "-")[0]

		outs = append(outs, class+"/"+strings.Join(p, "."))
	}

	for _, o := range outs {
		println("OUT   ", o)
	}

	return strings.Join(outs, "/")
}
