package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLongestSubstring(s string) int {
    bucket := map[byte]int{} // 哈希表用于记录每个字符是否出现过
    lengh := len(s)
	maxSub := 0
	lp := -1 // 左指针，初始值-1，表示还没向右移动过

	for i := 0; i < lengh; i++ {
		if i != 0 {
			delete(bucket, s[i - 1]) // 左指针向右移动一格，删除一个字符
		}
		for lp + 1 < lengh && bucket[s[lp + 1]] == 0 {
			bucket[s[lp + 1]] ++
			lp ++ // 不断移动右指针
		}
        maxSub = max(lp - i + 1, maxSub)
	}

	return maxSub
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcdacd"
	cnt := lengthOfLongestSubstring(s)
	assert.Equal(t, 4, cnt)
}