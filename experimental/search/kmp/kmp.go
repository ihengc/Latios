package kmp

/*
* @author: Chen Chiheng
* @date: 2023/4/11 0011 13:21
* @description:
**/

// search 在txt文本中暴力查找pat子字符串。
func search(pat string, txt string) int {
	var (
		i int = 0 // i为文本指针。
		j int = 0 // j为子字符串（模式）指针。
	)
	// 在文本txt中查找pat，i不可能超过txt的长度，相应的j不可能超过pat的长度。
	for ; i < len(txt) && j < len(pat); i++ {
		// 若i和j处的字符不相等，则需要将i向前移动i-j位，j需要放到pat的第一位。
		// 若i和j出的字符相等，则向后移动j。
		if pat[j] == txt[i] {
			j++
		} else {
			i -= j
			j = 0
		}
	}
	// 若j已经等于pat的长度，说明pat已经查找完成并且已经找到pat。
	// 在pat查找完后，j仍然会加1，因此会等于pat的长度。
	if j == len(pat) {
		return i - len(pat) // 返回pat在txt中的起始位置。
	}
	return -1 // 返回-1表示未找到pat子串。
}

// 上述暴力查找算法中，需要回退的情况，以及如何回退：
// 文本指针以前的字符表示已经比较过的字符，后面的字符表示未比较的字符。
// 子字符串指针以前的字符表示相等的子字符串，后面的字符表示未比较的字符。
// 在遇到不相等的字符时，子字符串指针回退到子字符串的第一位。文本指针向前移动子字符串指针移动的距离。
// 		  0	          J		  M
// 		  | - - - - - - | - - - - |
// | - - - | - - - - - - | - - - - - - - |
// 0      (I-J)          I				 N

// Search kmp子字符串查找算法。
func Search(pat string, txt string) int {
	var (
		i, j = 0, 0
	)
	next := buildNext(pat)
	for i < len(txt) {
		// 若指针i和j处字符相等，则向后移动i和j指针。
		if txt[i] == pat[j] {
			i++
			j++
			// 若在子串中的某个位置处出现字符不相等（也就说j指针前的子串是相等的），
			// 计算j需要回退的步长，j回退。
		} else if j > 0 {
			j = next[j-1]
		} else {
			// 此时j等于0，说明子串pat的第一个字符与txt[j]处的字符就不相等。
			i++
		}
		if j == len(pat) {
			return i - j
		}
	}
	return -1
}

// buildNext 计算回退步长。
func buildNext(pat string) []int {
	// 我们只需要查找当前pat[:j]中最长的公共前后缀。
	next := make([]int, 1)
	prefixLen := 0
	i := 0
	for i < len(pat) {
		if pat[prefixLen] == pat[i] {
			prefixLen++
			next = append(next, prefixLen)
			i++
		} else {
			if prefixLen == 0 {
				next = append(next, 0)
				i++
			} else {
				prefixLen = next[prefixLen-1]
			}
		}
	}
	return next
}

// 当在子字符串中的某个位置出现不匹配时，说明当前不匹配的位置前的所有字符
// 是匹配的。
