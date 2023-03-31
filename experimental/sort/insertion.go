package sort

/*
* @author: Chen Chiheng
* @date: 2023/3/31 0031 10:27
* @description: 插入排序。
**/

// 本文件主要对插入排序进行实现和原理解析以及复杂度分析。

// 插入排序基本思想是将集合（A）划分为两个部分（已排序A1和未排序A2），从A2中获取元素，插入到A1中
// 并且保证A1依然有序。

// 虽然golang标准库sort定义了什么是可排序集合，为了直接展示抽象过程，在
// 这里再抽象一遍。

// Interface 表示一个可以排序的集合。
type Interface interface {
	// Len 返回集合中元素的个数。
	Len() int
	// Swap 交换集合中两个位置处的元素。
	Swap(i, j int)
	// Less 比较集合中两个位置处元素的大小。
	Less(i, j int) bool
}

// Insertion 插入排序。
func Insertion(l Interface) {
	// 将集合l分隔为[0,i-1]（A1）和[i, l.Len()-1]（A2）两个区间。我们需要在
	// A2中取元素与A1中的元素进行比较找到插入位置，然后插入从A2中取出
	// 的元素。
	for i := 1; i < l.Len(); i++ {
		k := i // k永远指向从A2取出元素的位置。
		// 比较k位置（从A2取出的元素）与其前一位置处元素的大小。
		// k位于集合第一个元素时说明我们将A1中的元素全部比较了一遍。
		for k-1 >= 0 && l.Less(k, k-1) {
			// 若k位置处的元素比其前一位置处元素小，说明A1不满足有序，并且其前一位置处的元素要在k位置。
			// 我们交换两位置处的元素，记得k需要永远指向A2取出元素的位置。
			l.Swap(k, k-1)
			// 移动k。
			k--
		}
	}
}

// 最好情况下复杂度为O(n),最坏为O(n²)。
// 其实从上述边分析边实现的过程中可以发现上述插入排序的实现是具有可以优化的地方。
