package net

import "net"

// virtualAddr is a fake, in-memory address representation.
type virtualAddr struct {
	network string
	addr    string
}

// ResolveVirtualAddr returns a new in-memory Addr.
func ResolveVirtualAddr(network, addr string) net.Addr {
	return &virtualAddr{network, addr}
}

// Network returns the address's network name.
func (a *virtualAddr) Network() string {
	return a.network
}

// String returns the address's string representation.
func (a *virtualAddr) String() string {
	return a.addr
}
