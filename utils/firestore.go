package utils

import "strings"

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
