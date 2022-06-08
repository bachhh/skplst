package skplst

import "math/rand"

const MAXLEVEL = 16 // this should cap num nodes at 2^16
const NOTFOUND = -1

var NIL *Node = nil

// SkipList header of the skiplist
type SkipList struct {
	maxLevel uint // max level of skiplist
	forward  []*Node
	count    uint
}

func NewSkipList() *SkipList {
	return &SkipList{forward: []*Node{NIL}}
}

type Node struct {
	// forward[i] store pointers to all level i-th nodes
	// at level 0, we have the standard linked list
	forward []*Node
	key     int
}

// level 0 is the leaf node, so
func (s *SkipList) generateLevel() uint {
	r := rand.Int31()

	// instead of calling rand n times, call once and count num of consecutive 1s
	var n uint = 0
	for r&(1<<n) > 0 {
		n++
	}
	return n
}

// Search for key k
// @return true if key k is found
func (s *SkipList) Search(k int) bool {
	curNode := s.forward[0]
	for i := s.maxLevel; i > 0; i-- {
		// "skip" to largest node with key < k
		for curNode.forward[i] != NIL && curNode.forward[i].key < k {
			curNode = curNode.forward[i]
		}
		// if curNode.forward[i].key >= k, we skipped too much, descend to lower level
	}

	// invariant: curNode.key < k
	curNode = curNode.forward[0] // check the next node
	if curNode != NIL && curNode.key == k {
		return true
	}
	return false
}

// Insert
// @return true insert not duplicate, false if duplicate key
func (s *SkipList) Insert(k int) bool {

	// this is our level for the node
	updateList := []*Node{}

	curNode := s.forward[0]
	for i := s.maxLevel; i > 0; i-- {
		for curNode.forward[i] != NIL && curNode.forward[i].key < k {
			curNode = curNode.forward[i]
		}
		if updateList[len(updateList)-1] != curNode {
			updateList = append(updateList, curNode.forward[i])
		}
	}

	curNode = curNode.forward[0]
	if curNode != NIL && curNode.key == k {
		return false
	}
	n := s.generateLevel()
	return true
}
