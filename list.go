package skplst

import "math/rand"

const MAXLEVEL = 10

var NIL *Node = nil

// SkipList header of the skiplist
type SkipList struct {
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
	var n uint = 0
	for rand.Int()%2 == 0 && n < MAXLEVEL {
		n++
	}
	return n
}

func (s *SkipList) Search(key int) {

}
