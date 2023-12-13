package models

import (
	"net/http"
)

type Job struct {
	Meta   Internals
	Status string `json:"status" firestore:"status"`
	Stages Stages
}

type Stages struct {
	IsPreview bool
	Prepare   interface{}
	Generate  interface{}
}

func (collection *Collection) NewJob(name, status string) *Job {
	c := &Job{
		Meta:   collection.Meta.NewInternals("jobs"),
		Status: status,
	}
	return c
}

func (job *Job) Validate(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool
	job.Stages.IsPreview, exists = AssertKeyValueBool(w, m, "preview")
	if !exists {
		return false
	}
	job.Status, exists = AssertKeyValue(w, m, "status")
	return exists
}
