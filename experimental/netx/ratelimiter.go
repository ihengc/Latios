package netx

import (
	"context"
	"time"
)

/*
* @author: Chen Chiheng
* @date: 2023/4/4 0004 14:33
* @description:
**/

// 本文件对速率控制进行分析实现。

// Limit 速率。
type Limit float64

// RateLimiter 速率控制接口。
type RateLimiter interface {
	Wait(context.Context) error
	WaitN(context.Context, int) error
	Allow(time.Time) bool
	AllowN(time.Time, int) bool
	Reserve(time.Time)
	ReserveN(time.Time, int)
	// SetLimit 设置速率。
	SetLimit(Limit)
	// SetBurst 设置桶大小。
	SetBurst(int)
}

type Reservation struct {
}
