package eblance

import (
	"fmt"
	"time"
)

type Center struct {
	jobPool       *JobPool
	agentPool     *AgentPool
	JobChannel    chan *Job
	ResultChannel chan *JobResult
}

func (center *Center) AddJob(job *Job) bool {
	return center.jobPool.AddJob(job)
}

func (center *Center) GetJob(id JobID) *Job {
	job, _ := center.jobPool.GetJob(id)
	return job
}

func (center *Center) Start() {
	// listen for results
	go func() {
		for result := range center.ResultChannel {
			fmt.Printf("Job %d finished: %d\n", result.Job.ID, result.Code)
		}
	}()

	// Keep scheduling job
	go func() {
		for {
			if center.jobPool.HasJob() {
				job := center.jobPool.NextJob()
				if job != nil {
					center.JobChannel <- job
				}
			} else {
				time.Sleep(time.Second * 1)
			}
		}
	}()
}

func (center *Center) Stop() {
	close(center.ResultChannel)
}

func NewCenter() *Center {
	jobPool := JobPool{}
	agentPool := AgentPool{}
	return &Center{
		&jobPool,
		&agentPool,
		make(chan *Job),
		make(chan *JobResult),
	}
}
