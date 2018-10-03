package irc

import (
	"errors"
	"strings"
)

// ErrCTCPEmptyCommand is returned by EncodeCTCP when the command is empty.
var ErrCTCPEmptyCommand = errors.New("empty ctcp command")

// ParseCTCP parses the commands and args out of a CTCP string.
func ParseCTCP(s string) (command, args string, ok bool) {
	if len(s) < 2 || s[0] != '\x01' || s[len(s)-1] != '\x01' {
		return "", "", false
	}

	s = s[1 : len(s)-1]

	i := strings.IndexByte(s, ' ')
	if i == -1 {
		return s, "", true
	}

	return s[:i], s[i+1:], true
}

// EncodeCTCP encodes a command and its arguments into a string.
func EncodeCTCP(command, args string) (string, error) {
	if command == "" {
		return "", ErrCTCPEmptyCommand
	}

	if args == "" {
		return "\x01" + command + "\x01", nil
	}

	return "\x01" + command + " " + args + "\x01", nil
}
