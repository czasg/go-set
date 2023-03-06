package set

import (
	"testing"
)

func TestStringSet(t *testing.T) {
	set123 := NewStringSet("1", "2", "3")
	{
		// Intersect
		set1 := NewStringSet("1", "2")
		set2 := NewStringSet("2", "3")
		if !set1.Intersect(set2).Equals(NewStringSet("2")) {
			t.Error("Intersect异常")
		}
	}
	{
		// Union
		set1 := NewStringSet("1", "2")
		set2 := NewStringSet("2", "3")
		if !set1.Union(set2).Equals(set123) {
			t.Error("Union异常")
		}
	}
	{
		// Expect
		set1 := NewStringSet("1", "2")
		set2 := NewStringSet("2", "3")
		if !set1.Expect(set2).Equals(NewStringSet("1")) {
			t.Error("Expect异常")
		}
	}
	{
		// Contains
		set1 := NewStringSet("1", "2")
		set2 := NewStringSet("2", "3")
		if !set123.Contains(set1) || !set123.Contains(set2) {
			t.Error("Contains异常")
		}
	}
	{
		// Add
		set1 := NewStringSet("1")
		if !set1.Add("1", "2", "3").Equals(set123) {
			t.Error("Add异常")
		}
	}
	{
		// Del
		set1 := NewStringSet("1", "2", "3", "4")
		if !set1.Del("4", "5", "6").Equals(set123) {
			t.Error("Del异常")
		}
	}
	{
		// Has
		set1 := NewStringSet("1", "2", "3", "4")
		if !set1.Has("1", "2", "3") {
			t.Error("Has异常")
		}
		if set1.Has("3", "9") {
			t.Error("Has异常")
		}
	}
	{
		// Any
		set1 := NewStringSet("1", "2", "3", "4")
		if !set1.Any("2", "3", "4", "5") {
			t.Error("Any异常")
		}
		if set1.Any() {
			t.Error("Any异常")
		}
	}
	{
		// Equals
		if NewStringSet("1").Equals(NewStringSet("1", "2")) {
			t.Error("Equals异常")
		}
		if NewStringSet("1", "3").Equals(NewStringSet("1", "2")) {
			t.Error("Equals异常")
		}
	}
}
