package eblance

import (
	"fmt"
)

type JobPool struct {
	name        string
	jobs        map[JobID]*Job
	pendingJobs JobQueue
}

func (pool *JobPool) addJob(job *Job) error {
	if _, ok := pool.jobs[job.ID]; ok {
		return fmt.Errorf("Can not add Job due to one with same jobID %d already exists in pool %s", job.ID, pool.name)
	}

	pool.jobs[job.ID] = job
	return nil
}

func (pool *JobPool) AddJob(job *Job) bool {
	err := pool.addJob(job)
	if err == nil {
		return false
	}
	pool.pendingJobs.AddJob(job)
	return true
}

func (pool *JobPool) GetJob(id JobID) (*Job, error) {
	if job, ok := pool.jobs[id]; ok {
		return job, nil
	} else {
		return nil, fmt.Errorf("Job with id %d is not in pool %s.", id, pool.name)
	}
}

func (pool *JobPool) HasJob() bool {
	return pool.pendingJobs.HasJob()
}

func (pool *JobPool) NextJob() *Job {
	return pool.pendingJobs.NextJob()
}

type AgentPool struct {
	name   string
	agents []*Agent
}
