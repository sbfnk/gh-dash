package prview

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func newBlockedModel(t *testing.T, resolved ...bool) Model {
	t.Helper()
	m := newTestModelForChecks(t, checksTestOptions{})
	m.pr.Data.Primary.MergeStateStatus = "BLOCKED"
	threads := &m.pr.Data.Enriched.ReviewThreads
	threads.Nodes = slices.Grow(threads.Nodes, len(resolved))[:len(resolved)]
	for i, isResolved := range resolved {
		threads.Nodes[i].IsResolved = isResolved
	}
	return m
}

func TestViewMergeStatus_BlockedCountsUnresolvedThreads(t *testing.T) {
	m := newBlockedModel(t, false, true, false)
	got, status := m.viewMergeStatus()

	require.Equal(t, statusFailure, status)
	require.Contains(t, got, "Merging is blocked")
	require.Contains(t, got, "2 unresolved review threads")
}

func TestViewMergeStatus_BlockedSingleUnresolvedThread(t *testing.T) {
	m := newBlockedModel(t, true, false)
	got, _ := m.viewMergeStatus()

	require.Contains(t, got, "1 unresolved review thread")
	require.NotContains(t, got, "threads")
}

func TestViewMergeStatus_BlockedAllThreadsResolved(t *testing.T) {
	m := newBlockedModel(t, true, true)
	got, _ := m.viewMergeStatus()

	require.Contains(t, got, "Merging is blocked")
	require.NotContains(t, got, "unresolved")
}
