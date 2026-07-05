package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{p.l - 1, p.c},
		{p.l + 1, p.c},
		{p.l, p.c - 1},
		{p.l, p.c + 1},
	}
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	queue := []Pos{startPos}

	visitados := map[Pos]bool{startPos: true}

	caminho := make(map[Pos]Pos)

	fimEncontrado := false

	for len(queue) > 0 {
		atual := queue[0]
		queue = queue[1:]

		if atual == endPos {
			fimEncontrado = true
			break
		}

		for _, vizinho := range atual.getNeig(){
			if inside(grid, vizinho) && grid[vizinho.l][vizinho.c] != '#' && !visitados[vizinho]{
				visitados[vizinho] = true
				queue = append(queue, vizinho)
				caminho[vizinho] = atual
			}
		}
	}

	 if fimEncontrado{
		var voltar func(atual Pos)
		voltar = func(atual Pos) {
			grid[atual.l][atual.c] = '.'

			if atual == startPos{
				return
			}
			voltar(caminho[atual])
		}
		voltar(endPos)
	 }
}

func voltar(grid [][]rune, caminho map[Pos]Pos, atual Pos, startPos Pos){
	grid[atual.l][atual.c] = '.'

	if atual == startPos{
		return
	}

	anterior := caminho[atual]
	voltar(grid, caminho, anterior, startPos)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
