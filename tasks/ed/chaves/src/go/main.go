package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var fila []string

	for i := 0; i < 16; i++ {
		fila = append(fila, string(rune('A'+i)))
	}

	for len(fila) > 1 {

		timeA := fila[0]
		timeB := fila[1]
		fila = fila[2:]

		if !scanner.Scan() {
			break
		}

		var golsA, golsB int
		fmt.Sscanf(scanner.Text(), "%d %d", &golsA, &golsB)

		if golsA > golsB {
			fila = append(fila, timeA)
		} else {
			fila = append(fila, timeB)
		}
	}

	fmt.Println(fila[0])
}
