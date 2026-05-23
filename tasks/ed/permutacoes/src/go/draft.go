package main
import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)
func gerarPermutacoes(chars []byte, visitado []bool, atual []byte){

    if len(atual) == len(chars){
        fmt.Println(string(atual))
        return
    }

    for i := 0; i < len(chars); i++ {
      if !visitado[i] {
          visitado[i] = true

			atual = append(atual, chars[i])
            gerarPermutacoes(chars, visitado, atual)

            atual = atual[:len(atual)-1]
			visitado[i] = false
		}
	}
}

func main() {

    scanner := bufio.NewScanner(os.Stdin)

    if !scanner.Scan() {
		return
	}

    input := strings.TrimSpace(scanner.Text())
	chars := []byte(input)

    sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

    visitado := make([]bool, len(chars))
	atual := make([]byte, 0, len(chars))

    gerarPermutacoes(chars, visitado, atual)
}
