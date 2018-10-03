package irchandle

func Chain(base Handler, middlewares ...func(Handler) Handler) Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		base = middlewares[i](base)
	}
	return base
}
