package main
import ("bufio"
	"fmt"
	"os"
	"strconv"
	"strings")

func ehValido(s []byte, digito byte, pos int, L int) bool{
    
    n := len(s)
    
    inicio := pos - L
    if inicio < 0 {
        inicio = 0
    }

    for i := inicio; i < pos; i++ {
		if s[i] == digito {
			return false
		}
	}
    
    fim := pos + L
	if fim >= n {
		fim = n - 1
	}

	for i := pos + 1; i <= fim; i++ {
		if s[i] == digito {
			return false
		}
	}

    return true
}

func resolver(s []byte, index int, L int) bool {
	if index == len(s) {
		return true
	}

	if s[index] != '.' {
		// Garante que o número fixo inicial não quebra as regras
		if !ehValido(s, s[index], index, L) {
			return false
		}
		return resolver(s, index+1, L)
	}
    for d := 0; d <= L; d++ {
		digitoChar := byte('0' + d)

		if ehValido(s, digitoChar, index, L) {
			s[index] = digitoChar 

			if resolver(s, index+1, L) {
				return true
			}
            s[index] = '.'
		}
	}

	return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)


    if !scanner.Scan() {
		return
	}

    linhaSeq := strings.TrimSpace(scanner.Text())
	s := []byte(linhaSeq)

    if !scanner.Scan() {
		return
	}
    linhaL := strings.TrimSpace(scanner.Text())
	L, _ := strconv.Atoi(linhaL)

    if resolver(s, 0, L) {
		fmt.Println(string(s))
	}
}
