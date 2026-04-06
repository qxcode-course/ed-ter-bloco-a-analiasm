package main

import (
	"bufio"
	"fmt"
	"os"
)

func e_valor(grid [][]rune, lin, c int, value rune) bool{
	numlin := len(grid) 
	numc := len(grid[0])
	if c < 0 || c >= numc || lin < 0 || lin >= numlin{
		return false
	}
	return grid[lin][c] == value
}

func burnTrees(grid [][]rune, lin, c int) {
	if !e_valor(grid, lin, c, '#'){
		return 
	}
	grid[lin][c] = 'o'
	burnTrees(grid, lin, c - 1)
	burnTrees(grid, lin, c + 1)
	burnTrees(grid, lin - 1, c)
	burnTrees(grid, lin + 1, c)


	// se estiver fora da matriz, retorne
	// se o elemento atual não for uma arvore, retorne
	// queime a arvore colocando o caractere 'o' na posição atual
	// chame a recursão para todos os 4 vizinhos
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

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
