package linked_list

type Node struct {
	value any
	prev  *Node
	next  *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func CreateNode(value any) *Node {
	return &Node{value: value}
}

func (n *LinkedList) Prepend(value any) {

	new_node := CreateNode(value)

	n.length++
	if n.head == nil {
		n.head = new_node
		n.tail = new_node
		return
	}

	n.head.prev = new_node
	new_node.next = n.head
	n.head = new_node

}

func (n *LinkedList) Append(value any) {
	new_node := CreateNode(value)
	n.length++
	if n.tail == nil {
		n.head = new_node
		n.tail = new_node
		return
	}

	n.tail.next = new_node
	new_node.prev = n.tail
	n.tail = new_node
}

func (n *LinkedList) get(idx int) *Node {
	if idx >= n.length || idx < 0 {
		panic("Out of bounds")
	}
	node := n.head
	for i := 0; i < idx; i++ {
		node = node.next
	}
	return node
}

func (n *LinkedList) Get(idx int) any {
	return n.get(idx).value
}

func (n *LinkedList) InsertAt(value any, idx int) {

	if idx > n.length || idx < 0 {
		panic("Out of bounds")
	} else if idx == 0 {
		n.Prepend(value)
		return
	} else if idx == n.length {
		n.Append(value)
		return
	}

	node := n.get(idx)
	new_node := CreateNode(value)
	n.length++

	new_node.prev = node.prev
	new_node.next = node
	node.prev.next = new_node
	node.prev = new_node

}

func (n *LinkedList) RemoveAt(idx int) any {

	if idx >= n.length || idx < 0 {
		panic("Out of bounds")
	} else if idx == 0 {
		return n.Unshift()
	} else if idx == n.length-1 {
		return n.Pop()
	}

	node := n.get(idx)
	n.length--

	node.prev.next = node.next
	node.next.prev = node.prev

	return node.value

}

func (n *LinkedList) PeekHead() any {

	if n.head == nil {
		return nil
	}

	return n.head.value

}

func (n *LinkedList) PeekTail() any {

	if n.tail == nil {
		return nil
	}

	return n.tail.value

}

func (n *LinkedList) Pop() any {

	if n.tail == nil {
		return nil
	}

	node := n.tail
	n.tail = node.prev
	n.tail.next = nil

	return node.value

}

func (n *LinkedList) Unshift() any {

	if n.head == nil {
		return nil
	}

	node := n.head
	n.head = node.next
	n.head.prev = nil

	return node.value

}
