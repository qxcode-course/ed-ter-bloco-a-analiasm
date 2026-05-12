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
}

type LList struct{
	root *Node
	size int
}

func NewLList() *LList{
	root := &Node{}
	root.next = root
	root.prev = root
	return &LList{root: root, size: 0}
}

func(l *LList)insert(value int, prev, next *Node){
	newNode := &Node{Value: value, prev: prev, next: next}
	prev.next = newNode
	next.prev = newNode
	l.size++
}
func (l *LList) remove(node *Node){
	if node == l.root{
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

	func (l*LList) Size() int{
	return l.size
	}

	func (l *LList) Clear(){
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
	}

	func (l *LList) PushFront(value int){
	l.insert(value, l.root, l.root.next)
	}

	func (l *LList) PushBack(value int){
	l.insert(value, l.root.prev, l.root)
	}

	func (l *LList) PopFront(){
	if l.size > 0{
		l.remove(l.root.next)
	}
}

func (l *LList) PopBack(){
	if l.size > 0{
		l.remove(l.root.prev)
	}
}

func (l *LList) String() string{
	var elements []string
	for it := l.root.next; it != l.root; it = it.next{
		elements = append(elements, fmt.Sprint(it.Value))
	}
	return "[" + strings.Join(elements, ", ") + "]"
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
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
