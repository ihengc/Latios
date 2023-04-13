package brute_force_search

/*
 * @author: Chen Chiheng
 * @date: 2023/4/13 20:37
 * @description:
 */

// Search 暴力搜索算法。
func Search(pat string, txt string) int {
	pIndex := 0
	tIndex := 0
	for pIndex < len(pat) && tIndex < len(txt) {
		if pat[pIndex] == txt[tIndex] {
			pIndex++
			tIndex++
		} else if pIndex > 0 {
			tIndex -= pIndex
			pIndex = 0
		} else {
			tIndex++
		}
	}
	if pIndex == len(pat) {
		return tIndex - pIndex
	}
	return -1
}
