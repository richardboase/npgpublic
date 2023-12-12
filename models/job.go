package models

import (
	"net/http"
)

type Job struct {
	Meta    Internals
	Name    string          `json:"name" firestore:"name"`
	Status  string          `json:"status" firestore:"status"`
	Jobs    []string        `json:"jobs" firestore:"jobs"`
	SubJobs map[string]*Job `json:"subjobs" firestore:"subjobs"`
	JSON    string          `json:"json" firestore:"json"`
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

func (job *Job) Validate(w http.ResponseWriter, m map[string]interface{}) bool {

	var exists bool
	job.Name, exists = AssertKeyValue(w, m, "name")
	if !exists {
		return false
	}
	job.Status, exists = AssertKeyValue(w, m, "status")
	return exists
}
