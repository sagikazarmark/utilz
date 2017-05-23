package net_test

import (
	"testing"

	"net"
	"sync"

	_net "github.com/sagikazarmark/utilz/net"
)

func TestPipeListen(t *testing.T) {
	addr := _net.ResolveVirtualAddr("network", "addr")

	listener, dialer := _net.PipeListen(addr)

	var wg sync.WaitGroup

	var clientConn, serverConn net.Conn

	writtenBytes := []byte("piped")
	var readBytes = make([]byte, len(writtenBytes))
	var written, read int

	wg.Add(2)

	go func() {
		defer wg.Done()

		var err error

		clientConn, err = dialer.Dial()
		if err != nil {
			t.Fatalf("cannot dial: %v", err)
		}

		written, err = clientConn.Write(writtenBytes)
		if err != nil {
			t.Fatalf("cannot write: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		var err error

		serverConn, err = listener.Accept()
		if err != nil {
			t.Fatalf("cannot accept: %v", err)
		}

		read, err = serverConn.Read(readBytes)
		if err != nil {
			t.Fatalf("cannot write: %v", err)
		}
	}()

	wg.Wait()

	if written != read {
		t.Errorf("data size mismatch, written %d bytes, read %d bytes", written, read)
	}

	if string(writtenBytes) != string(readBytes) {
		t.Errorf("data mismatch, written %d, read %d", writtenBytes, readBytes)
	}
}
