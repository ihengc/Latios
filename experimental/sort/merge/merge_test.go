package merge

import (
	"testing"
)

/*
* @author: Chen Chiheng
* @date: 2023/3/31 0031 17:25
* @description:
**/

type testMergeList []int

func (t testMergeList) Len() int {
	return len(t)
}

func (t testMergeList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t testMergeList) Less(i, j int) bool {
	return t[i] < t[j]
}

func TestMerge(t *testing.T) {
	il := []int{9, 0, 1, 8, 4, 7, 2, 6, 3}
	tl := testMergeList(il)
	MergeSort(tl, 0, len(tl)-1)
	for i := 1; i < tl.Len(); i++ {
		if tl.Less(i, i-1) {
			t.Fatal("Merge error.")
		}
	}
}
