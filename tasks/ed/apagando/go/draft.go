package main
import "fmt"
func main() {

	var n, m, id int
	fmt.Scan(&n)
	filaOrig := make([]int, n)

	for i := 0; i < n; i++{
		fmt.Scan(&filaOrig[i])
	}

    fmt.Scan(&m)
	saiu := make(map[int]bool)
	for i := 0; i < m; i++{
		fmt.Scan(&id)
		saiu[id] = true
	}
	
	for _, pessoa := range filaOrig{
		if !saiu[pessoa]{
			fmt.Printf("%d ", pessoa)
		}
	}

	fmt.Println()

}