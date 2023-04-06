package acceptor

import "net"

/*
* @author: Chen Chiheng
* @date: 2023/4/6 0006 10:29
* @description:
**/

type UDPAcceptor struct {
	conn     net.Conn
	running  bool
	connChan chan net.Conn
}

func NewUDPAcceptor(network, address string) (*UDPAcceptor, error) {
	addr, err := net.ResolveUDPAddr(network, address)
	if err != nil {
		return nil, err
	}
	udpConn, err := net.ListenUDP(network, addr)
	if err != nil {
		return nil, err
	}
	return &UDPAcceptor{
		conn:     udpConn,
		connChan: make(chan net.Conn, 1),
	}, nil
}
