package t19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}

	var sliceNodes []*ListNode
	current := head
	for {
		sliceNodes = append(sliceNodes, current)
		current = current.Next
		if current.Next == nil {
			sliceNodes = append(sliceNodes, current)
			break
		}
	}

	removeIndex := len(sliceNodes) - n

	if removeIndex == 0 {
		return sliceNodes[removeIndex+1]
	}

	prev := sliceNodes[removeIndex-1]

	if removeIndex+1 >= len(sliceNodes) {
		prev.Next = nil
	} else {
		prev.Next = sliceNodes[removeIndex+1]
	}

	return head
}
