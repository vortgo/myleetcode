package t24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sub := ListNode{
		Next: head,
	}

	sliceNodes := make([]*ListNode, 0)

	current := sub.Next
	for current != nil {
		sliceNodes = append(sliceNodes, current)

		current = current.Next
	}

	i := 0
	for {
		if i+1 >= len(sliceNodes) {
			break
		}

		sliceNodes[i], sliceNodes[i+1] = sliceNodes[i+1], sliceNodes[i]
		i += 2
	}

	for i := 0; i < len(sliceNodes)-1; i++ {
		sliceNodes[i].Next = sliceNodes[i+1]
	}

	sliceNodes[len(sliceNodes)-1].Next = nil

	return sliceNodes[0]
}
