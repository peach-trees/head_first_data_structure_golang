package exercises

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	sentry := &ListNode{}
	cur := sentry
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		} // else>>
		cur = cur.Next
		cur.Next = nil
	} // for>
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return sentry.Next
}
