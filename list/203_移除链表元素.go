package list


// https://leetcode.cn/problems/remove-linked-list-elements/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	curr := head.Next
	prev := head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		}else {
			prev = curr
		}
		curr = curr.Next
	}
	if head.Val == val {
		head = head.Next
	}
	return head
}