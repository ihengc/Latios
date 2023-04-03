package netx

import "net"

/*
* @author: Chen Chiheng
* @date: 2023/4/3 0003 11:31
* @description:
**/

// 本文件对bind,listen,accept系统调用进行抽象，封装。

// Acceptor 此接口将封装被动（监听）套接字的基本功能，
type Acceptor interface {
	// ListenAndServe 执行accept系统调用。
	ListenAndServe()
	// GetConnChan 获取已经建立成功连接队列。
	// 设计连接队列的目的为了将连接的获取与处理解耦。
	GetConnChan() <-chan net.Conn
	// Addr 获取监听套接字绑定的地址。
	Addr() net.Addr
	// Close 关闭监听套接字。不会关闭（或清空）已连接队列。
	// 仅仅停止接收新的连接。
	Close() error
}

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
