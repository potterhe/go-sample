package algorithm

import (
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

func TestBinaryTreeInorderTraversal(t *testing.T) {
	/*
			 1
		    / \
		   2   3
	*/
	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	/*
			 1
		    /
		   2
		  / \
		 4   3
	*/
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

	tests := []struct {
		input     *TreeNode
		preorder  []int
		inorder   []int
		postorder []int
	}{
		{
			input:     &tree,
			preorder:  []int{1, 2, 3},
			inorder:   []int{2, 1, 3},
			postorder: []int{2, 3, 1},
		},
		{
			input:     &tree2,
			preorder:  []int{1, 2, 4, 3},
			inorder:   []int{4, 2, 3, 1},
			postorder: []int{4, 3, 2, 1},
		},
		{
			input:     &tree3,
			preorder:  []int{1, 2, 4, 3, 5, 6, 7, 8},
			inorder:   []int{4, 2, 5, 3, 6, 1, 7, 8},
			postorder: []int{4, 5, 6, 3, 2, 8, 7, 1},
		},
	}

	var r []int
	for _, test := range tests {
		r = preorderTraversal(test.input)
		if !intSliceEqual(r, test.preorder) {
			t.Errorf("preorder fail: except: %v, %v given\n", test.preorder, r)
		}
		r = preorderTraversal2(test.input)
		if !intSliceEqual(r, test.preorder) {
			t.Errorf("preorder fail: except: %v, %v given\n", test.preorder, r)
		}

		r = inorderTraversal(test.input)
		if !intSliceEqual(r, test.inorder) {
			t.Errorf("inorder fail: except: %v, %v given\n", test.inorder, r)
		}
		r = inorderTraversal2(test.input)
		if !intSliceEqual(r, test.inorder) {
			t.Errorf("inorder fail: except: %v, %v given\n", test.inorder, r)
		}

		r = postorderTraversal(test.input)
		if !intSliceEqual(r, test.postorder) {
			t.Errorf("postorder fail: except: %v, %v given\n", test.postorder, r)
		}
		r = postorderTraversal2(test.input)
		if !intSliceEqual(r, test.postorder) {
			t.Errorf("postorder fail: except: %v, %v given\n", test.postorder, r)
		}
	}
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
	//	return preorderTraversal2(root)

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

		r = append(r, curr.Val)
		if curr.Right != nil {
			stack.push(curr.Right)
		}
		if curr.Left != nil {
			stack.push(curr.Left)
		}
	}

	return r
}

// 左中右
func inorderTraversal(root *TreeNode) []int {
	//	return inorderTraversal2(root)
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

func inorderTraversal3(root *TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}

	return r
}

// 后序 左 右 中
func postorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}

	if root.Left != nil {
		r = append(r, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		r = append(r, postorderTraversal(root.Right)...)
	}
	r = append(r, root.Val)

	return r
}

func postorderTraversal2(root *TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}

	stack := treeNodeStack{}
	curr := root
	var last *TreeNode
	for stack.len() > 0 || curr != nil {
		for curr != nil {
			stack.push(curr)
			curr = curr.Left
		}
		curr = stack.peek()
		if curr.Right != nil && curr.Right != last {
			curr = curr.Right
		} else {
			r = append(r, curr.Val)
			stack.pop()
			last = curr
			curr = nil
		}
	}

	return r
}

func intSliceEqual(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}
