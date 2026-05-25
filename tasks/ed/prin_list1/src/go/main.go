package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	var elements []string
	for it := l.Front(); it != l.End(); it = it.Next(){
		valueStr := fmt.Sprint(it.Value)
		if it == sword{
			elements = append(elements, ">"+valueStr)
		}else {
			elements = append(elements, valueStr)
		}
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	proximo := it.next

	if proximo == l.root{
		proximo = proximo.next
	}

	return proximo
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {

		
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
