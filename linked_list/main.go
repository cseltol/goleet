package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Len  int
}

func (list *LinkedList) prepend(n *Node) {
	list.Len++
	// list.Head, list.Head.Next = n, list.Head

	sec := list.Head
	list.Head = n
	list.Head.Next = sec
}

func (l LinkedList) printList() {
	p := l.Head
	for l.Len > 0 {
		fmt.Printf("%d ", p.Val)
		p = p.Next
		l.Len--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) deleteByVal(val int) {
	if l.Len == 0 {
		return
	}

	if val == l.Head.Val {
		l.Head = l.Head.Next
		l.Len--
	}
	
	prevDel := l.Head
	for prevDel.Next.Val != val {
		if prevDel.Next.Next == nil {
			return
		}
		prevDel = prevDel.Next
	}
	prevDel.Next = prevDel.Next.Next
	l.Len--
}

func main() {
	list := LinkedList{}
	node_1 := &Node{Val: 92}
	node_2 := &Node{Val: 22}
	node_3 := &Node{Val: 32}
	node_4 := &Node{Val: 42}
	node_5 := &Node{Val: 52}

	list.prepend(node_1)
	list.prepend(node_2)
	list.prepend(node_3)
	list.prepend(node_4)
	list.prepend(node_5)

	list.printList()

	list.deleteByVal(22)
	list.printList()

	list.deleteByVal(100)
	list.deleteByVal(52)
	list.printList()

	empty := LinkedList{}
	empty.deleteByVal(1)
}
