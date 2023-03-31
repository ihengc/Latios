package sort

import "testing"

/*
* @author: Chen Chiheng
* @date: 2023/3/31 0031 10:55
* @description:
**/

type testList []int

func (t testList) Len() int {
	return len(t)
}

func (t testList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t testList) Less(i, j int) bool {
	return t[i] < t[j]
}

func TestInsertion(t *testing.T) {
	il := []int{9, 0, 1, 8, 4, 7, 2, 6, 3}
	tl := testList(il)
	Insertion(tl)
	for i := 1; i < tl.Len(); i++ {
		if tl.Less(i, i-1) {
			t.Fatal("Insertion error.")
		}
	}
}
