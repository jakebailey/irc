package irc

import (
	"errors"
	"strings"
)

var (
	// ErrEmptyMessage is returned when the parser encounters an empty message.
	ErrEmptyMessage = errors.New("empty message")

	// ErrInvalidMessage is returned when the parser encounters an invalid message.
	// This error is likely to be replaced with a more helpful error (encoding something
	// like where the error occurred).
	ErrInvalidMessage = errors.New("invalid message")
)

// Parse parses a string into a message. All fields of the message struct
// will be set (default as needed), so there is no need to zero before parsing.
//
// Parse does not check the input for newlines, which are normally invalid.
func (m *Message) Parse(raw string) error {
	return parseMessage(raw, m)
}

// ParseMessage parses a string and returns a new Message. A string is
// used as the input, as it was found that they're more performant than
// using a byte slice, even with an extra initial copy.
//
// ParseMessage does not check the input for newlines, which are normally invalid.
func ParseMessage(raw string) (*Message, error) {
	m := &Message{}

	if err := parseMessage(raw, m); err != nil {
		return nil, err
	}

	return m, nil
}

func parseMessage(raw string, m *Message) error {
	if raw == "" {
		return ErrEmptyMessage
	}

	// Easier and no slower to just zero out the message early.
	// It may be better to not do this selectively and have an option to
	// do something like keeping the tags map around (to be reused).
	*m = Message{
		Raw: raw,
	}

	var err error

	raw, err = m.parseTags(raw)
	if err != nil {
		return err
	}

	raw, err = m.parsePrefix(raw)
	if err != nil {
		return err
	}

	raw, err = m.parseCommand(raw)
	if err != nil {
		return err
	}

	if raw == "" {
		return nil
	}

	m.parseParamsAndTrailing(raw)

	return nil
}

func (m *Message) parseTags(raw string) (string, error) {
	if raw == "" {
		return "", ErrInvalidMessage
	}

	if raw[0] != '@' {
		return raw, nil
	}

	raw = raw[1:]

	i := strings.IndexByte(raw, ' ')
	if i == -1 {
		return "", ErrInvalidMessage
	}

	tags := raw[:i]
	raw = raw[i+1:]

	// <SPACE> can be many spaces, but the above stopped at the first space,
	// so trim off any other spaces.
	raw = strings.TrimLeftFunc(raw, isSpace)

	if tags == "" {
		m.ForcedTags = true
		return raw, nil
	}

	numTags := strings.Count(tags, ";") + 1
	m.Tags = make(map[string]string, numTags)

	for tags != "" {
		var pair string

		end := strings.IndexByte(tags, ';')
		if end == -1 {
			pair = tags
			tags = ""
		} else {
			pair = tags[:end]
			tags = tags[end+1:]
		}

		i := strings.IndexByte(pair, '=')

		if i == -1 {
			m.Tags[pair] = ""
			continue
		}

		k := pair[:i]
		v := pair[i+1:]

		m.Tags[k] = tagUnescape(v)
	}

	return raw, nil
}

func (m *Message) parsePrefix(raw string) (string, error) {
	if raw == "" {
		return "", ErrInvalidMessage
	}

	if raw[0] != ':' {
		return raw, nil
	}

	raw = raw[1:]

	i := strings.IndexByte(raw, ' ')
	if i == -1 {
		return "", ErrInvalidMessage
	}

	prefix := raw[:i]
	raw = raw[i+1:]

	user := strings.IndexByte(prefix, '!')
	host := strings.IndexByte(prefix, '@')

	if user > 0 && host > user {
		m.Prefix.Name = prefix[:user]
		m.Prefix.User = prefix[user+1 : host]
		m.Prefix.Host = prefix[host+1:]
	} else if user > 0 {
		m.Prefix.Name = prefix[:user]
		m.Prefix.User = prefix[user+1:]
	} else if host > 0 {
		m.Prefix.Name = prefix[:host]
		m.Prefix.Host = prefix[host+1:]
	} else {
		m.Prefix.Name = prefix
	}

	// <SPACE> can be many spaces, but the above stopped at the first space,
	// so trim off any other spaces.
	raw = strings.TrimLeftFunc(raw, isSpace)

	return raw, nil
}

func (m *Message) parseCommand(raw string) (string, error) {
	if raw == "" {
		return "", ErrInvalidMessage
	}

	i := strings.IndexByte(raw, ' ')
	if i == -1 {
		m.Command = raw
		return "", nil
	}

	m.Command = raw[:i]
	raw = raw[i+1:]

	// No space trimming is needed here, as parseParamsAndTrailing uses
	// stringFields to pick apart the params.

	return raw, nil
}

func (m *Message) parseParamsAndTrailing(raw string) {
	if i := strings.IndexByte(raw, ':'); i != -1 {
		m.Trailing = raw[i+1:]
		raw = raw[:i]

		// Set this up in case the message is reencoded.
		m.ForcedTrailing = m.Trailing == ""
	}

	params := stringFields(raw, ' ')

	if len(params) != 0 {
		m.Params = params
	}
}

func isSpace(r rune) bool {
	return r == ' '
}
