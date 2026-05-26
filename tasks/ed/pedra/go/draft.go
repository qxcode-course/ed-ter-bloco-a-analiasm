package main
import "fmt"

func main(){

var competidores int
var pedra1, pedra2 int

iddocampeao := -1
menordiferenca := 0

fmt.Scanf("%d", &competidores)
	    
for i := 0; i < competidores ; i++{
    fmt.Scanf("%d", &pedra1)
    fmt.Scanf("%d", &pedra2)

	if pedra1 >= 10 && pedra2 >= 10{
		diferenca := pedra1 - pedra2

		if diferenca < 0 {
			diferenca = -diferenca
		}

		if iddocampeao == -1 || diferenca < menordiferenca{
			menordiferenca = diferenca
			iddocampeao = i
		}
	}
}
 
    if iddocampeao == -1 {
		fmt.Printf("sem ganhador\n")
	}else{
		fmt.Printf("%d\n", iddocampeao)
	}

	
}