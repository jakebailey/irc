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
}

// Prefix is an IRC prefix.
type Prefix struct {
	Name string
	User string
	Host string
}
