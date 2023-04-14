package bm

/*
 * @author: Chen Chiheng
 * @date: 2023/4/13 20:42
 * @description: Boyer-Moore字符串查找算法。
 */

func Search(pat string, txt string) int {
	pLen := len(pat)
	tLen := len(txt)
	pMap := make(map[byte]int)
	for i := 0; i < pLen; i++ {
		pMap[pat[i]] = i
	}
	skip := 0
	for i := 0; i < tLen-pLen; i += skip {
		for j := pLen - 1; j >= 0; j-- {
			if pat[j] == txt[i+j] {
			}
		}
	}
	return -1
}

// 从右向左与模式进行匹配。
// 1.首先比较位置5处的E和N（不匹配）但N出现在了模式字符串中。
// F I N D I N A H A Y S T A C K N E E D L E
// N E E D L E
// 2.比较位置10处的E和S（不匹配）。
// F I N D I N A H A Y S T A C K N E E D L E
// 			 N E E D L E
// 3.比较位置16处的E和E（匹配），再比较前一位（位置15）L和N（不匹配）。
// F I N D I N A H A Y S T A C K N E E D L E
//						 N E E D L E
// 4.同1。
// F I N D I N A H A Y S T A C K N E E D L E
// 			                     N E E D L E
