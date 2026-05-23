package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	rows := len(grid)
	if rows == 0 {
		return false
	}

	cols := len(grid[0])
	length := len(word)

	if length > rows * cols{
		return false
	}

	gridCounts := make(map[byte]int)
	for r := 0; r < rows; r++{
		for c := 0; c < cols; c++{
			gridCounts[grid[r][c]]++
		}
	}

	wordCounts := make(map[byte]int)
	for i := 0; i < length; i++ {
		wordCounts[word[i]]++
	}

	for char, count := range wordCounts{
		if gridCounts[char] < count{
			return false
		}
	}

	if gridCounts[word[length-1]] < gridCounts[word[0]]{
		reversed := make([]byte, length)
		for i := 0; i < length; i++ {
			reversed[i] = word[length-1-i]
		}
		word = string(reversed)
	}

	var dfs func(r, c, index int)bool
	dfs = func(r, c, index int)bool{

		if index == length{
			return true
		}

		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != word[index] {
			return false
		}

		temp := grid[r][c]
		grid[r][c] = '#'

		found := dfs(r+1, c, index+1) ||
			dfs(r-1, c, index+1) ||
			dfs(r, c+1, index+1) ||
			dfs(r, c-1, index+1)

		grid[r][c] = temp

		return found
	}
	for r := 0; r < rows; r++{
		for c := 0; c < cols; c++{
			if dfs(r, c, 0){
				return true
			}
		}
	}
	return false

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
