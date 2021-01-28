package tencent

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode = nil

	for head.Next != nil {
		next := head.Next
		head.Next = prev

		prev = head
		head = next
	}

	head.Next = prev

	return head
}

// 递归
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil

	return cur
}
