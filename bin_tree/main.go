package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value   int32
	Left  *Node
	Right *Node
}

// type BinTree struct {
// 	root *Node
// }

func NewNode(val int32) *Node {
	return &Node{
		Value:   val,
		Left:  nil,
		Right: nil,
	}
}

func (n *Node) Insert(val int32) {
	if n.Value < val {
		if n.Right == nil {
			n.Right = NewNode(val)
		} else {
			n.Right.Insert(val)
		}
	} else if n.Value > val {
		if n.Left == nil {
			n.Left = NewNode(val)
		} else {
			n.Left.Insert(val)
		}
	}
}

func (n *Node) Search(val int32) bool {
	if n == nil {
		return false
	}

	if n.Value < val {
		return n.Right.Search(val)
	} else if n.Value > val {
		return n.Left.Search(val)
	}
	return true
}

func PrintInorder(root *Node) {
	if root != nil {
		PrintInorder(root.Left)
		fmt.Println(root.Value)
		PrintInorder(root.Right)
	}
}

func invert(root *Node) *Node {
	if root != nil {
		root.Left, root.Right = invert(root.Right), invert(root.Left)
	}
	return root
}

func isBalanced(root *Node) bool {
	h := height(root) 
	return h == - 1
}

func height(root *Node) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	if leftHeight == -1 { return -1 }

	rightHeight := height(root.Left.Right)
	if rightHeight == -1 { return -1 }

	if math.Abs(float64(leftHeight) - float64(rightHeight)) > 1 {
		return -1
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

func main() {
	tree := NewNode(100)
	tree.Insert(200)
	tree.Insert(300)

	fmt.Println(tree)
	fmt.Println("isBalanced:", isBalanced(tree))
	fmt.Println("Inverted:", invert(tree))
	fmt.Println("Inverted isBalanced:", isBalanced(tree))
}