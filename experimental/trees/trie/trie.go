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
	// Contains 查询前缀树中是否保存key对应的值。
	Contains(key string) bool
	// IsEmpty 前缀树是否为空。
	IsEmpty() bool
	// LongestPrefixOf prefix的前缀中最长的键。
	LongestPrefixOf(prefix string) string
	// KeysWithPrefix 所有以prefix为前缀的键。
	KeysWithPrefix(prefix string) []string
	// KeysThatMatch 所有和s匹配的键（其中“."能匹配任意字符）。
	KeysThatMatch(s string) []string
	// Size 键值对数量。
	Size() int
	// Keys 返回前缀树中所有的键。
	Keys() []string
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
	trie.root = remove(trie.root, key, 0)
}

func remove(n *node, key string, d int) *node {
	// 先查找到目标键key，后进行删除。
	// 若当前节点n为nil，则说明我们不需要进行删除。
	if n == nil {
		return nil
	}
	// 当d等于key的长度时，key已经查找完成，我们要将当前节点n的值value
	// 置为nil。
	if len(key) == d {
		n.value = nil
	} else {
		// 当前key还未查找完，需要向当前节点n的子节点继续进行查询。
		char := []byte(key)[d]
		// 当前节点n的子节点指针，若其子节点被删除，应该指向nil；否则
		// 继续指向原子节点。
		n.children[char] = remove(n.children[char], key, d+1)
	}
	if n.value != nil {
		return n
	}
	// 当前节点n的值value已经为nil，若其无子节点，则我们需要删除当前节点n。
	// 若其有子节点说明当前节点n是其子节点的一个前缀节点。我们可以不做任何处理。
	for i := 0; i < R; i++ {
		if n.children[i] != nil {
			return n
		}
	}
	return nil
}

// Contains 查询前缀树中是否保存key对应的值。
func (trie *Trie) Contains(key string) bool {
	n := get(trie.root, key, 0)
	if n != nil && n.value != nil {
		return true
	}
	return false
}

// LongestPrefixOf 给的一个前缀prefix，查询前缀树中与prefix有最长公共前缀的键。
func (trie *Trie) LongestPrefixOf(prefix string) string {
	length := search(trie.root, prefix, 0, 0)
	return string([]byte(prefix)[:length])
}

func search(n *node, s string, d int, length int) int {
	if n == nil {
		return length
	}
	if n.value != nil {
		length = d
	}
	if len(s) == d {
		return length
	}
	char := []byte(s)[d]
	return search(n.children[char], s, d+1, length)
}

// KeysWithPrefix 所有以prefix为前缀的键。
func (trie *Trie) KeysWithPrefix(prefix string) []string {
	queue := make([]string, 0)
	collect(trie.root, prefix, queue)
	return queue
}

func collect(n *node, prefix string, queue []string) {
	if n == nil {
		return
	}
	if n.value != nil {
		queue = append(queue, prefix)
	}
	for i := 0; i < R; i++ {
		collect(n, string(append([]byte(prefix), byte(i))), queue)
	}
}

// KeysThatMatch 所有和s匹配的键（其中“."能匹配任意字符）。
func (trie *Trie) KeysThatMatch(pat string) []string {
	queue := make([]string, 0)
	matchCollect(trie.root, "", pat, queue)
	return queue
}

func matchCollect(n *node, pre string, pat string, queue []string) {
	d := len(pre)
	if n == nil {
		return
	}
	if d == len(pat) && n.value != nil {
		queue = append(queue, pre)
	}
	if d == len(pat) {
		return
	}
	char := []byte(pat)[d]
	for i := 0; i < R; i++ {
		if char == '.' || int(char) == i {
			matchCollect(n.children[char], string(append([]byte(pre), byte(i))), pat, queue)
		}
	}
}

// Keys 返回前缀树中所有的键。
func (trie *Trie) Keys() []string {
	return trie.KeysWithPrefix("")
}

// IsEmpty 前缀树是否为空。
func (trie *Trie) IsEmpty() bool {
	for i := 0; i < R; i++ {
		if trie.root.children[i] != nil {
			return false
		}
	}
	return true
}

// Size 返回键值对数量。
func (trie *Trie) Size() int {
	return size(trie.root)
}

func size(n *node) int {
	if n == nil {
		return 0
	}
	counts := 0
	if n.value != nil {
		counts++
	}
	for i := 0; i < R; i++ {
		counts += size(n.children[i])
	}
	return counts
}

// New 创建前缀树。
func New() *Trie {
	return &Trie{
		root: &node{},
	}
}
