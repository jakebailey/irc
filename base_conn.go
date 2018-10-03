package irc

import (
	"bufio"
	"io"
	"net"
	"sync"
)

// BaseConn is a simple IRC connection.
type BaseConn struct {
	conn    net.Conn
	scanner *bufio.Scanner
	mu      sync.Mutex
}

var _ Conn = (*BaseConn)(nil)

// BaseDial dials the provided address and returns a BaseConn.
func BaseDial(addr string) (*BaseConn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	b := &BaseConn{
		conn:    conn,
		scanner: bufio.NewScanner(conn),
	}
	return b, nil
}

// Close closes the underlying connection.
func (b *BaseConn) Close() error {
	return b.conn.Close()
}

// Encode encodes a message over the connection.
func (b *BaseConn) Encode(m *Message) error {
	_, err := m.WriteToWithNewline(b.conn)
	return err
}

// Decode decodes a message into the argument, which cannot be nil.
func (b *BaseConn) Decode(m *Message) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.scanner.Scan() {
		if err := b.scanner.Err(); err != nil {
			return err
		}
		return io.EOF
	}
	return m.Parse(b.scanner.Text())
}
