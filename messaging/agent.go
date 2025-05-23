package messaging

import "strings"

const (
	ChannelSize          = 16
	AssignmentIdentifier = "#"
)

// NewAgent - agent constructor
type NewAgent func() Agent

// Agent - agent
type Agent interface {
	Uri() string
	Message(m *Message)
}

func Shutdown(agent Agent) {
	if agent != nil {
		agent.Message(ShutdownMessage)
	}
}

func Startup(agent Agent) {
	if agent != nil {
		agent.Message(StartupMessage)
	}
}

func Name(agent Agent) string {
	if agent == nil {
		return ""
	}
	return getName(agent.Uri())
}

func getName(uri string) string {
	i := strings.Index(uri, AssignmentIdentifier)
	if i == -1 {
		return uri
	}
	return uri[:i]
}
