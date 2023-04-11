package kmp

import (
	"testing"
)

/*
* @author: Chen Chiheng
* @date: 2023/4/11 0011 14:09
* @description:
**/

func TestKMP(t *testing.T) {
	testSearch(t)
	testKMPSearch(t)
}

// kmp子字符串查找测试。
func testKMPSearch(t *testing.T) {
	txt := "name"
	pat := "nam"
	ret := Search(pat, txt)
	if ret != 0 {
		t.Fatal("kmp search error")
	}
	txt = "xName"
	pat = "me"
	ret = Search(pat, txt)
	if ret != 3 {
		t.Fatal("kmp search error")
	}
	txt = "xName"
	pat = "jack"
	ret = Search(pat, txt)
	if ret != -1 {
		t.Fatal("kmp search error")
	}
}

// 暴力查找测试。
func testSearch(t *testing.T) {
	txt := "name"
	pat := "nam"
	ret := search(pat, txt)
	if ret != 0 {
		t.Fatal("search error")
	}
	txt = "xName"
	pat = "me"
	ret = search(pat, txt)
	if ret != 3 {
		t.Fatal("search error")
	}
	txt = "xName"
	pat = "jack"
	ret = search(pat, txt)
	if ret != -1 {
		t.Fatal("search error")
	}
}
