package algorithm

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type treeNodeStack struct {
	stack []*TreeNode
}

func (s *treeNodeStack) push(node *TreeNode) {
	s.stack = append(s.stack, node)
}

func (s *treeNodeStack) pop() *TreeNode {
	r := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return r
}

func (s *treeNodeStack) len() int {
	return len(s.stack)
}

func (s *treeNodeStack) peek() *TreeNode {
	return s.stack[len(s.stack)-1]
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	return preorderTraversal2(root)

	r := make([]int, 0)
	if root == nil {
		return r
	}

	r = append(r, root.Val)

	if root.Left != nil {
		rl := preorderTraversal(root.Left)
		r = append(r, rl...)
	}
	if root.Right != nil {
		rr := preorderTraversal(root.Right)
		r = append(r, rr...)
	}
	return r
}

func preorderTraversal2(root *TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}

	stack := treeNodeStack{}
	stack.push(root)
	for stack.len() > 0 {
		curr := stack.pop()
		if curr.Right != nil {
			stack.push(curr.Right)
		}
		if curr.Left != nil {
			stack.push(curr.Left)
		}
		r = append(r, curr.Val)
	}

	return r
}

func TestBinaryTreeInorderTraversal(t *testing.T) {

	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	//

	r := preorderTraversal(&tree)
	fmt.Println(r)
	//r1 := []int{1, 2, 3}

	tree2 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	r2 := preorderTraversal(&tree2)
	fmt.Println(r2)
	//r1 := []int{1, 2, 3}

	r3 := inorderTraversal(&tree)
	r4 := inorderTraversal(&tree2)
	fmt.Println("inorder:", r3, r4)
	//inorderTraversal2(tree)

	/*
					 1
				    / \
				   2   7
				  / \   \
				 4   3   8
		     		/ \
		           5   6
	*/
	tree3 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 7,
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	// 4 2 5 3 6 1 7 8
	r = preorderTraversal(&tree3)
	fmt.Println("preorder:", r)
	r5 := inorderTraversal2(&tree3)
	r6 := inorderTraversal3(&tree3)
	fmt.Println("inorder:", r5, r6)

	r = postorderTraversal(&tree3)
	fmt.Println("postorderTraversal:", r)
}

// 左中右
func inorderTraversal(root *TreeNode) []int {
	return inorderTraversal2(root)
	r := make([]int, 0)
	if root == nil {
		return r
	}

	if root.Left != nil {
		r = append(r, inorderTraversal(root.Left)...)
	}
	r = append(r, root.Val)
	if root.Right != nil {
		r = append(r, inorderTraversal(root.Right)...)
	}
	return r
}

func inorderTraversal2(root *TreeNode) []int {
	r := make([]int, 0)

	if root == nil {
		return r
	}

	/*
			 1
		    /
		   2
		  / \
		 4   3
	*/
	// 4 2 3 1

	/*
					 1
				    / \
				   2   7
				  / \   \
				 4   3   8
		     		/ \
		           5   6
	*/
	// 前序 1 2 4 3 5 6 7 8
	// 中序 4 2 5 3 6 1 7 8
	// 后序 4 5 6 3 2 8 7 1
	stack := treeNodeStack{}
	stack.push(root)
	for stack.len() > 0 {
		curr := stack.pop()

		// 标记位
		if curr == nil {
			r = append(r, stack.pop().Val)
			continue
		}

		if curr.Right != nil {
			stack.push(curr.Right)
		}
		stack.push(curr)
		stack.push(nil) // 标记位
		if curr.Left != nil {
			stack.push(curr.Left)
		}
	}

	return r
}

func inorderTraversal3(root *TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}

	stack := treeNodeStack{}
	curr := root
	for curr != nil || stack.len() > 0 {
		for curr != nil {
			stack.push(curr)
			curr = curr.Left
		}

		curr = stack.pop()
		r = append(r, curr.Val)
		curr = curr.Right
	}
	return r
}

// 后序 左 右 中
func postorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}

	stack := treeNodeStack{}
	curr := root
	var last *TreeNode
	for curr != nil || stack.len() > 0 {
		for curr != nil {
			stack.push(curr)
			curr = curr.Left
		}

		curr = stack.peek()
		if curr.Right == nil || curr.Right == last {
			r = append(r, curr.Val)
			stack.pop()
			last = curr
			curr = nil
		} else {
			curr = curr.Right
		}
	}

	return r
}
