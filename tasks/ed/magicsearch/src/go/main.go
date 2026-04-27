package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	left := 0
	right := len(slice) - 1

	for left <= right{
		mid := left + (right-left)/2

		if slice[mid] == value{

			lastPos := mid
			for i := mid + 1; i < len(slice); i++ {
				if slice[i] == value{
					lastPos = i
				} else{
					break
				}
			}
			return lastPos
		}

		if slice[mid] < value {
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return left
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan(){
		return
	}
	line := scanner.Text()

	line = strings.ReplaceAll(line, "[", "")
	line = strings.ReplaceAll(line, "]", "")
	parts := strings.Fields(line)

	slice := make([]int, 0)
	for _, elem := range parts{
		val, err := strconv.Atoi(elem)
		if err == nil{
			slice = append(slice, val)
		}
	}

	if !scanner.Scan(){
		return
	}
	value, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	fmt.Println(MagicSearch(slice, value))
	
}
