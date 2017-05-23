package net

import (
	"errors"
	"net"
)

// pipeListener listens to a virtual (or any) address, receives and accepts connections from PipeDialer.
type pipeListener struct {
	addr   net.Addr
	ch     chan net.Conn
	closer chan struct{}
}

// PipeDialer creates Pipe connections and passes them in a channel.
type PipeDialer struct {
	ch     chan net.Conn
	closer chan struct{}
}

// PipeListen creates a Listener-Dialer pair.
func PipeListen(a net.Addr) (net.Listener, *PipeDialer) {
	ch := make(chan net.Conn)
	closer := make(chan struct{})

	return &pipeListener{a, ch, closer}, &PipeDialer{ch, closer}
}

// Addr returns the listener's address.
func (l *pipeListener) Addr() net.Addr {
	return l.addr
}

// Accept returns a server connection (if any).
func (l *pipeListener) Accept() (net.Conn, error) {
	select {
	case conn := <-l.ch:
		return conn, nil
	case <-l.closer:
		return nil, errors.New("Listener closed")
	}
}

// Close closes the listener.
func (l *pipeListener) Close() error {
	close(l.closer)

	return nil
}

// Dial creates a new Pipe connection.
func (l *PipeDialer) Dial() (net.Conn, error) {
	select {
	case <-l.closer:
		return nil, errors.New("Listener closed")
	default:
	}

	serverConn, clientConn := net.Pipe()
	l.ch <- serverConn

	return clientConn, nil
}
