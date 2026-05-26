package main
import(
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func totalNQueens(n int) int {
    cols := make([]bool, n)
    diagonal1 := make([]bool, 2*n) 
    diagonal2 := make([]bool, 2*n)

    count := 0

    var backtrack func(r int)
	backtrack = func(r int) {
        if r == n {
			count++
			return
		}

        for c := 0; c < n; c++ {
			idDiag1 := r - c + n
			idDiag2 := r + c

			if cols[c] || diagonal1[idDiag1] || diagonal2[idDiag2] {
				continue
			}

            cols[c] = true
			diagonal1[idDiag1] = true
			diagonal2[idDiag2] = true

            backtrack(r + 1)

            cols[c] = false
			diagonal1[idDiag1] = false
			diagonal2[idDiag2] = false
		}
	}

    backtrack(0)
	return count
}



func main() {

    scanner := bufio.NewScanner(os.Stdin)

    if !scanner.Scan() {
        return
    }

    line := strings.TrimSpace(scanner.Text())

    if line == "" {
        return
	}

    n, _ := strconv.Atoi(line)

    fmt.Println(totalNQueens(n))



}