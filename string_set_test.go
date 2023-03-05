package set

import (
	"testing"
)

func TestStringSet(t *testing.T) {
	set1 := NewStringSet("1", "2")
	set2 := NewStringSet("2", "3")
	// 交集
	if !set1.Intersect(set2).Equals(NewStringSet("2")) {
		t.Error("交集异常")
	}
	// 并集
	if !set1.Union(set2).Equals(NewStringSet("1", "2", "3")) {
		t.Error("并集异常")
	}
	// 差集
	if !set1.Expect(set2).Equals(NewStringSet("1")) {
		t.Error("差集异常")
	}
	// 包含
	if !NewStringSet("1", "2", "3").Contains(set1) ||
		!NewStringSet("1", "2", "3").Contains(set2) {
		t.Error("包含异常")
	}
	// 删除
	if NewStringSet("1", "2", "3").Del("3").Equals(set1) {
		t.Error("删除异常")
	}
	// 包含
	if NewStringSet("1", "2", "3").Has("3") {
		t.Error("包含异常")
	}
	// 存在
	if NewStringSet("1", "2", "3").Any("3", "4") {
		t.Error("存在异常")
	}
}
