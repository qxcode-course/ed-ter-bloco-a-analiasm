package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {

	nl := len(grid)

	if nl == 0 {
		return
	}

    nc := len(grid[0])

	stack := []Pos{{l: l, c: c}}

	dRow := []int{-1, 1, 0, 0}
	dCol := []int{0, 0, -1, 1}

	for len(stack) > 0 {
	
		topoIdx := len(stack) - 1
		atual := stack[topoIdx]
		stack = stack[:topoIdx]

		if grid[atual.l][atual.c] == '#' {
			grid[atual.l][atual.c] = 'o'

			for i := 0; i < 4; i++ {
				novoL := atual.l + dRow[i]
				novoC := atual.c + dCol[i]

				if novoL >= 0 && novoL < nl && novoC >= 0 && novoC < nc {
					if grid[novoL][novoC] == '#' {
						stack = append(stack, Pos{l: novoL, c: novoC})
					}
				}
			}
		}
	}
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
