package irchandle

import (
	"context"

	"github.com/jakebailey/irc"
)

// Handler handles an IRC message. Its definition mirrors http.Handler.
type Handler interface {
	HandleMessage(ctx context.Context, e irc.Encoder, m *irc.Message)
}

// HandlerFunc is a type adapter to use handler functions as Handlers. Its
// definition mirrors http.HandlerFunc.
type HandlerFunc func(ctx context.Context, e irc.Encoder, m *irc.Message)

// HandleMessage calls f(ctx, e, m).
func (h HandlerFunc) HandleMessage(ctx context.Context, e irc.Encoder, m *irc.Message) {
	h(ctx, e, m)
}
