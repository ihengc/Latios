package bst

/*
 * @author: Chen Chiheng
 * @date: 2023/4/10 20:01
 * @description:
 */

// Interface 二叉查找树接口。
type Interface interface {
	// Get 查找key对应的值，若key在二叉查找树中不存在返回nil。
	Get(key Key) interface{}
	// Put 存放或者更新key对应的值。
	Put(key Key, value interface{})
	Size() int
	Max() Key
	Min() Key
	Floor(key Key) Key
	Select(k int) Key
	Rank(key Key) int
	DeleteMin()
	Delete(key Key)
	Keys() []Key
	RangeKeys(low Key, high Key) []Key
}

// Key 用来二叉查找树节点之间进行比较。
type Key interface {
	Less(key Key) bool
}

// node 二叉查找树的节点。
type node struct {
	// key表示可比较的键。
	key Key
	// value 存储值。
	value interface{}
	// left, right 左右子节点。
	left, right *node
	// n 以当节点为根节点的子树中节点的总数。
	n int
}
