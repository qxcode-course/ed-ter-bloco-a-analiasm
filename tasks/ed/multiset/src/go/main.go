package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type MultiSet struct{
	data []int
}

func NewMultiSet(capacity int) *MultiSet{
	return &MultiSet{
		data: make([]int, 0, capacity),
	}
}

func (ms *MultiSet) betterSearch(value int) (bool, int) {
	low := 0
	high := len(ms.data)

	for low < high {
		mid := low + (high-low)/2
		if ms.data[mid] == value{
			return true, mid
		} else if ms.data[mid] < value{
			low = mid + 1
		} else{
			high = mid
		}
	}
	return false, low
}

func (ms *MultiSet) Insert(value int){
	_, pos := ms.betterSearch(value)

	ms.data = append(ms.data, 0)

	copy(ms.data[pos+1:], ms.data[pos:])

	ms.data[pos] = value
}

func (ms *MultiSet) Contains(value int) bool{
	found, _ := ms.betterSearch(value)
	return found
}

func (ms *MultiSet) Erase(value int) {
	found, pos := ms.betterSearch(value)
	if !found{
		fmt.Println("value not found")
		return
	}
	ms.data = append(ms.data[:pos], ms.data[pos+1:]...)
}

func (ms *MultiSet) Count(value int) int{
	found, pos := ms.betterSearch(value)

	if !found {
		return 0
	}
	count := 0

	for i := pos; i >= 0 && ms.data[i] == value; i--{
		count++
	}

	for i := pos + 1; i < len(ms.data) && ms.data[i] == value; i++{
		count++
	}

	return count
}

func (ms *MultiSet) Unique() int{
	if len(ms.data) == 0{
		return 0
	}
	uniques := 1
	for i := 1; i < len(ms.data); i++{
		if ms.data[i] != ms.data[i-1]{
			uniques++
		}
	}
	return uniques
}

func (ms *MultiSet) Clear(){
	ms.data = ms.data[:0]
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var sb strings.Builder

	for i, v := range slice{
		sb.WriteString(strconv.Itoa(v))

		if i < len(slice)-1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ms *MultiSet

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}
		args := strings.Fields(line)
		cmd := args[0]

		switch cmd {

		case "end", "$end":
			return

		case "init", "$init":
			cap, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(cap)

		case "insert", "$insert":
			for _, part := range args[1:]{
				val, _ := strconv.Atoi(part)
				ms.Insert(val)
			}
		case "show", "$show":
			fmt.Printf("[%s]\n", Join(ms.data, ", "))

		case "contains", "$contains":
			val, _ := strconv.Atoi(args[1])
			if ms.Contains(val){
				fmt.Println("true")
			}else{
				fmt.Println("false")
			}

		case "erase", "$erase":
			val, _ := strconv.Atoi(args[1])
			ms.Erase(val)
		case "count", "$count":
			val, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(val))
		case "unique", "$unique":
			fmt.Println(ms.Unique())
		case "clear", "$clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
