package messaging

import (
	"errors"
)

type controlAgent struct {
	running bool
	uri     string
	name    string
	ch      chan *Message
	handler Handler
}

// NewControlAgent - create an agent that only listens on a control channel, and has a default AgentRun func
func NewControlAgent(uri string, handler Handler) (Agent, error) {
	if handler == nil {
		return nil, errors.New("error: control agent handler is nil")
	}
	return newControlAgent(uri, make(chan *Message, ChannelSize), handler), nil
	//return NewAgentWithChannels(uri, nil, nil, controlAgentRun, ctrlHandler)
}

func newControlAgent(uri string, ch chan *Message, handler Handler) *controlAgent {
	c := new(controlAgent)
	c.uri = uri
	c.ch = ch
	c.handler = handler
	return c
}

// Uri - identity
func (c *controlAgent) Uri() string { return c.uri }

// String - identity
func (c *controlAgent) String() string { return c.Uri() }

// Message - message an agent
func (c *controlAgent) Message(msg *Message) {
	if msg == nil {
		return
	}
	if msg.Event() == StartupEvent {
		c.run()
		return
	}
	if !c.running {
		return
	}
	switch msg.Channel() {
	case ChannelControl:
		if c.ch != nil {
			c.ch <- msg
		}
	default:
	}
}

// Run - run the agent
func (c *controlAgent) run() {
	if c.running {
		return
	}
	c.running = true
	go controlAgentRun(c)
}

// Shutdown - shutdown the agent
/*
func (c *controlAgent) Shutdown() {
	if !c.running {
		return
	}
	c.running = false
	c.Message(Shutdown)
}
*/

func (c *controlAgent) shutdown() {
	close(c.ch)
}

// controlAgentRun - a simple run function that only handles control messages, and dispatches via a message handler
func controlAgentRun(c *controlAgent) {
	if c == nil || c.handler == nil {
		return
	}
	// ctrlHandler Handler
	//if h, ok := state.(Handler); ok {
	//	ctrlHandler = h
	//} else {
	//	return
	//}
	for {
		select {
		case msg, open := <-c.ch:
			if !open {
				return
			}
			switch msg.Event() {
			case ShutdownEvent:
				c.handler(NewMessageWithError(ChannelControl, msg.Event(), nil))
				c.shutdown()
				return
			default:
				c.handler(msg)
			}
		default:
		}
	}
}
