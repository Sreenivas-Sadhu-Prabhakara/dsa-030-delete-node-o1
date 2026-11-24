package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// DeleteNode removes the given node from its list in O(1), without access to
// the head: it copies the successor's value into the node and splices the
// successor out. Returns false for a nil node or a tail node — a tail cannot
// be removed without its predecessor.
func DeleteNode(n *Node) bool {
	if n == nil || n.Next == nil {
		return false
	}
	n.Val = n.Next.Val
	n.Next = n.Next.Next
	return true
}

func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	head := fromSlice([]int{4, 5, 1, 9})
	DeleteNode(head.Next) // delete the node holding 5
	fmt.Println(toSlice(head))
}
