package main
import "fmt"
func main() {
    var competidores, pedraA, pedraB int

    campid := -1
    menordiferenca := 0

    fmt.Scanf("%d", &competidores)

    for i := 0; i < competidores; i++{
        fmt.Scanf("%d %d", &pedraA, pedraB)

        if pedraA >= 10 && pedraB >= 10 {
            diferenca := pedraA - pedraB

            if diferenca <0{
                diferenca = -diferenca
            }

            if campid == -1 || diferenca < menordiferenca{
                menordiferenca = diferenca
                campid = 1
            }
        }
    }
        if campid == -1{
            fmt.Printf("sem ganhador\n")
        }else{
            fmt.Printf("%d\n", &campid)
        }
}
