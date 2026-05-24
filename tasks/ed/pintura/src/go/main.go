package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)


func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	corOriginal := image[sr][sc]

	
	if corOriginal == color {
		return image
	}

	nl := len(image)
	ncMax := len(image[0]) 

	type Pixel struct {
		r, c int
	}

	stack := []Pixel{{r: sr, c: sc}}
	image[sr][sc] = color

	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	for len(stack) > 0 {

		atual := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for i := 0; i < 4; i++ {
			nr := atual.r + dr[i]
			nc := atual.c + dc[i]

			
			if nr >= 0 && nr < nl && nc >= 0 && nc < ncMax {

				if image[nr][nc] == corOriginal {
					image[nr][nc] = color
					stack = append(stack, Pixel{r: nr, c: nc})
				}
			}
		}
	}

	return image
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])

	result := floodFill(image, sr, sc, color)

	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}