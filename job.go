package eblance

import (
	"sync/atomic"
	"time"
)

type JobID uint64
type Command string

type JobStatus int8
type ExitCode int

var jobCount uint64 = 0

const (
	Pending JobStatus = iota
	Running
	Success
	Failed
)

var NullTime time.Time = time.Unix(0, 0)

type JobResult struct {
	Job  *Job
	Out  string
	Err  string
	Code ExitCode
}

type Job struct {
	ID         JobID
	Status     JobStatus
	Command    Command
	Agent      *Agent
	Result     *JobResult
	SubmitTime time.Time
	StartTime  time.Time
	EndTime    time.Time
}

func NewJob(command string) *Job {
	id := atomic.AddUint64(&jobCount, 1)
	return &Job{
		JobID(id),
		Pending,
		Command(command),
		nil, // agent not assigned yet
		nil, // result unknow yet
		time.Now(),
		NullTime,
		NullTime,
	}
}

func IsNull(t time.Time) bool {
	return t == NullTime
}
