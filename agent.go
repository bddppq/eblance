package eblance

import (
	"golang.org/x/crypto/ssh"
	"net"
)

type Agent struct {
	host *remoteHost
}

type remoteHost struct {
	hostname   string
	center     *Center
	connection *net.Conn
}

func (h *remoteHost) Start() {
	startAgent(h)
}

func (h *remoteHost) GetCenter() *Center {
	return h.center
}

func (h *remoteHost) Execute(job *Job) {
	session, err := h.connection.NewSession()
	defer session.Close()
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)
	}
	session.Run(string(job.Command))
}

func NewAgent(name string, center *Center) Agent {
	host := &remoteHost{
		name,
		center,
	}
	return &Agent{
		host,
	}
}
