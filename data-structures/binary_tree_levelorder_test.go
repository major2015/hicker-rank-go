package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
	var results [][]int
	if root == nil {
		return results
	}

	var stack = []*TreeNode{}
	stack = append(stack, root) // initialize stack with root

	for len(stack) > 0 {
		sub := []int{}
        length := len(stack)
		for i := 0; i < length; i++ {
			ele := stack[0]
			stack = stack[1:len(stack)]

            sub = append(sub, ele.Val)
			
			if ele.Left != nil {
				stack = append(stack, ele.Left)
			}
			if ele.Right != nil {
				stack = append(stack, ele.Right)
			}
		}
		results = append(results, sub)
	}
	return results
}

func TestLevelOrder(t *testing.T) {
	res := levelOrder(root)
	var experts = [][]int{{7}, {4, 10}, {2, 8, 12}, {9, 11}}
	assert.Equal(t, 4, len(res))

	for i := 0; i < len(res); i++ {
		assert.Equal(t, len(res[i]), len(experts[i]))
		for j := 0; j < len(res); j++ {
			assert.Equal(t, res[j], experts[j])
		}
	}
}