package irc

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeIgnoreRaw(t *testing.T) {
	raw := `@test=>\:=\:< :jake!jake@jake.com PRIVMSG #jake :Hello, World!`
	m, err := ParseMessage(raw)
	if assert.NoError(t, err) {
		m.Raw = "THIS SHOULD NOT BE USED"

		ms := m.String()
		assert.Equal(t, raw, ms, "encoded message should be equal")
		assert.Equal(t, len(ms), m.Len(), "encoded message length should be equal")
	}
}

func BenchmarkMessageEncode(b *testing.B) {
	m, err := ParseMessage(rawTwitch)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := m.buffer()
		bufferPool.Put(buf)
	}
}

func BenchmarkMessageWriteTo(b *testing.B) {
	m, err := ParseMessage(rawTwitch)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.WriteToWithNewline(ioutil.Discard) //nolint:errcheck
	}
}

func BenchmarkMessageLen(b *testing.B) {
	m, err := ParseMessage(rawTwitch)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Len()
	}
}
