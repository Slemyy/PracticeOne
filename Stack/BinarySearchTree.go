package Stack

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (bst *BinarySearchTree) Insert(value int) {
	newNode := &TreeNode{data: value}

	if bst.root == nil {
		bst.root = newNode
		return
	} else {
		insertNode(bst.root, newNode)
	}
}

func insertNode(node, newNode *TreeNode) {
	if newNode.data < node.data {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else if newNode.data > node.data {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (bst *BinarySearchTree) Print() {
	printTree(bst.root, "", true, true)
}

func printTree(node *TreeNode, indent string, hasRightSibling bool, last bool) {
	if node != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("┗╺")
			indent += "  "
		} else {
			if hasRightSibling {
				fmt.Print("┣╸")
				indent += "┃ "
			} else {
				fmt.Print("┗╺")
				indent += "   "
			}
		}
		fmt.Println(node.data)

		printTree(node.left, indent, node.right != nil, false)
		printTree(node.right, indent, hasRightSibling, true)
	}
}

func (bst *BinarySearchTree) DepthFirstSearch() {
	if bst.root == nil {
		return
	}

	stack := Create()
	current := bst.root

	for current != nil || !stack.IsEmpty() {
		for current != nil {
			stack.Push(current)
			current = current.left
		}

		pop, _ := stack.Pop()
		fmt.Printf("%d ", pop.(*TreeNode).data)

		current = pop.(*TreeNode).right
	}
}
