package _100_Same_Tree

import "testing"

func TestSameTreeNegativeCase1(t *testing.T) {
	var root1, root2 TreeNode
	root1.Val = 1
	root1.Left = new(TreeNode)
	root1.Left.Val = 2
	root1.Right = new(TreeNode)
	root1.Right.Val = 3

	root2.Val = 1
	root2.Right = new(TreeNode)
	root2.Right.Val = 2
	root2.Left = new(TreeNode)
	root2.Left.Val = 3

	b := isSameTree(&root1, &root2)
	if b != false {
		t.Fail()
	}
}

func TestSameTreeNegativeCase2(t *testing.T) {
	var root1, root2 TreeNode
	root1.Val = 1
	root1.Left = new(TreeNode)
	root1.Left.Val = 2

	root2.Val = 1
	root2.Right = new(TreeNode)
	root2.Right.Val = 2

	b := isSameTree(&root1, &root2)
	if b != false {
		t.Fail()
	}
}

func TestSameTreePositvieCase1(t *testing.T) {
	var root1, root2 TreeNode
	root1.Val = 1
	root1.Left = new(TreeNode)
	root1.Left.Val = 2
	root1.Right = new(TreeNode)
	root1.Right.Val = 3

	root2.Val = 1
	root2.Left = new(TreeNode)
	root2.Left.Val = 2
	root2.Right = new(TreeNode)
	root2.Right.Val = 3

	b := isSameTree(&root1, &root2)
	if b != true {
		t.Fail()
	}
}
