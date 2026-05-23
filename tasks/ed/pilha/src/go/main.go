package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	data []T
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, capacity),
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() error {
	if s.IsEmpty() {
		return fmt.Errorf("stack is empty")
	}
	s.data = s.data[:len(s.data)-1]
	return nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, fmt.Errorf("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack[T]) Clear() {
	s.data = s.data[:0]
}

func (s *Stack[T]) String() string {
	capacidadeTotal := cap(s.data)
	tamanhoAtual := len(s.data)
	
	parts := make([]string, capacidadeTotal)
	for i := 0; i < capacidadeTotal; i++ {
		if i < tamanhoAtual {
			parts[i] = fmt.Sprintf("%v", s.data[i])
		} else {
			parts[i] = "_"
		}
	}
	return strings.Join(parts, ", ")
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewStack[int](10)

	for scanner.Scan() {
		line = scanner.Text()
		
		fmt.Println(line)
		
		if !strings.HasPrefix(line, "$") {
			fmt.Println("$" + line)
		} else {
			fmt.Println(line)
		}
		
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		
		cmd = parts[0]
		if strings.HasPrefix(cmd, "$") {
			cmd = cmd[1:]
		}

		switch cmd {
		case "end":
			return
		case "init":
			capacidade, _ := strconv.Atoi(parts[1])
			v = NewStack[int](capacidade)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Push(value)
			}
		case "debug":
			fmt.Println(v)
		case "top":
			top, err := v.Peek()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "size":
			fmt.Println(v.Size())
		case "pop":
			err := v.Pop()
			if err != nil {
				fmt.Println(err)
			}
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}