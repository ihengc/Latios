package pool

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"sync"
)

/*
* @author: Chen Chiheng
* @date: 2023/3/13 0013 15:14
* @description:
**/

// Interface 数据库连接池接口。将原先在数据库连接上的操作（接口）移动到数据库连接池上。
// 直接通过数据库连接池执行SQL，数据库连接池内部自动管理连接。
// TODO 还有一种实现方式时，保持数据库连接上的接口不变，数据库连接池负责管理连接。
type Interface interface {
	PingContext(ctx context.Context) error
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Close() error
}

// maxBadConnRetries 最大重试次数
const maxBadConnRetries = 2

// Pool 数据库连接池。
type Pool struct {
	// connector 数据库连接器，负责创建连接
	// TODO 若创建连接失败，是否需要重试；重试是否设置最大重试次数，每次重试时间间隔采取何种策略。
	connector driver.Connector
	// connRequests 获取连接的实体（连接池相当于生产者，获取连接的实体相当于消费者）。
	// 若此时连接池的链接不够，对于获取连接的实体，它们将会被阻塞；在池中有可用连接后，会
	// 存在一个连接分配的问题，这里将获取连接的实体进行排队，总是将连接分配给队列头部的元素。
	mu           sync.Mutex
	connRequests map[uint64]chan connRequest
	// nextRequest 指向下一个被分配连接的“获取连接的实体”。
	nextRequest uint64
	// createSignal 创建连接信号，若需要创建连接，通过此通道通知。
	createSignal chan struct{}
	// closed 标识数据库连接池是否已经被关闭（关闭后将会影响其他接口，所以设置此字段，还有
	// 一个原因就是数据库连接池的关闭和使用不一定在一个goroutine中，被关闭后需要通知另一个
	// 正在使用的goroutine）。
	// TODO 对于一个已经关闭的数据库连接池，但有goroutine持有池中的一条连接，需要如何处理此种情况。
	// TODO 可以在数据库连接中加入是否关闭字段，当连接池被关闭后，更新该字段，在执行SQL时抛错。
	closed bool
	// maxOpen 最大打开的连接数。

}

// connectionCreator 连接创建器。
func (pool *Pool) connectionCreator(ctx context.Context) {
	// 这里传入ctx的原因是因为此方法将在一个单独的goroutine中运行（控制该goroutine）。
	// 监听退出和创建新连接的通知。
	for {
		select {
		case <-ctx.Done():
			return
		case <-pool.createSignal:
			pool.createConnection(ctx)
		}
	}
}

// createConnection 新建一条连接。
func (pool *Pool) createConnection(ctx context.Context) {
	// 新建连接。
	raw, err := pool.connector.Connect(ctx)
	// 需要考虑连接池是否已经关闭，连接池已经关闭，则不需要再创建新的连接。
	pool.mu.Lock()
	defer pool.mu.Unlock()
	// 连接池已经关闭。
	if pool.closed {
		if err == nil {
			_ = raw.Close()
		}
		return
	}
	// 创建连接失败，需要执行创建失败后的策略。
	if err != nil {
		return
	}
	// 创建连接成功。需要把此连接放入池中。
}

// NewWithConnector 通过数据库连接器创建一个数据库连接池。
func NewWithConnector(connector driver.Connector) *Pool {
	// TODO 是在创建数据库连接池时去建立所有连接，还是只建立一条连接，或者说在真正执行SQL时建立。
	return nil
}
