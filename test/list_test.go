package test

import (
	"math/rand"
	"skplst"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkipList(t *testing.T) {
	const count int = 100
	rand.Seed(int64(time.Now().Nanosecond()))
	testCases := make([]int, count)
	for i := 0; i < count; i++ {
		testCases[i] = int(rand.Int31n(int32(count * 10)))
	}

	list := skplst.NewSkipList()
	for i := range testCases {
		list.Insert(testCases[i])
	}

	lastVal := -1
	for cur := list.Head; cur.GetType() != skplst.TAILTYPE; cur = cur.Forward[0] {
		assert.Truef(t, cur.Key >= lastVal, "error: %d>%d key should be increasing", cur.Key, lastVal)
		lastVal = cur.Key
	}

	// test search
	for i := range testCases {
		require.Truef(t, list.Search(testCases[i]), "did not found %d", testCases[i])
	}

}
