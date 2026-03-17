package main
import "fmt"

func main(){

var competidores int
var pedraA, pedraB int

campeaoid := -1
menordiferenca := 0

fmt.Scanf("%d", &competidores)
	    
for i := 0; i < competidores ; i++{
    fmt.Scanf("%d", &pedraA)
    fmt.Scanf("%d", &pedraB)

	if pedraA >= 10 && pedraB >= 10{
		diferenca := pedraA - pedraB

		if diferenca < 0 {
			diferenca = -diferenca
		}

		if campeaoid == -1 || diferenca < menordiferenca{
			menordiferenca = diferenca
			campeaoid = i
		}
	}
}
 
    if campeaoid == -1 {
		fmt.Printf("sem ganhador\n")
	}else{
		fmt.Printf("%d\n", campeaoid)
	}

}