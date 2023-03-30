package rpc

import (
	"Latios/experimental/rpc/connectivity"
	"net"
)

/*
 * @author: Chen Chiheng
 * @date: 2023/3/29 21:30
 * @description:
 */

// addrConn 表示一个网络连接。
type addrConn struct {
	// raw 表示一个底层的网络连接（也可以是经过抽象后的底层连接)）。
	raw net.Conn
	// state 表示此连接的状态，使用 updateConnectivityState 更新此连接的状态。
	state connectivity.State
}
