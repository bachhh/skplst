package skplst

import "math/rand"

const MAXLEVEL = 16 // this should cap num nodes at 2^16
const NOTFOUND = -1

var NIL *Node = nil

// SkipList header of the skiplist
type SkipList struct {
	MaxLevel uint // max level of skiplist
	Forward  []*Node
	Count    uint
}

// TODO:
// OPTIMIZE1: forward []*Node doesn't need to store MAXLEVEL, only the current max level
// OPTIMIZE2: faster roll for random level
func NewSkipList() *SkipList {
	return &SkipList{Forward: []*Node{NIL}}
}

type Node struct {
	// Forward[i] store pointers to all level i-th nodes
	// at level 0, we have the standard linked list
	Forward []*Node
	Key     int
	Level   uint
}

// level 0 is the leaf node, so
func (s *SkipList) generateLevel() uint {
	r := rand.Int31()

	// instead of calling rand n times, call once and count num of consecutive 1s
	var n uint = 0
	for r&(1<<n) > 0 && n < MAXLEVEL {
		n++
	}
	return n
}

// Search for key k
// @return true if key k is found
func (s *SkipList) Search(k int) bool {
	curNode := s.Forward[0]
	for i := MAXLEVEL; i > 0; i-- {
		// "skip" to largest node with key < k
		for curNode.Forward[i] != NIL && curNode.Forward[i].Key < k {
			curNode = curNode.Forward[i]
		}
		// if curNode.forward[i].key >= k, we skipped too much, descend to lower level
	}

	// invariant: curNode.key < k
	curNode = curNode.Forward[0] // check the next node
	if curNode != NIL && curNode.Key == k {
		return true
	}
	return false
}

// Insert
// @return true insert not duplicate, false if duplicate key
func (s *SkipList) Insert(k int) bool {

	// this is our level for the node
	updateList := make([]*Node, MAXLEVEL)

	curNode := s.Forward[0]
	for i := MAXLEVEL; i > 0; i-- {
		for curNode.Forward[i] != NIL && curNode.Forward[i].Key < k {
			curNode = curNode.Forward[i]
		}
		updateList[i] = curNode
	}

	curNode = curNode.Forward[0]
	if curNode != NIL && curNode.Key == k {
		return false
	}
	n := s.generateLevel()
	newNode := &Node{Key: k, Forward: make([]*Node, MAXLEVEL), Level: n}
	// if n >= s.MaxLevel+1 {
	// 	n = s.MaxLevel + 1
	// }

	for i := n; i >= 0; i-- {
		newNode.Forward[i] = updateList[i].Forward[i]
		updateList[i].Forward[i] = newNode
	}
	return true
}

func (s *SkipList) Delete(k int) bool {
	updateList := make([]*Node, MAXLEVEL)

	curNode := s.Forward[0]
	for i := MAXLEVEL; i > 0; i-- {
		for curNode.Forward[i] != NIL && curNode.Forward[i].Key < k {
			curNode = curNode.Forward[i]
		}
		updateList[i] = curNode
	}

	curNode = curNode.Forward[0]
	if curNode != NIL && curNode.Key == k {
		return false
	}

	n := curNode.Level
	for i := n; i >= 0; i-- {
		updateList[i].Forward[i] = curNode.Forward[i]
	}
	return true
}
