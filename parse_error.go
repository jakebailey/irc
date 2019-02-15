package irc

var _ error = (*ParseError)(nil)

// ParseError is returned when parsing a message. It's useful for
// distinguishing parse errors from network errors when using Conn.Encode.
type ParseError struct {
	message string
}

func (p ParseError) Error() string {
	return p.message
}
