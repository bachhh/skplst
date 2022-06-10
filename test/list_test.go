package test

import (
	"math/rand"
	"skplst"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkipList(t *testing.T) {
	const count int = 100
	testCases := make([]int, count)
	for i := 0; i < count; i++ {
		testCases[i] = int(rand.Int31n(int32(count * 10)))
	}

	list := skplst.NewSkipList()
	for i := range testCases {
		require.True(t, list.Insert(testCases[i]))
	}

}
