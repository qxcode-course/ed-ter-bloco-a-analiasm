package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Node struct{
	Value int
	next *Node
	prev *Node
	root *Node
}

func (n *Node) Next() *Node{
	if n.next == n.root{
		return nil
	}
	return n.next
}
func (n *Node) Prev() *Node{
	if n.prev == n.root{
		return nil
	}
	return n.prev
}

type LList struct{
	root *Node
	size int
}

func NewLList() *LList{
	sentinel:= &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	sentinel.root = sentinel
	return &LList{root: sentinel, size: 0}
}

func (l *LList) Size() int{
	return l.size
}

func (l *LList) Clear(){
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}
	
func (l *LList) Front() *Node{
	if l.size == 0{
		return nil
	}
	return l.root.next
}
	
func (l *LList) Back() *Node {
	if l.size == 0 {
		return nil
	}
	return l.root.prev
}

func (l *LList) PushBack(value int) {
	l.Insert(l.root, value) 
}

func (l *LList) PushFront(value int) {
	l.Insert(l.root.next, value) 
}

func (l *LList) PopBack() {
	if l.size > 0 {
		l.Remove(l.root.prev)
	}
}

func (l *LList) PopFront() {
	if l.size > 0 {
		l.Remove(l.root.next)
	}
}

func (l *LList) Search(value int) *Node {
	for it := l.root.next; it != l.root; it = it.next {
		if it.Value == value {
			return it
		}
	}
	return nil
}

func (l *LList) Insert(node *Node, value int) {
	newNode := &Node{Value: value, root: l.root}
	
	newNode.next = node
	newNode.prev = node.prev
	
	node.prev.next = newNode
	node.prev = newNode
	
	l.size++
}

func (l *LList) Remove(node *Node) *Node {
	if node == l.root {
		return nil
	}
	
	next := node.next
	prev := node.prev
	
	prev.next = next
	next.prev = prev
	
	l.size--
	
	if next == l.root {
		return nil
	}
	return next
}

func (l *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for it := l.Front(); it != nil; it = it.Next() {
		sb.WriteString(strconv.Itoa(it.Value))
		if it.Next() != nil {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
	func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			 fmt.Println(ll.String())
		case "size":
			 fmt.Println(ll.Size())
		case "push_back":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			 }
		case "push_front":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushFront(num)
			 }
		case "pop_back":
			 ll.PopBack()
		case "pop_front":
			 ll.PopFront()
		case "clear":
			 ll.Clear()
		case "walk":
			 fmt.Print("[ ")
			 for node := ll.Front(); node != nil; node = node.Next() {
			 	fmt.Printf("%v ", node.Value)
			 }
			 fmt.Print("]\n[ ")
			 for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.Value)
			 }
			 fmt.Println("]")
		case "replace":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	node.Value = newvalue
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "insert":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	ll.Insert(node, newvalue)
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "remove":
			 oldvalue, _ := strconv.Atoi(args[1])
		 node := ll.Search(oldvalue)
			 if node != nil {
			 	ll.Remove(node)
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
