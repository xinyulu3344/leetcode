package list


// https://leetcode.cn/problems/middle-of-the-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val in *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	size := 0
	curr := head
	for curr != nil {
		size++
		curr = curr.Next
	}
	count := size >> 1
	for i := 0; i < count; i++ {
		head = head.Next
	}
	return head
}

func middleNode2(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		} else if fast.Next != nil {
			return slow.Next
		} else {
			return slow
		}
	}
}
