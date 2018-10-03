package irchandle

import (
	"context"
	"sync"

	"github.com/jakebailey/irc"
)

// Client reads messages from the specified Conn and handles them using the
// provided Handler.
type Client struct {
	Conn    irc.Conn
	Handler Handler
	Context context.Context

	// Sync instructs the client to handle messages synchronously, such that
	// a message will not be handled until the previous message has been.
	Sync bool

	// Pool enables pooling of messsages. This lowers memory usage, but
	// messages given to the handler cannot be used once the handler returns.
	// Sync overrides this option.
	Pooled bool
}

// Run starts the client.
func (c *Client) Run() error {
	if c.Context == nil {
		c.Context = context.Background()
	}

	if c.Sync {
		return c.runSync()
	}

	if c.Pooled {
		return c.runPooled()
	}

	for {
		m := &irc.Message{}
		if err := c.Conn.Decode(m); err != nil {
			return err
		}

		go c.Handler.HandleMessage(c.Context, c.Conn, m)
	}
}

func (c *Client) runSync() error {
	for {
		var m irc.Message
		if err := c.Conn.Decode(&m); err != nil {
			return err
		}

		c.Handler.HandleMessage(c.Context, c.Conn, &m)
	}
}

var messagePool = sync.Pool{
	New: func() interface{} { return &irc.Message{} },
}

func (c *Client) runPooled() error {
	for {
		m := messagePool.Get().(*irc.Message)
		if err := c.Conn.Decode(m); err != nil {
			return err
		}

		go func(m *irc.Message) {
			c.Handler.HandleMessage(c.Context, c.Conn, m)
			messagePool.Put(m)
		}(m)
	}
}
