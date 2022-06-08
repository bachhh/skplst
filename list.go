package skplst

import "math/rand"

const MAXLEVEL = 10
const NOTFOUND = -1

var NIL *Node = nil

// SkipList header of the skiplist
type SkipList struct {
	level   uint
	forward []*Node
}

func NewSkipList() *SkipList {
	return &SkipList{forward: []*Node{NIL}}
}

type Node struct {
	forward []*Node
	key     int
}

// level 0 is the leaf node, so
func (this *SkipList) generateLevel() uint {
	r := rand.Int31()
	// instead of calling rand n times, call once and count num of consecutive 1s
	var n uint = 0
	for r&(1<<n) > 0 {
		n++
	}
	return n
}

func (s *SkipList) Search(k int) bool {
	x := s.forward[0]
	for i := s.level; i > 0; i-- {
		for x.forward[i].key < k {
			x = x.forward[i]
		}
	}

	x = x.forward[1]
	if x.key == k {
		return true
	}

	return false
}
