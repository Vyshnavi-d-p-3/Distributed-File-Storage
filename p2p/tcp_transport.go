package p2p

import "net"

type TCPTransport struct {
	listenAddress string
	listener net.Listener

	mu sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport{
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAccept() error{
	var err error

	t.listner
	ln, err := net.Listen("tcp",t.listenAddress)
}