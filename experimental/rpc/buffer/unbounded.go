package buffer

import "sync"

/*
* @author: Chen Chiheng
* @date: 2023/3/30 0030 14:12
* @description:
**/

// Unbounded 无边界的缓冲。
type Unbounded struct {
	// c 用户将从此通道获取缓存的数据。
	c   chan interface{}
	mu  sync.Mutex
	buf []interface{}
}

func (u *Unbounded) Put(v interface{}) {
	u.mu.Lock()
	// 若读通道无数据会阻塞获取者，此时放入数据时直接放入读通道中
	// 可以实现唤醒（结束阻塞）的效果。
	if len(u.buf) == 0 {
		select {
		case u.c <- v:
			u.mu.Unlock()
			return
		default:
		}
	}
	u.buf = append(u.buf, v)
	u.mu.Unlock()
}

func (u *Unbounded) Load() {
	u.mu.Lock()
	if len(u.buf) > 0 {
		select {
		case u.c <- u.buf[0]:
			u.buf[0] = nil
			u.buf = u.buf[1:]
		default:
		}
	}
	u.mu.Unlock()
}

func (u *Unbounded) Get() <-chan interface{} {
	return u.c
}

func NewUnbounded() *Unbounded {
	return &Unbounded{c: make(chan interface{}, 1)}
}
