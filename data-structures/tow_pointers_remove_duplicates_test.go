package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeDuplicates(nums []int) int {
    if nums == nil {
        return 0
    }
    i, j := 0, 0
    for j < len(nums) {
        if i == 0 || nums[j] != nums[i - 1] {
            nums[i] = nums[j]
            i++
            j++
        } else {
            j++
        }
    }
    return i
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1,1,2}
	l := removeDuplicates(nums)
	assert.Equal(t, 2, l)
}