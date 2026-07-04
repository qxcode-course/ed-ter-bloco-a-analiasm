package main

import (
	"fmt"
)

type Point struct{
    row int 
    col int
}

func orangeRotting(grid [][]int)int{
    if len(grid) == 0{
        return 0
    }

    rows := len(grid)
    cols := len(grid[0])
    var queue []Point
    freshOranges := 0
    minutes := 0

    for r := 0; r < rows; r++{
        for c := 0; c < cols; c++{
            if grid[r][c] == 2{
                queue = append(queue, Point{row: r, col: c})
            }else if grid[r][c] == 1{
                freshOranges++
            }
        }
    }

    if freshOranges == 0{
        return  0
    }

    directions := [][] int{
        {-1, 0},
        {1, 0},
        {0, -1},
        {0, 1},
    }

    for len(queue) > 0 && freshOranges > 0{
        minutes++
        currentLevelSize := len(queue)

        for i := 0; i < currentLevelSize; i++{
            curr := queue[0]
            queue = queue[1:]

            for _, d := range directions{
                nextRow := curr.row + d[0]
                nextCol := curr.col + d[1]

                if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols && grid[nextRow][nextCol] == 1{
                    grid[nextRow][nextCol] = 2
                    freshOranges--

                    queue = append(queue, Point{row: nextRow, col: nextCol})
                }
            }
        }
    }

    if freshOranges == 0{
        return minutes
    }

    return -1

}

func main() {
    var rows, cols int

    if _, err := fmt.Scan(&rows, &cols); err != nil{
        return
    }

   grid := make([][]int, rows)

   for i := range grid{
    grid[i] = make([]int, cols)

    for j := 0; j < cols; j++{
        fmt.Scan(&grid[i][j])
    }
   }

    fmt.Println(orangeRotting(grid))   
}