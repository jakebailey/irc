package irc

import (
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseConn(t *testing.T) {
	m, err := ParseMessage(rawTwitch)
	assert.NoError(t, err)
	m.Raw = "" // Don't check raw.

	sender, receiver := net.Pipe()

	go testBaseConnSender(t, NewBaseConn(sender), m)
	testBaseConnReceiver(t, NewBaseConn(receiver), m)
}

func TestBaseDial(t *testing.T) {
	m, err := ParseMessage(rawTwitch)
	assert.NoError(t, err)
	m.Raw = "" // Don't check raw.

	laddr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	assert.NoError(t, err)

	listener, err := net.ListenTCP("tcp", laddr)
	assert.NoError(t, err)
	defer listener.Close()

	go func() {
		sender, err := listener.Accept()
		assert.NoError(t, err)
		testBaseConnSender(t, NewBaseConn(sender), m)
	}()

	receiver, err := BaseDial(listener.Addr().String())
	assert.NoError(t, err)
	testBaseConnReceiver(t, receiver, m)
}

func testBaseConnSender(t *testing.T, sConn *BaseConn, m *Message) {
	defer assertClose(t, sConn)

	err := sConn.Encode(m)
	assert.NoError(t, err)
}

func testBaseConnReceiver(t *testing.T, rConn *BaseConn, m *Message) {
	defer assertClose(t, rConn)

	var got Message

	err := rConn.Decode(&got)
	assert.NoError(t, err)

	got.Raw = "" // Don't check raw.

	assert.Equal(t, m, &got)

	err = rConn.Decode(&got)
	assert.Equal(t, io.EOF, err)
}

func TestBaseConnDecodeErr(t *testing.T) {
	sender, receiver := net.Pipe()
	assertClose(t, sender)
	assertClose(t, receiver)

	rConn := NewBaseConn(receiver)

	err := rConn.Decode(&Message{})
	assert.Equal(t, io.ErrClosedPipe, err)
}

func assertClose(t *testing.T, closer interface{ Close() error }) {
	err := closer.Close()
	assert.NoError(t, err)
}

func TestBaseConnDialErr(t *testing.T) {
	_, err := BaseDial("")
	assert.EqualError(t, err, "dial tcp: missing address")
}
