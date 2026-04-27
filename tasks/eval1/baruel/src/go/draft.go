package main
import "fmt"
func main() {
    var qtdtalb, qtdbfig, i int
    var seq[100]int
    var duplicada = 0

    fmt.Scan(qtdtalb)
    fmt.Scan(qtdbfig)


    for i = 0; i < qtdbfig; i++{
    fmt.Scan(seq)
    }
        

    for i = 0; i < qtdbfig; i++{

        if seq[i] == seq[i + 1]{
            duplicada += 1
        }
        
        if duplicada == 0{
            fmt.Print("N")
        }
    }
        
    
}
