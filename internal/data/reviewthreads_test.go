package data

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func makeReviewThreads(resolved ...bool) ReviewThreads {
	threads := ReviewThreads{TotalCount: len(resolved)}
	threads.Nodes = slices.Grow(threads.Nodes, len(resolved))[:len(resolved)]
	for i, isResolved := range resolved {
		threads.Nodes[i].IsResolved = isResolved
	}
	return threads
}

func TestReviewThreadsHasUnresolved(t *testing.T) {
	require.False(t, makeReviewThreads().HasUnresolved())
	require.False(t, makeReviewThreads(true, true).HasUnresolved())
	require.True(t, makeReviewThreads(true, false).HasUnresolved())
}
