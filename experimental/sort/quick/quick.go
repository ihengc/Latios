package quick

/*
* @author: Chen Chiheng
* @date: 2023/4/17 0017 14:02
* @description: 快速排序
**/

// 本文件对快速排序进行讨论分析并实现。

// QuickSort 快速排序。
func QuickSort(x []int) {
	sort(x, 0, len(x)-1)
}

func sort(x []int, low int, high int) {
	p := partition(x, low, len(x)-1)
	sort(x, low, p-1)
	sort(x, p+1, high)
}

func partition(x []int, low int, high int) int {
	value := x[low]
	i, j := low+1, high
	for {
		// 从数组x的左边向右边移动，比较当前指针i处的元素
		// 是否比value值小。
		for x[i] < value { // 若比value值小，指针i继续向右移动；否则停止移动。
			if i == high {
				break
			}
			i++
		}
		for x[j] > value {
			if j == low {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		x[i], x[j] = x[j], x[i]
	}
	x[low], x[j] = x[j], x[low]
	return j
}
