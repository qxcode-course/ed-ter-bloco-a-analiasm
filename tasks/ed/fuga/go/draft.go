package main
import "fmt"
func main() {
   
var H, P, F, D int

fmt.Scanf("%d", &H)
fmt.Scanf("%d", &P)
fmt.Scanf("%d", &F)
fmt.Scanf("%d", &D)


for {
    F += D

    if F > 15{
        F = 0
    }else if F < 0{
        F = 15
    }

    if F == H{
        fmt.Printf("S\n")
        break
    }

    if F == P{
        fmt.Printf("N\n")
        break
    }
}

}

