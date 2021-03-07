package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* 双指针同向移动
 */
func moveZeroes(nums []int)  {
    i, j := 0, 0
    for j < len(nums) {
        if nums[j] != 0 {
            nums[i] = nums[j]
            i++
        }
        j++
    }

    for ; i < len(nums); i++ {
        nums[i] = 0
    }
}

func TestMoveZeros(t *testing.T) {
	nums := []int{0,1,0,3,2}
	moveZeroes(nums)
	assert.Equal(t, 0, nums[3])
	assert.Equal(t, 0, nums[4])
}