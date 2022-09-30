package linked_list

import "testing"

func HelperCreateTestLinkedList() *LinkedList {
	list := &LinkedList{}
	nodeA := CreateNode("A")
	nodeB := CreateNode("B")
	nodeA.next = nodeB
	nodeB.prev = nodeA
	list.head = nodeA
	list.tail = nodeB
	list.length = 2
	return list
}

func HelperCreateTestLinkedListMore() *LinkedList {
	list := HelperCreateTestLinkedList()
	nodeC := CreateNode("C")
	list.tail.next = nodeC
	nodeC.prev = list.tail
	list.tail = nodeC
	list.length = 3
	return list
}

func TestDoublyLinkedCreateNode(t *testing.T) {

	node := CreateNode(1)
	if node.value != 1 {
		t.Errorf("Expected 1 got %v", node.value)
	}

}

func TestDoublyLinkedListPrepend(t *testing.T) {

	// prepending to a new list
	list := &LinkedList{}
	list.Prepend(4)
	if list.head.value != 4 || list.tail.value != 4 {
		t.Errorf("Expected 4 got %v", list.head.value)
	}

	// prepending to an already created list
	list.Prepend("test")
	if list.head.value != "test" {
		t.Errorf("Expected %q got %v", "test", list.head.value)
	}
	if list.head.next.value != 4 || list.tail.value != 4 {
		t.Errorf("Expected 4 got %v", list.head.value)
	}

}

func TestDoublyLinkedListAppend(t *testing.T) {

	// appending to a new list
	list := &LinkedList{}
	list.Append(4)
	if list.tail.value != 4 || list.head.value != 4 {
		t.Errorf("Expected 4 got %v", list.head.value)
	}

	// appending to an already created list
	list.Append("test")
	if list.tail.value != "test" {
		t.Errorf("Expected %q got %v", "test", list.tail.value)
	}
	if list.tail.prev.value != 4 || list.head.value != 4 {
		t.Errorf("Expected 4 got %v", list.head.value)
	}

}

func TestDoublyLinkedListGet(t *testing.T) {

	// Get from an empty list

	t.Run("Get from Empty List", func(s *testing.T) {
		defer func() {
			if err := recover(); err != "Out of bounds" {
				s.Errorf("Expected error: Out of bounds, Got: %v", err)
			}
		}()
		list := &LinkedList{}
		list.Get(0)
	})

	t.Run("Get from a new list", func(s *testing.T) {
		list := HelperCreateTestLinkedList()
		computed := list.Get(0)
		expected := list.head.value
		if expected != computed {
			t.Errorf("Expected %q got %v", expected, computed)
		}
	})

}

func TestDoublyLinkedListPop(t *testing.T) {

	t.Run("pop an empty list", func(sub *testing.T) {
		list := &LinkedList{}
		if list.Pop() != nil {
			t.Errorf("Expected nil got %v", list.head.value)
		}
	})

	t.Run("pop a filled list", func(sub *testing.T) {
		list := HelperCreateTestLinkedList()
		tail := list.tail

		computed := list.Pop()
		if computed != tail.value {
			t.Errorf("Expected %q got %v", tail.value, computed)
		}
	})

}

func TestDoublyLinkedListUnshift(t *testing.T) {

	t.Run("unshift an empty list", func(sub *testing.T) {
		list := &LinkedList{}
		if list.Unshift() != nil {
			t.Errorf("Expected nil got %v", list.head.value)
		}
	})

	t.Run("unshift a filled list", func(sub *testing.T) {
		list := HelperCreateTestLinkedList()
		head := list.head

		computed := list.Unshift()
		if computed != head.value {
			t.Errorf("Expected %q got %v", head.value, computed)
		}
	})

}

func TestDoublyLinkedListInsertAt(t *testing.T) {

	t.Run("insert on out of bounds index", func(sub *testing.T) {
		defer func() {
			// Test if got panic from test
			if err := recover(); err != "Out of bounds" {
				sub.Errorf("Expected error: Out of bounds, Got: %v", err)
			}
		}()

		list := &LinkedList{}
		list.InsertAt(4, -1)
	})

	t.Run("insert on an empty list", func(t *testing.T) {
		list := &LinkedList{}
		list.InsertAt(4, 0)
		if list.head.value != 4 {
			t.Errorf("Expected 4 got %v", list.head.value)
		}
	})

	t.Run("insert at the end", func(t *testing.T) {
		list := HelperCreateTestLinkedList()
		expected := "test"
		list.InsertAt(expected, list.length)
		if list.tail.value != expected {
			t.Errorf("Expected %q got %v", expected, list.tail.value)
		}
	})

	t.Run("insert at the center", func(t *testing.T) {
		list := HelperCreateTestLinkedList()
		expected := "test"
		list.InsertAt(expected, 1)
		computed := list.head.next.value

		if computed != expected {
			t.Errorf("Expected %q got %v", expected, computed)
		}
	})

}

func TestDoublyLinkedListRemoveAt(t *testing.T) {

	t.Run("remove at on empty list", func(sub *testing.T) {
		defer func() {
			// Test if got panic from test
			if err := recover(); err != "Out of bounds" {
				sub.Errorf("Expected error: Out of bounds, Got: %v", err)
			}
		}()

		list := &LinkedList{}
		list.RemoveAt(0)
	})

	t.Run("remove at start index", func(t *testing.T) {
		list := HelperCreateTestLinkedList()
		expected := list.head.value
		computed := list.RemoveAt(0)
		if expected != computed {
			t.Errorf("Expected %v, Got %v", expected, computed)
		}
	})

	t.Run("remove at end index", func(t *testing.T) {
		list := HelperCreateTestLinkedList()
		expected := list.tail.value
		computed := list.RemoveAt(list.length - 1)
		if expected != computed {
			t.Errorf("Expected %v, Got %v", expected, computed)
		}
	})

	t.Run("remove at center index", func(t *testing.T) {
		list := HelperCreateTestLinkedListMore()
		expected := list.head.next.value
		computed := list.RemoveAt(list.length - 2)
		if expected != computed {
			t.Errorf("Expected %v, Got %v", expected, computed)
		}
	})
}
