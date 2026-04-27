package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data  []int
	index int
}

type ReverseIterator struct {
	data  []int
	index int
}

type CyclicIterator struct {
	data  []int
	index int
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: -1}
}

func (l *MyList) ReverseIterador() *ReverseIterator {
	return &ReverseIterator{data: l.data, index: len(l.data)}
}

func (l *MyList) CyclicIterator() *CyclicIterator {
	return &CyclicIterator{data: l.data, index: -1}
}

func (i *Iterator) HasNext() bool {
	return i.index < len (i.data)-1
}

func (i *Iterator) Next() int {
	i.index++
	return i.data[i.index]
}

func (i *ReverseIterator) HasNext() bool{
	return i.index > 0
}

func (i *ReverseIterator) Next() int{
	i.index--
	return i.data[i.index]
}

func (i *CyclicIterator) HasNext() bool {
	return len(i.data ) > 0
}

func (i *CyclicIterator) Next() int {
	i.index = (i.index + 1) % len(i.data)
	return i.data[i.index]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					val, _ := strconv.Atoi(value)
					slice[i] = val
				}
				mylist = NewMyList(slice)
			}
			
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")

		case "reverse":
			fmt.Print("[ ")
            for it := mylist.ReverseIterador(); it.HasNext();{
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")

		case "cyclic":
			if len(mylist.data) == 0 {continue}
			qtd, _ := strconv.Atoi(args[1])
			fmt.Print("[ ")
			it := mylist.CyclicIterator()
			for n := 0; n < qtd; n++{
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")

		case "end":
			return
		}
	}

}
