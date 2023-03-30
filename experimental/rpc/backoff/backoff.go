package backoff

import (
	"math/rand"
	"sync"
	"time"
)

/*
 * @author: Chen Chiheng
 * @date: 2023/3/29 22:09
 * @description:
 */

// backoff 包实现了回退策略（在建立失败后需要此策略以合理发起重试）。

// Strategy 回退策略。
type Strategy interface {
	// Backoff 给定连续失败的次数，返回下一次重试前等待的时间量。
	Backoff(retries int) time.Duration
}

// Config 定义回退选项。
type Config struct {
	// BaseDelay 表示发生第一次失败后回退的时间量。
	BaseDelay time.Duration
	// Multiplier
	Multiplier float64
	// Jitter
	Jitter float64
	// MaxDelay 最大回退延迟时间。
	MaxDelay time.Duration
}

// 通常在很多实现中都使用“指数退避”的策略。

type Exponential struct {
	Config Config
}

var (
	r  = rand.New(rand.NewSource(time.Now().UnixNano()))
	mu sync.Mutex
)

func Float64() float64 {
	mu.Lock()
	defer mu.Unlock()
	return r.Float64()
}

func (e Exponential) Backoff(retries int) time.Duration {
	// 第一次回退基础时间量。
	if retries == 0 {
		return e.Config.BaseDelay
	}
	backoff := float64(e.Config.BaseDelay)
	max := float64(e.Config.MaxDelay)
	// 根据当前连续失败次数计算回退时延。
	for backoff < max && retries > 0 {
		backoff *= e.Config.Multiplier
		retries--
	}
	if backoff > max {
		backoff = max
	}
	// 经过“指数回退”策略计算出的时延会在一个因子上随机进行浮动。
	// 产生随机回退时延的重要原因时避免同时请求一个服务。
	backoff = backoff + backoff*(e.Config.Jitter*Float64()*2-1)
	if backoff < 0 {
		return 0
	}
	return time.Duration(backoff)
}
