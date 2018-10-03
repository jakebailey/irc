package irc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:lll
const rawTwitch = `@badges=;color=#1E90FF;display-name=atmao;emotes=360:39-46,48-55;id=b4bf69df-abf5-474b-9f1f-db441720bfb2;mod=0;room-id=54706574;subscriber=0;tmi-sent-ts=1493710725463;turbo=0;user-id=37012175;user-type= :atmao!atmao@atmao.tmi.twitch.tv PRIVMSG #joshog :I'm a scruncher but I'm ashamed off it FailFish FailFish`

func TestParseEncodeGood(t *testing.T) {
	t.Run("twitch message", func(t *testing.T) {
		expected := &Message{
			Tags: map[string]string{
				"badges":       "",
				"color":        "#1E90FF",
				"display-name": "atmao",
				"emotes":       "360:39-46,48-55",
				"id":           "b4bf69df-abf5-474b-9f1f-db441720bfb2",
				"mod":          "0",
				"room-id":      "54706574",
				"subscriber":   "0",
				"tmi-sent-ts":  "1493710725463",
				"turbo":        "0",
				"user-id":      "37012175",
				"user-type":    "",
			},
			Prefix: Prefix{
				Name: "atmao",
				User: "atmao",
				Host: "atmao.tmi.twitch.tv",
			},
			Command:  "PRIVMSG",
			Params:   []string{"#joshog"},
			Trailing: "I'm a scruncher but I'm ashamed off it FailFish FailFish",
		}

		m, err := ParseMessage(rawTwitch)
		if assert.NoError(t, err) {
			assert.Equal(t, expected, m, "messages should be equal")
		}
	})

	t.Run("no tags", func(t *testing.T) {
		raw := `:jake!jake@jake.com PRIVMSG #jake :Hello, World!`
		expected := &Message{
			Prefix: Prefix{
				Name: "jake",
				User: "jake",
				Host: "jake.com",
			},
			Command:  "PRIVMSG",
			Params:   []string{"#jake"},
			Trailing: "Hello, World!",
		}

		m, err := ParseMessage(raw)
		if assert.NoError(t, err) {
			assert.Equal(t, expected, m, "messages should be equal")

			ms := m.String()
			assert.Equal(t, raw, ms, "encoded message should be equal")
			assert.Equal(t, len(ms), m.Len(), "encoded message length should be equal")
		}
	})

	t.Run("single tag", func(t *testing.T) {
		raw := `@test=>\:=\:< :jake!jake@jake.com PRIVMSG #jake :Hello, World!`
		expected := &Message{
			Tags: map[string]string{
				"test": ">;=;<",
			},
			Prefix: Prefix{
				Name: "jake",
				User: "jake",
				Host: "jake.com",
			},
			Command:  "PRIVMSG",
			Params:   []string{"#jake"},
			Trailing: "Hello, World!",
		}

		m, err := ParseMessage(raw)
		if assert.NoError(t, err) {
			assert.Equal(t, expected, m, "messages should be equal")

			ms := m.String()
			assert.Equal(t, raw, ms, "encoded message should be equal")
			assert.Equal(t, len(ms), m.Len(), "encoded message length should be equal")
		}
	})

	t.Run("empty tag with equal", func(t *testing.T) {
		raw := `@test= :jake!jake@jake.com PRIVMSG #jake :Hello, World!`
		rawEnc := `@test :jake!jake@jake.com PRIVMSG #jake :Hello, World!`
		expected := &Message{
			Tags: map[string]string{
				"test": "",
			},
			Prefix: Prefix{
				Name: "jake",
				User: "jake",
				Host: "jake.com",
			},
			Command:  "PRIVMSG",
			Params:   []string{"#jake"},
			Trailing: "Hello, World!",
		}

		m, err := ParseMessage(raw)
		if assert.NoError(t, err) {
			assert.Equal(t, expected, m, "messages should be equal")

			ms := m.String()
			assert.Equal(t, rawEnc, ms, "encoded message should be equal")
			assert.Equal(t, len(ms), m.Len(), "encoded message length should be equal")
		}
	})

	t.Run("empty tag without equal", func(t *testing.T) {
		raw := `@test :jake!jake@jake.com PRIVMSG #jake :Hello, World!`
		expected := &Message{
			Tags: map[string]string{
				"test": "",
			},
			Prefix: Prefix{
				Name: "jake",
				User: "jake",
				Host: "jake.com",
			},
			Command:  "PRIVMSG",
			Params:   []string{"#jake"},
			Trailing: "Hello, World!",
		}

		m, err := ParseMessage(raw)
		if assert.NoError(t, err) {
			assert.Equal(t, expected, m, "messages should be equal")

			ms := m.String()
			assert.Equal(t, raw, ms, "encoded message should be equal")
			assert.Equal(t, len(ms), m.Len(), "encoded message length should be equal")
		}
	})

	t.Run("many tags", func(t *testing.T) {
		possible := []string{`@a;b=c TEST`, `@b=c;a TEST`}
		expected := &Message{
			Tags: map[string]string{
				"a": "",
				"b": "c",
			},
			Command: "TEST",
		}

		for _, raw := range possible {
			m, err := ParseMessage(raw)
			if assert.NoError(t, err) {
				assert.Equal(t, expected, m, "messages should be equal")
				assert.Contains(t, possible, m.String(), "encoded message should be possible")
			}
		}
	})

	t.Run("tag escaping", func(t *testing.T) {
		raw := `@a=\:\s\r\n\\ TEST`
		expected := &Message{
			Tags: map[string]string{
				"a": "; \r\n\\",
			},
			Command: "TEST",
		}

		m, err := ParseMessage(raw)
		if assert.NoError(t, err) {
			assert.Equal(t, expected, m, "messages should be equal")
			assert.Equal(t, raw, m.String(), "encoded message should be equal")
		}
	})

	t.Run("prefix formats", func(t *testing.T) {
		tests := []struct {
			raw string
			m   Message
		}{
			{
				raw: `:jake FOO`,
				m: Message{
					Prefix:  Prefix{Name: "jake"},
					Command: "FOO",
				},
			},
			{
				raw: `:jake!bake FOO`,
				m: Message{
					Prefix:  Prefix{Name: "jake", User: "bake"},
					Command: "FOO",
				},
			},
			{
				raw: `:jake@rake.bar FOO`,
				m: Message{
					Prefix:  Prefix{Name: "jake", Host: "rake.bar"},
					Command: "FOO",
				},
			},
			{
				raw: `:jake!bake@rake.bar FOO`,
				m: Message{
					Prefix:  Prefix{Name: "jake", User: "bake", Host: "rake.bar"},
					Command: "FOO",
				},
			},
		}

		for _, test := range tests {
			m, err := ParseMessage(test.raw)
			if assert.NoError(t, err) {
				assert.Equal(t, &test.m, m, "messages should be equal")
			}
		}
	})

	t.Run("only command", func(t *testing.T) {
		expected := &Message{Command: "TEST"}

		m, err := ParseMessage("TEST")
		if assert.NoError(t, err) {
			assert.Equal(t, expected, m, "messages should be equal")
		}
	})

	t.Run("empty trailing", func(t *testing.T) {
		expected := &Message{Command: "TEST", ForcedTrailing: true}

		m, err := ParseMessage("TEST :")
		if assert.NoError(t, err) {
			assert.Equal(t, expected, m, "messages should be equal")
		}
	})

	t.Run("parse function", func(t *testing.T) {
		m1, err1 := ParseMessage(`@test= :jake!jake@jake.com PRIVMSG #jake :Hello, World!`)
		var m2 Message
		err2 := m2.Parse(`@test= :jake!jake@jake.com PRIVMSG #jake :Hello, World!`)

		if assert.NoError(t, err1) && assert.NoError(t, err2) {
			assert.Equal(t, m1, &m2, "result of ParseMessage and Message.Parse should be equal")
		}
	})

	t.Run("parse overwrite", func(t *testing.T) {
		m := &Message{
			Tags: map[string]string{
				"a": "a",
				"b": "b",
			},
			Prefix: Prefix{
				Name: "name",
				User: "user",
				Host: "host",
			},
			Command:  "COMMAND",
			Params:   []string{"a", "b"},
			Trailing: "trailing",
		}

		err := m.Parse("@ PRIVMSG :")
		assert.NoError(t, err)

		expected := &Message{
			Command:        "PRIVMSG",
			ForcedTags:     true,
			ForcedTrailing: true,
		}

		assert.Equal(t, expected, m)
	})

	t.Run("backslash at end", func(t *testing.T) {
		raw := `@=\ `
		_, err := ParseMessage(raw)
		assert.Equal(t, ErrInvalidMessage, err)
	})
}

func BenchmarkParseMessageTwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ParseMessage(rawTwitch); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParseMessageEscaping(b *testing.B) {
	//nolint:lll
	raw := `@badges=;color=#1E90FF;display-name=\:atmao;emotes=360:39-46,48-55;id=\\\rb4bf69df-abf5-474b-9f1f-db441720bfb2;mod=0;room-id=54706574;subscriber=0;tmi-sent-ts=1493710725463;turbo=0;user-id=37012175;user-type= :atmao!atmao@atmao.tmi.twitch.tv PRIVMSG #joshog :I'm a scruncher but I'm ashamed off it FailFish FailFish`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ParseMessage(raw); err != nil {
			b.Fatal(err)
		}
	}
}