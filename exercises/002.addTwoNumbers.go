package exercises

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	sentry := new(ListNode)
	cur := sentry
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		curSum := carry
		if l1 != nil {
			curSum += l1.Val
			l1 = l1.Next
		} // if>>
		if l2 != nil {
			curSum += l2.Val
			l2 = l2.Next
		} // if>>
		cur.Next = &ListNode{Val: curSum % 10}
		cur = cur.Next
		carry = curSum / 10
	} // for>
	return sentry.Next
}
