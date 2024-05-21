package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	empty := Slice[int]{make([]int, 0)}
	copiedEmpty := empty.Copy()
	require.Equalf(t, empty, copiedEmpty, "empty")

	padding := Slice[int]{make([]int, 10)}
	copiedPadding := padding.Copy()
	require.Equalf(t, padding, copiedPadding, "padding")

	list := Slice[int]{[]int{1, 2, 3, 4, 5}}
	copiedList := list.Copy()
	require.Equalf(t, list, copiedList, "list")
}

func TestDelete(t *testing.T) {
	list := Slice[int]{[]int{1, 2, 3, 4, 5}}

	list.Delete(-1)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5}, "delete out of range from left")

	list.Delete(5)
	require.Equalf(t, list.internal, []int{1, 2, 3, 4, 5}, "delete out of range from right")

	list.Delete(0)
	require.Equalf(t, list.internal, []int{2, 3, 4, 5}, "delete at index 0")

	list.Delete(len(list.internal) - 1)
	require.Equalf(t, list.internal, []int{2, 3, 4}, "delete at last element index")

	list.Delete(1)
	require.Equalf(t, list.internal, []int{2, 4}, "delete at index")
}