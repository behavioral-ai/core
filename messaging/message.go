package messaging

import (
	"errors"
	"fmt"
	"net/http"
)

// TODO : add support for control messages or restart, apply-changes, rollback-changes

const (
	StartupEvent  = "event:startup"
	ShutdownEvent = "event:shutdown"
	PauseEvent    = "event:pause"  // disable data channel receive
	ResumeEvent   = "event:resume" // enable data channel receive

	ObservationEvent = "event:observation"
	TickEvent        = "event:tick"
	DataChangeEvent  = "event:data-change"

	XTo      = "x-to"
	XFrom    = "x-from"
	XEvent   = "x-event"
	XChannel = "x-channel"

	//XAgentId           = "x-agent-id"
	//XForwardTo         = "x-forward-to"
	//ContentTypeStatus  = "application/status"
	//ContentTypeConfig  = "application/config"
	DataChannelType    = "DATA"
	ControlChannelType = "CTRL"

	//ChannelRight      = "RIGHT"
	//ChannelLeft       = "LEFT"
	//ChannelNone = "NONE"

	ContentType      = "Content-Type"
	ContentTypeError = "application/error"

	//XRelatesTo         = "x-relates-to"
	//XMessageId         = "x-message-id"
)

// Handler - uniform interface for message handling
type Handler func(msg *Message)

// Message - message
type Message struct {
	Header http.Header
	Body   any
}

func NewControlMessage(to, from, event string) *Message {
	return NewMessage(ControlChannelType, to, from, event)
}

func NewMessageWithError(channel, to, from, event string, err error) *Message {
	m := NewMessage(channel, to, from, event)
	m.SetContent(ContentTypeError, err)
	m.Body = err
	return m
}

func NewMessage(channel, to, from, event string) *Message {
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
}

func (m *Message) To() string {
	return m.Header.Get(XTo)
}

func (m *Message) SetTo(uri string) {
	m.Header.Set(XTo, uri)
}

func (m *Message) From() string {
	return m.Header.Get(XFrom)
}

func (m *Message) SetFrom(uri string) {
	m.Header.Set(XFrom, uri)
}

func (m *Message) Event() string {
	return m.Header.Get(XEvent)
}

func (m *Message) Channel() string {
	return m.Header.Get(XChannel)
}

func (m *Message) SetContentType(contentType string) {
	if len(contentType) == 0 {
		return //errors.New("error: content type is empty")
	}
	m.Header.Add(ContentType, contentType)
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
