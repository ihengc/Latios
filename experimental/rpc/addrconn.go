package rpc

import (
	"Latios/experimental/rpc/connectivity"
	"context"
	"errors"
	"net"
	"sync"
	"time"
)

/*
 * @author: Chen Chiheng
 * @date: 2023/3/29 21:30
 * @description:
 */

// addrConn 表示一个网络连接。
type addrConn struct {
	ctx   context.Context
	dOpts dialOptions
	// raw 表示一个底层的网络连接（也可以是经过抽象后的底层连接)）。
	raw net.Conn
	mu  sync.Mutex
	// state 表示此连接的状态，使用 updateConnectivityState 更新此连接的状态。
	state connectivity.State
	// backoffIdx 回退次数。
	backoffIdx int
}

// connect 建立连接。
// 连接的初始状态必须是IDLE，此连接的状态将尝试发生改变。
func (ac *addrConn) connect() error {
	ac.mu.Lock()
	if ac.state == connectivity.Shutdown {
		return errors.New("addrconn:connect is closed")
	}
	if ac.state != connectivity.Idle {
		ac.mu.Unlock()
		return nil
	}
	ac.updateConnectivityState(connectivity.Connecting, nil) // 更新连接状态。
	ac.mu.Unlock()
	ac.resetConnection() // 重置连接。
	return nil
}

func (ac *addrConn) updateConnectivityState(connecting connectivity.State, lastErr error) {

}

// resetConnection 重置连接。
func (ac *addrConn) resetConnection() {
	// 建立连接时，应该指定一个超时时间或者使用默认的超时时间。
	backoffDuration := ac.dOpts.bs.Backoff(ac.backoffIdx)
	timer := time.NewTimer(backoffDuration)
	select {
	case <-timer.C:
		ac.mu.Lock()
		ac.backoffIdx++
		ac.mu.Unlock()
	}
}

// createConnection 创建连接。
func (ac *addrConn) createConnection() {

}
