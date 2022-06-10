package skplst

import (
	"math/rand"
)

const MAXLEVEL = 15 // this should cap num nodes at 2^16
const NOTFOUND = -1

const (
	HEADTYPE = 1
	TAILTYPE = 2
)

// SkipList header of the skiplist
type SkipList struct {
	MaxLevel uint  // max level of skiplist
	Head     *Node // point to header
	Count    uint
}

func NewForward() []*Node { return make([]*Node, MAXLEVEL+1) }

// TODO:
// OPTIMIZE1: forward []*Node doesn't need to store MAXLEVEL, only the current max level
// OPTIMIZE2: faster roll for random level
func NewSkipList() *SkipList {
	tail := &Node{nodeType: TAILTYPE}
	head := &Node{nodeType: HEADTYPE, Forward: NewForward()}
	for i := range head.Forward {
		head.Forward[i] = tail
	}
	return &SkipList{Head: head}
}

type Node struct {
	// Forward[i] store pointers to all level i-th nodes
	// at level 0, we have the standard linked list
	Forward []*Node
	Key     int
	Level   uint

	nodeType int // 0 = normal, 1 = head, 2 = tail
}

func (n *Node) GetType() int {
	return n.nodeType
}

// level 0 is the leaf node, so
func (s *SkipList) generateLevel() int {
	r := rand.Int31()

	// instead of calling rand n times, call once and count num of consecutive 1s
	var n int = 0
	for r&(1<<n) > 0 && n < MAXLEVEL {
		n++
	}
	return n
}

// return true if k is less than key of node n
func (s *SkipList) lessThan(node *Node, k int) bool {
	if node == nil {
		return false
	}

	if node.nodeType == TAILTYPE { // everything is less than TAIL
		return true
	} else if node.nodeType == HEADTYPE {
		return false
	}
	return k < node.Key

}

// Search for key k
// @return true if key k is found
func (s *SkipList) Search(k int) bool {

	curNode := s.Head
	for i := MAXLEVEL; i >= 0; i-- {
		for !s.lessThan(curNode.Forward[i], k) {
			curNode = curNode.Forward[i]
		}
	}

	if curNode.GetType() != TAILTYPE && curNode.Key == k {
		return true
	}
	return false
}

// Insert
// @return true insert not duplicate, false if duplicate key
func (s *SkipList) Insert(k int) bool {

	// this is our level for the node
	updateList := make([]*Node, MAXLEVEL+1)

	curNode := s.Head
	for i := MAXLEVEL; i >= 0; i-- {
		for !s.lessThan(curNode.Forward[i], k) {
			curNode = curNode.Forward[i]
		}
		updateList[i] = curNode
	}

	if curNode.GetType() != TAILTYPE && curNode.Key == k {
		return false
	}

	n := s.generateLevel()
	newNode := &Node{Key: k, Forward: make([]*Node, MAXLEVEL), Level: uint(n)}

	for i := n; i >= 0; i-- {
		newNode.Forward[i] = updateList[i].Forward[i]
		updateList[i].Forward[i] = newNode
	}
	return true
}

func (s *SkipList) Delete(k int) bool {
	updateList := make([]*Node, MAXLEVEL+1)

	curNode := s.Head
	for i := MAXLEVEL; i > 0; i-- {
		for s.lessThan(curNode, k) {
			curNode = curNode.Forward[i]
		}
		updateList[i] = curNode
	}

	curNode = curNode.Forward[0]
	if curNode.nodeType != TAILTYPE && curNode.Key == k {
		return false
	}

	n := curNode.Level
	for i := n; i >= 0; i-- {
		updateList[i].Forward[i] = curNode.Forward[i]
	}
	return true
}
