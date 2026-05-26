package main
import "fmt"
type Ponto struct{
    x, y int
}
func main() {
    var q int
    var direcao string
    fmt.Scan(&q, &direcao)

    corpinhodebolinha := make([]Ponto, q)
    
    for i := 0; i < q; i++{
        fmt.Scan(&corpinhodebolinha[i].x, &corpinhodebolinha[i].y)
    }

posicoesanteriores := make([]Ponto, q)
copy(posicoesanteriores, corpinhodebolinha)

switch direcao{
case "U":
    corpinhodebolinha[0].y--
case "D":
    corpinhodebolinha[0].y++
case "L":
    corpinhodebolinha[0].x--
case "R":
    corpinhodebolinha[0].x++
}
      
for i := 1; i < q; i++{        
    corpinhodebolinha[i] = posicoesanteriores[i-1]                                     
}

for _, c := range corpinhodebolinha{
    fmt.Printf("%d %d\n", c.x, c.y)
}

}
