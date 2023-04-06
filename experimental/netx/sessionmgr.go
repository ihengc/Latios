package netx

import (
	"errors"
	"fmt"
	"sync"
)

/*
* @author: Chen Chiheng
* @date: 2023/4/6 0006 13:47
* @description:
**/

// SessionMgr 负责管理会话。
// 为每个客户端连接创建一个会话，不同的场景可能需要区分会话创建的时机。
type SessionMgr struct {
	mu               sync.Mutex
	sessions         map[userIDType]*Session
	onCloseCallbacks []func(*Session)
}

// IsOnline 报告用户是否在线。
func (sm *SessionMgr) IsOnline(useID userIDType) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	_, ok := sm.sessions[useID]
	return ok
}

// Count 报告在线用户数量。
func (sm *SessionMgr) Count() int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return len(sm.sessions)
}

// Multicast 多播。
func (sm *SessionMgr) Multicast(userIDs []userIDType, data []byte) error {
	return nil
}

// Broadcast 广播。
func (sm *SessionMgr) Broadcast(data []byte) error {
	return nil
}

// KickOut 踢出指定玩家，若玩家对应的会话不存在则返回错误（或者返回bool表示是否
// 成功踢出）。
func (sm *SessionMgr) KickOut(userID userIDType, data []byte) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	session, ok := sm.sessions[userID]
	if !ok {
		return errors.New(fmt.Sprintf("sessionmgr:the session of user %d does not exist ", userID))
	}
	// 发送踢出数据帧。
	session.Close()
	return nil
}

// Close 关闭指定会话。
func (sm *SessionMgr) Close(userID userIDType) error {
	session, ok := sm.sessions[userID]
	if !ok {
		return errors.New(fmt.Sprintf("sessionmgr:the session of user %d does not exist ", userID))
	}
	session.Close()
	return nil
}

// CloseAll 关闭所有会话。
func (sm *SessionMgr) CloseAll() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	for _, session := range sm.sessions {
		session.Close()
	}
}

// Clear 清空所有会话。
func (sm *SessionMgr) Clear() {
	sm.mu.Lock()
	sm.sessions = make(map[userIDType]*Session)
	defer sm.mu.Unlock()
}
