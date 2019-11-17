package server

import (
	"net"
)

// Listener ???
type Listener struct {
	net.Listener
	h *Handler
}

// NewListener creates a new Listener.
func NewListener(protocol, address string, handler *Handler) (*Listener, error) {
	l, err := net.Listen(protocol, address)
	if err != nil {
		return nil, err
	}
	return &Listener{l, handler}, nil
}

// Accept ????
func (l *Listener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	l.h.AddNetConnection(&conn)
	return conn, err
}
