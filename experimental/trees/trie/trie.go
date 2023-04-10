package trie

/*
* @author: Chen Chiheng
* @date: 2023/4/10 0010 10:16
* @description:
**/

// R 基数。
const R = 128

// Interface 前缀树接口。
type Interface interface {
	// Put 向前缀树中插入键值对，若值为nil则删除键key。
	Put(key string, value interface{})
	// Get 返回键key对应的值；若值不存在则返回nil。
	Get(key string) interface{}
	// Delete 删除键key（以及其对应的值）。
	Delete(key string)
}

// node 前缀树节点。
type node struct {
	// value 存放字符串对应的值。
	value interface{}
	// children 当前节点的子节点。
	children [R]*node
}

// Trie 前缀树（基数树）。
type Trie struct {
	// root 前缀树根节点，不存储任何数据。
	root *node
	// size 键值对个数。
	size int
}

// Put 插入键值对到前缀树中。
func (trie *Trie) Put(key string, value interface{}) {
	put(trie.root, key, value, 0)
}

func put(n *node, key string, value interface{}, d int) {
	// 从根节点开始，若当前节点为nil说明需要创建新的节点存放key中的某个字符。
	if n == nil {
		n = &node{}
	}
	// 若d等于key的长度说明key中的每一个字符都已经存放在前缀树中，不需要
	// 再进行插入。
	if len(key) == d {
		return
	}
	// 取出key中当前需要插入的字符，计算出当前字符char需要放在哪个
	// 子节点中，因为基数R取的128，所以这里直接为children数组的下标。
	char := []byte(key)[d]
	// 执行插入。
	put(n.children[char], key, value, d+1)
}

// Get 获取键对应的值，若键不存在则返回nil。
func (trie *Trie) Get(key string) interface{} {
	n := get(trie.root, key, 0)
	if n == nil {
		return nil
	}
	return n.value
}

func get(n *node, key string, d int) *node {
	// 从根节点开始，向子节点进行查找。若当前节点为空说明key在前缀树中
	// 不存在，我们返回nil。
	if n == nil {
		return nil
	}
	// 若d等于key的长度，说明key中的每个字符都已经查找完，会出现两种情况：
	// 1.当前节点的value有值，说明当前节点为我们要查找的节点。
	// 2.当前节点的value为空，说明当前key在前缀树中不存在。
	// 上述两种情况我们都只需要返回当前节点n即可。
	if len(key) == d {
		return n
	}
	// 继续在子节点中查询key中的下一个字符。
	char := []byte(key)[d]
	return get(n.children[char], key, d+1)
}

// Delete 删除键key（以及其对应的值）。
func (trie *Trie) Delete(key string) {

}

func remove(n *node, key string, d int) {

}
