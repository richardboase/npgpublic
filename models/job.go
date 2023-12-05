package models

type Job struct {
	Meta   Internals
	Status string `json:"status" firestore:"status"`
}

func (job *Job) NewJob(name, status string) *Job {
	c := &Job{
		Meta:   job.Meta.NewInternals("jobs"),
		Status: status,
	}
	return c
}

func (collection *Collection) NewJob(name, status string) *Job {
	c := &Job{
		Meta:   collection.Meta.NewInternals("jobs"),
		Status: status,
	}
	return c
}
