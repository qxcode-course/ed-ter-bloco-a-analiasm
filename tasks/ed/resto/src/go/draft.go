package main
import "fmt"
func divisaocresto(n int){
    if n == 0{
        return
    }
    
    divisaocresto(n/2)
    fmt.Printf("%d %d\n", n/2, n%2)
}

func main() {
    var n int
    fmt.Scan(&n)

    divisaocresto(n)
}
