package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeDuplicates(nums []int) int {
    if nums == nil {
        return 0
    }
    i, j := 0, 0 // 双指针解题思维，同向而行，[0, i)是有效保留数据，[i, j)是已过滤无效数据，[j, len(nums))是未执行数据
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