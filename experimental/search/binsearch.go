package search

/*
* @author: Chen Chiheng
* @date: 2023/4/3 0003 16:04
* @description:
**/

// Binsearch 折半查找（在有序数组x中查找k）。
// 若查询到k则返回其索引，否则返回-1。
func Binsearch(x []int, k int) int {
	low := 0
	high := len(x) - 1
	for low <= high {
		middle := low + (high-low)/2
		if x[middle] == k {
			return middle
		} else if x[middle] > k {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}
