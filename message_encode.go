package irc

import (
	"bytes"
	"io"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func (m *Message) String() string {
	buf := m.buffer()
	s := buf.String()
	bufferPool.Put(buf)
	return s
}

// Bytes returns the message encoded as a byte slice. This slice is safe for
// reuse.
func (m *Message) Bytes() []byte {
	buf := m.buffer()
	arr := buf.Bytes()
	arr2 := make([]byte, len(arr))
	copy(arr2, arr)
	bufferPool.Put(buf)
	return arr2
}

// WriteToWithNewline writes the message to a writer with a terminating `\r\n`.
func (m *Message) WriteToWithNewline(w io.Writer) (n int64, err error) {
	buf := m.buffer()
	buf.WriteString("\r\n")
	n, err = buf.WriteTo(w)
	bufferPool.Put(buf)
	return n, err
}

// MarshalText implements TextMarshaler for Message.
func (m *Message) MarshalText() (text []byte, err error) {
	return m.Bytes(), nil
}

// UnmarshalText implements TextUnmarshaler for Message.
func (m *Message) UnmarshalText(text []byte) error {
	return m.Parse(string(text))
}

func (m *Message) buffer() *bytes.Buffer {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	if len(m.Tags) != 0 {
		buf.WriteByte('@')

		sep := false
		for k, v := range m.Tags {
			if sep {
				buf.WriteByte(';')
			} else {
				sep = true
			}

			buf.WriteString(k)

			if v != "" {
				buf.WriteByte('=')

				// If the value doesn't have a character that needs escaping,
				// then don't run the Replacer, as it appears to do a lot of
				// work even when replacing nothing. The order of chars to
				// check is generally ordered by how often those characters
				// are used in a tag.
				if containsEscapable(v) {
					// buf is a bytes.Buffer, so writing to it cannot fail.
					tagEscape.WriteString(buf, v) //nolint:errcheck
				} else {
					buf.WriteString(v)
				}
			}
		}

		buf.WriteByte(' ')
	}

	// A prefix must always have a servername/nick, so check this first,
	// then decide to include the user and host.
	if m.Prefix.Name != "" {
		buf.WriteByte(':')
		buf.WriteString(m.Prefix.Name)

		if m.Prefix.User != "" {
			buf.WriteByte('!')
			buf.WriteString(m.Prefix.User)
		}

		if m.Prefix.Host != "" {
			buf.WriteByte('@')
			buf.WriteString(m.Prefix.Host)
		}

		buf.WriteByte(' ')
	}

	buf.WriteString(m.Command)

	for _, p := range m.Params {
		buf.WriteByte(' ')
		buf.WriteString(p)
	}

	if m.Trailing != "" || m.ForcedTrailing {
		buf.WriteString(" :")
		buf.WriteString(m.Trailing)
	}

	return buf
}

// Len returns the length of the encoded message. This message does not
// actually encode the message, instead simulating encoding and only calculates
// the length.
func (m *Message) Len() int {
	length := 0

	if len(m.Tags) > 0 || m.ForcedTags {
		length += len(m.Tags) + 1

		for k, v := range m.Tags {
			length += len(k)

			if v != "" {
				length += len(v) + 1
				length += countEscapeable(v)
			}
		}
	}

	if m.Prefix.Name != "" {
		length += len(m.Prefix.Name) + 2

		if m.Prefix.User != "" {
			length += len(m.Prefix.User) + 1
		}

		if m.Prefix.Host != "" {
			length += len(m.Prefix.Host) + 1
		}
	}

	length += len(m.Command) + 1

	if len(m.Params) > 0 {
		length += len(m.Params)
		for _, p := range m.Params {
			length += len(p)
		}
	}

	if m.Trailing != "" || m.ForcedTrailing {
		length += len(m.Trailing) + 1
	}

	return length
}
