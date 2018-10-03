package irc

// Encoder takes a message and encodes it, sending it to wherever the
// encoder chooses.
type Encoder interface {
	Encode(*Message) error
}

// Decoder decodes a message into the provided Message struct. Decode
// must set every field of the message struct.
type Decoder interface {
	Decode(*Message) error
}

// Conn is both an Encoder and Decoder, but also includes a Close function.
type Conn interface {
	Encoder
	Decoder
	Close() error
}
