package bst

import binary_tree "github.com/karthikvempati/data-structures-go/tree"

type BinarySearchTree struct {
	root *binary_tree.BinaryNode
}

//Builds a Binary tree from a sorited array
//Assumption: nums is a sorted array
func (bst *BinarySearchTree) BuildBST(nums []int) *binary_tree.BinaryNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &binary_tree.BinaryNode{Data: nums[0]}
	}
	center := len(nums) / 2
	root := binary_tree.BinaryNode{Data: nums[center], Left: bst.BuildBST(nums[:center]), Right: bst.BuildBST(nums[center+1:])}
	return &root
}

func InsertElement(root *binary_tree.BinaryNode, number int) *binary_tree.BinaryNode {
	if root == nil {
		root = &binary_tree.BinaryNode{Data: number}
	} else if root.Data > number {
		InsertElement(root.Left, number)
	} else if root.Data < number {
		InsertElement(root.Right, number) 
	} else if root.Data == number {
		panic("The number already exists")
	}
	return root
}

func (bst *BinarySearchTree) BstToArray(root *binary_tree.BinaryNode) []int {
	nums := []int{}
	if root != nil {
		nums = append(nums, root.Data)
		nums = append(nums, bst.BstToArray(root.Left)...)
		nums = append(nums, bst.BstToArray(root.Right)...)
	}

	return nums
}

func (bst *BinarySearchTree) InOrderTraversal(root *binary_tree.BinaryNode) []int {
	nums := []int{}
	if root != nil {
		nums = append(nums, bst.InOrderTraversal(root.Left)...)
		nums = append(nums, root.Data)
		nums = append(nums, bst.InOrderTraversal(root.Right)...)
	}
	return nums
}

func (bst *BinarySearchTree) PreOrderTraversal(root *binary_tree.BinaryNode) []int {
	nums := []int{}
	if root != nil {
		nums = append(nums, root.Data)
		nums = append(nums, bst.InOrderTraversal(root.Left)...)
		nums = append(nums, bst.InOrderTraversal(root.Right)...)
	}
	return nums
}
