package models

type Design struct {
	Meta       Internals
	Index      int               `json:"index" firestore:"index"`
	Name       string            `json:"name" firestore:"name"`
	Attributes []*Attribute      `json:"attributes" firestore:"attributes"`
	Data       map[string]string `json:"data" firestore:"data"`
}

func (job *Job) NewDesign(index int) *Design {
	c := &Design{
		Meta:  job.Meta.NewInternals("designs"),
		Index: index,
	}
	return c
}
