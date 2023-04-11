package merge

/*
* @author: Chen Chiheng
* @date: 2023/3/31 0031 17:19
* @description:
**/

// 本文件对归并排序算法进行描述分析和实现。

// MergeSort 归并排序。
// low 为子集合的下界，high为子集合的上界。
func MergeSort(x []int, low, high int) {
	if low >= high {
		return
	}
	middle := low + (high-low)/2 // middle 为进一步分隔子集合（分解为子问题）。
	MergeSort(x, low, middle)    // 对子问题进行求解。
	MergeSort(x, middle+1, high)
	merge(x, low, middle, high) // 合并子问题求得的解。
}

func merge(x []int, low int, middle int, high int) {
	// 数组x分解已经被分解为两个有序子集合[low, middle],[middle+1, high]
	// 合并两个有序子集合[low, middle]和[middle+1, high]
	tLow := low
	tMiddle := middle + 1
	tX := make([]int, 0)
	for tLow <= middle && tMiddle <= high {
		if x[tLow] > x[tMiddle] {
			tX = append(tX, x[tMiddle])
			tMiddle++
		} else {
			tX = append(tX, x[tLow])
			tLow++
		}
	}
	if tLow > middle && tMiddle <= high {
		tX = append(tX, x[tMiddle:high+1]...)
	}
	if tMiddle > high && tLow <= middle {
		tX = append(tX, x[tLow:middle+1]...)
	}
	for i := 0; i < len(tX); i++ {
		x[low+i] = tX[i]
	}
}
