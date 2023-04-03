package search

import "testing"

/*
* @author: Chen Chiheng
* @date: 2023/4/3 0003 16:09
* @description:
**/

func TestBinsearch(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 20, 30, 33}
	for i := 0; i < len(x)-1; i++ {
		ret := Binsearch(x, x[i])
		if ret != i {
			t.Fatalf("Binsearch error for key:%d, index:%d ret:%d", x[i], i, ret)
		}
	}
}
