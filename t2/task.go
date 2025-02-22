package t2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode

	addNext := false
	var currentNode *ListNode
	for l1 != nil || l2 != nil || addNext {
		firstArg := 0
		secondArg := 0
		fromPrev := 0
		if addNext {
			fromPrev = 1
			addNext = false
		}
		if l1 != nil {
			firstArg = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			secondArg = l2.Val
			l2 = l2.Next
		}
		num := firstArg + secondArg + fromPrev
		if num > 9 {
			addNext = true
			num = num % 10
		}

		if result == nil {
			currentNode = &ListNode{Val: num}
			result = currentNode
		} else {
			node := &ListNode{Val: num}
			currentNode.Next = node
			currentNode = node
		}
	}

	return result
}
