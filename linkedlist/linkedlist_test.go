package linkedlist

import (
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	n := NewNode(1)

	ok := n.Add(2)
	if !ok {
		t.Error("Expected true, got false")
	}
}

func TestDelete(t *testing.T) {
	n := NewNode(1)

	n.Add(2)
	n.Add(3)

	ok := n.Delete(2)
	if !ok {
		t.Error("Expected true, got false")
	}
}

func TestHead(t *testing.T) {
	n := NewNode(1)

	n.Add(2)
	n.Add(3)

	res := n.Head()

	if res.Value != 1 {
		t.Error("Expected 1, got " + strconv.Itoa(res.Value))
	}
}

func TestTail(t *testing.T) {
	n := NewNode(1)

	n.Add(2)
	n.Add(3)

	res := n.Tail()
	if res.Value != 3 {
		t.Error("Expected 3, got " + strconv.Itoa(res.Value))
	}
}

func TestPrint(t *testing.T) {
	n := NewNode(1)
	n.Add(2)
	n.Add(3)

	res := n.Print()
	extd := ">1->2->3"

	if res != extd {
		t.Error("Expected: " + extd + " Got: " + res)
	}

}

func TestReplace(t *testing.T) {
	n := NewNode(1)
	n.Add(2)
	n.Add(3)

	ok := n.Replace(2, 4)
	if !ok {
		t.Error("replace should be true")
	}

}

func TestLength(t *testing.T) {
	n := NewNode(1)
	n.Add(2)
	n.Add(3)

	res := n.Length()
	if res != 3 {
		t.Error("Expected 3, got " + strconv.Itoa(res))
	}
}

func TestReplaceFail(t *testing.T) {
	n := NewNode(1)
	n.Add(2)
	n.Add(3)

	ok := n.Replace(5, 4)

	if ok {
		t.Error("replace should be false")
	}

}
