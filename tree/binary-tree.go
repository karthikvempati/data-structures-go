package binary_tree

type BinaryNode struct {
	Data  int
	Left  *BinaryNode
	Right *BinaryNode
}

type BinaryTree struct {
	Root *BinaryNode
}

func CreateBinaryTree(nums []int) BinaryTree {
	tree := BinaryTree{Root: &BinaryNode{Data: nums[0], Left: nil, Right: nil}}
	for i := 1; i < len(nums); i++ {
		tree.Root.InsertElement(nums[i])
		//fmt.Println(tree.Root.Left)
	}

	return tree
}

func (root *BinaryNode) InsertElement(val int) {

	if root == nil {
		root.Data = val
		return
	}

	queue := Queue{}
	queue.EnQueue(root)
	for ok := true; ok; ok = !(len(queue.items) == 0) {

		node := queue.DeQueue()

		if node.Left != nil {
			queue.EnQueue(node.Left)
		} else if node.Left == nil {
			node.Left = &BinaryNode{Data: val, Left: nil, Right: nil}
			break
		}

		if node.Right != nil {
			queue.EnQueue(node.Right)
		} else {
			node.Right = &BinaryNode{Data: val, Left: nil, Right: nil}
			break
		}
	}
}

func (root *BinaryNode) LevelOrderTraversal() []int {
	traversal := []int{}
	queue := Queue{}
	queue.EnQueue(root)
	for ok := true; ok; ok = !(len(queue.items) == 0) {

		node := queue.DeQueue()

		if node.Left != nil {
			queue.EnQueue(node.Left)
		}

		if node.Right != nil {
			queue.EnQueue(node.Right)
		}

		traversal = append(traversal, node.Data)
	}

	return traversal
}
