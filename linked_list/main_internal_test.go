package main

import (
	"testing"

	"github.com/reedobrien/checkers"
)

func TestInsertBeforeVal(t *testing.T) {

	t.Run("non-empty list", func(t *testing.T) {
		startList := &ListNode{Val: 1}

		toInsert := &ListNode{Val: 10}
		insertBeforeVal(1, startList, toInsert)

		checkers.Equals(t, startList.Val, 10)
		checkers.Equals(t, startList.Next.Val, 1)
	})

	t.Run("third member alter", func(t *testing.T) {
		startList := &ListNode{
			Val:  1,
			Next: &ListNode{Val: 10},
		}

		toInsert := &ListNode{Val: 10000}

		insertBeforeVal(10, startList, toInsert)

		checkers.Equals(t, startList.Val, 1)
		checkers.Equals(t, startList.Next.Val, 10000)
	})
}

func TestInsertAfterVal(t *testing.T) {
	t.Run("after first value", func(t *testing.T) {
		startList := &ListNode{
			Val:  1,
			Next: &ListNode{Val: 10},
		}

		toInsert := &ListNode{Val: 10000}

		insertAfterVal(1, startList, toInsert)

		checkers.Equals(t, startList.Val, 1)
		checkers.Equals(t, startList.Next.Val, 10000)
	})

	t.Run("after second value", func(t *testing.T) {
		startList := &ListNode{
			Val:  1,
			Next: &ListNode{Val: 10},
		}

		toInsert := &ListNode{Val: 10000}

		insertAfterVal(10, startList, toInsert)

		checkers.Equals(t, startList.Val, 1)
		checkers.Equals(t, startList.Next.Val, 10)
		checkers.Equals(t, startList.Next.Next.Val, 10000)
	})
}
