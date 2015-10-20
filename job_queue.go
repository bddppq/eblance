package eblance

import (
	"github.com/oleiade/lane"
)

type JobQueue struct {
	*lane.Queue
}

func NewJobQueue(name string) *JobQueue {
	return &JobQueue{
		lane.NewQueue(),
	}
}

func (h *JobQueue) AddJob(job *Job) {
	h.Enqueue(job)
}

func (h *JobQueue) HasJob() bool {
	return h.Head() != nil
}

func (h *JobQueue) NextJob() *Job {
	return h.Dequeue().(*Job)
}
