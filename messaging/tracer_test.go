package messaging

import (
	"fmt"
)

type agent struct {
	agentId string
	ch      *Channel
}

func NewAgent(uri string, ch *Channel) Agent {
	a := new(agent)
	a.agentId = uri
	a.ch = ch
	return a
}

func (t *agent) Uri() string           { return t.agentId }
func (t *agent) Name() string          { return t.agentId }
func (t *agent) Message(m *Message)    { fmt.Printf("test: opsAgent.Message() -> %v\n", m) }
func (t *agent) Notify(status *Status) { fmt.Printf("%v", status) }

// func (t *agent) IsFinalized() bool     { return t.ch.IsFinalized() }
func (t *agent) Run() {}
func (t *agent) Shutdown() {
	if t.ch != nil {
		t.ch.Close()
		t.ch = nil
	}
}

func _ExampleDefaultTracer_Trace() {
	a := NewAgent("agent:test", NewEmissaryChannel())
	DefaultTracer.Trace(nil, Emissary, "event:shutdown", "agent shutdown")
	fmt.Printf("\n")

	DefaultTracer.Trace(a, Master, "event:shutdown", "agent shutdown")
	fmt.Printf("\n")

	//Output:
	//<nil> : emissary event:shutdown agent shutdown
	//agent:test : master event:shutdown agent shutdown

}

func ExampleAccess_No_Filter() {
	filter := NewTraceFilter("", "")
	channel := ""
	event := ""

	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))
	channel = "channel"
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))
	channel = ""
	event = "event"
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))

	//Output:
	//test: Access(,) -> [channel:] [event:] [access:true]
	//test: Access(,) -> [channel:channel] [event:] [access:true]
	//test: Access(,) -> [channel:] [event:event] [access:true]

}

func ExampleAccess_Channel() {
	filter := NewTraceFilter(Emissary, "")
	channel := ""
	event := ""

	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))
	channel = Emissary
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))

	channel = Master
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))

	channel = ""
	event = "event"
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))

	//Output:
	//test: Access(emissary,) -> [channel:] [event:] [access:false]
	//test: Access(emissary,) -> [channel:emissary] [event:] [access:true]
	//test: Access(emissary,) -> [channel:master] [event:] [access:false]
	//test: Access(emissary,) -> [channel:] [event:event] [access:false]

}

func ExampleAccess_Event() {
	filter := NewTraceFilter("", ShutdownEvent)
	channel := ""
	event := ""

	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))
	event = ShutdownEvent
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))
	event = StartupEvent
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))

	channel = Emissary
	event = StartupEvent
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))

	event = ShutdownEvent
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))

	//Output:
	//test: Access(,event:shutdown) -> [channel:] [event:] [access:false]
	//test: Access(,event:shutdown) -> [channel:] [event:event:shutdown] [access:true]
	//test: Access(,event:shutdown) -> [channel:] [event:event:startup] [access:false]
	//test: Access(,event:shutdown) -> [channel:emissary] [event:event:startup] [access:false]
	//test: Access(,event:shutdown) -> [channel:emissary] [event:event:shutdown] [access:true]

}

func ExampleAccess() {
	filter := NewTraceFilter(Emissary, ShutdownEvent)
	channel := ""
	event := ""

	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))
	event = ShutdownEvent
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))
	event = StartupEvent
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))

	channel = Emissary
	event = StartupEvent
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))

	channel = Emissary
	event = ShutdownEvent
	fmt.Printf("test: Access(%v,%v) -> [channel:%v] [event:%v] [access:%v]\n", filter.Channel, filter.Event, channel, event, filter.Access(channel, event))

	//Output:
	//test: Access(emissary,event:shutdown) -> [channel:] [event:] [access:false]
	//test: Access(emissary,event:shutdown) -> [channel:] [event:event:shutdown] [access:false]
	//test: Access(emissary,event:shutdown) -> [channel:] [event:event:startup] [access:false]
	//test: Access(emissary,event:shutdown) -> [channel:emissary] [event:event:startup] [access:false]
	//test: Access(emissary,event:shutdown) -> [channel:emissary] [event:event:shutdown] [access:true]

}
