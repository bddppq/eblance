package eblance

type Agent interface {
	AddToCenter(center *Center)
	GetJobChannel() chan<- *Job
}

type remoteHost struct {
}

type localHost struct {
}

func NewAgent(name string) *Agent {
	return nil
}
