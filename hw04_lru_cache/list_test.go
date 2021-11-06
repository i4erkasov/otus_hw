package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [<nil>][10][<nil>]
		l.PushBack(20)  // [<nil>][10][20][<nil>]
		l.PushBack(30)  // [<nil>][10][20][30][<nil>]

		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 30
		l.Remove(middle)         // [10, 30]

		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}

		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)

		// проверим в обратном порядке (reverse)
		elems = make([]int, 0, l.Len())
		for i := l.Back(); i != nil; i = i.Prev {
			elems = append(elems, i.Value.(int))
		}

		require.Equal(t, []int{50, 30, 10, 40, 60, 80, 70}, elems)
	})

	t.Run("remove all", func(t *testing.T) {
		l := NewList()

		// заполняем наш список
		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		}

		for i := l.Front(); i != nil; i = i.Next {
			l.Remove(i)
		}

		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})
}
