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
    for i := len(m) - 1; i >= 0; i-- {
        if tmp1Head == nil {
            break
        }
        tmp1Head.Val = m[i]
        tmp1Head = tmp1Head.Next
    }
    return tmp2Head
}

// 迭代解法
func reverseList1(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}

// 递归解法
func reverseList2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseList2(head.Next)
    head.Next.Next = head
    head.Next = nil
    
    return newHead
}
