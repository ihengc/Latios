package connectivity

import (
	"sync"
)

/*
 * @author: Chen Chiheng
 * @date: 2023/3/29 21:59
 * @description:
 */

// connectivityStateManager 负责管理连接状态。
type connectivityStateManager struct {
	mu         sync.Mutex
	state      State
	notifyChan chan struct{}
}

// updateState 更新连接状态。
func (csm *connectivityStateManager) updateState(state State) {
	csm.mu.Lock()
	defer csm.mu.Unlock()
	// 当连接状态变更到 connectivity.Shutdown 后不能变更到任何状态。
	if csm.state == Shutdown {
		return
	}
	// 当前连接的状态与需要变更的状态一致，不要做任何操作。
	if csm.state == state {
		return
	}
	// 变更状态。
	csm.state = state
	// 成功变更状态后发送状态变更通知。
	if csm.notifyChan != nil {
		// 通过关闭通道的方式来发送通知。
		// 此通道的读取方会阻塞在此通道上，此时关闭通道则会结束读取方的阻塞，从而
		// 达到通知的目的。
		close(csm.notifyChan)
		csm.notifyChan = nil
	}
}

// getState 获取连接的状态。
func (csm *connectivityStateManager) getState() State {
	csm.mu.Lock()
	defer csm.mu.Unlock()
	return csm.state
}

// getNotifyChan 获取状态变更消息通知通道。
// 我们有需要监听连接状态变更的需求，提供一种状态
// 变更后主动通知的机制（而非我们去轮询一个连接的状态）。
// 比如轮询getState。
func (csm *connectivityStateManager) getNotifyChan() <-chan struct{} {
	// 这里限制的通道的操作，只允许读。
	csm.mu.Lock()
	defer csm.mu.Unlock()
	return csm.notifyChan
}
