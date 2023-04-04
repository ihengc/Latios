package netx

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
* @author: Chen Chiheng
* @date: 2023/4/4 0004 10:33
* @description:
**/

// Server 服务端。
type Server struct {
	// acceptors 监听套接字。
	acceptors []Acceptor
	// pingStrikes ping异常数量。
	pingStrikes uint8
	// lastPingAt 最近一次ping的时间。
	lastPingAt time.Time
	// onExitCallbacks 服务退出回调函数。
	onExitCallbacks []func()
	fr              FrameReader
}

// Start 启动服务。
func (s *Server) Start() {
	// 启动监听。
	for _, acceptor := range s.acceptors {
		go acceptor.ListenAndServe()
	}
	// 处理客户端连接。
	for _, acceptor := range s.acceptors {
		connChan := acceptor.GetConnChan()
		go s.handleConn(connChan)
	}
	// 监听退出信号。
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigChan:

	}
	// 调用回调函数。
	// TODO 处理callback报错。
	for _, callback := range s.onExitCallbacks {
		if err := fnWrapper(callback); err != nil {
		}
	}
}

// handleConn 处理客户端连接。
func (s *Server) handleConn(connChan <-chan net.Conn) {
	for conn := range connChan {
		go func(conn net.Conn) {
			// TODO 读写数据。
			for {
				frame, err := s.fr.ReadFrame()
				if err != nil {
					continue
				}
				switch frame.Header().Type {
				case FrameSYN:
				case FrameACK:
				case FramePing:
				case FrameData:
				case FramePush:
				case FrameGoAway:
				}
			}
		}(conn)
	}
}
