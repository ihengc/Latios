package pool

import (
	"database/sql/driver"
	"sync"
)

/*
* @author: Chen Chiheng
* @date: 2023/3/13 0013 15:14
* @description:
**/

// conn 封装原始的数据库链接，Golang标准库中的数据库链接，在database/driver/driver.go中被定义。
// 在连接池中当连接不再被使用时，并不是去关闭此连接。而是归还给连接池。
type conn struct {
	// pool 连接池指针，保存连接池，方便归还此连接，
	pool *Pool
	// createAt 此连接被创建的时间，若对连接有使用时间上的限制，则通过此字段计算使用时间。
	sync.Mutex
	// raw 原始数据库连接（标准库）。
	raw driver.Conn
	// closed 若在使用此连接时，此连接已经被关闭，需要去重新建立连接。
	closed bool
}

// 检测此连接当前是否可以被使用。

// connRequest 获取连接的实体。
type connRequest struct {
	conn *conn
	err  error
}
