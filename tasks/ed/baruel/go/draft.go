package main

import "fmt"

func main() {
	var todoAlbum, totalPossuidas int
	fmt.Scan(&todoAlbum)
	fmt.Scan(&totalPossuidas)

	temnoalbum := make(map[int]bool)
	var repetidas []int
	ultimo := -1

	for i := 0; i < totalPossuidas; i++ {
		var num int
		fmt.Scan(&num) 

		if num == ultimo {
			repetidas = append(repetidas, num)
		}
		temnoalbum[num] = true
		ultimo = num
	}

	if len(repetidas) == 0 {
		fmt.Println("N")
	} else {
		for i, f := range repetidas {
			fmt.Print(f)
			if i < len(repetidas)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	falta := false
	pfalta := true

	for i := 1; i <= todoAlbum; i++ {
		if !temnoalbum[i] {
			if !pfalta {
				fmt.Print(" ")
			}
			fmt.Print(i)
			falta = true
			pfalta = false
		}
	}

	if !falta {
		fmt.Print("N")
	}
	fmt.Println()
}