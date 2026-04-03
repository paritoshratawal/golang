package main

import "fmt"

type Node struct {
	data int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) addAtFront(value int) {
	new_node := &Node{data: value}
	if l.head == nil {
		l.head = new_node
		l.size++
	} else {
		new_node.next = l.head
		l.head = new_node
		l.size++
	}
	return
}

func (l *LinkedList) addAtEnd(value int) {
	fmt.Println("addAtEnd", value)
	current_node := l.head
	for current_node.next != nil {
		current_node = current_node.next
	}

	new_node := &Node{data: value}
	current_node.next = new_node
	l.size++
}

// func (l *LinkedList) InsertAt(index, v int) error {
// 	if index < 0 || index > l.size {
// 		return fmt.Errorf("index out of range")
// 	}
// 	if index == 0 {
// 		l.addAtFront(v)
// 		return nil
// 	}

// 	cur := l.head
// 	for i := 0; i < index-1; i++ {
// 		cur = cur.next	
// 	}
// 	newNode := &Node{Value: v, Next: cur.Next}
// 	cur.Next = newNode
// 	l.Size++
// 	return nil
// }

// func (l *LinkedList) RemoveFront() error {
// 	if l.Head == nil {
// 		return fmt.Errorf("list is empty")
// 	}
// 	l.Head = l.Head.Next
// 	l.Size--
// 	return nil
// }

// func (l *LinkedList) RemoveBack() error {
// 	if l.Head == nil {
// 		return fmt.Errorf("list is empty")
// 	}
// 	if l.Head.Next == nil {
// 		l.Head = nil
// 		l.Size--
// 		return nil
// 	}
// 	prev, cur := l.Head, l.Head.Next
// 	for cur.Next != nil {
// 		prev = cur
// 		cur = cur.Next
// 	}
// 	prev.Next = nil
// 	l.Size--
// 	return nil
// }

// func (l *LinkedList) RemoveAt(index int) error {
// 	if index < 0 || index >= l.Size {
// 		return fmt.Errorf("index out of range")
// 	}
// 	if index == 0 {
// 		return l.RemoveFront()
// 	}
// 	cur := l.Head
// 	for i := 0; i < index-1; i++ {
// 		cur = cur.Next
// 	}
// 	cur.Next = cur.Next.Next
// 	l.Size--
// 	return nil
// }

// func (l *LinkedList) Find(v int) (int, bool) {
// 	cur := l.Head
// 	idx := 0
// 	for cur != nil {
// 		if cur.Value == v {
// 			return idx, true
// 		}
// 		cur = cur.Next
// 		idx++
// 	}
// 	return -1, false
// }

// func (l *LinkedList) ToSlice() []int {
// 	res := make([]int, 0, l.Size)
// 	cur := l.Head
// 	for cur != nil {
// 		res = append(res, cur.Value)
// 		cur = cur.Next
// 	}
// 	return res
// }

// func (l *LinkedList) String() string {
// 	return fmt.Sprint(l.ToSlice())
// }

func (l *LinkedList) display() {
	if l.head == nil {
		fmt.Println("Linked List is empty")
		return
	}
	current_node := l.head
	for current_node != nil {
		fmt.Println("Data", current_node.data)
		current_node = current_node.next
	}
}

func main() {
	list := &LinkedList{}
	list.addAtFront(30)
	list.addAtFront(20)
	list.addAtFront(10)
	list.display()

	list.addAtEnd(40)
	list.display()

	// list.AddFront(5)
	// fmt.Println("After AddFront(5):", list)

	// _ = list.InsertAt(2, 15)
	// fmt.Println("After InsertAt(2,15):", list)

	// idx, found := list.Find(20)
	// fmt.Printf("Find(20): idx=%d found=%t\n", idx, found)

	// _ = list.RemoveAt(2)
	// fmt.Println("After RemoveAt(2):", list)

	// _ = list.RemoveFront()
	// fmt.Println("After RemoveFront():", list)

	// _ = list.RemoveBack()
	// fmt.Println("After RemoveBack():", list)

	// fmt.Println("Final size:", list.Size)
}

