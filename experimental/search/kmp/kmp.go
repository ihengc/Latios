package kmp

/*
* @author: Chen Chiheng
* @date: 2023/4/11 0011 13:21
* @description:
**/

// Search kmp子字符串查找算法。
func Search(pat string, txt string) int {
	pIndex := 0
	tIndex := 0
	next := buildNext(pat)
	for tIndex < len(txt) {
		if pat[pIndex] == txt[tIndex] {
			pIndex++
			tIndex++
		} else if pIndex > 0 {
			pIndex = next[pIndex]
		} else {
			tIndex++
		}
		if pIndex == len(pat) {
			return tIndex - pIndex
		}
	}
	return -1
}

func buildNext(pat string) []int {
	pIndex := 1
	prefixLen := 0
	next := make([]int, 0)
	next = append(next, 0)
	for pIndex < len(pat) {
		if pat[pIndex] == pat[prefixLen] {
			pIndex++
			prefixLen++
			next = append(next, prefixLen)
		} else {
			if prefixLen == 0 {
				pIndex++
			} else {
				pIndex = next[prefixLen-1]
			}
		}
	}
	return next
}
