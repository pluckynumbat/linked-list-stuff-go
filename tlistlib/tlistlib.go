package tlistlib

type TailedList struct {
	head, tail *Node
}

type Node struct {
	data string
	next *Node
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	return n.data
}

func (tl *TailedList) String() string {
	if tl == nil {
		return ""
	}

	str := ""

	runner := tl.head
	for runner != nil {
		str += runner.String()
		str += "->"
		runner = runner.next
	}
	str += "nil"
	return str
}

func constructListFromNode(node *Node) (tlist TailedList) {
	tlist = TailedList{}
	tlist.head = node
	tlist.tail = node

	return
}
