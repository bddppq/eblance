package eblance

import (
	"time"
)

type JobID int64
type Command string

type JobStatus int8

const (
	Pending JobStatus = iota
	Running
	Success
	Failed
)

type JobResult struct {
	job  *Job
	out  string
	err  string
	code int
}

type Job struct {
	ID         JobID
	Status     JobStatus
	Command    Command
	Agent      *Agent
	Result     JobResult
	SubmitTime time.Time
	StartTime  time.Time
	EndTime    time.Time
}
