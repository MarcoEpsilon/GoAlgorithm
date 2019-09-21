package strings

import (
	"testing"
)
func TestNormalSubString(t *testing.T) {
	text := "abcdeabceda"
	pattern := "dea"
	i := NormalSubString(pattern, text)
	if i != 3 {
		t.Errorf("error index: %d", i)
	}
}

func TestNext(t *testing.T) {
	pattern := "ababaaababaa"
	next := Next(pattern)
	result := []int{0, 0, 0, 1, 2, 3, 1, 1, 2, 3, 4, 5}
	errMsg := "func Next's result is (%d = %d) is unexpected"
	if len(result) != len(next) {
		t.Errorf(errMsg, len(result), len(next))
	}
	for index, value := range next {
		if result[index] != next[index] {
			t.Errorf(errMsg, value, result[index])
		}
	}
}
func TestKmpSubString(t *testing.T) {
	text := "ababcabcacbab"
	pattern := "abcac"
	i := KmpSubString(pattern, text)
	if i != 5 {
		t.Errorf("error index: %d", i)
	}
}

func TestPermutation(t *testing.T) {
	text := "abc"
	permutation := Permutation(text)
	result := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	errMsg := "func Permutation's result (%q = %q) is unexpected"
	if len(result) != len(permutation) {
		t.Errorf(errMsg, len(result), len(permutation))
	}
	for i, p := range permutation {
		if result[i] != p {
			t.Errorf(errMsg, result[i], p)
		}
	}
}

func TestMaxCommonSubString(t *testing.T) {
	left := "abcdefgh"
	right := "mnqpdefxh"
	max := MaxCommonSubString(left, right)
	if max != "def" {
		t.Error(max)
	}
}