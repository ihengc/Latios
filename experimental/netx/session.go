package netx

import (
	"errors"
	"fmt"
	"net"
)

/*
* @author: Chen Chiheng
* @date: 2023/4/6 0006 11:57
* @description:
**/

// sessionIDType 会话ID类型。
type sessionIDType uint32

// userIDType 用户ID类型。
type userIDType uint64

type Session struct {
	id         sessionIDType
	userID     userIDType // userID 0表示无效用户ID。
	conn       net.Conn
	sessionMgr *SessionMgr
}

// BindUser 绑定用户ID。
func (s *Session) BindUser(userID userIDType) error {
	if s.userID != 0 {
		return errors.New(fmt.Sprintf("session:  the user %d has been bound.", s.userID))
	}
	s.userID = userID
	return nil
}

// IsBindUser 报告此会话是否绑定了用户。
func (s *Session) IsBindUser() bool {
	return s.userID != 0
}

// Push 发送消息给客户端（帧类型为 FramePush）。
func (s *Session) Push(data []byte) error {
	return nil
}

// GoAway 关闭连接，但在关闭前会发送用户指定的消息（帧类型为 FrameGoAway）。
func (s *Session) GoAway(data []byte) error {
	return nil
}

// Close 关闭会话。
func (s *Session) Close() {
	_ = s.conn.Close()
}
