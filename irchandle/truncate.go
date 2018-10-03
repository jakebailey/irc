package irchandle

import (
	"context"
	"errors"

	"github.com/jakebailey/irc"
)

var ErrMessageTooLong = errors.New("message too long to truncate")

func Truncate(length int) func(Handler) Handler {
	return func(handler Handler) Handler {
		return HandlerFunc(func(ctx context.Context, e irc.Encoder, m *irc.Message) {
			e = truncator{length, e}
			handler.HandleMessage(ctx, e, m)
		})
	}
}

type truncator struct {
	length int
	e      irc.Encoder
}

func (t truncator) Encode(m *irc.Message) error {
	mLen := m.Len()
	if mLen <= t.length {
		return t.e.Encode(m)
	}

	diff := mLen - t.length

	i := len(m.Trailing) - diff

	if i > 0 {
		m.Trailing = m.Trailing[:i]
		return t.e.Encode(m)
	}

	return ErrMessageTooLong
}
