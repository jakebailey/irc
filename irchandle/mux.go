package irchandle

import (
	"context"

	"github.com/jakebailey/irc"
)

type Mux struct {
	m        map[string]Handler
	notFound Handler
}

func NewMux() *Mux {
	return &Mux{
		m: make(map[string]Handler),
	}
}

var _ Handler = (*Mux)(nil)

func (mux *Mux) Handle(command string, handler Handler) {
	if _, ok := mux.m[command]; ok {
		panic("irc: multiple registrations for " + command)
	}
	mux.m[command] = handler
}

func (mux *Mux) HandleFunc(command string, handler func(context.Context, irc.Encoder, *irc.Message)) {
	if _, ok := mux.m[command]; ok {
		panic("irc: multiple registrations for " + command)
	}
	mux.m[command] = HandlerFunc(handler)
}

func (mux *Mux) HandleMessage(ctx context.Context, e irc.Encoder, m *irc.Message) {
	handler, ok := mux.m[m.Command]
	if !ok {
		handler = mux.notFound
	}

	if handler != nil {
		handler.HandleMessage(ctx, e, m)
	}
}

func (mux *Mux) NotFound(handler Handler) {
	mux.notFound = handler
}
