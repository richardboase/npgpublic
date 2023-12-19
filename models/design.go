package models

import (
	"strconv"
	"strings"
)

type Design struct {
	Meta       Internals
	Index      int               `json:"index" firestore:"index"`
	Name       string            `json:"name" firestore:"name"`
	Elements   []*Element        `json:"elements" firestore:"elements"`
	Attributes map[string]int    `json:"attributes" firestore:"attributes"`
	Data       map[string]string `json:"data" firestore:"data"`
}

func (job *Job) NewDesign(index int) *Design {
	c := &Design{
		Meta:  job.Meta.NewInternals("designs"),
		Index: index,
	}
	parts := strings.Split(c.Meta.ID, "-")
	parts[len(parts)-1] = strconv.Itoa(index)
	c.Meta.ID = strings.Join(parts, "-")
	return c
}
