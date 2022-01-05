package linkedlist

import (
	"fmt"
	"strconv"
)

// llrepr is the string representation of the linked list
var llrepr string

// counter is a counter used to keep track of the length of the linked list
var counter int

// Node is the linked list node implementation
type Node struct {
	Value int
	Next  *Node
}

// NewNode is the constructor for the linked list
// It resets the linked list string and counter as well
func NewNode(v int) *Node {
	// setup
	counter = 0
	llrepr = ""

	// build
	return &Node{
		Value: v,
	}
}

// Add adds a new node to the linked list
func (n *Node) Add(v int) bool {
	if n.Next == nil {
		n.Next = &Node{Value: v}
		return true
	}

	return n.Next.Add(v)
}

// Delete deletes a node from the linked list
func (n *Node) Delete(v int) bool {
	if n.Next == nil {
		return false
	}

	if n.Next.Value == v {
		n.Next = n.Next.Next
		return true
	}

	return n.Next.Delete(v)
}

// Head returns the head of the linked list
func (n *Node) Head() *Node {
	return n
}

// Tail returns the tail of the linked list
func (n *Node) Tail() *Node {
	if n.Next == nil {
		return n
	}

	return n.Next.Tail()
}

// Replace replaces a node with another node
func (n *Node) Replace(v, w int) bool {
	counter++
	tot := n.Length()

	if counter > tot {
		return false
	}

	if n.Value == v {
		n.Value = w
		return true
	}

	return n.Next.Replace(v, w)
}

// Length returns the length of the linked list
func (n *Node) Length() int {
	if n.Next == nil {
		return 1
	}

	return 1 + n.Next.Length()
}

// Print prints the linked list
func (n *Node) Print() string {
	prefix := ">"

	if len(llrepr) != 0 {
		prefix = "->"
	}

	llrepr += fmt.Sprintf("%s%v", prefix, strconv.Itoa(n.Value))

	if n.Next != nil {
		n.Next.Print()
	}

	return llrepr
}
