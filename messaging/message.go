package messaging

import (
	"errors"
	"fmt"
	"net/http"
)

//StartEvent    = "event:start"
//StopEvent     = "event:stop"

const (
	StartupEvent  = "event:startup"
	ShutdownEvent = "event:shutdown"
	PauseEvent    = "event:pause"  // disable data channel receive
	ResumeEvent   = "event:resume" // enable data channel receive
	ConfigEvent   = "event:config"
	StatusEvent   = "event:status"

	ObservationEvent = "event:observation"
	TickEvent        = "event:tick"
	DataChangeEvent  = "event:data-change"

	Master   = "master"
	Emissary = "emissary"
	Control  = "ctrl"
	Data     = "data"

	XTo        = "x-to"
	XFrom      = "x-from"
	XEvent     = "x-event"
	XChannel   = "x-channel"
	XRelatesTo = "x-relates-to"

	ContentType           = "Content-Type"
	ContentTypeError      = "application/error"
	ContentTypeMap        = "application/map"
	ContentTypeStatus     = "application/status"
	ContentTypeEventing   = "application/eventing"
	ContentTypeDispatcher = "application/dispatcher"
)

var (
	StartupMessage  = NewMessage(Control, StartupEvent)
	ShutdownMessage = NewMessage(Control, ShutdownEvent)
	PauseMessage    = NewMessage(Control, PauseEvent)
	ResumeMessage   = NewMessage(Control, ResumeEvent)

	EmissaryShutdownMessage = NewMessage(Emissary, ShutdownEvent)
	MasterShutdownMessage   = NewMessage(Master, ShutdownEvent)
)

// Handler - uniform interface for message handling
type Handler func(msg *Message)

// Message - message
type Message struct {
	Header http.Header
	Body   any
	Reply  Handler
}

func NewMessage(channel, event string) *Message {
	m := new(Message)
	m.Header = make(http.Header)
	m.Header.Add(XChannel, channel)
	m.Header.Add(XEvent, event)
	return m
}

func NewMessageWithError(channel, event string, err error) *Message {
	m := NewMessage(channel, event)
	m.SetContent(ContentTypeError, err)
	return m
}

func NewAddressableMessage(channel, to, from, event string) *Message {
	m := new(Message)
	m.Header = make(http.Header)
	m.Header.Add(XChannel, channel)
	m.Header.Add(XTo, to)
	m.Header.Add(XFrom, from)
	m.Header.Add(XEvent, event)
	return m
}

func (m *Message) String() string {
	return fmt.Sprintf("[chan:%v] [from:%v] [to:%v] [%v]", m.Channel(), m.From(), m.To(), m.Event())
	//return fmt.Sprintf("[chan:%v] [%v]", m.Channel(), m.Event())
}

func (m *Message) RelatesTo() string {
	return m.Header.Get(XRelatesTo)
}

func (m *Message) To() string {
	return m.Header.Get(XTo)
}

func (m *Message) SetTo(uri string) *Message {
	m.Header.Set(XTo, uri)
	return m
}

func (m *Message) From() string {
	return m.Header.Get(XFrom)
}

func (m *Message) SetFrom(uri string) *Message {
	m.Header.Set(XFrom, uri)
	return m
}

func (m *Message) Event() string {
	return m.Header.Get(XEvent)
}

func (m *Message) Channel() string {
	return m.Header.Get(XChannel)
}

func (m *Message) SetChannel(channel string) *Message {
	m.Header.Set(XChannel, channel)
	return m
}

func (m *Message) SetContentType(contentType string) *Message {
	if len(contentType) == 0 {
		return m //errors.New("error: content type is empty")
	}
	m.Header.Add(ContentType, contentType)
	return m
}

func (m *Message) ContentType() string {
	return m.Header.Get(ContentType)
}

func (m *Message) SetContent(contentType string, content any) error {
	if len(contentType) == 0 {
		return errors.New("error: content type is empty")
	}
	if content == nil {
		return errors.New("error: content is nil")
	}
	m.Body = content
	m.Header.Add(ContentType, contentType)
	return nil
}

func NewConfigMapMessage(cfg map[string]string) *Message {
	m := NewMessage(Control, ConfigEvent)
	m.SetContent(ContentTypeMap, cfg)
	return m
}

func ConfigMapContent(m *Message) map[string]string {
	if m.Event() != ConfigEvent || m.ContentType() != ContentTypeMap {
		return nil
	}
	if cfg, ok := m.Body.(map[string]string); ok {
		return cfg
	}
	return nil
}

func NewStatusMessage(status *Status, relatesTo string) *Message {
	m := NewMessage(Control, StatusEvent)
	m.SetContent(ContentTypeStatus, status)
	if relatesTo != "" {
		m.Header.Add(XRelatesTo, relatesTo)
	}
	return m
}

func StatusContent(m *Message) (*Status, string) {
	if m.Event() != StatusEvent || m.ContentType() != ContentTypeStatus {
		return nil, ""
	}
	if s, ok := m.Body.(*Status); ok {
		return s, m.RelatesTo()
	}
	return nil, ""
}

func NewEventingHandlerMessage(agent Agent) *Message {
	m := NewMessage(Control, ConfigEvent)
	m.SetContent(ContentTypeEventing, agent)
	return m
}

func EventingHandlerContent(m *Message) (Agent, bool) {
	if m.Event() != ConfigEvent || m.ContentType() != ContentTypeEventing {
		return nil, false
	}
	if v, ok := m.Body.(Agent); ok {
		return v, true
	}
	return nil, false
}

func NewDispatcherMessage(dispatcher Dispatcher) *Message {
	m := NewMessage(Control, ConfigEvent)
	m.SetContent(ContentTypeDispatcher, dispatcher)
	return m
}

func DispatcherContent(m *Message) (Dispatcher, bool) {
	if m.Event() != ConfigEvent || m.ContentType() != ContentTypeDispatcher {
		return nil, false
	}
	if v, ok := m.Body.(Dispatcher); ok {
		return v, true
	}
	return nil, false
}

// Reply - function used by message recipient to reply with a Status
func Reply(msg *Message, status *Status, from string) {
	if msg == nil || status == nil || msg.Reply == nil {
		return
	}
	m := NewStatusMessage(status, msg.Event())
	m.Header.Set(XFrom, from)
	msg.Reply(m)
}
