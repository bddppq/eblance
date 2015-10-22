package eblance

type Agent interface {
	Start()
	GetCenter() *Center
	Execute(job *Job)
}

type remoteHost struct {
	center *Center
}

type localHost struct {
	center *Center
}

func (h *remoteHost) Start() {
	startAgent(h)
}

func (h *localHost) Start() {
	startAgent(h)
}

func (h *remoteHost) GetCenter() *Center {
	return h.center
}

func (h *localHost) GetCenter() *Center {
	return h.center
}

func (h *remoteHost) Execute(job *Job) {
	execute(h, job)
}

func (h *localHost) Execute(job *Job) {
	execute(h, job)

}

func startAgent(agent Agent) {
	for {
		// TODO
	}
}

func execute(agent Agent, job *Job) {
	go func() {
		// TODO
	}()
}

func NewAgent(name string, center *Center) Agent {
	// TODO: add more check here (e.g. ip address) to identify localhost
	// Or do we really want to support localhost?
	if name == "localhost" {
		return &localHost{
			center,
		}
	} else {
		return &remoteHost{
			center,
		}
	}
}
