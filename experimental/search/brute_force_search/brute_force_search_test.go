package brute_force_search

import "testing"

/*
 * @author: Chen Chiheng
 * @date: 2023/4/13 20:39
 * @description:
 */

func TestSearch(t *testing.T) {
	txt := "MyNameIsLiKu"
	for i := 0; i < len(txt); i++ {
		if Search(string(txt[i]), txt) != i {
			t.Fatal("search err")
		}
	}
}
