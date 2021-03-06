package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b - 'a']++
		}
		m[cnt] = append(m[cnt], str)
	}
	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(strs)

	experts := [][]string{
		{"ate","eat","tea"},
		{"nat","tan"},
		{"bat"},
	}

	for i := 0; i < len(ans); i++ {
		a := ans[i]
		e := experts[i]
		assert.Equal(t, len(e), len(a))
	}
}