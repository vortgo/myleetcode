package m1

type Node struct {
	next *Node
	v    int
}

func revert(head *Node) *Node {
	current := head
	var prev *Node
	var next *Node
	for current != nil {
		next = current.next
		current.next = prev
		prev = current

		current = next
	}

	return prev
}
