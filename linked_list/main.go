package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func recursivePrint(list *ListNode) {
	if list == nil {
		return
	}

	if list.Next == nil {
		fmt.Println(list.Val)

		return
	}

	fmt.Printf("%d => ", list.Val)

	recursivePrint(list.Next)
}

func insertBeforeVal(val int, list *ListNode, toInsert *ListNode) {
	if list == nil || toInsert == nil {
		return
	}

	cur := list
	var prevNode *ListNode

	for cur != nil {

		if cur.Val == val {
			if prevNode != nil {
				prevNode.Next = toInsert
				toInsert.Next = cur

				break
			}

			curPtr := *cur

			toInsert.Next = &curPtr

			*list = *toInsert

			break
		}
		prevNode = cur
		cur = cur.Next
	}
}

func insertAfterVal(val int, list *ListNode, toInsert *ListNode) {
	if list == nil || toInsert == nil {
		return
	}

	cur := list

	for cur != nil {
		if cur.Val == val {
			toInsert.Next = cur.Next

			cur.Next = toInsert

			break
		}

		cur = cur.Next
	}
}

func insertAtEnd(list *ListNode, toInsert *ListNode) {
	if list == nil || toInsert == nil {

		return
	}

	cur := list

	for cur != nil {
		if cur.Next == nil {
			cur.Next = toInsert

			break
		}

		cur = cur.Next
	}
}

func insertAtStart(list *ListNode, toInsert *ListNode) {
	if list == nil || toInsert == nil {

		return
	}

	curList := *list

	toInsert.Next = &curList

	*list = *toInsert
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	fmt.Println("Start")
	recursivePrint(l)
	fmt.Println()

	fmt.Println("Second")
	insertAfterVal(3, l, &ListNode{Val: 6})
	recursivePrint(l)
	fmt.Println()

	fmt.Println("Third")
	insertBeforeVal(6, l, &ListNode{Val: 10})
	recursivePrint(l)
	fmt.Println()

	fmt.Println("Fourth")
	insertAtEnd(l, &ListNode{Val: 100, Next: nil})
	recursivePrint(l)
	fmt.Println()

	fmt.Println("5th")
	insertAtStart(l, &ListNode{Val: 100000, Next: nil})
	recursivePrint(l)
	fmt.Println()

	fmt.Println("6th")
	l2 := &ListNode{Val: 99}
	insertBeforeVal(99, l2, &ListNode{Val: 9999})
	recursivePrint(l2)
}
