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
 func postorderTraversalRecursion(root *TreeNode) []int {
	var results = []int{}
	if root == nil {
		return results
	}

	tmp := postorderTraversalRecursion(root.Left)
	for _, i := range tmp {
		results = append(results, i)
	}

	tmp = postorderTraversalRecursion(root.Right)
	for _, i := range tmp {
		results = append(results, i)
	}
	
	results = append(results, root.Val)
	return results
}

func postorderTraversalTraverse(root *TreeNode) []int {
	var results = []int{}
	if root == nil {
		return results
	}
	var prev *TreeNode
	var stack = []*TreeNode{}

	for root != nil || len(stack) > 0 { // root != nil condition for first step.
		for root != nil {
			stack = append(stack, root)
			root = root.Left // move to the left node of subtree
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if root.Right == nil || root.Right == prev {
			results = append(results, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}

	return results
}

func TestPostorder(t *testing.T) {
	res := postorderTraversalRecursion(root)
	var experts = []int{2, 4, 9, 8, 11, 12, 10, 7}
	assert.Equal(t, len(res), len(experts))

	for i := 0; i < len(res); i++ {
		assert.Equal(t, res[i], experts[i])
	}

	res = postorderTraversalTraverse(root)
	assert.Equal(t, len(res), len(experts))

	for i := 0; i < len(res); i++ {
		assert.Equal(t, res[i], experts[i])
	}
}