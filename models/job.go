package models

import (
	"net/http"

	"github.com/golangdaddy/gethealthy/models"
)

type Job struct {
	Meta    Internals
	Name    string   `json:"name" firestore:"name"`
	Status  string   `json:"status" firestore:"status"`
	Jobs    []string `json:"subjobs" firestore:"subjobs"`
	SubJobs map[string]*Job
	JSON    string
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
	job.Name, exists = models.AssertKeyValue(w, m, "name")
	if !exists {
		return false
	}
	job.Status, exists = models.AssertKeyValue(w, m, "status")
	return exists
}
