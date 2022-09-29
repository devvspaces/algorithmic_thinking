package queues

type Node struct {
	value any
	next  *Node
}

type Queue struct {
	length int
	head   *Node
	tail   *Node
}

func CreateQueue() Queue {
	return Queue{
		length: 0,
	}
}

func (q *Queue) Dequeue() any {

	if q.length == 0 {
		return nil
	}

	q.length--

	head := q.head
	q.head = head.next
	head.next = nil

	return head.value
}

func (q *Queue) Enqueue(value any) {
	q.length++

	node := Node{value: value}

	if q.head == nil {
		q.head = &node
		q.tail = &node
		return
	}

	q.tail.next = &node
	q.tail = &node
}

func (q *Queue) Peek() any {
	return q.head.value
}

// func main() {
// 	q := CreateQueue()
// 	q.Enqueue(5)
// 	q.Enqueue('a')
// 	q.Enqueue("Futa")
// 	fmt.Println(q.head, q.tail, q.length)
// 	oldHeadValue := q.Dequeue()
// 	fmt.Println(oldHeadValue)
// 	fmt.Println(q.head, q.tail, q.length)
// }
