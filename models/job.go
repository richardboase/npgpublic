package models

type Job struct {
	Meta   Internals
	Status string `json:"status" firestore:"status"`
}

func (collection *Collection) NewJob(status string) *Job {
	c := &Job{
		Meta:   collection.Meta.NewInternals("jobs"),
		Status: status,
	}
	return c
}
