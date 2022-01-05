package linkedlist

// Node is the linked list node implementation
type Node struct {
	Value int
	Next  *Node
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
	if n.Next == nil {
		return n
	}
	return n.Next.Head()
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
	if n.Value == v {
		n.Value = w
		return true
	}
	return n.Next.Replace(v, w)
}
