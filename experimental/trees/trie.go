package trees

/*
 * @author: Chen Chiheng
 * @date: 2023/4/7 21:11
 * @description:
 */

// 本文分析单词查找树（前缀树）的性质，以及如何实现。

// TrieInterface 定义前缀树接口。
type TrieInterface interface {
	// Put 向表中插入键值对（若值为nil，则删除键key）。
	Put(key string, value interface{})
	// Get 获取键key对应的值，若键不存在则返回nil.
	Get(key string) interface{}
	// Delete 删除键key以及其对应的值。
	Delete(key string)
	// Contains 报告表中是否包含键key的值。
	Contains(key string) bool
	// IsEmpty 报告表是否为空。
	IsEmpty() bool
	// LongestPrefixOf 返回s的前缀中最长的键。
	LongestPrefixOf(s string) string
	// KeysWithPrefix 返回所有以s为前缀的键。
	KeysWithPrefix(prefix string) []string
	// KeysThatMatch 返回所有和s匹配的键（“.”可以匹配任意字符）。
	KeysThatMatch(s string) []string
	// Size 返回键值对的数量。
	Size() int
	// Keys 返回表中所有的键。
	Keys() []string
}

// R 基数（ascii码表中全部字符）。
const R = 128

// node 节点。
type node struct {
	// children 当前节点的子节点。
	children [R]*node
	// value 当前节点绑定的值。
	value interface{}
}

// Trie 前缀树。
type Trie struct {
	// root 根节点。
	root *node
	// size 保存当前表中键的数量。
	size int
}

// Put 向表中插入键值对（若值为nil，则删除键key）。
func (trie *Trie) Put(key string, value interface{}) {
	trie.root = put(trie.root, key, value, 0)
}

// put 递归调用。d表示当前插入到key中哪个字符。
func put(n *node, key string, value interface{}, d int) *node {
	// 若n为nil，则表明需要创建新节点。
	if n == nil {
		n = &node{}
	}
	// 当d等于key的长度时，表明key中的字符以及全部放入到表中（递归出口）。
	if d == len(key) {
		n.value = value
		return n
	}
	//  查找d在key中对应的字符。
	char := []byte(key)[d]
	// 查找char在children中的位置。
	n.children[char] = put(n.children[char], key, value, d+1)
	return n
}

func (trie *Trie) Get(key string) interface{} {
	return get(trie.root, key, 0)
}

func get(n *node, key string, d int) *node {
	if n == nil {
		return nil
	}
	if d == len(key) {
		return n
	}
	char := []byte(key)[d]
	return get(n.children[char], key, d+1)
}

func (trie *Trie) Delete(key string) {
	trie.root = remove(trie.root, key, 0)
}

func remove(n *node, key string, d int) *node {
	if n == nil {
		return nil
	}
	if d == len(key) {
		n.value = nil
	} else {
		char := []byte(key)[d]
		n.children[char] = remove(n.children[char], key, d+1)
	}
	if n.value != nil {
		return n
	}
	for i := 0; i < R; i++ {
		if n.children[i] != nil {
			return n
		}
	}
	return nil
}

func (trie *Trie) LongestPrefixOf(s string) string {
	length := search(trie.root, s, 0, 0)
	return string([]byte(s)[:length])
}

func search(n *node, s string, d int, length int) int {
	if n == nil {
		return length
	}
	if n.value != nil {
		length = d
	}
	if d == len(s) {
		return length
	}
	char := []byte(s)[d]
	return search(n.children[char], s, d+1, length)
}

func (trie *Trie) Size() int {
	return trie.size
}

func (trie *Trie) IsEmpty() bool {
	return trie.size == 0
}

func (trie *Trie) Contains(key string) bool {
	n := get(trie.root, key, 0)
	if n != nil && n.value != nil {
		return true
	}
	return false
}

func (trie *Trie) KeysWithPrefix(prefix string) []string {

	return nil
}

func collect(n *node, prefix string, queue []string) {
	if n == nil {
		return
	}
	if n.value != nil {
		queue = append(queue, prefix)
	}
}

func (trie *Trie) KeysThatMatch(s string) []string {
	return nil
}

func (trie *Trie) Keys() []string {
	return nil
}

// lazySize 延迟计算表中键的数量。
func lazySize(n *node) int {
	if n == nil {
		return 0
	}
	counts := 0
	if n.value != nil {
		counts++
	}
	for i := 0; i < R; i++ {
		counts += lazySize(n.children[i])
	}
	return counts
}
