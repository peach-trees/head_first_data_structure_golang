package exercises

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentry := &ListNode{Next: head}
	var pre, cur *ListNode
	cur = sentry
	count := 1
	for head != nil {
		if count >= n {
			pre = cur
			cur = cur.Next
		} // if>
		head = head.Next
		count++
	} // for>
	if pre != nil && pre.Next != nil {
		pre.Next = pre.Next.Next
	}
	return sentry.Next
}
