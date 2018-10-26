package irc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCTCP(t *testing.T) {
	t.Run("bad inputs", func(t *testing.T) {
		tests := []string{
			"",
			"\x01",
			"\x01a",
			"\x01\x01a",
			"a\x01",
			"aa",
		}

		for _, test := range tests {
			command, args, ok := ParseCTCP(test)
			assert.Empty(t, command)
			assert.Empty(t, args)
			assert.False(t, ok)
		}
	})

	t.Run("valid inputs", func(t *testing.T) {
		tests := []struct {
			input   string
			command string
			args    string
		}{
			{
				input: "\x01\x01",
			},
			{
				input:   "\x01ACTION\x01",
				command: "ACTION",
			},
			{
				input:   "\x01ACTION \x01",
				command: "ACTION",
			},
			{
				input:   "\x01ACTION foo\x01",
				command: "ACTION",
				args:    "foo",
			},
			{
				input:   "\x01ACTION foo bar\x01",
				command: "ACTION",
				args:    "foo bar",
			},
		}

		for _, test := range tests {
			command, args, ok := ParseCTCP(test.input)
			assert.Equal(t, test.command, command)
			assert.Equal(t, test.args, args)
			assert.True(t, ok)
		}
	})
}

func TestEncodeCTCP(t *testing.T) {
	t.Run("empty command", func(t *testing.T) {
		s, err := EncodeCTCP("", "")
		assert.Equal(t, s, "")
		assert.Equal(t, err, ErrCTCPEmptyCommand)
	})

	t.Run("valid inputs", func(t *testing.T) {
		tests := []struct {
			command  string
			args     string
			expected string
		}{
			{
				command:  "ACTION",
				expected: "\x01ACTION\x01",
			},
			{
				command:  "ACTION",
				args:     "says hi",
				expected: "\x01ACTION says hi\x01",
			},
		}

		for _, test := range tests {
			s, err := EncodeCTCP(test.command, test.args)
			assert.Nil(t, err)
			assert.Equal(t, test.expected, s)
		}
	})
}

func BenchmarkParseCTCP(b *testing.B) {
	raw := "\x01ACTION says hi\x01"

	b.SetBytes(int64(len(raw)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseCTCP(raw) //nolint:errcheck
	}
}

func BenchmarkEncodeCTCP(b *testing.B) {
	command := "ACTION"
	args := "says hi"

	b.SetBytes(int64(len(command) + len(args) + 3))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeCTCP(command, args) //nolint:errcheck
	}
}

func BenchmarkEncodeCTCPNoArgs(b *testing.B) {
	command := "ACTION"

	b.SetBytes(int64(len(command) + 2))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeCTCP(command, "") //nolint:errcheck
	}
}
