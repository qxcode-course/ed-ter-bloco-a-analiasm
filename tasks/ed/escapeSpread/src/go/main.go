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
	cleanGrid := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++{
		var row []byte
		for j := 0; j < len(grid[i]); j++{
			if grid[i][j] != ' '{
				row = append(row, grid[i][j])
			}
		}
		cleanGrid[i] = row
	}
	grid = cleanGrid

	m, n := len(grid), len(grid[0])

	fireTime := make([][]int, m)

	for i := 0; i < m; i++ {
		fireTime[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fireTime[i][j] = 1e9 
		}
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

	dirs := [][]int{
		{-1,0}, {1,0}, {0, -1}, {0, 1},
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

	canEscape := func(wait int) bool{
		if fireTime[0][0] != 1e9 && fireTime[0][0] <= wait {
			return false
		}

		visited := make([][]bool, m)
		for i := 0; i < m; i++{
			visited[i] = make([]bool, n)
		}
		queuePerson := list.New()
		queuePerson.PushBack([]int{0, 0, wait})
		visited[0][0] = true

		for queuePerson.Len() > 0 {
			curr := queuePerson.Remove(queuePerson.Front()).([]int)
			r, c, t := curr[0], curr[1], curr[2]

			if r == m-1 && c == n-1{
				return true
			}

			for _, d := range dirs{
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] != '2' && !visited[nr][nc]{
					
					nextTime := t + 1

					if nr == m-1 && nc == n-1{
						if fireTime[nr][nc] == 1e9 || nextTime <= fireTime[nr][nc]{
							return true
						}
					}else{
						if fireTime[nr][nc] == 1e9 || nextTime < fireTime[nr][nc]{
							visited[nr][nc] = true
							queuePerson.PushBack([]int{nr, nc, nextTime})
						}
					}
				}
			}
		}
		return false
	}

	low, high := 0, int(1e9)
	ans := -1

	for low <= high{
		mid := low + (high-low)/2

		if canEscape(mid){
			ans = mid
			low = mid + 1
		}else{
			high = mid - 1
		}
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
