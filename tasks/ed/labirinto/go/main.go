package main

import (
	"bufio"
	"fmt"
	"os"
)

	type Pos struct{
		l, c int
	}

	func search(grid [][]rune, momento Pos, fimPos Pos) bool{
		l, c := momento.l, momento.c

		if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]){
		return false
	   } 

	   if grid[l][c] != ' '{
		return false
	   }

	   grid[l][c] = '.'

	   if l == fimPos.l && c == fimPos.c{
		return true
	   }

	   vizinhos := []Pos{
		{l - 1, c}, {l + 1, c}, {l, c + 1}, {l, c - 1}, 
	   }

	   for _, v := range vizinhos{
		if search(grid, v, fimPos){ 
              return true
		}
	   }

	   grid[l][c] = ' '
	   return false

	}
	
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
