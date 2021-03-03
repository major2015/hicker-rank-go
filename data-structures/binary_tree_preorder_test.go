package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
    Val int
    Left *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversalRecursion(root *TreeNode) []int {
	var results []int
	if root == nil {
		return results
	}
	
	results = append(results, root.Val)
	tmp := preorderTraversalRecursion(root.Left)
	for _, i := range tmp {
		results = append(results, i)
	}

	tmp = preorderTraversalRecursion(root.Right)
	for _, i := range tmp {
		results = append(results, i)
	}
	
	return results
}

func preorderTraversalTraverse(root *TreeNode) []int {
	var results []int
	if root == nil {
		return results
	}

	var stack = []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		top := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		results = append(results, top.Val)

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}

	return results
}

var root = &TreeNode{
	Val: 7,
	Left: &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 8,
			Left: nil,
			Right: &TreeNode{
				Val: 9,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 12,
			Right: nil,
			Left: &TreeNode{
				Val: 11,
				Left: nil, 
				Right: nil,
			},
		},
	},
}

func TestPreorder(t *testing.T) {
	res := preorderTraversalRecursion(root)
	var experts = []int{7, 4, 2, 10, 8, 9, 12, 11}
	assert.Equal(t, len(res), len(experts))

	for i := 0; i < len(res); i++ {
		assert.Equal(t, res[i], experts[i])
	}

	res = preorderTraversalTraverse(root)
	assert.Equal(t, len(res), len(experts))

	for i := 0; i < len(res); i++ {
		assert.Equal(t, res[i], experts[i])
	}
}