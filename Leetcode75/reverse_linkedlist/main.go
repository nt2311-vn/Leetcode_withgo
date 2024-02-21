package main

import "fmt"

type ListNode struct {
	Val  int
	next *ListNode
}

func main() {
	linkedList := ListNode{
		Val: 1,
		next: &ListNode{
			Val: 2,
			next: &ListNode{
				Val: 3,
				next: &ListNode{
					Val: 4,
					next: &ListNode{
						Val: 5, next: nil,
					},
				},
			},
		},
	}

	reverse := reverseList(&linkedList)

	fmt.Println(reverse)
}

func reverseList(head *ListNode) *ListNode {
	var temp *ListNode
	curr := head
	for curr != nil {
		newNode := curr.next
		curr.next = temp
		temp = curr
		curr = newNode
	}
	return temp
}

func printList(head *ListNode) {
	curr := head

	for curr != nil {
		fmt.Printf("%d", curr.Val)
		curr = curr.next
	}
	fmt.Println()
}
