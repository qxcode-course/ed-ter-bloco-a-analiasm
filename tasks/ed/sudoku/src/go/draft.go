package main
import (
    "bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func ehValido(matriz [][]rune, lin, col int, num rune) bool {
	n := len(matriz)

    for i := 0; i < n; i++ {
		if matriz[lin][i] == num || matriz[i][col] == num {
			return false
		}
	}

    tamanhoBloco := int(math.Sqrt(float64(n)))

    lInicio := (lin / tamanhoBloco) * tamanhoBloco
	cInicio := (col / tamanhoBloco) * tamanhoBloco

    for r := lInicio; r < lInicio+tamanhoBloco; r++ {
		for c := cInicio; c < cInicio+tamanhoBloco; c++ {
			if matriz[r][c] == num {
				return false
			}
		}
	}

	return true
}

func resolver(matriz [][]rune, index int) bool {
	n := len(matriz)

    if index == n*n {
		return true
	}

    l := index / n
	c := index % n


    if matriz[l][c] != '.' {
		return resolver(matriz, index+1)
	}

    for d := 1; d <= n; d++ {
		numChar := rune('0' + d)
        
        if ehValido(matriz, l, c, numChar) {
			matriz[l][c] = numChar

           if resolver(matriz, index+1) {
				return true
			}
            
            matriz[l][c] = '.' 
		}
	}

    return false 
}

func main() {

    scanner := bufio.NewScanner(os.Stdin)

    if !scanner.Scan() {
		return
	}

    var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)

    matriz := make([][]rune, n)
	for i := 0; i < n; i++ {

        if scanner.Scan() {
			linha := strings.TrimSpace(scanner.Text())
			matriz[i] = []rune(linha)
		}
	}

    if resolver(matriz, 0) {

		for i := 0; i < n; i++ {
			fmt.Println(string(matriz[i]))
		}
	}
}
