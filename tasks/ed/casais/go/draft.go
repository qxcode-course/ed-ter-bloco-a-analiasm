package main
import "fmt"
func main() {
var animaist int
fmt.Scan(&animaist)

naocasais := [50]int{}
qtdnaocasais := 0
casais := 0

for i := 0; i < animaist; i++{
    var animal int
    fmt.Scan(&animal)
    achou := false

    for j := 0; j < qtdnaocasais; j++{
        if naocasais[j] == -animal && naocasais[j] != 0{
            casais++
            naocasais[j] = 0
            achou = true
            break
        }
    }
    
    if !achou{
        naocasais[qtdnaocasais] = animal
        qtdnaocasais++
    }
}
    fmt.Printf("%d\n", casais)
}
