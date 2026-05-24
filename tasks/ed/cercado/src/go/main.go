package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	nl := len(board)
	if nl == 0 {
		return
	}
	nc := len(board[0])

	dRow := []int{-1, 1, 0, 0}
	dCol := []int{0, 0, -1, 1}

	marcarBorda := func(startL, startC int) {
		if board[startL][startC] != 'O' {
			return
		}

		stack := []Pos{{l: startL, c: startC}}
		board[startL][startC] = '*' 

		for len(stack) > 0 {
			atual := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for i := 0; i < 4; i++ {
				novoL := atual.l + dRow[i]
				novoC := atual.c + dCol[i]

				if novoL >= 0 && novoL < nl && novoC >= 0 && novoC < nc {
					if board[novoL][novoC] == 'O' {
						board[novoL][novoC] = '*'
						stack = append(stack, Pos{l: novoL, c: novoC})
					}
				}
			}
		}
	}

	for i := 0; i < nl; i++ {
		marcarBorda(i, 0)
		marcarBorda(i, nc-1)
	}

	for j := 0; j < nc; j++ {
		marcarBorda(0, j)
		marcarBorda(nl-1, j)
	}

	for i := 0; i < nl; i++ {
		for j := 0; j < nc; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}