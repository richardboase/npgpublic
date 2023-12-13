package models

import (
	"net/http"
)

type Job struct {
	Meta Internals
	// pending:started:completed:failed
	Status string `json:"status" firestore:"status"`
	Stage  string `json:"stage" firestore:"stage"`
	Stages Stages
}

type Stages struct {
	IsPreview bool
	Prepare   Stage
	Generate  Stage
}

type Stage struct {
	Data      interface{}
	Notes     []string
	Completed bool
}

func (collection *Collection) NewJob(name, status string) *Job {
	c := &Job{
		Meta:   collection.Meta.NewInternals("jobs"),
		Status: "pending",
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
