package irc

// Message is an IRC message.
type Message struct {
	Tags           map[string]string
	Prefix         Prefix
	Command        string
	Params         []string
	Trailing       string
	ForcedTags     bool
	ForcedTrailing bool

	// Raw contains the raw unparsed message. This is not used for encoding,
	// and is included for users which want the exact original message.
	// Including this incurs no extra overhead, since the other parts of the
	// message are references to parts of this string, meaning that the
	// lifetime of this string is just as long as the message as a whole.
	Raw string
}

// Prefix is an IRC prefix.
type Prefix struct {
	Name string
	User string
	Host string
}
