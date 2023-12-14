package models

type Design struct {
	Meta       Internals
	Name       string            `json:"name" firestore:"name"`
	Index      string            `json:"index" firestore:"index"`
	Attributes []*Attribute      `json:"attributes" firestore:"attributes"`
	Data       map[string]string `json:"data" firestore:"data"`
}

func (job *Job) NewDesign(index string) *Design {
	c := &Design{
		Meta:  job.Meta.NewInternals("designs"),
		Index: index,
	}
	return c
}