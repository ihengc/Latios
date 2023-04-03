package netx

import "time"

/*
 * @author: Chen Chiheng
 * @date: 2023/4/3 21:47
 * @description:
 */

// KeepaliveParameters 保活参数。
type KeepaliveParameters struct {
	// MaxConnectionAge 表示在发送RST包之前连接的最大存活时间。
	// 默认为无限大。此值应该是一个随机浮动的值。
	MaxConnectionAge  time.Duration
	MaxConnectionIdle time.Duration
	// Time 经过一段时间若客户端无任何活动，则会ping客户端。
	Time time.Duration
	// Timeout 在进行ping检查后，若在 Timeout 时间段内连接
	// 没有任何活动则关闭此连接。
	Timeout time.Duration
	// 是否设置最小ping时间间隔，超过此时间后关闭连接。
}
