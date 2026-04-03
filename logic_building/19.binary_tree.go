package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Left  *Node
	Value int
	Right *Node
}

func (n *Node) insert_node(value int) {
	if value < n.Value {
		//Insert in Left
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.insert_node(value)
		}
	} else {
		//Insert in Right
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.insert_node(value)
		}
	}
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Println(n.Value)
	n.Right.InOrder()
}

func main() {
	depth := 6
	var root *Node
	is_root := true

	for i := 1; i <= depth; i++ {
		value := rand.Intn(20) + 1
		if is_root {
			root = &Node{
				Value: value,
			}
			is_root = false
		} else {
			root.insert_node(value)
		}
	}

	fmt.Println("In-order traversal:")
	root.InOrder()
}
