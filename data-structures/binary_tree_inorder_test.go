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
func inorderTraversalRecursion(root *TreeNode) []int {
	var results = []int{}
	if root == nil {
		return results
	}

	tmp := inorderTraversalRecursion(root.Left)
	for _, i := range tmp {
		results = append(results, i)
	}

	results = append(results, root.Val)

	tmp = inorderTraversalRecursion(root.Right)
	for _, i := range tmp {
		results = append(results, i)
	}

	return results
}

func inorderTraversalTraverse(root *TreeNode) []int {
	var results = []int{}
	if root == nil {
		return results
	}

	var stack = []*TreeNode{}
	for root != nil || len(stack) > 0 { // root != nil condition for first step.
		for root != nil {
			stack = append(stack, root)
			root = root.Left // move to the left node
		}

		top := stack[len(stack) - 1]
		results = append(results, top.Val)
		root = top.Right
		stack = stack[:len(stack) - 1]
	}

	return results
}

func TestInorder(t *testing.T) {
	res := inorderTraversalRecursion(root)
	var experts = []int{2, 4, 7, 8, 9, 10, 11, 12}
	assert.Equal(t, len(res), len(experts))

	for i := 0; i < len(res); i++ {
		assert.Equal(t, res[i], experts[i])
	}

	res = inorderTraversalTraverse(root)
	assert.Equal(t, len(res), len(experts))

	for i := 0; i < len(res); i++ {
		assert.Equal(t, res[i], experts[i])
	}
}

