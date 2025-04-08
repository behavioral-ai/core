package messaging

import (
	"errors"
)

const (
	SubscriptionCreateEvent = "event:subscription-create"
	SubscriptionCancelEvent = "event:subscription-cancel"
	ContentTypeSubscription = "application/x-subscription"
)

type Subscription struct {
	Channel string
	Event   string
	From    string
}

type Catalog struct {
	subs []Subscription
}

func (c *Catalog) Lookup(event string) (subs []Subscription, ok bool) {
	for _, item := range c.subs {
		if event == item.Event {
			subs = append(subs, item)
			ok = true
		}
	}
	return
}

func (c *Catalog) Create(s Subscription) error {
	if s.From == "" || s.Event == "" || s.Channel == "" {
		return errors.New("invalid subscription: from or event is empty")
	}
	for _, item := range c.subs {
		// Check if already subscribed
		if s.From == item.From && s.Event == item.Event {
			return nil
		}
	}
	c.subs = append(c.subs, s)
	return nil
}

func (c *Catalog) CreateWithMessage(m *Message) error {
	if m == nil {
		return nil
	}
	if s, ok := SubscriptionCreateContent(m); ok {
		err := c.Create(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Catalog) Cancel(s Subscription) {
	for i, item := range c.subs {
		if s.From == item.From && s.Event == item.Event {
			if len(c.subs) == 1 {
				c.subs = nil
			} else {
				if i == len(c.subs)-1 {
					c.subs = c.subs[:i]
				} else {
					first := c.subs[:i]
					last := c.subs[i+1:]
					c.subs = nil
					c.subs = append(c.subs, first...)
					c.subs = append(c.subs, last...)
				}
			}
			return
		}
	}
}

func (c *Catalog) CancelWithMessage(m *Message) {
	if m == nil {
		return
	}
	if s, ok := SubscriptionCancelContent(m); ok {
		c.Cancel(s)
	}
}

func NewSubscriptionCreateMessage(to, from, channel, event string) *Message {
	if to == "" || from == "" || event == "" {
		return nil
	}
	// Send to publishers control channel
	m := NewMessage(ChannelControl, SubscriptionCreateEvent)
	m.SetTo(to)
	m.SetFrom(from)
	// Allow subscriber to determine receive channel
	if channel == "" {
		channel = ChannelControl
	}
	m.SetContent(ContentTypeSubscription, Subscription{Channel: ChannelControl, From: from, Event: event})
	return m
}

func SubscriptionCreateContent(m *Message) (Subscription, bool) {
	if m == nil || m.Event() != SubscriptionCreateEvent || m.ContentType() != ContentTypeSubscription {
		return Subscription{}, false
	}
	if v, ok := m.Body.(Subscription); ok {
		return v, true
	}
	return Subscription{}, false
}

func NewSubscriptionCancelMessage(to, from, event string) *Message {
	if to == "" || from == "" || event == "" {
		return nil
	}
	m := NewMessage(ChannelControl, SubscriptionCancelEvent)
	m.SetTo(to)
	m.SetFrom(from)
	m.SetContent(ContentTypeSubscription, Subscription{From: from, Event: event})
	return m
}

func SubscriptionCancelContent(m *Message) (Subscription, bool) {
	if m == nil || m.Event() != SubscriptionCancelEvent || m.ContentType() != ContentTypeSubscription {
		return Subscription{}, false
	}
	if v, ok := m.Body.(Subscription); ok {
		return v, true
	}
	return Subscription{}, false
}
