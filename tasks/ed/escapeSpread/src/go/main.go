package main

import (
	"bufio"
	"fmt"
	"os"
	"container/list"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	fireTime := make([][]int, m)
	personTime := make([][]int, m)

	for i := 0; i < m; i++ {
		fireTime[i] = make([]int, n)
		personTime[i] = make([]int, n)

		for j := 0; j < n; j++ {
			fireTime[i][j] = 1e9 
			personTime[i][j] = 1e9
		}
	}

	dirs := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	queueFire := list.New()

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' { 
				fireTime[i][j] = 0
				queueFire.PushBack([]int{i, j})
			}
		}
	}

	for queueFire.Len() > 0 {
		curr := queueFire.Remove(queueFire.Front()).([]int)
		r, c := curr[0], curr[1]

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] != '2' { 
				if fireTime[nr][nc] == 1e9 { 
					fireTime[nr][nc] = fireTime[r][c] + 1
					queueFire.PushBack([]int{nr, nc})
				}
			}
		}
	}

	queuePerson := list.New()
	personTime[0][0] = 0
	queuePerson.PushBack([]int{0, 0})

	for queuePerson.Len() > 0 {
		curr := queuePerson.Remove(queuePerson.Front()).([]int) 
		r, c := curr[0], curr[1]
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]

			if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] != '2' {
				if personTime[nr][nc] == 1e9 {
					personTime[nr][nc] = personTime[r][c] + 1
					queuePerson.PushBack([]int{nr, nc})
				}
			}
		}
	}

	if personTime[m-1][n-1] == 1e9 {
		return -1
	}

	if fireTime[m-1][n-1] == 1e9 {
		return 1e9 
	}

	ans := fireTime[m-1][n-1] - personTime[m-1][n-1]
	if ans < 0{
		return -1
	}
	return ans
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
