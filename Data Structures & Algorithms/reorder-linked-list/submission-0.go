/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	border := slow
	prev := border
	node := border.Next
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	left := head
	right := prev
	for right != border {
		nextRight := right.Next
		nextLeft := left.Next
		right.Next = left.Next
		left.Next = right
		right = nextRight
		left = nextLeft
	}
	border.Next = nil
}
