package bm

/*
 * @author: Chen Chiheng
 * @date: 2023/4/13 20:42
 * @description: Boyer-Moore字符串查找算法。
 */

func Search(pat string, txt string) int {
	pMap := make(map[byte]int)
	for i := 0; i < len(pat); i++ {
		pMap[pat[i]] = i
	}
	tIndex := 0
	pIndex := len(pat) - 1
	// 从右向左开始比较，若出现字符不同的情况：
	// 先确定文本字符是否在模式字符串中存在，再确定最右边文本字符的位置。
	for pIndex >= 0 {
		if pat[pIndex] == txt[tIndex+len(pat)] {

		} else {
			if skip, ok := pMap[pat[pIndex]]; ok {
				tIndex += len(pat) - skip
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
