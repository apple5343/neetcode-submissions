/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l := head
	r := head
	for i := 0; i < n; i++ {
		r = r.Next
	}
	var prev *ListNode
	for r != nil {
		prev = l
		l = l.Next
		r = r.Next
	}
	if prev == nil {
		return l.Next
	}
	prev.Next = l.Next
	return head
}
