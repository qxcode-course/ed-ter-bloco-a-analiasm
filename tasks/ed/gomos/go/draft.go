package main
import "fmt"
type Ponto struct{
    x, y int
}
func main() {
    var q int
    var direcao string
    fmt.Scan(&q, &direcao)

    corpinho := make([]Ponto, q)
    
    for i := 0; i < q; i++{
        fmt.Scan(&corpinho[i].x, &corpinho[i].y)
    }

posicoesanteriores := make([]Ponto, q)
copy(posicoesanteriores, corpinho)

switch direcao{
case "U":
    corpinho[0].y--
case "D":
    corpinho[0].y++
case "L":
    corpinho[0].x--
case "R":
    corpinho[0].x++
}
      
for i := 1; i < q; i++{        
    corpinho[i] = posicoesanteriores[i-1]                                     
}

for _, c := range corpinho{
    fmt.Printf("%d %d\n", c.x, c.y)
}

}
