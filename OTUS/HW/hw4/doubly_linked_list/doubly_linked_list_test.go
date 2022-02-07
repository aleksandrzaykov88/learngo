package doublylinkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	list := NewList()
	list.PushFront(5)
	list.PushFront(6)
	fmt.Println(list.LastElem.Value())

	require.Equal(t, 6, list.LastElem.Value())

	list.Remove(*list.FirstElem)

	require.Equal(t, 1, list.Len())

	list.PushBack(55)
	list.PushFront(666)

	require.Equal(t, 666, list.LastElem.Value())
	require.Equal(t, 55, list.FirstElem.Value())
	require.Equal(t, 3, list.Len())

	i := list.FirstElem.Next()
	list.Remove(*i)

	require.Equal(t, 2, list.Len())
}
