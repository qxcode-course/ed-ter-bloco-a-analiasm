package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
}


func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}

func (l *LList) String() string {
	var elements []string
	for it := l.root.next; it != l.root; it = it.next{
		elements = append(elements, fmt.Sprint(it.Value))
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

func equals(a, b *LList) bool{
	itA := a.root.next
	itB := b.root.next

	for itA != a.root && itB != b.root{
		if itA.Value != itB.Value{
			return false
		}
		itA = itA.next
		itB = itB.next
	}
	return itA == a.root && itB == b.root
}

func addsorted(l *LList, value int){
	it := l.root.next

	for it != l.root && it.Value < value{
		it  = it.next
	}

	l.insertBefore(it, value)
}

func reverse(l *LList){
	if l.root.next == l.root{
		return
	}
	curr := l.root

	for {
		curr.next, curr.prev = curr.prev, curr.next
		curr = curr.prev
		if curr == l.root{
			break
		}
	}
}

func merge(a, b *LList) *LList{
	itA := a.root.next
	itB := b.root.next

	for itB != b.root{
		if itA == a.root || itB.Value < itA.Value{
			a.insertBefore(itA, itB.Value)
			itB = itB.next
		} else{
			itA = itA.next
		}
	}
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			 lla := str2list(args[1])
			 llb := str2list(args[2])
			 if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
			 	fmt.Println("diferentes")
			 }
		case "addsorted":
			 lla := NewLList()
			 for i := 1; i < len(args); i++ {
			 	value, _ := strconv.Atoi(args[i])
			 	addsorted(lla, value)
			 }
			 fmt.Println(lla)
		case "reverse":
			 lla := str2list(args[1])
			 reverse(lla)
			 fmt.Println(lla)
		case "merge":
			 lla := str2list(args[1])
			 llb := str2list(args[2])
			 merged := merge(lla, llb)
			 fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
