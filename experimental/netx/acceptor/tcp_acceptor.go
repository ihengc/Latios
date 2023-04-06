package acceptor

import "net"

/*
* @author: Chen Chiheng
* @date: 2023/4/6 0006 10:32
* @description:
**/

// tcpAcceptor tcp监听套接字。
type tcpAcceptor struct {
	ln net.Listener
	// running tcp监听套接字是否运行标识。
	running bool
	// connChan 连接（阻塞）通道。
	connChan chan net.Conn
}

// ListenAndServe 启动tcp监听。
func (tc *tcpAcceptor) ListenAndServe() {
	// 保证tcp监听套接字正在运行。
	for tc.running {
		// TODO accept 调用会发生何种错误，对于特定错误我们需要如何处理？
		conn, err := tc.ln.Accept()
		if err != nil {
			continue
		}
		tc.connChan <- conn
	}
}

// GetConnChan 获取已经建立成功连接只读队列。
func (tc *tcpAcceptor) GetConnChan() <-chan net.Conn {
	return tc.connChan
}

// Addr 获取监听套接字绑定的地址。
func (tc *tcpAcceptor) Addr() net.Addr {
	return tc.ln.Addr()
}

// Close 停止接收新连接。
func (tc *tcpAcceptor) Close() {
	if tc.running {
		_ = tc.ln.Close()
		tc.running = false
	}
}

// newTCPAcceptor 创建tcp监听套接字。若给定非法地址将返回错误。
func newTCPAcceptor(network, address string) (*tcpAcceptor, error) {
	addr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return nil, err
	}
	ln, err := net.ListenTCP(network, addr)
	if err != nil {
		return nil, err
	}
	return &tcpAcceptor{
		ln:       ln,
		connChan: make(chan net.Conn, 1),
	}, nil
}
