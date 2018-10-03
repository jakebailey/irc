package irc

import (
	"io/ioutil"
	"testing"
)

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
