package exercises

import "container/heap"

type IntHeap []*ListNode

func (i IntHeap) Len() int {
	return len(i)
}

func (i IntHeap) Less(m, n int) bool {
	return i[m].Val < i[n].Val
}

func (i IntHeap) Swap(m, n int) {
	i[m], i[n] = i[n], i[m]
}

func (i *IntHeap) Push(x interface{}) {
	*i = append(*i, x.(*ListNode))
}

func (i *IntHeap) Pop() interface{} {
	n := len(*i)
	x := (*i)[n-1]
	*i = (*i)[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var checkHeap IntHeap
	heap.Init(&checkHeap)
	for i := range lists {
		if lists[i] != nil {
			heap.Push(&checkHeap, lists[i])
		} // if>>
	} // for>

	sentry := &ListNode{}
	cur := sentry
	for checkHeap.Len() > 0 {
		curMin := heap.Pop(&checkHeap).(*ListNode)
		cur.Next = curMin
		cur = cur.Next
		if curMin.Next != nil {
			heap.Push(&checkHeap, curMin.Next)
		} // if>>
	} // for>
	return sentry.Next
}
