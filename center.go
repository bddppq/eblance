package eblance

type Center struct {
	agents        []*Agent
	jobPool       JobPool
	agentPool     AgentPool
	ResultChannel chan<- JobResult
}

func (center *Center) AddJob(job *Job) bool {
	return center.jobPool.AddJob(job)
}

func (center *Center) GetJob(id JobID) *Job {
	job, _ := center.jobPool.GetJob(id)
	return job
}
