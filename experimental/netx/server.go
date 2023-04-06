package netx

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

/*
* @author: Chen Chiheng
* @date: 2023/4/6 0006 10:33
* @description:
**/

// ServerMode 服务模式（集群，独立）。
type ServerMode uint8

const (
	Cluster    ServerMode = 0x1
	Standalone ServerMode = 0x2
)

// ServerType 服务类型。
type ServerType uint8

const (
	Frontend ServerType = 0x1
	Backend  ServerType = 0x2
)

// Server 服务。
type Server struct {
	mode      ServerMode
	servType  ServerType
	acceptors []Acceptor
	// onExitCallbacks 可以设计为固定大小，并直接初始化。
	onExitCallbacks map[string]func()
	// closeNotify 服务关闭通知通道。由 Shutdown 发送通知。
	closeNotify chan struct{}
	frw         FrameReaderWriter
	sessionMgr  *SessionMgr
}

func funcWrapper(callback func()) (err error) {
	callback()
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%s", r))
		}
	}()
	return
}

// AddAcceptor 添加  acceptor 。
func (serv *Server) AddAcceptor(acceptor Acceptor) {
	if serv.servType == Backend {
		panic("server:the backend server do not need acceptor ")
	}
	if serv.acceptors == nil {
		serv.acceptors = make([]Acceptor, 0)
	}
	serv.acceptors = append(serv.acceptors, acceptor)
}

// processConn 处理客户端连接。
func (serv *Server) processConn(connChan <-chan net.Conn) {
	// connChan 中为已经接入的连接。在服务关闭时，关闭刚接入的连接，
	// 正在接收消息，处理消息，写出消息的连接，应该等待全部消息处理完后
	// 再关闭处于此状态下的连接（或者清空接收缓冲区中的数据，只等待消息处理
	// ，写出缓冲区写完）。
}

// Start 运行服务。
func (serv *Server) Start() {
	if serv.servType == Frontend {
		for _, acceptor := range serv.acceptors {
			// 当acceptor停止的时候，连接可能要继续被处理。
			go acceptor.ListenAndServe()
			connChan := acceptor.GetConnChan()
			go serv.processConn(connChan)
		}

	}
	if serv.servType == Backend {

	}
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigChan:
	case <-serv.closeNotify:

	}
	// 停止接入客户端连接，
	for _, acceptor := range serv.acceptors {
		acceptor.Close()
	}
	serv.sessionMgr.CloseAll() // 关闭所有会话。
	// 执行服务退出前的清理工作。
	for name, callback := range serv.onExitCallbacks {
		if err := funcWrapper(callback); err != nil {
			fmt.Println(fmt.Sprintf("server:invoke callback function %s err %s", name, err.Error()))
		}
	}
}

// OnExitCallback 注册服务退出回调函数。
func (serv *Server) OnExitCallback(name string, callback func()) {
	if serv.onExitCallbacks == nil {
		serv.onExitCallbacks = make(map[string]func())
	}
	_, ok := serv.onExitCallbacks[name]
	if ok {
		panic(fmt.Sprintf("server:the callback function name %s is repeated", name))
	}
	serv.onExitCallbacks[name] = callback
}

// Shutdown 关闭服务。
func (serv *Server) Shutdown() {
	select {
	case <-serv.closeNotify:
	default:
		close(serv.closeNotify)
	}
}
