package list


// https://leetcode.cn/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	m := make([]int, 0)
	tmp1Head := head
	tmp2Head := head
	for {
		if head == nil {
			break
		}
		m = append(m, head.Val)
		head = head.Next
	}
	for i := len(m)-1; i >= 0; i-- {
		if tmp1Head == nil {
			break
		}
		tmp1Head.Val = m[i]
		tmp1Head = tmp1Head.Next
	}
	return tmp2Head
}

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}