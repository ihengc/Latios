package acceptor

import "net"

/*
* @author: Chen Chiheng
* @date: 2023/4/6 0006 10:29
* @description:
**/

type UDPAcceptor struct {
	ln      net.Listener
	running bool
	// connChan 连接（阻塞）通道。
	connChan chan net.Conn
}
