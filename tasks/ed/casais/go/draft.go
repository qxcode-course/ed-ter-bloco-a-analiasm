package main
import "fmt"
func main() {
var animaist int
fmt.Scan(&animaist)

naosaocasais := [50]int{}
qtdnaosaocasais := 0
casais := 0

for i := 0; i < animaist; i++{
    var animal int
    fmt.Scan(&animal)
    achou := false

    for j := 0; j < qtdnaosaocasais; j++{
        if naosaocasais[j] == -animal && naosaocasais[j] != 0{
            casais++
            naosaocasais[j] = 0
            achou = true
            break
        }
    }
    
    if !achou{
        naosaocasais[qtdnaosaocasais] = animal
        qtdnaosaocasais++
    }
}
    fmt.Printf("%d\n", casais)
}
