package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Point armazena uma coordenada na matriz do labirinto
type Point struct {
	r, c int
}

// Stack simplificada baseada nas orientações nativas do Go
type Stack struct {
	data []Point
}

func (s *Stack) Push(p Point) {
	s.data = append(s.data, p)
}

func (s *Stack) Pop() Point {
	if len(s.data) == 0 {
		return Point{-1, -1}
	}
	p := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return p
}

func (s *Stack) Top() Point {
	if len(s.data) == 0 {
		return Point{-1, -1}
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}

	// Lê as dimensões da matriz
	dimensions := strings.Fields(scanner.Text())
	rows, _ := strconv.Atoi(dimensions[0])
	cols, _ := strconv.Atoi(dimensions[1])

	// Monta a matriz do labirinto e localiza o início 'I' e fim 'F'
	grid := make([][]rune, rows)
	var start, end Point

	for i := 0; i < rows; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			grid[i] = []rune(line)
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 'I' {
					start = Point{r: i, c: j}
				} else if grid[i][j] == 'F' {
					end = Point{r: i, c: j}
				}
			}
		}
	}

	// Pilhas para controle do percurso conforme o algoritmo
	caminho := &Stack{}
	becos := &Stack{}

	caminho.Push(start)

	// Direções de movimento: Cima, Baixo, Esquerda, Direita
	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	for !caminho.IsEmpty() {
		atual := caminho.Top()

		// Se o atual for o ponto final, encontramos o caminho
		if atual.r == end.r && atual.c == end.c {
			break
		}

		// Encontra vizinhos válidos (não paredes e não visitados)
		var validos []Point
		for i := 0; i < 4; i++ {
			nr := atual.r + dr[i]
			nc := atual.c + dc[i]

			// Verifica se está dentro dos limites da matriz
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				char := grid[nr][nc]
				// Espaço vazio ou o próprio Fim são destinos válidos
				if char == ' ' || char == 'F' {
					validos = append(validos, Point{r: nr, c: nc})
				}
			}
		}

		if len(validos) > 0 {
			proximo := validos[0]
			// Se não for o destino final, marcamos visualmente o avanço com o ponto '.'
			if grid[proximo.r][proximo.c] != 'F' {
				grid[proximo.r][proximo.c] = '.'
			}
			caminho.Push(proximo)
		} else {
			// Não há caminhos válidos a partir daqui: encontramos um beco sem saída
			beco := caminho.Pop()
			// Evitamos desmarcar as posições fixas de Início 'I' e Fim 'F'
			if grid[beco.r][beco.c] == '.' {
				grid[beco.r][beco.c] = 'x' // Marcação temporária de beco
				becos.Push(beco)
			}
		}
	}

	// Limpa as marcações temporárias de beco ('x') voltando-as para espaços vazios (' ')
	for !becos.IsEmpty() {
		b := becos.Pop()
		grid[b.r][b.c] = ' '
	}

	// Imprime o labirinto resolvido exatamente no formato esperado
	for i := 0; i < rows; i++ {
		fmt.Println(string(grid[i]))
	}
}