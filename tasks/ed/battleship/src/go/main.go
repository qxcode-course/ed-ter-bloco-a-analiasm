package main

import (
	"bufio"
	"fmt"
	"os"
)

// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {
	
	if len(board) == 0 {
		return 0
	}

	count := 0
	linhas := len(board)
	colunas := len(board[0])

	visitados := make([][]bool, linhas)

	for i := range visitados{
		visitados[i] = make([]bool, colunas)
	}

	for i := 0; i < linhas; i++{
		for j := 0; j < colunas; j++{
			if board[i][j] == 'X' && !visitados[i][j]{
				count++
			}

			k := j

			for k < linhas && board[i][k] == 'X'{
				visitados[i][k] = true
				k++
			}

			k = i

			for k < linhas && board[k][j] == 'X'{
				visitados[k][j] = true
				k++
			}
		}
	}
	return count
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
