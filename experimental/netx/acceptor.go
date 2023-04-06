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
	// 仅仅停止接收新连接。
	Close() error
}

// 每个Acceptor单独运行在goroutine中，关闭则再另一个goroutine中。
// 是否需要同步状态，若需要同步，则是采用共享内存，还是消息传递。
