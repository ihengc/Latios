package netx

import (
	"net"
	"time"
)

/*
 * @author: Chen Chiheng
 * @date: 2023/4/3 21:31
 * @description:
 */

// 本文件将对常用的运输层协议（udp，tcp）进行抽象，封装。

// Connection 表示一个网络运输通道。
type Connection interface {
	Read([]byte) (int, error)
	Write([]byte)
	Close()
	GracefulClose()
	RemoteAddr() net.Addr
}

// streamConn 流式运输通道。
type streamConn struct {
	conn net.Conn
	// readTimeout 读超时间隔。
	// TODO 选项类的参数可单独抽象为一个选项结构。
	readTimeout time.Duration
	// pingStrikes 防御ping攻击，记录违规ping的次数。
	pingStrikes uint8
	// lastPingAt 最近ping时间。
	lastPingAt time.Time
}

func (sc *streamConn) handlePing() {

}
