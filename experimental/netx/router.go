package netx

/*
* @author: Chen Chiheng
* @date: 2023/4/6 0006 14:57
* @description:
**/

// 本文件对路由表进行抽象和实现。

// 1.map路由表、
// 2.radix-tree路由表。

// RouteIDType 路由类型。
type RouteIDType uint32

// LocalHandlerFunc 当服务作为独立服务运行时使用此路由函数。
type LocalHandlerFunc func(c *Context)

type Router struct {
	routes map[*Route]LocalHandlerFunc
}
