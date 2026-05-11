package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
	"strings"
)
type OrderedSet struct{
	items []int
}

func NewSet(capacity int) *OrderedSet{
	return &OrderedSet{items: make([]int, 0, capacity)}
}

func (s *OrderedSet) findIndex(value int) int{
	return sort.Search(len(s.items), func(i int) bool{
		return s.items[i] >= value
	})
}

func (s *OrderedSet) Contains(value int) bool{
	idx := s.findIndex(value)
	return idx < len(s.items) && s.items[idx] == value
}

func (s *OrderedSet) Insert(value int){
	idx := s.findIndex(value)

	if idx < len(s.items) && s.items[idx] == value{
		return
	}

	s.items = append(s.items, 0)
	copy(s.items[idx+1:], s.items[idx:])
	s.items[idx] = value
}

func (s *OrderedSet) Erase(value int) bool{
	idx := s.findIndex(value)
	if idx < len(s.items) && s.items[idx] == value{
		s.items = append(s.items[idx:], s.items[idx+1:]... )
		return true
	}
	return false
}

func (s *OrderedSet) String() string{
	if len(s.items) == 0 {
		return "[]"
	}
	res := make([]string, len(s.items))
	for i, v := range s.items{
		res[i] = strconv.Itoa(v)
	}
	return "[" + strings.Join(res, ", ") + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    var set *OrderedSet
	
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd := parts[0]
		fmt.Printf("$%s\n", line)

		switch cmd {
		case "end":
			return
		case "init":
			cap, _ := strconv.Atoi(parts[1])
			set = NewSet(cap)
		case "insert":
		    for _, part := range parts[1:] {
		 	val, _ := strconv.Atoi(part)
			set.Insert(val)
			}
		case "show":
			fmt.Println(set.String())
		case "erase":
		   val, _ := strconv.Atoi(parts[1])
		   if !set.Erase(val){
			fmt.Println("value not found")
		   }
		case "contains":
		 val, _ := strconv.Atoi(parts[1])
		 fmt.Printf("%t\n", set.Contains(val))
	    default:
			fmt.Println("fail: comando invalido")
		}
	}
}
